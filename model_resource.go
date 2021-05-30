/*
 * JAMS REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: V6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package jams

import (
	"time"
)

type Resource struct {
	ACL               *ACL             `json:"acl,omitempty"`
	LastChangeUTC     time.Time        `json:"lastChangeUTC,omitempty"`
	ResourceName      string           `json:"resourceName,omitempty"`
	ResourceID        int              `json:"resourceID,omitempty"`
	QuantityAvailable int              `json:"quantityAvailable,omitempty"`
	QuantityInUse     int              `json:"quantityInUse,omitempty"`
	RoutineName       string           `json:"routineName,omitempty"`
	NodeSpecific      bool             `json:"nodeSpecific,omitempty"`
	Description       string           `json:"description,omitempty"`
	ResourceDetails   []ResourceDetail `json:"resourceDetails,omitempty"`
}
