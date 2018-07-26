/*
 * cdn
 *
 * <br/>https://ncloud.apigw.ntruss.com/cdn/v2
 *
 * API version: 2018-07-26T02:23:29Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cdn

type CdnPlusInstance struct {

	// CDN인스턴스번호
CdnInstanceNo *string `json:"cdnInstanceNo,omitempty"`

	// CDN인스턴스상태
CdnInstanceStatus **CommonCode `json:"cdnInstanceStatus,omitempty"`

	// CDN인스턴스OP
CdnInstanceOperation **CommonCode `json:"cdnInstanceOperation,omitempty"`

	// CDN인스턴스상태명
CdnInstanceStatusName *string `json:"cdnInstanceStatusName,omitempty"`

	// 생성일자
CreateDate *string `json:"createDate,omitempty"`

	// UPTIME
LastModifiedDate *string `json:"lastModifiedDate,omitempty"`

	// CDN인스턴스설명
CdnInstanceDescription *string `json:"cdnInstanceDescription,omitempty"`

	// 서비스이름
ServiceName *string `json:"serviceName,omitempty"`

	// 라이브트랜스코더여부
IsForLiveTranscoder *bool `json:"isForLiveTranscoder,omitempty"`

	// 라이브트랜스코더인스턴스번호리스트
LiveTranscoderInstanceNoList []string `json:"liveTranscoderInstanceNoList,omitempty"`

	// Image Optimizer여부
IsForImageOptimizer *bool `json:"isForImageOptimizer,omitempty"`

	// Image Optimizer인스턴스번호
ImageOptimizerInstanceNo *string `json:"imageOptimizerInstanceNo,omitempty"`

IsAvailablePartialDomainPurge *bool `json:"isAvailablePartialDomainPurge,omitempty"`

	// CDN+서비스도메인리스트
CdnPlusServiceDomainList []CdnPlusServiceDomain `json:"cdnPlusServiceDomainList,omitempty"`

CdnPlusRule **CdnPlusRule `json:"cdnPlusRule,omitempty"`
}
