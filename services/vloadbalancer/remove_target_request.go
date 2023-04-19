/*
 * vloadbalancer
 *
 * VPC Load Balancer 관련 API<br/>https://ncloud.apigw.ntruss.com/vloadbalancer/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vloadbalancer

type RemoveTargetRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 타겟그룹번호
TargetGroupNo *string `json:"targetGroupNo"`

	// 타겟번호리스트
TargetNoList []*string `json:"targetNoList"`
}
