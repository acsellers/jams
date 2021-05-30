/*
 * JAMS REST API

 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)

 * API version: V6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package jams

type NestedMenuEntry struct {
	Items    []NestedMenuEntry `json:"items,omitempty"`
	MenuName string            `json:"menuName,omitempty"`
	MenuType string            `json:"menuType,omitempty"`
	MenuText string            `json:"menuText,omitempty"`
	Id       int               `json:"id,omitempty"`
}
