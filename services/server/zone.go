/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-09-19T10:11:03Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type Zone struct {

	// 존(Zone)번호
ZoneNo *string `json:"zoneNo,omitempty"`

	// 존(Zone)명
ZoneName *string `json:"zoneName,omitempty"`

	// 존(Zone) 코드
ZoneCode *string `json:"zoneCode,omitempty"`

	// 존(Zone)설명
ZoneDescription *string `json:"zoneDescription,omitempty"`

	// 리전번호
RegionNo *string `json:"regionNo,omitempty"`
}
