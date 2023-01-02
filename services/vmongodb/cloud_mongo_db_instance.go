/*
 * vmongodb
 *
 * <br/>https://ncloud.apigw.ntruss.com/vmongodb/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vmongodb

type CloudMongoDbInstance struct {

	// CloudMongoDb인스턴스번호
	CloudMongoDbInstanceNo *string `json:"cloudMongoDbInstanceNo,omitempty"`

	// CloudMongoDb서비스이름
	CloudMongoDbServiceName *string `json:"cloudMongoDbServiceName,omitempty"`

	// CloudMongoDb인스턴스상태이름
	CloudMongoDbInstanceStatusName *string `json:"cloudMongoDbInstanceStatusName,omitempty"`

	// CloudMongoDb인스턴스상태
	CloudMongoDbInstanceStatus *CommonCode `json:"cloudMongoDbInstanceStatus,omitempty"`

	// CloudMongoDb인스턴스OP
	CloudMongoDbInstanceOperation *CommonCode `json:"cloudMongoDbInstanceOperation,omitempty"`

	// CloudMongoDb이미지상품코드
	CloudMongoDbImageProductCode *string `json:"cloudMongoDbImageProductCode,omitempty"`

	// CloudMongoDb엔진버전
	EngineVersion *string `json:"engineVersion,omitempty"`

	// CloudMongoDb라이선스
	License *CommonCode `json:"license,omitempty"`

	// CloudMongoDb Member 포트
	MemberPort *int32 `json:"memberPort,omitempty"`

	// CloudMongoDb Mongos 포트
	MongosPort *int32 `json:"mongosPort,omitempty"`

	// CloudMongoDb Config 포트
	ConfigPort *int32 `json:"configPort,omitempty"`

	// CloudMongoDb Arbiter 포트
	ArbiterPort *int32 `json:"arbiterPort,omitempty"`

	// 백업파일보관기간
	BackupFileRetentionPeriod *int32 `json:"backupFileRetentionPeriod,omitempty"`

	// 백업시간
	BackupTime *string `json:"backupTime,omitempty"`

	// 생성일자
	CreateDate *string `json:"createDate,omitempty"`

	// ACG번호리스트
	AccessControlGroupNoList []string `json:"accessControlGroupNoList,omitempty"`

	// 샤드수
	ShardCount *int32 `json:"shardCount,omitempty"`

	// Cloud DB for MongoDB 데이터 압축 알고리즘
	Compress *CommonCode `json:"compress,omitempty"`

	// Cloud DB for MongoDB 클러스터 타입
	ClusterType *CommonCode `json:"clusterType,omitempty"`

	// CloudMongoDb서버인스턴스리스트
	CloudMongoDbServerInstanceList []*CloudMongoDbServerInstance `json:"CloudMongoDbServerInstanceList,omitempty"`
}
