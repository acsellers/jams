/*
 * JAMS REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: V6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ReleaseEntry struct {
	AuditComment string `json:"auditComment,omitempty"`
	ReleaseType string `json:"releaseType,omitempty"`
	ForcePresent bool `json:"forcePresent,omitempty"`
	DependOk bool `json:"dependOk,omitempty"`
}