/*
 * JAMS REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: V6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type TopQueryOption struct {
	Context *ODataQueryContext `json:"context,omitempty"`
	RawValue string `json:"rawValue,omitempty"`
	Value int32 `json:"value,omitempty"`
	Validator *TopQueryValidator `json:"validator,omitempty"`
}