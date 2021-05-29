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

type SubmitAPI service

/*
SubmitAPI Get the submit information for a Job or Setup with the specified name.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param name The qualified name of the Job or Setup.

@return SubmitInfo
*/
func (a *SubmitAPI) SubmitInfoByName(ctx context.Context, name string) (SubmitInfo, *http.Response, error) {
	var localVarReturnValue SubmitInfo

	// create path and map variables
	apiPath := a.client.cfg.BasePath + "/api/submit"

	headers := make(map[string]string)
	queryParams := url.Values{}

	queryParams.Add("name", parameterToString(name, ""))
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
			var v SubmitInfo
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
SubmitAPI Get the submit info for a Setup with the specified ID
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The ID of the Setup.

@return SubmitInfo
*/
func (a *SubmitAPI) SubmitInfoBySetupID(ctx context.Context, id int32) (SubmitInfo, *http.Response, error) {
	var localVarReturnValue SubmitInfo

	// create path and map variables
	apiPath := a.client.cfg.BasePath + "/api/submit/setup/{id}"
	apiPath = strings.Replace(apiPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

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
			var v SubmitInfo
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
SubmitAPI Get the submit info for a Job with the specified ID
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The ID of the Job.

@return SubmitInfo
*/
func (a *SubmitAPI) SubmitInfoByJobID(ctx context.Context, id int32) (SubmitInfo, *http.Response, error) {
	var localVarReturnValue SubmitInfo

	// create path and map variables
	apiPath := a.client.cfg.BasePath + "/api/submit/job/{id}"
	apiPath = strings.Replace(apiPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

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
			var v SubmitInfo
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
SubmitAPI Submit the job specified in the SubmitInfo object. Usually, you create the SubmitInfo object              by calling GetSubmit which returns a SubmitInfo. You can then modify values in the SubmitInfo              before calling PutSubmit.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param si A SubmitInfo object

@return SubmitInfo
*/
func (a *SubmitAPI) CreateSubmitInfo(ctx context.Context, si SubmitInfo) (SubmitInfo, *http.Response, error) {
	var localVarReturnValue SubmitInfo

	// create path and map variables
	apiPath := a.client.cfg.BasePath + "/api/submit"

	headers := make(map[string]string)
	queryParams := url.Values{}

	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.prepareRequest(ctx, apiPath, "POST", &si, headers, queryParams)
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
			var v SubmitInfo
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
