/*
 * cdn
 *
 * <br/>https://ncloud.apigw.ntruss.com/cdn/v2
 *
 * API version: 2018-07-26T02:23:29Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cdn

type GlobalCdnRule struct {

	// 프로토콜구분코드
ProtocolTypeCode *string `json:"protocolTypeCode,omitempty"`

	// 서비스도메인구분코드
ServiceDomainTypeCode *string `json:"serviceDomainTypeCode,omitempty"`

	// 원본URL
OriginUrl *string `json:"originUrl,omitempty"`

	// 원본경로
OriginPath *string `json:"originPath,omitempty"`

	// 원본HTTP포트
OriginHttpPort *int32 `json:"originHttpPort,omitempty"`

	// 원본HTTPS포트
OriginHttpsPort *int32 `json:"originHttpsPort,omitempty"`

	// 포워드호스트헤더구분코드
ForwardHostHeaderTypeCode *string `json:"forwardHostHeaderTypeCode,omitempty"`

	// 포워드호스트헤더
ForwardHostHeader *string `json:"forwardHostHeader,omitempty"`

	// 캐쉬키호스트명구분코드
CacheKeyHostNameTypeCode *string `json:"cacheKeyHostNameTypeCode,omitempty"`

	// GZIP압축사용여부
IsGzipCompressionUse *bool `json:"isGzipCompressionUse,omitempty"`

	// 캐싱옵션구분코드
CachingOptionTypeCode *string `json:"cachingOptionTypeCode,omitempty"`

	// 오류내용응답사용여부
IsErrorContentsResponseUse *bool `json:"isErrorContentsResponseUse,omitempty"`

	// TTL캐싱
CachingTtlTime *int32 `json:"cachingTtlTime,omitempty"`

	// 쿼리스트링무시여부
IsQueryStringIgnoreUse *bool `json:"isQueryStringIgnoreUse,omitempty"`

	// 헤더제거사용여부
IsRemoveVaryHeaderUse *bool `json:"isRemoveVaryHeaderUse,omitempty"`

	// 대용량파일최적화사용여부
IsLargeFileOptimizationUse *bool `json:"isLargeFileOptimizationUse,omitempty"`

	// GZIP응답구분코드
GzipResponseTypeCode *string `json:"gzipResponseTypeCode,omitempty"`

	// 레퍼러도메인사용여부
IsReferrerDomainUse *bool `json:"isReferrerDomainUse,omitempty"`

	// 레퍼러도메인리스트
ReferrerDomainList []string `json:"referrerDomainList,omitempty"`

	// 레퍼러도메인제한사용여부
IsReferrerDomainRestrictUse *bool `json:"isReferrerDomainRestrictUse,omitempty"`

	// 보안토큰사용여부
IsSecureTokenUse *bool `json:"isSecureTokenUse,omitempty"`

	// 보안토큰비밀번호
SecureTokenPassword *string `json:"secureTokenPassword,omitempty"`

	// 보안토큰재발급여부
IsReissueSecureTokenPassword *bool `json:"isReissueSecureTokenPassword,omitempty"`

	// 인증서이름
CertificateName *string `json:"certificateName,omitempty"`

	// 엑세스로그사용여부
IsAccessLogUse *bool `json:"isAccessLogUse,omitempty"`

	// 엑세스로그파일스토리지인스턴스이름
AccessLogFileStorageContainerName *string `json:"accessLogFileStorageContainerName,omitempty"`
}
