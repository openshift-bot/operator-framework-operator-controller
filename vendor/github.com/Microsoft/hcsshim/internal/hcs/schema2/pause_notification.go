/*
 * HCS API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package hcsschema

//  Notification data that is indicated to components running in the Virtual Machine.
type PauseNotification struct {
	Reason string `json:"Reason,omitempty"`
}
