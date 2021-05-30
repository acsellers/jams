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
)

// Linger please
var (
	_ context.Context
)

type UserSecurityAPI service

/*
UserSecurityAPI Deletes a UserSecurity object.

 * @param name The name of the UserSecurity to delete.

returns string
*/
func (a *UserSecurityAPI) DeleteUser(ctx context.Context, name string) (string, error) {
	var returnValue string

	// create path and map variables
	apiPath := fmt.Sprintf("%s/api/usersecurity/%s", a.client.cfg.BasePath, name)

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.buildRequest(apiPath, "DELETE", nil, headers, url.Values{})
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
UserSecurityAPI Gets a collection of all users


returns []UserSecurity
*/
func (a *UserSecurityAPI) Users(ctx context.Context) ([]UserSecurity, error) {
	var returnValue []UserSecurity

	// create path and map variables
	apiPath := a.client.cfg.BasePath + "/api/usersecurity"

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
			var v []UserSecurity
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
UserSecurityAPI Get the UserSecurity with the specified name.

 * @param name The name of the UserSecurity.

returns UserSecurity
*/
func (a *UserSecurityAPI) UserByName(ctx context.Context, name string) (UserSecurity, error) {
	var returnValue UserSecurity

	// create path and map variables
	apiPath := fmt.Sprintf("%s/api/usersecurity/%s", a.client.cfg.BasePath, name)

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
			var v UserSecurity
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
UserSecurityAPI Get the User with the specified ID.

 * @param id The id of the UserSecurity.

returns UserSecurity
*/
func (a *UserSecurityAPI) UserByID(ctx context.Context, id int) (UserSecurity, error) {
	var returnValue UserSecurity

	// create path and map variables
	apiPath := fmt.Sprintf("%s/api/usersecurity/%d", a.client.cfg.BasePath, id)

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
			var v UserSecurity
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
UserSecurityAPI Creates a new UserSecurity object.

 * @param user A UserSecurity object

returns UserSecurity
*/
func (a *UserSecurityAPI) CreateUser(ctx context.Context, user UserSecurity) (UserSecurity, error) {
	var returnValue UserSecurity

	// create path and map variables
	apiPath := a.client.cfg.BasePath + "/api/usersecurity/user"

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.buildRequest(apiPath, "POST", &user, headers, url.Values{})
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
			var v UserSecurity
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
UserSecurityAPI Updates an existing UserSecurity object or creates new definition if it doesn't exist.

 * @param user A UserSecurity object

returns UserSecurity
*/
func (a *UserSecurityAPI) UpdateUser(ctx context.Context, user UserSecurity) (UserSecurity, error) {
	var returnValue UserSecurity

	// create path and map variables
	apiPath := a.client.cfg.BasePath + "/api/usersecurity/user"

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	r, err := a.client.buildRequest(apiPath, "PUT", &user, headers, url.Values{})
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
			var v UserSecurity
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
