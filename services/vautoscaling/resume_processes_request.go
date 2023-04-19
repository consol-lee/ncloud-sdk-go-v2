/*
 * vautoscaling
 *
 * VPC Auto Scaling 관련 API<br/>https://ncloud.apigw.ntruss.com/vautoscaling/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vautoscaling

type ResumeProcessesRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 오토스케일링그룹번호
AutoScalingGroupNo *string `json:"autoScalingGroupNo"`

	// 스케일링프로세스코드리스트
ScalingProcessCodeList []*string `json:"scalingProcessCodeList"`
}
