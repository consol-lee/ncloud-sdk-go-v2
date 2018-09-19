/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-09-19T10:11:03Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type DeleteInstanceTagsRequest struct {

	// 인스턴스번호리스트
InstanceNoList []*string `json:"instanceNoList"`

	// 인스턴스태그리스트
InstanceTagList []*InstanceTagParameter `json:"instanceTagList,omitempty"`
}
