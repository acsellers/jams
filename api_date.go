/*
 * JAMS REST API

 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)

 * API version: V6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package jams

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/url"
	"time"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type DateAPI service

/*
DateAPI Evaluate a natural language date and return the date that it identifies.

 * @param date The natural language date specification.
 * @param optional nil or *DateApiDateGetEvaluateDateOpts - Optional Parameters:
     * @param "Today" (optional.Time) -  The date that should be considered \&quot;today\&quot; when evaluating              the date specification.
     * @param "StartDate" (optional.Time) -  The starting date to be used for interval specification (i.e. every other monday)

returns time.Time
*/

type DateApiDateGetEvaluateDateOpts struct {
	Today     optional.Time
	StartDate optional.Time
}

func (a *DateAPI) EvaluateDate(ctx context.Context, date string, localVarOptionals *DateApiDateGetEvaluateDateOpts) (time.Time, error) {
	var returnValue time.Time

	// create path and map variables
	apiPath := a.client.cfg.BasePath + "/api/date/evaluate"

	headers := make(map[string]string)
	queryParams := url.Values{}

	queryParams.Add("date", parameterToString(date, ""))
	if localVarOptionals != nil && localVarOptionals.Today.IsSet() {
		queryParams.Add("today", parameterToString(localVarOptionals.Today.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StartDate.IsSet() {
		queryParams.Add("startDate", parameterToString(localVarOptionals.StartDate.Value(), ""))
	}
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
			var v time.Time
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
DateAPI Adds a date to an existing DateType or creates a new DateType if it doesn't exist.

 * @param dateTypeName The name of the DateType to add the date to.
 * @param date The date to add.

returns Object
*/
func (a *DateAPI) UpdateDate(ctx context.Context, dateTypeName string, date time.Time) (Object, error) {
	var returnValue Object

	// create path and map variables
	apiPath := fmt.Sprintf("%s/api/date/%s/%s", a.client.cfg.BasePath, dateTypeName, date.Format(time.RFC3339))

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.buildRequest(apiPath, "PUT", nil, headers, url.Values{})
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
