/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-09-19T10:11:03Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type GetServerProductListRequest struct {

	// 제외할상품코드
ExclusionProductCode *string `json:"exclusionProductCode,omitempty"`

	// 인터넷라인구분코드
InternetLineTypeCode *string `json:"internetLineTypeCode,omitempty"`

	// 조회할상품코드
ProductCode *string `json:"productCode,omitempty"`

	// 리전번호
RegionNo *string `json:"regionNo,omitempty"`

	// 서버이미지상품코드
ServerImageProductCode *string `json:"serverImageProductCode"`

	// ZONE번호
ZoneNo *string `json:"zoneNo,omitempty"`
}
