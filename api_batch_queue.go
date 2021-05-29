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

type BatchQueueApiService service

/*
BatchQueueApiService Delete the batch queue definition by name.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param name Name of the batch queue to delete.

@return string
*/
func (a *BatchQueueApiService) BatchQueueDeleteBatchQueue(ctx context.Context, name string) (string, *http.Response, error) {
	var (
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue string
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/batchqueue/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	headers := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
			var v string
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
BatchQueueApiService Gets all batch queue definitions
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return []BatchQueue
*/
func (a *BatchQueueApiService) BatchQueueGetBatchQueue(ctx context.Context) ([]BatchQueue, *http.Response, error) {
	var (
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []BatchQueue
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/batchqueue"

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
			var v []BatchQueue
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
BatchQueueApiService Gets a batch queue definition by ID
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id ID of the BatchQueue

@return BatchQueue
*/
func (a *BatchQueueApiService) BatchQueueGetBatchQueueByID(ctx context.Context, id int32) (BatchQueue, *http.Response, error) {
	var (
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue BatchQueue
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/batchqueue/{id}"
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
			var v BatchQueue
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
BatchQueueApiService Gets a batch queue definition by name
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param name Name of the BatchQueue

@return BatchQueue
*/
func (a *BatchQueueApiService) BatchQueueGetBatchQueueByName(ctx context.Context, name string) (BatchQueue, *http.Response, error) {
	var (
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue BatchQueue
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/batchqueue/{name}"
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
			var v BatchQueue
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
BatchQueueApiService Adds a new batch queue definition
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param batchqueue The BatchQueue object.

@return BatchQueue
*/
func (a *BatchQueueApiService) BatchQueuePostBatchQueue(ctx context.Context, batchqueue BatchQueue) (BatchQueue, *http.Response, error) {
	var (
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue BatchQueue
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/batchqueue"

	headers := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	// body params
	localVarPostBody = &batchqueue
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
			var v BatchQueue
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
BatchQueueApiService Updates an existing batch queue
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param batchqueue Batch queue definition to be updated

@return BatchQueue
*/
func (a *BatchQueueApiService) BatchQueuePutBatchQueue(ctx context.Context, batchqueue BatchQueue) (BatchQueue, *http.Response, error) {
	var (
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue BatchQueue
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/batchqueue"

	headers := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	// body params
	localVarPostBody = &batchqueue
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
			var v BatchQueue
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
BatchQueueApiService Starts a  batch queue
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param name Name of the queue to start

@return string
*/
func (a *BatchQueueApiService) BatchQueueStartBatchQueue(ctx context.Context, name string) (string, *http.Response, error) {
	var (
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue string
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/batchqueue/start/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	headers := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

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
			var v string
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
BatchQueueApiService Stops a batch queue
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param name Name of the queue to stop

@return string
*/
func (a *BatchQueueApiService) BatchQueueStopBatchQueue(ctx context.Context, name string) (string, *http.Response, error) {
	var (
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue string
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/batchqueue/stop/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	headers := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

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
			var v string
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