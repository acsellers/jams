/*
 * JAMS REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: V6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package jams

type Version struct {
	Major         int `json:"major,omitempty"`
	Minor         int `json:"minor,omitempty"`
	Build         int `json:"build,omitempty"`
	Revision      int `json:"revision,omitempty"`
	MajorRevision int `json:"majorRevision,omitempty"`
	MinorRevision int `json:"minorRevision,omitempty"`
}
