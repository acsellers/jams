/*
 * JAMS REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: V6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package jams

type Tenant struct {
	CompanyName   string   `json:"companyName,omitempty"`
	Title         string   `json:"title,omitempty"`
	Theme         string   `json:"theme,omitempty"`
	ServerVersion *Version `json:"serverVersion,omitempty"`
}
