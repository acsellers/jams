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

type TriggerEvent struct {
	EventID          int       `json:"eventID,omitempty"`
	EventName        string    `json:"eventName,omitempty"`
	Description      string    `json:"description,omitempty"`
	LastChangeUTC    time.Time `json:"lastChangeUTC,omitempty"`
	EventType        string    `json:"eventType,omitempty"`
	EventHappenedUTC time.Time `json:"eventHappenedUTC,omitempty"`
	EventIsTrue      bool      `json:"eventIsTrue,omitempty"`
}
