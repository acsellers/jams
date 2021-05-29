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

type Report struct {
	Copies          int32     `json:"copies,omitempty"`
	Description     string    `json:"description,omitempty"`
	FullFilename    string    `json:"fullFilename,omitempty"`
	LastChangeUTC   time.Time `json:"lastChangeUTC,omitempty"`
	LogicalName     string    `json:"logicalName,omitempty"`
	PrintForm       string    `json:"printForm,omitempty"`
	PrintQualifiers string    `json:"printQualifiers,omitempty"`
	PrintQueue      string    `json:"printQueue,omitempty"`
	ReportID        string    `json:"reportID,omitempty"`
	RetentionDays   int32     `json:"retentionDays,omitempty"`
}
