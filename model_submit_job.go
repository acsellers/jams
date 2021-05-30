/*
 * JAMS REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: V6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package jams

type SubmitJob struct {
	AgentNode            string `json:"agentNode,omitempty"`
	AllowEditSource      bool   `json:"allowEditSource,omitempty"`
	AltUserName          string `json:"altUserName,omitempty"`
	BatchQueueID         int    `json:"batchQueueID,omitempty"`
	BatchQueueName       string `json:"batchQueueName,omitempty"`
	Description          string `json:"description,omitempty"`
	DetailID             int    `json:"detailID,omitempty"`
	FolderID             int    `json:"folderID,omitempty"`
	FolderName           string `json:"folderName,omitempty"`
	Hold                 bool   `json:"hold,omitempty"`
	InSchedule           bool   `json:"inSchedule,omitempty"`
	JobID                int    `json:"jobID,omitempty"`
	KeepLogs             string `json:"keepLogs,omitempty"`
	LogFile              bool   `json:"logFile,omitempty"`
	LogFilename          string `json:"logFilename,omitempty"`
	LogLocation          string `json:"logLocation,omitempty"`
	MinimumSeverity      string `json:"minimumSeverity,omitempty"`
	MissedWindowAction   string `json:"missedWindowAction,omitempty"`
	Name                 string `json:"name,omitempty"`
	NotifyOfMissedWindow bool   `json:"notifyOfMissedWindow,omitempty"`
	NotifyUser           bool   `json:"notifyUser,omitempty"`
	OverrideName         string `json:"overrideName,omitempty"`
	PrecheckInterval     string `json:"precheckInterval,omitempty"`
	PrecheckJobID        int    `json:"precheckJobID,omitempty"`
	RequiresSnapshot     bool   `json:"requiresSnapshot,omitempty"`
	Restartable          bool   `json:"restartable,omitempty"`
	RunawayElapsed       string `json:"runawayElapsed,omitempty"`
	RunPriority          int    `json:"runPriority,omitempty"`
	ScheduleFromTime     string `json:"scheduleFromTime,omitempty"`
	ScheduleToTime       string `json:"scheduleToTime,omitempty"`
	ScheduleWindowID     int    `json:"scheduleWindowID,omitempty"`
	SchedulingPriority   int    `json:"schedulingPriority,omitempty"`
	ShortElapsed         string `json:"shortElapsed,omitempty"`
	ShortSeverity        string `json:"shortSeverity,omitempty"`
	Source               string `json:"source,omitempty"`
	SourceModified       bool   `json:"sourceModified,omitempty"`
	StalledTime          string `json:"stalledTime,omitempty"`
	Step                 int    `json:"step,omitempty"`
	SubmitMethodID       int    `json:"submitMethodID,omitempty"`
	TimestampLogs        string `json:"timestampLogs,omitempty"`
	UserID               int    `json:"userID,omitempty"`
	UserName             string `json:"userName,omitempty"`
	WaitFor              bool   `json:"waitFor,omitempty"`
}
