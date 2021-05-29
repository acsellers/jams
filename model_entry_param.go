/*
 * JAMS REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: V6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type EntryParam struct {
	DataType string `json:"dataType,omitempty"`
	DefaultFormat string `json:"defaultFormat,omitempty"`
	JamsEntry int32 `json:"jamsEntry,omitempty"`
	LoadedFrom string `json:"loadedFrom,omitempty"`
	ParameterOrigin string `json:"parameterOrigin,omitempty"`
	ParamName string `json:"paramName,omitempty"`
	Sequence int32 `json:"sequence,omitempty"`
	Value string `json:"value,omitempty"`
}