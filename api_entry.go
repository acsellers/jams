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

func (a *EntryAPI) EntryDownload(ctx context.Context, entryId int, localVarOptionals *EntryApiEntryDownloadOpts) (Object, error) {
	var returnValue Object

	// create path and map variables
	apiPath := fmt.Sprintf("%s/api/entry/log/%d", a.client.cfg.BasePath, entryId)

	queryParams := url.Values{}
	if localVarOptionals != nil && localVarOptionals.IsPreview.IsSet() {
		queryParams.Add("isPreview", parameterToString(localVarOptionals.IsPreview.Value(), ""))
	}

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
func (a *EntryAPI) EntryParameter(ctx context.Context, id int, name string) (EntryParam, error) {
	var returnValue EntryParam

	// create path and map variables
	apiPath := fmt.Sprintf("%s/api/entry/%d/parameter/%s", a.client.cfg.BasePath, id, name)

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
func (a *EntryAPI) EntryParameters(ctx context.Context, id int) ([]EntryParam, error) {
	var returnValue []EntryParam

	// create path and map variables
	apiPath := fmt.Sprintf("%s/api/entry/%d/parameter", a.client.cfg.BasePath, id)

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
func (a *EntryAPI) EntryByID(ctx context.Context, id int) (Entry, error) {
	var returnValue Entry

	// create path and map variables
	apiPath := fmt.Sprintf("%s/api/entry/%d", a.client.cfg.BasePath, id)

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
func (a *EntryAPI) CancelEntry(ctx context.Context, id int, cancelEntry CancelEntry) error {
	var ()

	// create path and map variables
	apiPath := fmt.Sprintf("%s/api/entry/%d/cancel", a.client.cfg.BasePath, id)

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.prepareRequest(ctx, apiPath, "PUT", &cancelEntry, headers, url.Values{})
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
func (a *EntryAPI) UpdateEntryParameter(ctx context.Context, id int, name, value string) error {
	var ()

	// create path and map variables
	apiPath := fmt.Sprintf("%s/api/entry/%d/parameter/%s/%s", a.client.cfg.BasePath, id, name, value)

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.prepareRequest(ctx, apiPath, "PUT", nil, headers, url.Values{})
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
func (a *EntryAPI) RescheduleEntry(ctx context.Context, id int, rescheduleEntry RescheduleEntry) error {
	var ()

	// create path and map variables
	apiPath := fmt.Sprintf("%s/api/entry/%d/reschedule", a.client.cfg.BasePath, id)

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.prepareRequest(ctx, apiPath, "PUT", &rescheduleEntry, headers, url.Values{})
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
func (a *EntryAPI) RestartEntry(ctx context.Context, id int, restartEntry RestartEntry) error {
	var ()

	// create path and map variables
	apiPath := fmt.Sprintf("%s/api/entry/%d/restart", a.client.cfg.BasePath, id)

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.prepareRequest(ctx, apiPath, "PUT", &restartEntry, headers, url.Values{})
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
func (a *EntryAPI) UpdateStatus(ctx context.Context, id int, status string, icon string, message string, permanent bool) error {
	var ()

	// create path and map variables
	apiPath := fmt.Sprintf("%s/api/entry/%d/status", a.client.cfg.BasePath, id)

	queryParams := url.Values{}
	queryParams.Add("status", parameterToString(status, ""))
	queryParams.Add("icon", parameterToString(icon, ""))
	queryParams.Add("message", parameterToString(message, ""))
	queryParams.Add("permanent", parameterToString(permanent, ""))

	headers := make(map[string]string)
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
func (a *EntryAPI) HoldEntry(ctx context.Context, id int, holdEntry HoldEntry) error {
	var ()

	// create path and map variables
	apiPath := fmt.Sprintf("%s/api/entry/%d/hold", a.client.cfg.BasePath, id)

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.prepareRequest(ctx, apiPath, "PUT", &holdEntry, headers, url.Values{})
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
func (a *EntryAPI) ReleaseEntry(ctx context.Context, id int, releaseEntry ReleaseEntry) error {
	var ()

	// create path and map variables
	apiPath := fmt.Sprintf("%s/api/entry/%d/release", a.client.cfg.BasePath, id)

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.prepareRequest(ctx, apiPath, "PUT", &releaseEntry, headers, url.Values{})
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
