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

type JobAPI service

/*
JobAPI Deletes a Job by name.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param name The name of the Job to delete

@return string
*/
func (a *JobAPI) DeleteJob(ctx context.Context, name string) (string, error) {
	var returnValue string

	// create path and map variables
	apiPath := a.client.cfg.BasePath + "/api/job"

	queryParams := url.Values{}
	queryParams.Add("name", parameterToString(name, ""))

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.prepareRequest(ctx, apiPath, "DELETE", nil, headers, queryParams)
	if err != nil {
		return returnValue, err
	}

	response, err := a.client.Call(r)
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
			var v string
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
JobAPI Get the Job with the specified ID.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The ID of the Job.

@return Job
*/
func (a *JobAPI) JobByID(ctx context.Context, id int) (Job, error) {
	var returnValue Job

	// create path and map variables
	apiPath := fmt.Sprintf("%s/api/job/%d", a.client.cfg.BasePath, id)

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.prepareRequest(ctx, apiPath, "GET", nil, headers, url.Values{})
	if err != nil {
		return returnValue, err
	}

	response, err := a.client.Call(r)
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
			var v Job
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
JobAPI Get the Job with the specified name.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param name The name of the Job

@return Job
*/
func (a *JobAPI) JobByName(ctx context.Context, name string) (Job, error) {
	var returnValue Job

	// create path and map variables
	apiPath := fmt.Sprintf("%s/api/job/%s", a.client.cfg.BasePath, name)

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.prepareRequest(ctx, apiPath, "GET", nil, headers, url.Values{})
	if err != nil {
		return returnValue, err
	}

	response, err := a.client.Call(r)
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
			var v Job
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
JobAPI Gets a list of all jobs
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return []Job
*/
func (a *JobAPI) Jobs(ctx context.Context) ([]Job, error) {
	var returnValue []Job

	// create path and map variables
	apiPath := a.client.cfg.BasePath + "/api/job"

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.prepareRequest(ctx, apiPath, "GET", nil, headers, url.Values{})
	if err != nil {
		return returnValue, err
	}

	response, err := a.client.Call(r)
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
			var v []Job
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
JobAPI Gets Jobs in the Folder with the specified name
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param folderName The name of the Folder to return Jobs from.
 * @param name

@return []Job
*/
func (a *JobAPI) JobsByFolder(ctx context.Context, folderName string, name string) ([]Job, error) {
	var returnValue []Job

	// create path and map variables
	apiPath := fmt.Sprintf("%s/api/job/folder/%s", a.client.cfg.BasePath, name)

	queryParams := url.Values{}
	queryParams.Add("folderName", parameterToString(folderName, ""))

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.prepareRequest(ctx, apiPath, "GET", nil, headers, queryParams)
	if err != nil {
		return returnValue, err
	}

	response, err := a.client.Call(r)
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
			var v []Job
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
JobAPI Gets all jobs in the Folder with the specified ID
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id ID of the folder definition

@return []Job
*/
func (a *JobAPI) JobsByFolderID(ctx context.Context, id int) ([]Job, error) {
	var returnValue []Job

	// create path and map variables
	apiPath := fmt.Sprintf("%s/api/job/folder/%d", a.client.cfg.BasePath, id)

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.prepareRequest(ctx, apiPath, "GET", nil, headers, url.Values{})
	if err != nil {
		return returnValue, err
	}

	response, err := a.client.Call(r)
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
			var v []Job
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
JobAPI Creates a new Job object.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param job A Models.Job object

@return Job
*/
func (a *JobAPI) CreateJob(ctx context.Context, job Job) (Job, error) {
	var returnValue Job

	// create path and map variables
	apiPath := a.client.cfg.BasePath + "/api/job"

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.prepareRequest(ctx, apiPath, "POST", &job, headers, url.Values{})
	if err != nil {
		return returnValue, err
	}

	response, err := a.client.Call(r)
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
			var v Job
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
JobAPI Updates an existing Job object or creates new definition if it doesn't exist.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param job A Models.Job object

@return Job
*/
func (a *JobAPI) UpdateJob(ctx context.Context, job Job) (Job, error) {
	var returnValue Job

	// create path and map variables
	apiPath := a.client.cfg.BasePath + "/api/job"

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.prepareRequest(ctx, apiPath, "PUT", &job, headers, url.Values{})
	if err != nil {
		return returnValue, err
	}

	response, err := a.client.Call(r)
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
			var v Job
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
