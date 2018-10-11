//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api ULB DescribeVServer

package ulb

import (
	"github.com/ucloud/ucloud-sdk-go/sdk"
	"github.com/ucloud/ucloud-sdk-go/sdk/request"
	"github.com/ucloud/ucloud-sdk-go/sdk/response"
)

// DescribeVServerRequest is request schema for DescribeVServer action
type DescribeVServerRequest struct {
	request.CommonBase

	// 负载均衡实例的Id
	ULBId *string `required:"true"`

	// VServer实例的Id；若指定则返回指定的VServer实例的信息； 若不指定则返回当前负载均衡实例下所有VServer的信息
	VServerId *string `required:"false"`
}

// DescribeVServerResponse is response schema for DescribeVServer action
type DescribeVServerResponse struct {
	response.CommonBase

	// 满足条件的VServer总数
	TotalCount int

	// VServer列表，每项参数详见 ULBVServerSet
	DataSet []ULBVServerSet
}

// NewDescribeVServerRequest will create request of DescribeVServer action.
func (c *ULBClient) NewDescribeVServerRequest() *DescribeVServerRequest {
	cfg := c.client.GetConfig()

	return &DescribeVServerRequest{
		CommonBase: request.CommonBase{
			Region:    sdk.String(cfg.Region),
			ProjectId: sdk.String(cfg.ProjectId),
		},
	}
}

// DescribeVServer - 获取ULB下的VServer的详细信息
func (c *ULBClient) DescribeVServer(req *DescribeVServerRequest) (*DescribeVServerResponse, error) {
	var err error
	var res DescribeVServerResponse

	err = c.client.InvokeAction("DescribeVServer", req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}