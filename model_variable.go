/*
 * JAMS REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: V6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

type Variable struct {
	Acl *Acl `json:"acl,omitempty"`
	CurrentLength int32 `json:"currentLength,omitempty"`
	DataType string `json:"dataType,omitempty"`
	Description string `json:"description,omitempty"`
	LastChangedBy string `json:"lastChangedBy,omitempty"`
	LastChangeUTC time.Time `json:"lastChangeUTC,omitempty"`
	ParentFolderId int32 `json:"parentFolderId,omitempty"`
	ParentFolderName string `json:"parentFolderName,omitempty"`
	Value string `json:"value,omitempty"`
	VariableId int32 `json:"variableId,omitempty"`
	VariableName string `json:"variableName,omitempty"`
}
