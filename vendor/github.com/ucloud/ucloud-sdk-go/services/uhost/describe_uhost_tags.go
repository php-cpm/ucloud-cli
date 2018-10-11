//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UHost DescribeUHostTags

package uhost

import (
	"github.com/ucloud/ucloud-sdk-go/sdk"
	"github.com/ucloud/ucloud-sdk-go/sdk/request"
	"github.com/ucloud/ucloud-sdk-go/sdk/response"
)

// DescribeUHostTagsRequest is request schema for DescribeUHostTags action
type DescribeUHostTagsRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone *string `required:"false"`
}

// DescribeUHostTagsResponse is response schema for DescribeUHostTags action
type DescribeUHostTagsResponse struct {
	response.CommonBase

	// 已有主机的业务组总个数
	TotalCount int

	// 业务组集合见 UHostTagSet
	TagSet []UHostTagSet
}

// NewDescribeUHostTagsRequest will create request of DescribeUHostTags action.
func (c *UHostClient) NewDescribeUHostTagsRequest() *DescribeUHostTagsRequest {
	cfg := c.client.GetConfig()

	return &DescribeUHostTagsRequest{
		CommonBase: request.CommonBase{
			Region:    sdk.String(cfg.Region),
			ProjectId: sdk.String(cfg.ProjectId),
		},
	}
}

// DescribeUHostTags - 获取指定数据中心的业务组列表。
func (c *UHostClient) DescribeUHostTags(req *DescribeUHostTagsRequest) (*DescribeUHostTagsResponse, error) {
	var err error
	var res DescribeUHostTagsResponse

	err = c.client.InvokeAction("DescribeUHostTags", req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}