/*
 * JAMS REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: V6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package jams

type IEdmNavigationSource struct {
	NavigationPropertyBindings []IEdmNavigationPropertyBinding `json:"navigationPropertyBindings,omitempty"`
	Path                       *IEdmPathExpression             `json:"path,omitempty"`
	Type_                      *IEdmType                       `json:"type,omitempty"`
}
