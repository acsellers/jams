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
	"strings"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type EntryAPI service

/*
EntryAPI Gets log file as download or preview of first 512 kB.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param entryId The Entry Id
 * @param optional nil or *EntryApiEntryDownloadOpts - Optional Parameters:
     * @param "IsPreview" (optional.Bool) -  If true, returns first 512 kB of log file. If false or omitted, returns log file as attachment

@return Object
*/

type EntryApiEntryDownloadOpts struct {
	IsPreview optional.Bool
}

func (a *EntryAPI) EntryDownload(ctx context.Context, entryId int32, localVarOptionals *EntryApiEntryDownloadOpts) (Object, error) {
	var returnValue Object

	// create path and map variables
	apiPath := a.client.cfg.BasePath + "/api/entry/log/{entryId}"
	apiPath = strings.Replace(apiPath, "{"+"entryId"+"}", fmt.Sprintf("%v", entryId), -1)

	headers := make(map[string]string)
	queryParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.IsPreview.IsSet() {
		queryParams.Add("isPreview", parameterToString(localVarOptionals.IsPreview.Value(), ""))
	}
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
			var v Object
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
EntryAPI Gets entry information for all entries in the current schedule.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return []Entry
*/
func (a *EntryAPI) Entries(ctx context.Context) ([]Entry, error) {
	var returnValue []Entry

	// create path and map variables
	apiPath := a.client.cfg.BasePath + "/api/entry"

	headers := make(map[string]string)
	queryParams := url.Values{}

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
			var v []Entry
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
EntryAPI Gets parameter value for a given entry.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The JAMS entry for which to get the parameter.
 * @param name The name of the parameter to return.

@return EntryParam
*/
func (a *EntryAPI) EntryParameter(ctx context.Context, id int32, name string) (EntryParam, error) {
	var returnValue EntryParam

	// create path and map variables
	apiPath := a.client.cfg.BasePath + "/api/entry/{id}/parameter/{name}"
	apiPath = strings.Replace(apiPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	apiPath = strings.Replace(apiPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	headers := make(map[string]string)
	queryParams := url.Values{}

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
			var v EntryParam
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
EntryAPI Gets parameter list for a given entry.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The JAMS entry for which to get parameters.

@return []EntryParam
*/
func (a *EntryAPI) EntryParameters(ctx context.Context, id int32) ([]EntryParam, error) {
	var returnValue []EntryParam

	// create path and map variables
	apiPath := a.client.cfg.BasePath + "/api/entry/{id}/parameter"
	apiPath = strings.Replace(apiPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	headers := make(map[string]string)
	queryParams := url.Values{}

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
			var v []EntryParam
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
EntryAPI Gets Entry information with the specified ID
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The JAMS entry to get.

@return Entry
*/
func (a *EntryAPI) EntryByID(ctx context.Context, id int32) (Entry, error) {
	var returnValue Entry

	// create path and map variables
	apiPath := a.client.cfg.BasePath + "/api/entry/{id}"
	apiPath = strings.Replace(apiPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	headers := make(map[string]string)
	queryParams := url.Values{}

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
			var v Entry
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
EntryAPI Cancels a JAMS entry.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The JAMS entry to cancel.
 * @param cancelEntry A Models.CancelEntry.


*/
func (a *EntryAPI) CancelEntry(ctx context.Context, id int32, cancelEntry CancelEntry) error {
	var ()

	// create path and map variables
	apiPath := a.client.cfg.BasePath + "/api/entry/{id}/cancel"
	apiPath = strings.Replace(apiPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	headers := make(map[string]string)
	queryParams := url.Values{}

	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.prepareRequest(ctx, apiPath, "PUT", &cancelEntry, headers, queryParams)
	if err != nil {
		return err
	}

	response, err := a.client.Call(r)
	if err != nil || response == nil {
		return err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return err
	}

	if response.StatusCode >= 300 {
		newErr := Error{
			Body: body,
			Err:  response.Status,
		}

		return newErr
	}

	return nil
}

/*
EntryAPI Sets parameter value for a given entry.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The JAMS entry to update.
 * @param name The name of the parameter to update.
 * @param value The new parameter value.


*/
func (a *EntryAPI) UpdateEntryParameter(ctx context.Context, id int32, name string, value string) error {
	var ()

	// create path and map variables
	apiPath := a.client.cfg.BasePath + "/api/entry/{id}/parameter/{name}/{value}"
	apiPath = strings.Replace(apiPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	apiPath = strings.Replace(apiPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	apiPath = strings.Replace(apiPath, "{"+"value"+"}", fmt.Sprintf("%v", value), -1)

	headers := make(map[string]string)
	queryParams := url.Values{}

	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.prepareRequest(ctx, apiPath, "PUT", nil, headers, queryParams)
	if err != nil {
		return err
	}

	response, err := a.client.Call(r)
	if err != nil || response == nil {
		return err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return err
	}

	if response.StatusCode >= 300 {
		newErr := Error{
			Body: body,
			Err:  response.Status,
		}

		return newErr
	}

	return nil
}

/*
EntryAPI Reschedules a JAMS entry.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The JAMS entry to reschedule.
 * @param rescheduleEntry A Models.RescheduleEntry.


*/
func (a *EntryAPI) RescheduleEntry(ctx context.Context, id int32, rescheduleEntry RescheduleEntry) error {
	var ()

	// create path and map variables
	apiPath := a.client.cfg.BasePath + "/api/entry/{id}/reschedule"
	apiPath = strings.Replace(apiPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	headers := make(map[string]string)
	queryParams := url.Values{}

	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.prepareRequest(ctx, apiPath, "PUT", &rescheduleEntry, headers, queryParams)
	if err != nil {
		return err
	}

	response, err := a.client.Call(r)
	if err != nil || response == nil {
		return err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return err
	}

	if response.StatusCode >= 300 {
		newErr := Error{
			Body: body,
			Err:  response.Status,
		}

		return newErr
	}

	return nil
}

/*
EntryAPI Restarts a JAMS entry.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The JAMS entry to restart.
 * @param restartEntry A Models.RestartEntry.


*/
func (a *EntryAPI) RestartEntry(ctx context.Context, id int32, restartEntry RestartEntry) error {
	var ()

	// create path and map variables
	apiPath := a.client.cfg.BasePath + "/api/entry/{id}/restart"
	apiPath = strings.Replace(apiPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	headers := make(map[string]string)
	queryParams := url.Values{}

	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.prepareRequest(ctx, apiPath, "PUT", &restartEntry, headers, queryParams)
	if err != nil {
		return err
	}

	response, err := a.client.Call(r)
	if err != nil || response == nil {
		return err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return err
	}

	if response.StatusCode >= 300 {
		newErr := Error{
			Body: body,
			Err:  response.Status,
		}

		return newErr
	}

	return nil
}

/*
EntryAPI Sets the status message for an entry.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The JAMS entry to update.
 * @param status The new status message.
 * @param icon The new icon.
 * @param message A tooltip message for this entry.
 * @param permanent When true, the status change is permanent.


*/
func (a *EntryAPI) UpdateStatus(ctx context.Context, id int32, status string, icon string, message string, permanent bool) error {
	var ()

	// create path and map variables
	apiPath := a.client.cfg.BasePath + "/api/entry/{id}/status"
	apiPath = strings.Replace(apiPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	headers := make(map[string]string)
	queryParams := url.Values{}

	queryParams.Add("status", parameterToString(status, ""))
	queryParams.Add("icon", parameterToString(icon, ""))
	queryParams.Add("message", parameterToString(message, ""))
	queryParams.Add("permanent", parameterToString(permanent, ""))
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.prepareRequest(ctx, apiPath, "PUT", nil, headers, queryParams)
	if err != nil {
		return err
	}

	response, err := a.client.Call(r)
	if err != nil || response == nil {
		return err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return err
	}

	if response.StatusCode >= 300 {
		newErr := Error{
			Body: body,
			Err:  response.Status,
		}

		return newErr
	}

	return nil
}

/*
EntryAPI Holds an Entry with the specified ID, with the              specified HoldEntry Audit Comment.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The JAMS entry to hold.
 * @param holdEntry A Models.HoldEntry.


*/
func (a *EntryAPI) HoldEntry(ctx context.Context, id int32, holdEntry HoldEntry) error {
	var ()

	// create path and map variables
	apiPath := a.client.cfg.BasePath + "/api/entry/{id}/hold"
	apiPath = strings.Replace(apiPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	headers := make(map[string]string)
	queryParams := url.Values{}

	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.prepareRequest(ctx, apiPath, "PUT", &holdEntry, headers, queryParams)
	if err != nil {
		return err
	}

	response, err := a.client.Call(r)
	if err != nil || response == nil {
		return err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return err
	}

	if response.StatusCode >= 300 {
		newErr := Error{
			Body: body,
			Err:  response.Status,
		}

		return newErr
	}

	return nil
}

/*
EntryAPI Releases a CurJob with the specified ID              to Run Again, with specified ReleaseEntry Audit Comment
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id The JAMS entry to release.
 * @param releaseEntry A Models.ReleaseEntry.


*/
func (a *EntryAPI) ReleaseEntry(ctx context.Context, id int32, releaseEntry ReleaseEntry) error {
	var ()

	// create path and map variables
	apiPath := a.client.cfg.BasePath + "/api/entry/{id}/release"
	apiPath = strings.Replace(apiPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	headers := make(map[string]string)
	queryParams := url.Values{}

	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.prepareRequest(ctx, apiPath, "PUT", &releaseEntry, headers, queryParams)
	if err != nil {
		return err
	}

	response, err := a.client.Call(r)
	if err != nil || response == nil {
		return err
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return err
	}

	if response.StatusCode >= 300 {
		newErr := Error{
			Body: body,
			Err:  response.Status,
		}

		return newErr
	}

	return nil
}
