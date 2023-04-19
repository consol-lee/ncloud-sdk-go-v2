/*
 * vpc
 *
 * VPC Network 관련 API<br/>https://ncloud.apigw.ntruss.com/vpc/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vpc

type AddRouteTableSubnetRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// 라우트테이블번호
RouteTableNo *string `json:"routeTableNo"`

	// 서브넷번호리스트
SubnetNoList []*string `json:"subnetNoList"`

	// VPC번호
VpcNo *string `json:"vpcNo"`
}
