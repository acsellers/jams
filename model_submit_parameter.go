/*
 * JAMS REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: V6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package jams

type SubmitParameter struct {
	AllowEntry      bool   `json:"allowEntry,omitempty"`
	DataType        string `json:"dataType,omitempty"`
	DefaultFormat   string `json:"defaultFormat,omitempty"`
	DetailID        int32  `json:"detailID,omitempty"`
	GlobalName      string `json:"globalName,omitempty"`
	GlobalOverride  bool   `json:"globalOverride,omitempty"`
	HelpText        string `json:"helpText,omitempty"`
	Hide            bool   `json:"hide,omitempty"`
	JobID           int32  `json:"jobID,omitempty"`
	Length          int32  `json:"length,omitempty"`
	MustFill        bool   `json:"mustFill,omitempty"`
	ParameterOrigin string `json:"parameterOrigin,omitempty"`
	ParamID         int32  `json:"paramID,omitempty"`
	ParamName       string `json:"paramName,omitempty"`
	ParamValue      string `json:"paramValue,omitempty"`
	Prompt          string `json:"prompt,omitempty"`
	Required        bool   `json:"required,omitempty"`
	Uppercase       bool   `json:"uppercase,omitempty"`
	ValidationData  string `json:"validationData,omitempty"`
	ValidationType  string `json:"validationType,omitempty"`
}
