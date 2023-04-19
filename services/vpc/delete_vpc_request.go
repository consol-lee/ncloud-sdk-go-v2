/*
 * vpc
 *
 * VPC Network 관련 API<br/>https://ncloud.apigw.ntruss.com/vpc/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vpc

type DeleteVpcRequest struct {

	// REGION코드
RegionCode *string `json:"regionCode,omitempty"`

	// VPC번호
VpcNo *string `json:"vpcNo"`
}
