/*
 * JAMS REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: V6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type CountQueryOption struct {
	Context *ODataQueryContext `json:"context,omitempty"`
	RawValue string `json:"rawValue,omitempty"`
	Value bool `json:"value,omitempty"`
	Validator *CountQueryValidator `json:"validator,omitempty"`
}