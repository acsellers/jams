/*
 * JAMS REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: V6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ODataQueryOptionsHistory struct {
	IfMatch *Object `json:"ifMatch,omitempty"`
	IfNoneMatch *Object `json:"ifNoneMatch,omitempty"`
	Context *ODataQueryContext `json:"context,omitempty"`
	Request *Object `json:"request,omitempty"`
	RawValues *ODataRawQueryOptions `json:"rawValues,omitempty"`
	SelectExpand *SelectExpandQueryOption `json:"selectExpand,omitempty"`
	Filter *FilterQueryOption `json:"filter,omitempty"`
	OrderBy *OrderByQueryOption `json:"orderBy,omitempty"`
	Skip *SkipQueryOption `json:"skip,omitempty"`
	Top *TopQueryOption `json:"top,omitempty"`
	Count *CountQueryOption `json:"count,omitempty"`
	Validator *ODataQueryValidator `json:"validator,omitempty"`
}