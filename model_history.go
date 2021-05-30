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

type History struct {
	BatchQueue        string    `json:"batchQueue,omitempty"`
	BufferedIOCount   int       `json:"bufferedIOCount,omitempty"`
	CompletionTimeUTC time.Time `json:"completionTimeUTC,omitempty"`
	CpuTime           string    `json:"cpuTime,omitempty"`
	DebugMode         bool      `json:"debugMode,omitempty"`
	DirectIOCount     int       `json:"directIOCount,omitempty"`
	ElapsedTime       string    `json:"elapsedTime,omitempty"`
	FinalSeverity     string    `json:"finalSeverity,omitempty"`
	FinalStatus       string    `json:"finalStatus,omitempty"`
	FinalStatusCode   int       `json:"finalStatusCode,omitempty"`
	FolderID          int       `json:"folderID,omitempty"`
	FolderName        string    `json:"folderName,omitempty"`
	HoldTimeUTC       time.Time `json:"holdTimeUTC,omitempty"`
	InitiatorID       int       `json:"initiatorID,omitempty"`
	InitiatorType     string    `json:"initiatorType,omitempty"`
	JamsEntry         int       `json:"jamsEntry,omitempty"`
	JamsId            string    `json:"jamsId,omitempty"`
	JobID             int       `json:"jobID,omitempty"`
	JobName           string    `json:"jobName,omitempty"`
	JobStatus         string    `json:"jobStatus,omitempty"`
	LogFilename       string    `json:"logFilename,omitempty"`
	MasterRON         int       `json:"masterRON,omitempty"`
	NodeName          string    `json:"nodeName,omitempty"`
	Note              string    `json:"note,omitempty"`
	OverrideJobName   string    `json:"overrideJobName,omitempty"`
	PageFaults        int       `json:"pageFaults,omitempty"`
	ProcessID         int       `json:"processID,omitempty"`
	RestartCount      int       `json:"restartCount,omitempty"`
	Ron               int       `json:"ron,omitempty"`
	ScheduledTimeUTC  time.Time `json:"scheduledTimeUTC,omitempty"`
	SetupID           int       `json:"setupID,omitempty"`
	SetupName         string    `json:"setupName,omitempty"`
	StartTimeUTC      time.Time `json:"startTimeUTC,omitempty"`
	SubmittedBy       string    `json:"submittedBy,omitempty"`
	UserName          string    `json:"userName,omitempty"`
	WorkingSetPeak    int64     `json:"workingSetPeak,omitempty"`
}
