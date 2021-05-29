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

type RescheduleEntry struct {
	BatchQueue   string    `json:"batchQueue,omitempty"`
	AgentNode    string    `json:"agentNode,omitempty"`
	HoldTime     time.Time `json:"holdTime,omitempty"`
	Held         bool      `json:"held,omitempty"`
	Priority     int32     `json:"priority,omitempty"`
	AuditComment string    `json:"auditComment,omitempty"`
}
