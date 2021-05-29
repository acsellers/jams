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

type EntryApiService service

/*
EntryApiService Gets log file as download or preview of first 512 kB.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param entryId The Entry Id
 * @param optional nil or *EntryApiEntryDownloadOpts - Optional Parameters:
     * @param "IsPreview" (optional.Bool) -  If true, returns first 512 kB of log file. If false or omitted, returns log file as attachment

@return Object
*/

type EntryApiEntryDownloadOpts struct {
	IsPreview optional.Bool
}

func (a *EntryApiService) EntryDownload(ctx context.Context, entryId int32, localVarOptionals *EntryApiEntryDownloadOpts) (Object, *http.Response, error) {
	var (
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue Object
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/entry/log/{entryId}"
	localVarPath = strings.Replace(localVarPath, "{"+"entryId"+"}", fmt.Sprintf("%v", entryId), -1)

	headers := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.IsPreview.IsSet() {
		localVarQueryParams.Add("isPreview", parameterToString(localVarOptionals.IsPreview.Value(), ""))
	}
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
EntryApiService Gets entry information for all entries in the current schedule.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return []Entry
*/
func (a *EntryApiService) EntryGetEntry(ctx context.Context) ([]Entry, *http.Response, error) {
	var (
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []Entry
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/entry"

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
			var v []Entry
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
EntryApiService Gets parameter value for a given entry.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The JAMS entry for which to get the parameter.
 * @param name The name of the parameter to return.

@return EntryParam
*/
func (a *EntryApiService) EntryGetEntryParameter(ctx context.Context, id int32, name string) (EntryParam, *http.Response, error) {
	var (
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue EntryParam
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/entry/{id}/parameter/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
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
			var v EntryParam
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
EntryApiService Gets parameter list for a given entry.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The JAMS entry for which to get parameters.

@return []EntryParam
*/
func (a *EntryApiService) EntryGetEntryParameters(ctx context.Context, id int32) ([]EntryParam, *http.Response, error) {
	var (
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []EntryParam
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/entry/{id}/parameter"
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
			var v []EntryParam
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
EntryApiService Gets Entry information with the specified ID
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The JAMS entry to get.

@return Entry
*/
func (a *EntryApiService) EntryGetEntryByID(ctx context.Context, id int32) (Entry, *http.Response, error) {
	var (
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue Entry
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/entry/{id}"
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
			var v Entry
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
EntryApiService Cancels a JAMS entry.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The JAMS entry to cancel.
 * @param cancelEntry A Models.CancelEntry.


*/
func (a *EntryApiService) EntryPutEntryCancel(ctx context.Context, id int32, cancelEntry CancelEntry) (*http.Response, error) {
	var (
		localVarPostBody  interface{}
		localVarFileName  string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/entry/{id}/cancel"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	headers := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	// body params
	localVarPostBody = &cancelEntry
	r, err := a.client.prepareRequest(ctx, localVarPath, "PUT", localVarPostBody, headers, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	response, err := a.client.callAPI(r)
	if err != nil || response == nil {
		return response, err
	}

	localVarBody, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return response, err
	}

	if response.StatusCode >= 300 {
		newErr := Error{
			Body: localVarBody,
			Err:  response.Status,
		}

		return response, newErr
	}

	return response, nil
}

/*
EntryApiService Sets parameter value for a given entry.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The JAMS entry to update.
 * @param name The name of the parameter to update.
 * @param value The new parameter value.


*/
func (a *EntryApiService) EntryPutEntryParameter(ctx context.Context, id int32, name string, value string) (*http.Response, error) {
	var (
		localVarPostBody  interface{}
		localVarFileName  string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/entry/{id}/parameter/{name}/{value}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"value"+"}", fmt.Sprintf("%v", value), -1)

	headers := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.prepareRequest(ctx, localVarPath, "PUT", localVarPostBody, headers, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	response, err := a.client.callAPI(r)
	if err != nil || response == nil {
		return response, err
	}

	localVarBody, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return response, err
	}

	if response.StatusCode >= 300 {
		newErr := Error{
			Body: localVarBody,
			Err:  response.Status,
		}

		return response, newErr
	}

	return response, nil
}

/*
EntryApiService Reschedules a JAMS entry.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The JAMS entry to reschedule.
 * @param rescheduleEntry A Models.RescheduleEntry.


*/
func (a *EntryApiService) EntryPutEntryReschedule(ctx context.Context, id int32, rescheduleEntry RescheduleEntry) (*http.Response, error) {
	var (
		localVarPostBody  interface{}
		localVarFileName  string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/entry/{id}/reschedule"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	headers := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	// body params
	localVarPostBody = &rescheduleEntry
	r, err := a.client.prepareRequest(ctx, localVarPath, "PUT", localVarPostBody, headers, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	response, err := a.client.callAPI(r)
	if err != nil || response == nil {
		return response, err
	}

	localVarBody, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return response, err
	}

	if response.StatusCode >= 300 {
		newErr := Error{
			Body: localVarBody,
			Err:  response.Status,
		}

		return response, newErr
	}

	return response, nil
}

/*
EntryApiService Restarts a JAMS entry.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The JAMS entry to restart.
 * @param restartEntry A Models.RestartEntry.


*/
func (a *EntryApiService) EntryPutEntryRestart(ctx context.Context, id int32, restartEntry RestartEntry) (*http.Response, error) {
	var (
		localVarPostBody  interface{}
		localVarFileName  string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/entry/{id}/restart"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	headers := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	// body params
	localVarPostBody = &restartEntry
	r, err := a.client.prepareRequest(ctx, localVarPath, "PUT", localVarPostBody, headers, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	response, err := a.client.callAPI(r)
	if err != nil || response == nil {
		return response, err
	}

	localVarBody, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return response, err
	}

	if response.StatusCode >= 300 {
		newErr := Error{
			Body: localVarBody,
			Err:  response.Status,
		}

		return response, newErr
	}

	return response, nil
}

/*
EntryApiService Sets the status message for an entry.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The JAMS entry to update.
 * @param status The new status message.
 * @param icon The new icon.
 * @param message A tooltip message for this entry.
 * @param permanent When true, the status change is permanent.


*/
func (a *EntryApiService) EntryPutEntryStatus(ctx context.Context, id int32, status string, icon string, message string, permanent bool) (*http.Response, error) {
	var (
		localVarPostBody  interface{}
		localVarFileName  string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/entry/{id}/status"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	headers := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("status", parameterToString(status, ""))
	localVarQueryParams.Add("icon", parameterToString(icon, ""))
	localVarQueryParams.Add("message", parameterToString(message, ""))
	localVarQueryParams.Add("permanent", parameterToString(permanent, ""))
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.prepareRequest(ctx, localVarPath, "PUT", localVarPostBody, headers, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	response, err := a.client.callAPI(r)
	if err != nil || response == nil {
		return response, err
	}

	localVarBody, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return response, err
	}

	if response.StatusCode >= 300 {
		newErr := Error{
			Body: localVarBody,
			Err:  response.Status,
		}

		return response, newErr
	}

	return response, nil
}

/*
EntryApiService Holds an Entry with the specified ID, with the              specified HoldEntry Audit Comment.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The JAMS entry to hold.
 * @param holdEntry A Models.HoldEntry.


*/
func (a *EntryApiService) EntryPutHoldEntry(ctx context.Context, id int32, holdEntry HoldEntry) (*http.Response, error) {
	var (
		localVarPostBody  interface{}
		localVarFileName  string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/entry/{id}/hold"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	headers := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	// body params
	localVarPostBody = &holdEntry
	r, err := a.client.prepareRequest(ctx, localVarPath, "PUT", localVarPostBody, headers, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	response, err := a.client.callAPI(r)
	if err != nil || response == nil {
		return response, err
	}

	localVarBody, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return response, err
	}

	if response.StatusCode >= 300 {
		newErr := Error{
			Body: localVarBody,
			Err:  response.Status,
		}

		return response, newErr
	}

	return response, nil
}

/*
EntryApiService Releases a CurJob with the specified ID              to Run Again, with specified ReleaseEntry Audit Comment
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The JAMS entry to release.
 * @param releaseEntry A Models.ReleaseEntry.


*/
func (a *EntryApiService) EntryPutReleaseEntry(ctx context.Context, id int32, releaseEntry ReleaseEntry) (*http.Response, error) {
	var (
		localVarPostBody  interface{}
		localVarFileName  string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/entry/{id}/release"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	headers := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	// body params
	localVarPostBody = &releaseEntry
	r, err := a.client.prepareRequest(ctx, localVarPath, "PUT", localVarPostBody, headers, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	response, err := a.client.callAPI(r)
	if err != nil || response == nil {
		return response, err
	}

	localVarBody, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return response, err
	}

	if response.StatusCode >= 300 {
		newErr := Error{
			Body: localVarBody,
			Err:  response.Status,
		}

		return response, newErr
	}

	return response, nil
}
