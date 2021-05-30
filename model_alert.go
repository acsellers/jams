/*
 * JAMS REST API

 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)

 * API version: V6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package jams

import (
	"time"
)

type Alert struct {
	AlertID         int       `json:"alertID,omitempty"`
	AlertName       string    `json:"alertName,omitempty"`
	Description     string    `json:"description,omitempty"`
	DetailText      string    `json:"detailText,omitempty"`
	EventList       []string  `json:"eventList,omitempty"`
	HandlerList     string    `json:"handlerList,omitempty"`
	InheritanceName string    `json:"inheritanceName,omitempty"`
	LastChangedBy   string    `json:"lastChangedBy,omitempty"`
	LastChangeUTC   time.Time `json:"lastChangeUTC,omitempty"`
	NotifyNames     string    `json:"notifyNames,omitempty"`
	SummaryText     string    `json:"summaryText,omitempty"`
}
