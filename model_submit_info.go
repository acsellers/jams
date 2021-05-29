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

type SubmitInfo struct {
	AfterTimeUTC time.Time `json:"afterTimeUTC,omitempty"`
	BatchQueueID int32 `json:"batchQueueID,omitempty"`
	BatchQueueName string `json:"batchQueueName,omitempty"`
	DebugMode bool `json:"debugMode,omitempty"`
	Description string `json:"description,omitempty"`
	FailShort bool `json:"failShort,omitempty"`
	FolderID int32 `json:"folderID,omitempty"`
	FolderName string `json:"folderName,omitempty"`
	Hold bool `json:"hold,omitempty"`
	JamsEntry int32 `json:"jamsEntry,omitempty"`
	JamsId string `json:"jamsId,omitempty"`
	Jobs []SubmitJob `json:"jobs,omitempty"`
	LogFile bool `json:"logFile,omitempty"`
	LogFilename string `json:"logFilename,omitempty"`
	LogLocation string `json:"logLocation,omitempty"`
	ManageAccess bool `json:"manageAccess,omitempty"`
	MissedWindowAction string `json:"missedWindowAction,omitempty"`
	Name string `json:"name,omitempty"`
	Note string `json:"note,omitempty"`
	NotifyOfMissedWindow bool `json:"notifyOfMissedWindow,omitempty"`
	NotifyUser bool `json:"notifyUser,omitempty"`
	OverrideName string `json:"overrideName,omitempty"`
	Parameters []SubmitParameter `json:"parameters,omitempty"`
	PrecheckInterval string `json:"precheckInterval,omitempty"`
	PrecheckJobID int32 `json:"precheckJobID,omitempty"`
	Restartable bool `json:"restartable,omitempty"`
	ResultText string `json:"resultText,omitempty"`
	RetainOption string `json:"retainOption,omitempty"`
	RetainTime string `json:"retainTime,omitempty"`
	Ron int32 `json:"ron,omitempty"`
	RunawayElapsed string `json:"runawayElapsed,omitempty"`
	ScheduleFromTime string `json:"scheduleFromTime,omitempty"`
	ScheduleToTime string `json:"scheduleToTime,omitempty"`
	ScheduleWindowID int32 `json:"scheduleWindowID,omitempty"`
	SchedulingPriority int32 `json:"schedulingPriority,omitempty"`
	SetSymbols bool `json:"setSymbols,omitempty"`
	ShortElapsed string `json:"shortElapsed,omitempty"`
	StalledTime string `json:"stalledTime,omitempty"`
	SubmitID int32 `json:"submitID,omitempty"`
	SubmitIsSimple bool `json:"submitIsSimple,omitempty"`
	SubmittedBy string `json:"submittedBy,omitempty"`
	SubmitType string `json:"submitType,omitempty"`
	SuppressNotify bool `json:"suppressNotify,omitempty"`
	TestSubmit bool `json:"testSubmit,omitempty"`
	TimestampLogs string `json:"timestampLogs,omitempty"`
	UserID int32 `json:"userID,omitempty"`
	UserName string `json:"userName,omitempty"`
	UseSymbols bool `json:"useSymbols,omitempty"`
}