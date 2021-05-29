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

type Resource struct {
	Acl *Acl `json:"acl,omitempty"`
	LastChangeUTC time.Time `json:"lastChangeUTC,omitempty"`
	ResourceName string `json:"resourceName,omitempty"`
	ResourceID int32 `json:"resourceID,omitempty"`
	QuantityAvailable int32 `json:"quantityAvailable,omitempty"`
	QuantityInUse int32 `json:"quantityInUse,omitempty"`
	RoutineName string `json:"routineName,omitempty"`
	NodeSpecific bool `json:"nodeSpecific,omitempty"`
	Description string `json:"description,omitempty"`
	ResourceDetails []ResourceDetail `json:"resourceDetails,omitempty"`
}
