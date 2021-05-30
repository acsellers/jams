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

type SubmitMenuAPI service

/*
SubmitMenuAPI Gets the full populated menu. This is starts from the Folders in the root.
 *

returns []NestedMenuEntry
*/
func (a *SubmitMenuAPI) NestedMenu(ctx context.Context) ([]NestedMenuEntry, error) {
	var returnValue []NestedMenuEntry

	// create path and map variables
	apiPath := a.client.cfg.BasePath + "/api/SubmitMenu/Nested"

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
			var v []NestedMenuEntry
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
SubmitMenuAPI Gets the default root level menu. This is created from the Folders in the root.
 *

returns []MenuEntry
*/
func (a *SubmitMenuAPI) SubmitMenu(ctx context.Context) ([]MenuEntry, error) {
	var returnValue []MenuEntry

	// create path and map variables
	apiPath := a.client.cfg.BasePath + "/api/SubmitMenu"

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
			var v []MenuEntry
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
SubmitMenuAPI Gets a submit menu.
 *
 * @param menuType Specifies which type of menu should be retrieved.
 * @param name Specifies the name of the folder or menu to be retrieved.

returns []MenuEntry
*/
func (a *SubmitMenuAPI) SubmitMenuByName(ctx context.Context, menuType, name string) ([]MenuEntry, error) {
	var returnValue []MenuEntry

	// create path and map variables
	apiPath := fmt.Sprintf("%s/api/SubmitMenu/%s/%s", a.client.cfg.BasePath, menuType, name)

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
			var v []MenuEntry
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
SubmitMenuAPI Gets the entries in a branch of a submit menu. Use GetSubmitMenuByName to get the top level menu              and use this to expand branches when they are selected.
 *
 * @param menuType Specifies which type of menu should be retrieved.
 * @param id The ID of the folder or menu that should be retrieved.

returns []MenuEntry
*/
func (a *SubmitMenuAPI) SubmitMenuByID(ctx context.Context, menuType string, id int) ([]MenuEntry, error) {
	var returnValue []MenuEntry

	// create path and map variables
	apiPath := fmt.Sprintf("%s/api/SubmitMenu/%s/%d", a.client.cfg.BasePath, menuType, id)

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
			var v []MenuEntry
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
