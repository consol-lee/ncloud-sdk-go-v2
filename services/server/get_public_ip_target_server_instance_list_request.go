/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type GetPublicIpTargetServerInstanceListRequest struct {

	// 리전번호
RegionNo *string `json:"regionNo,omitempty"`

	// ZONE번호
ZoneNo *string `json:"zoneNo,omitempty"`
}
