package jams

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

var (
	jsonCheck = regexp.MustCompile("(?i:(?:application|text)/json)")
	xmlCheck  = regexp.MustCompile("(?i:(?:application|text)/xml)")
)

// APIClient manages communication with the JAMS REST API API vV6
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.
	auth   *AccessToken

	// API Services
	Agent        *AgentAPI
	BatchQueue   *BatchQueueAPI
	Config       *ConfigAPI
	Date         *DateAPI
	DateType     *DateTypeAPI
	Entry        *EntryAPI
	Folder       *FolderAPI
	History      *HistoryAPI
	Job          *JobAPI
	Resource     *ResourceAPI
	Setup        *SetupAPI
	Submit       *SubmitAPI
	SubmitMenu   *SubmitMenuAPI
	Trigger      *TriggerAPI
	UserSecurity *UserSecurityAPI
	Variable     *VariableAPI
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.Agent = (*AgentAPI)(&c.common)
	c.BatchQueue = (*BatchQueueAPI)(&c.common)
	c.Config = (*ConfigAPI)(&c.common)
	c.Date = (*DateAPI)(&c.common)
	c.DateType = (*DateTypeAPI)(&c.common)
	c.Entry = (*EntryAPI)(&c.common)
	c.Folder = (*FolderAPI)(&c.common)
	c.History = (*HistoryAPI)(&c.common)
	c.Job = (*JobAPI)(&c.common)
	c.Resource = (*ResourceAPI)(&c.common)
	c.Setup = (*SetupAPI)(&c.common)
	c.Submit = (*SubmitAPI)(&c.common)
	c.SubmitMenu = (*SubmitMenuAPI)(&c.common)
	c.Trigger = (*TriggerAPI)(&c.common)
	c.UserSecurity = (*UserSecurityAPI)(&c.common)
	c.Variable = (*VariableAPI)(&c.common)

	return c
}

// parameterToString convert interface{} parameters to string, using a delimiter if format is provided.
func parameterToString(obj interface{}, collectionFormat string) string {
	var delimiter string

	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	}

	return fmt.Sprintf("%v", obj)
}

// Call do the request.
func (c *APIClient) Call(ctx context.Context, request *http.Request) (*http.Response, error) {
	if ctx != nil {
		request = request.WithContext(ctx)
	}
	if c.auth != nil {
		request.Header.Add("Authorization", "Bearer "+c.auth.AccessToken)
	}
	return c.cfg.HTTPClient.Do(request)
}

// buildRequest build the request
func (c *APIClient) buildRequest(
	path, method string,
	postBody interface{},
	headers map[string]string,
	queryParams url.Values) (localVarRequest *http.Request, err error) {

	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headers["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headers["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	url.Scheme = c.cfg.Scheme
	url.Host = c.cfg.Host

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = query.Encode()

	// Generate a new request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, url.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headers) > 0 {
		finalheaders := http.Header{}
		for h, v := range headers {
			finalheaders.Set(h, v)
		}
		localVarRequest.Header = finalheaders
	}

	// Override request host, if applicable
	if c.cfg.Host != "" {
		localVarRequest.Host = c.cfg.Host
	}

	// Add the user agent to the request.
	localVarRequest.Header.Add("User-Agent", c.cfg.UserAgent)

	for header, value := range c.cfg.DefaultHeader {
		localVarRequest.Header.Add(header, value)
	}

	return localVarRequest, nil
}

func (c *APIClient) decode(v interface{}, b []byte, contentType string) (err error) {
	if strings.Contains(contentType, "application/xml") {
		if err = xml.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	} else if strings.Contains(contentType, "application/json") {
		b = bytes.ReplaceAll(b, []byte(`T00:00:00"`), []byte(`T00:00:00Z"`))
		if err = json.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	}
	if sp, ok := v.(*string); ok && utf8.Valid(b) {
		*sp = string(b)
		return nil
	}

	return errors.New("undefined response type")
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if jsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if xmlCheck.MatchString(contentType) {
		xml.NewEncoder(bodyBuf).Encode(body)
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("invalid body type: %s", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		}
		expires = now.Add(lifetime)
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

// Error Provides access to the body, error and model on returned errors.
type Error struct {
	Body  []byte
	Err   string
	Model interface{}
}

// Error returns non-empty string if there was an error.
func (e Error) Error() string {
	return e.Err
}
