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

type Entry struct {
	AgentNode            string    `json:"agentNode,omitempty"`
	AltUserName          string    `json:"altUserName,omitempty"`
	BatchQueueID         int       `json:"batchQueueID,omitempty"`
	BatchQueueName       string    `json:"batchQueueName,omitempty"`
	CompletionTimeUTC    time.Time `json:"completionTimeUTC,omitempty"`
	CurrentState         string    `json:"currentState,omitempty"`
	CurrentStep          int       `json:"currentStep,omitempty"`
	DebugMode            bool      `json:"debugMode,omitempty"`
	DependOK             bool      `json:"dependOK,omitempty"`
	Description          string    `json:"description,omitempty"`
	DetailID             int       `json:"detailID,omitempty"`
	DisplayName          string    `json:"displayName,omitempty"`
	FinalSeverity        string    `json:"finalSeverity,omitempty"`
	FinalStatus          string    `json:"finalStatus,omitempty"`
	FinalStatusCode      int       `json:"finalStatusCode,omitempty"`
	Halted               bool      `json:"halted,omitempty"`
	Held                 bool      `json:"held,omitempty"`
	HoldTimeUTC          time.Time `json:"holdTimeUTC,omitempty"`
	Icon                 string    `json:"icon,omitempty"`
	IgnoreResReq         bool      `json:"ignoreResReq,omitempty"`
	InitiatorID          int       `json:"initiatorID,omitempty"`
	InitiatorType        string    `json:"initiatorType,omitempty"`
	JamsEntry            int       `json:"jamsEntry,omitempty"`
	JamsId               string    `json:"jamsId,omitempty"`
	JobID                int       `json:"jobID,omitempty"`
	JobName              string    `json:"jobName,omitempty"`
	JobStatus            string    `json:"jobStatus,omitempty"`
	LogFileActive        bool      `json:"logFileActive,omitempty"`
	LogFilename          string    `json:"logFilename,omitempty"`
	MasterEntry          int       `json:"masterEntry,omitempty"`
	MasterRON            int       `json:"masterRON,omitempty"`
	Message              string    `json:"message,omitempty"`
	MissedWindowAction   string    `json:"missedWindowAction,omitempty"`
	Monitor              bool      `json:"monitor,omitempty"`
	NodeName             string    `json:"nodeName,omitempty"`
	Note                 string    `json:"note,omitempty"`
	NotifyOfMissedWindow bool      `json:"notifyOfMissedWindow,omitempty"`
	NotifyUser           bool      `json:"notifyUser,omitempty"`
	OriginalHoldTimeUTC  time.Time `json:"originalHoldTimeUTC,omitempty"`
	ParentFolderID       int       `json:"parentFolderID,omitempty"`
	Permanent            bool      `json:"permanent,omitempty"`
	PrecheckCount        int       `json:"precheckCount,omitempty"`
	PrecheckEntry        int       `json:"precheckEntry,omitempty"`
	PrecheckInterval     string    `json:"precheckInterval,omitempty"`
	PrecheckJobID        int       `json:"precheckJobID,omitempty"`
	PrecheckTimeUTC      time.Time `json:"precheckTimeUTC,omitempty"`
	Preprocessed         bool      `json:"preprocessed,omitempty"`
	ProcessID            int       `json:"processID,omitempty"`
	Restartable          bool      `json:"restartable,omitempty"`
	RestartCount         int       `json:"restartCount,omitempty"`
	RetainOption         string    `json:"retainOption,omitempty"`
	RetainTime           string    `json:"retainTime,omitempty"`
	Ron                  int       `json:"ron,omitempty"`
	Runaway              bool      `json:"runaway,omitempty"`
	RunawayElapsed       string    `json:"runawayElapsed,omitempty"`
	RunPriority          int       `json:"runPriority,omitempty"`
	ScheduledTimeUTC     time.Time `json:"scheduledTimeUTC,omitempty"`
	ScheduleFromTime     string    `json:"scheduleFromTime,omitempty"`
	ScheduleToTime       string    `json:"scheduleToTime,omitempty"`
	ScheduleWindowID     int       `json:"scheduleWindowID,omitempty"`
	SchedulingPriority   int       `json:"schedulingPriority,omitempty"`
	SetupID              int       `json:"setupID,omitempty"`
	SetupOk              bool      `json:"setupOk,omitempty"`
	Stalled              bool      `json:"stalled,omitempty"`
	StalledTime          string    `json:"stalledTime,omitempty"`
	StartTimeUTC         time.Time `json:"startTimeUTC,omitempty"`
	Step                 int       `json:"step,omitempty"`
	StepWait             bool      `json:"stepWait,omitempty"`
	SubmitMethodName     string    `json:"submitMethodName,omitempty"`
	SubmitMethodID       int       `json:"submitMethodID,omitempty"`
	SubmittedBy          string    `json:"submittedBy,omitempty"`
	SuppressNotify       bool      `json:"suppressNotify,omitempty"`
	TempFilename         string    `json:"tempFilename,omitempty"`
	TodaysDate           time.Time `json:"todaysDate,omitempty"`
	UserID               int       `json:"userID,omitempty"`
	UserName             string    `json:"userName,omitempty"`
	Version              int       `json:"version,omitempty"`
	WaitFor              bool      `json:"waitFor,omitempty"`
}
