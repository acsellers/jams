/*
 * JAMS REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: V6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type CancelEntry struct {
	Force bool `json:"force,omitempty"`
	Reprocess bool `json:"reprocess,omitempty"`
	AuditComment string `json:"auditComment,omitempty"`
	Severity string `json:"severity,omitempty"`
}