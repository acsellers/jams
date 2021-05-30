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

type Folder struct {
	ACL                        *ACL        `json:"acl,omitempty"`
	BatchQueueName             string      `json:"batchQueueName,omitempty"`
	BatchQueueID               int         `json:"batchQueueID,omitempty"`
	Description                string      `json:"description,omitempty"`
	HomeDirectory              string      `json:"homeDirectory,omitempty"`
	KeepLogs                   string      `json:"keepLogs,omitempty"`
	LastChangedBy              string      `json:"lastChangedBy,omitempty"`
	LastChangeUTC              time.Time   `json:"lastChangeUTC,omitempty"`
	LogLocation                string      `json:"logLocation,omitempty"`
	NotifyEmail                string      `json:"notifyEmail,omitempty"`
	NotifyJobID                int         `json:"notifyJobID,omitempty"`
	NotifyJobName              string      `json:"notifyJobName,omitempty"`
	NotifyJobNameIsRelative    bool        `json:"notifyJobNameIsRelative,omitempty"`
	TimestampLogs              string      `json:"timestampLogs,omitempty"`
	UserID                     int         `json:"userID,omitempty"`
	NotifyOther                string      `json:"notifyOther,omitempty"`
	PrintLocation              string      `json:"printLocation,omitempty"`
	PrintQueue                 string      `json:"printQueue,omitempty"`
	Parameters                 []Param     `json:"parameters,omitempty"`
	Alerts                     []AlertLink `json:"alerts,omitempty"`
	SchedulingPriorityModifier int         `json:"schedulingPriorityModifier,omitempty"`
	NotifyUsers                string      `json:"notifyUsers,omitempty"`
	RecoveryInstructions       string      `json:"recoveryInstructions,omitempty"`
	RetainOption               string      `json:"retainOption,omitempty"`
	RetainTime                 string      `json:"retainTime,omitempty"`
	RunawayElapsedPer          int         `json:"runawayElapsedPer,omitempty"`
	TemplateLibrary            string      `json:"templateLibrary,omitempty"`
	ShortElapsedPer            int         `json:"shortElapsedPer,omitempty"`
	ShortSeverity              string      `json:"shortSeverity,omitempty"`
	StalledTime                string      `json:"stalledTime,omitempty"`
	FolderID                   int         `json:"folderID,omitempty"`
	FolderName                 string      `json:"folderName,omitempty"`
	Name                       string      `json:"name,omitempty"`
	ParentFolderID             int         `json:"parentFolderID,omitempty"`
	ParentFolderName           string      `json:"parentFolderName,omitempty"`
	UserName                   string      `json:"userName,omitempty"`
	Requirements               []ResReq    `json:"requirements,omitempty"`
	SearchPath                 string      `json:"searchPath,omitempty"`
}
