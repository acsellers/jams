/*
 * JAMS REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: V6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package jams

type PageResultHistory struct {
	Items        []History `json:"items,omitempty"`
	NextPageLink string    `json:"nextPageLink,omitempty"`
	Count        int64     `json:"count,omitempty"`
}
