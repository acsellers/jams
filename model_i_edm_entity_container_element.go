/*
 * JAMS REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: V6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package jams

type IEdmEntityContainerElement struct {
	ContainerElementKind string               `json:"containerElementKind,omitempty"`
	Container            *IEdmEntityContainer `json:"container,omitempty"`
}
