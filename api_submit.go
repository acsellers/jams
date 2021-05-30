/*
 * JAMS REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: V6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package jams

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/url"
)

// Linger please
var (
	_ context.Context
)

type SubmitAPI service

/*
SubmitAPI Get the submit information for a Job or Setup with the specified name.
 *
 * @param name The qualified name of the Job or Setup.

returns SubmitInfo
*/
func (a *SubmitAPI) SubmitInfoByName(ctx context.Context, name string) (SubmitInfo, error) {
	var returnValue SubmitInfo

	// create path and map variables
	apiPath := a.client.cfg.BasePath + "/api/submit"

	queryParams := url.Values{}
	queryParams.Add("name", parameterToString(name, ""))

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.buildRequest(apiPath, "GET", nil, headers, queryParams)
	if err != nil {
		return returnValue, err
	}

	response, err := a.client.Call(ctx, r)
	if err != nil || response == nil {
		return returnValue, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return returnValue, err
	}

	if response.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&returnValue, body, response.Header.Get("Content-Type"))
		return returnValue, err
	}

	if response.StatusCode >= 300 {
		newErr := Error{
			Body: body,
			Err:  response.Status,
		}

		if response.StatusCode == 200 {
			var v SubmitInfo
			err = a.client.decode(&v, body, response.Header.Get("Content-Type"))
			if err != nil {
				newErr.Err = err.Error()
				return returnValue, newErr
			}
			newErr.Model = v
			return returnValue, newErr
		}

		return returnValue, newErr
	}

	return returnValue, nil
}

/*
SubmitAPI Get the submit info for a Setup with the specified ID
 *
 * @param id The ID of the Setup.

returns SubmitInfo
*/
func (a *SubmitAPI) SubmitInfoBySetupID(ctx context.Context, id int) (SubmitInfo, error) {
	var returnValue SubmitInfo

	// create path and map variables
	apiPath := fmt.Sprintf("%s/api/submit/setup/%d", a.client.cfg.BasePath, id)

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.buildRequest(apiPath, "GET", nil, headers, url.Values{})
	if err != nil {
		return returnValue, err
	}

	response, err := a.client.Call(ctx, r)
	if err != nil || response == nil {
		return returnValue, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return returnValue, err
	}

	if response.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&returnValue, body, response.Header.Get("Content-Type"))
		return returnValue, err
	}

	if response.StatusCode >= 300 {
		newErr := Error{
			Body: body,
			Err:  response.Status,
		}

		if response.StatusCode == 200 {
			var v SubmitInfo
			err = a.client.decode(&v, body, response.Header.Get("Content-Type"))
			if err != nil {
				newErr.Err = err.Error()
				return returnValue, newErr
			}
			newErr.Model = v
			return returnValue, newErr
		}

		return returnValue, newErr
	}

	return returnValue, nil
}

/*
SubmitAPI Get the submit info for a Job with the specified ID
 *
 * @param id The ID of the Job.

returns SubmitInfo
*/
func (a *SubmitAPI) SubmitInfoByJobID(ctx context.Context, id int) (SubmitInfo, error) {
	var returnValue SubmitInfo

	// create path and map variables
	apiPath := fmt.Sprintf("%s/api/submit/job/%d", a.client.cfg.BasePath, id)

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.buildRequest(apiPath, "GET", nil, headers, url.Values{})
	if err != nil {
		return returnValue, err
	}

	response, err := a.client.Call(ctx, r)
	if err != nil || response == nil {
		return returnValue, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return returnValue, err
	}

	if response.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&returnValue, body, response.Header.Get("Content-Type"))
		return returnValue, err
	}

	if response.StatusCode >= 300 {
		newErr := Error{
			Body: body,
			Err:  response.Status,
		}

		if response.StatusCode == 200 {
			var v SubmitInfo
			err = a.client.decode(&v, body, response.Header.Get("Content-Type"))
			if err != nil {
				newErr.Err = err.Error()
				return returnValue, newErr
			}
			newErr.Model = v
			return returnValue, newErr
		}

		return returnValue, newErr
	}

	return returnValue, nil
}

/*
SubmitAPI Submit the job specified in the SubmitInfo object. Usually, you create the SubmitInfo object              by calling GetSubmit which returns a SubmitInfo. You can then modify values in the SubmitInfo              before calling PutSubmit.
 *
 * @param si A SubmitInfo object

returns SubmitInfo
*/
func (a *SubmitAPI) CreateSubmitInfo(ctx context.Context, si SubmitInfo) (SubmitInfo, error) {
	var returnValue SubmitInfo

	// create path and map variables
	apiPath := a.client.cfg.BasePath + "/api/submit"

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.buildRequest(apiPath, "POST", &si, headers, url.Values{})
	if err != nil {
		return returnValue, err
	}

	response, err := a.client.Call(ctx, r)
	if err != nil || response == nil {
		return returnValue, err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return returnValue, err
	}

	if response.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&returnValue, body, response.Header.Get("Content-Type"))
		return returnValue, err
	}

	if response.StatusCode >= 300 {
		newErr := Error{
			Body: body,
			Err:  response.Status,
		}

		if response.StatusCode == 200 {
			var v SubmitInfo
			err = a.client.decode(&v, body, response.Header.Get("Content-Type"))
			if err != nil {
				newErr.Err = err.Error()
				return returnValue, newErr
			}
			newErr.Model = v
			return returnValue, newErr
		}

		return returnValue, newErr
	}

	return returnValue, nil
}
