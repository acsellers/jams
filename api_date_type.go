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
)

// Linger please
var (
	_ context.Context
)

type DateTypeApiService service

/*
DateTypeApiService Deletes a DateType object.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param name The name of the DateType

@return Object
*/
func (a *DateTypeApiService) DateTypeDeleteDateType(ctx context.Context, name string) (Object, *http.Response, error) {
	var (
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue Object
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/datetype"

	headers := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("name", parameterToString(name, ""))
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.prepareRequest(ctx, localVarPath, "DELETE", localVarPostBody, headers, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
DateTypeApiService Gets a collection of all DateTypes
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return []DateType
*/
func (a *DateTypeApiService) DateTypeGetDateType(ctx context.Context) ([]DateType, *http.Response, error) {
	var (
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []DateType
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/datetype"

	headers := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.prepareRequest(ctx, localVarPath, "GET", localVarPostBody, headers, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
			var v []DateType
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
DateTypeApiService Get the DateType with the specified name.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param name The name of the DateType

@return DateType
*/
func (a *DateTypeApiService) DateTypeGetDateTypeByName(ctx context.Context, name string) (DateType, *http.Response, error) {
	var (
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue DateType
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/datetype/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	headers := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.prepareRequest(ctx, localVarPath, "GET", localVarPostBody, headers, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
			var v DateType
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
DateTypeApiService Get the DateType with the specified ID.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id

@return DateType
*/
func (a *DateTypeApiService) DateTypeGetDateTypeByID(ctx context.Context, id int32) (DateType, *http.Response, error) {
	var (
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue DateType
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/datetype/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	headers := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.prepareRequest(ctx, localVarPath, "GET", localVarPostBody, headers, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
			var v DateType
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
DateTypeApiService Creates a new DateType object.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dateType A Models.DateType object

@return Object
*/
func (a *DateTypeApiService) DateTypePostDateType(ctx context.Context, dateType DateType) (Object, *http.Response, error) {
	var (
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue Object
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/datetype/dateType"

	headers := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	// body params
	localVarPostBody = &dateType
	r, err := a.client.prepareRequest(ctx, localVarPath, "POST", localVarPostBody, headers, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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
DateTypeApiService Updates an existing DateType object or creates a new definition if it doesn&#39;t exist.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param dateType A Models.DateType object

@return Object
*/
func (a *DateTypeApiService) DateTypePutDateType(ctx context.Context, dateType DateType) (Object, *http.Response, error) {
	var (
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue Object
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/datetype/dateType"

	headers := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	// body params
	localVarPostBody = &dateType
	r, err := a.client.prepareRequest(ctx, localVarPath, "PUT", localVarPostBody, headers, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
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