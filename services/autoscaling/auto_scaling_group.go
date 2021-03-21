/*
 * autoscaling
 *
 * <br/>https://ncloud.apigw.ntruss.com/autoscaling/v2
 *
 * API version: 2018-08-07T06:47:31Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package autoscaling

type AutoScalingGroup struct {

	// 오토스케일링그룹번호
AutoScalingGroupNo *string `json:"autoScalingGroupNo,omitempty"`

	// 오토스케일링그룹명
AutoScalingGroupName *string `json:"autoScalingGroupName,omitempty"`

	// 론치설정명
LaunchConfigurationName *string `json:"launchConfigurationName,omitempty"`

	// 론치설정번호
LaunchConfigurationNo *string `json:"launchConfigurationNo,omitempty"`

	// 기대능력치
DesiredCapacity *int32 `json:"desiredCapacity,omitempty"`

	// 최소사이즈
MinSize *int32 `json:"minSize,omitempty"`

	// 최대사이즈
MaxSize *int32 `json:"maxSize,omitempty"`

DefaultCooldown *int32 `json:"defaultCooldown,omitempty"`

	// 로드밸런서인스턴스Summary리스트
LoadBalancerInstanceSummaryList []*LoadBalancerInstanceSummary `json:"loadBalancerInstanceSummaryList,omitempty"`

	// 헬스체크보류기간
HealthCheckGracePeriod *int32 `json:"healthCheckGracePeriod,omitempty"`

	// 헬스체크유형
HealthCheckType *CommonCode `json:"healthCheckType,omitempty"`

	// 생성일시
CreateDate *string `json:"createDate,omitempty"`

	// 오토스케일링그룹에속한서버인스턴스리스트
InAutoScalingGroupServerInstanceList []*InAutoScalingGroupServerInstance `json:"inAutoScalingGroupServerInstanceList,omitempty"`

	// 보류된프로세스리스트
SuspendedProcessList []*SuspendedProcess `json:"suspendedProcessList,omitempty"`

	// ZONE리스트
ZoneList []*Zone `json:"zoneList,omitempty"`
}
