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

type DateDetail struct {
	DateTypeID    int       `json:"dateTypeID,omitempty"`
	DateTypeType  string    `json:"dateTypeType,omitempty"`
	Description   string    `json:"description,omitempty"`
	LastChangeUTC time.Time `json:"lastChangeUTC,omitempty"`
	SpecificType  string    `json:"specificType,omitempty"`
	StartDate     time.Time `json:"startDate,omitempty"`
	WorkDay       string    `json:"workDay,omitempty"`
}
