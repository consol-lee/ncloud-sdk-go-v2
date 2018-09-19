/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-09-19T10:11:03Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type AddPortForwardingRulesRequest struct {

	// 포트포워딩설정번호
PortForwardingConfigurationNo *string `json:"portForwardingConfigurationNo"`

	// 포트포워딩RULE리스트
PortForwardingRuleList []*PortForwardingRuleParameter `json:"portForwardingRuleList"`
}
