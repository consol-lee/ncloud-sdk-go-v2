/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-09-19T10:11:03Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type TerminateServerInstancesResponse struct {

RequestId *string `json:"requestId,omitempty"`

ReturnCode *string `json:"returnCode,omitempty"`

ReturnMessage *string `json:"returnMessage,omitempty"`

TotalRows *int32 `json:"totalRows,omitempty"`

ServerInstanceList []*ServerInstance `json:"serverInstanceList,omitempty"`
}
