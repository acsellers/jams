/*
 * JAMS REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: V6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package jams

type AppMenuItem struct {
	Id      int32         `json:"id,omitempty"`
	Label   string        `json:"label,omitempty"`
	Icon    string        `json:"icon,omitempty"`
	Url     string        `json:"url,omitempty"`
	Route   string        `json:"route,omitempty"`
	SubMenu []AppMenuItem `json:"subMenu,omitempty"`
}
