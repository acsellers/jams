/*
 * JAMS REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: V6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type SelectExpandQueryOption struct {
	Context *ODataQueryContext `json:"context,omitempty"`
	RawSelect string `json:"rawSelect,omitempty"`
	RawExpand string `json:"rawExpand,omitempty"`
	Validator *SelectExpandQueryValidator `json:"validator,omitempty"`
	SelectExpandClause *SelectExpandClause `json:"selectExpandClause,omitempty"`
	LevelsMaxLiteralExpansionDepth int32 `json:"levelsMaxLiteralExpansionDepth,omitempty"`
}