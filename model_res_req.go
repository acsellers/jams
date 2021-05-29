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

type ResReq struct {
	AutoRelease bool `json:"autoRelease,omitempty"`
	LastChangeUTC time.Time `json:"lastChangeUTC,omitempty"`
	NetQuantityRequired int32 `json:"netQuantityRequired,omitempty"`
	QuantityInherited int32 `json:"quantityInherited,omitempty"`
	QuantityRequired int32 `json:"quantityRequired,omitempty"`
	ResourceID int32 `json:"resourceID,omitempty"`
	ResourceName string `json:"resourceName,omitempty"`
	RetainOnFailure bool `json:"retainOnFailure,omitempty"`
}