/*
 * JAMS REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: V6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type HistoryAPI service

/*
HistoryAPI Gets job execution history.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return []History
*/
func (a *HistoryAPI) History(ctx context.Context) ([]History, *http.Response, error) {
	var localVarReturnValue []History

	// create path and map variables
	apiPath := a.client.cfg.BasePath + "/api/history"

	headers := make(map[string]string)
	queryParams := url.Values{}

	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.prepareRequest(ctx, apiPath, "GET", nil, headers, queryParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	response, err := a.client.callAPI(r)
	if err != nil || response == nil {
		return localVarReturnValue, response, err
	}

	localVarBody, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return localVarReturnValue, response, err
	}

	if response.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, response.Header.Get("Content-Type"))
		return localVarReturnValue, response, err
	}

	if response.StatusCode >= 300 {
		newErr := Error{
			Body: localVarBody,
			Err:  response.Status,
		}

		if response.StatusCode == 200 {
			var v []History
			err = a.client.decode(&v, localVarBody, response.Header.Get("Content-Type"))
			if err != nil {
				newErr.Err = err.Error()
				return localVarReturnValue, response, newErr
			}
			newErr.Model = v
			return localVarReturnValue, response, newErr
		}

		return localVarReturnValue, response, newErr
	}

	return localVarReturnValue, response, nil
}

/*
HistoryAPI Gets job execution history using OData filters.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return PageResultHistory
*/
func (a *HistoryAPI) HistoryOdata(ctx context.Context) (PageResultHistory, *http.Response, error) {
	var localVarReturnValue PageResultHistory

	// create path and map variables
	apiPath := a.client.cfg.BasePath + "/api/history/odata"

	headers := make(map[string]string)
	queryParams := url.Values{}

	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.prepareRequest(ctx, apiPath, "GET", nil, headers, queryParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	response, err := a.client.callAPI(r)
	if err != nil || response == nil {
		return localVarReturnValue, response, err
	}

	localVarBody, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return localVarReturnValue, response, err
	}

	if response.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, response.Header.Get("Content-Type"))
		return localVarReturnValue, response, err
	}

	if response.StatusCode >= 300 {
		newErr := Error{
			Body: localVarBody,
			Err:  response.Status,
		}

		if response.StatusCode == 200 {
			var v PageResultHistory
			err = a.client.decode(&v, localVarBody, response.Header.Get("Content-Type"))
			if err != nil {
				newErr.Err = err.Error()
				return localVarReturnValue, response, newErr
			}
			newErr.Model = v
			return localVarReturnValue, response, newErr
		}

		return localVarReturnValue, response, newErr
	}

	return localVarReturnValue, response, nil
}

/*
HistoryAPI Gets log file as download or preview of first 512 kB.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param jobName The Job Name
 * @param ron The RON number
 * @param restartCount The Restart Count
 * @param optional nil or *HistoryApiHistoryGetJobLogOpts - Optional Parameters:
     * @param "IsPreview" (optional.Bool) -  If true, returns first 512 kB of log file. If false or omitted, returns log file as attachment

@return Object
*/

type HistoryApiHistoryGetJobLogOpts struct {
	IsPreview optional.Bool
}

func (a *HistoryAPI) JobLog(ctx context.Context, jobName string, ron int32, restartCount int32, localVarOptionals *HistoryApiHistoryGetJobLogOpts) (Object, *http.Response, error) {
	var localVarReturnValue Object

	// create path and map variables
	apiPath := a.client.cfg.BasePath + "/api/history/job/log/{jobName}/{ron}/{restartCount}"
	apiPath = strings.Replace(apiPath, "{"+"jobName"+"}", fmt.Sprintf("%v", jobName), -1)
	apiPath = strings.Replace(apiPath, "{"+"ron"+"}", fmt.Sprintf("%v", ron), -1)
	apiPath = strings.Replace(apiPath, "{"+"restartCount"+"}", fmt.Sprintf("%v", restartCount), -1)

	headers := make(map[string]string)
	queryParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.IsPreview.IsSet() {
		queryParams.Add("isPreview", parameterToString(localVarOptionals.IsPreview.Value(), ""))
	}
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.prepareRequest(ctx, apiPath, "GET", nil, headers, queryParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	response, err := a.client.callAPI(r)
	if err != nil || response == nil {
		return localVarReturnValue, response, err
	}

	localVarBody, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return localVarReturnValue, response, err
	}

	if response.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, response.Header.Get("Content-Type"))
		return localVarReturnValue, response, err
	}

	if response.StatusCode >= 300 {
		newErr := Error{
			Body: localVarBody,
			Err:  response.Status,
		}

		if response.StatusCode == 200 {
			var v Object
			err = a.client.decode(&v, localVarBody, response.Header.Get("Content-Type"))
			if err != nil {
				newErr.Err = err.Error()
				return localVarReturnValue, response, newErr
			}
			newErr.Model = v
			return localVarReturnValue, response, newErr
		}

		return localVarReturnValue, response, newErr
	}

	return localVarReturnValue, response, nil
}

/*
HistoryAPI Gets log file as download or preview of first 512 kB.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param setupName The Setup Name
 * @param ron The RON number
 * @param restartCount The Restart Count
 * @param optional nil or *HistoryApiHistoryGetSetupLogOpts - Optional Parameters:
     * @param "IsPreview" (optional.Bool) -  If true, returns first 512 kB of log file. If false or omitted, returns log file as attachment

@return Object
*/

type HistoryApiHistoryGetSetupLogOpts struct {
	IsPreview optional.Bool
}

func (a *HistoryAPI) SetupLog(ctx context.Context, setupName string, ron int32, restartCount int32, localVarOptionals *HistoryApiHistoryGetSetupLogOpts) (Object, *http.Response, error) {
	var localVarReturnValue Object

	// create path and map variables
	apiPath := a.client.cfg.BasePath + "/api/history/setup/log/{setupName}/{ron}/{restartCount}"
	apiPath = strings.Replace(apiPath, "{"+"setupName"+"}", fmt.Sprintf("%v", setupName), -1)
	apiPath = strings.Replace(apiPath, "{"+"ron"+"}", fmt.Sprintf("%v", ron), -1)
	apiPath = strings.Replace(apiPath, "{"+"restartCount"+"}", fmt.Sprintf("%v", restartCount), -1)

	headers := make(map[string]string)
	queryParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.IsPreview.IsSet() {
		queryParams.Add("isPreview", parameterToString(localVarOptionals.IsPreview.Value(), ""))
	}
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.prepareRequest(ctx, apiPath, "GET", nil, headers, queryParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	response, err := a.client.callAPI(r)
	if err != nil || response == nil {
		return localVarReturnValue, response, err
	}

	localVarBody, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return localVarReturnValue, response, err
	}

	if response.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, response.Header.Get("Content-Type"))
		return localVarReturnValue, response, err
	}

	if response.StatusCode >= 300 {
		newErr := Error{
			Body: localVarBody,
			Err:  response.Status,
		}

		if response.StatusCode == 200 {
			var v Object
			err = a.client.decode(&v, localVarBody, response.Header.Get("Content-Type"))
			if err != nil {
				newErr.Err = err.Error()
				return localVarReturnValue, response, newErr
			}
			newErr.Model = v
			return localVarReturnValue, response, newErr
		}

		return localVarReturnValue, response, newErr
	}

	return localVarReturnValue, response, nil
}
