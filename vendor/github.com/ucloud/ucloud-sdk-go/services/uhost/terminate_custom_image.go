//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UHost TerminateCustomImage

package uhost

import (
	"github.com/ucloud/ucloud-sdk-go/sdk"
	"github.com/ucloud/ucloud-sdk-go/sdk/request"
	"github.com/ucloud/ucloud-sdk-go/sdk/response"
)

// TerminateCustomImageRequest is request schema for TerminateCustomImage action
type TerminateCustomImageRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone *string `required:"false"`

	// 自制镜像ID 参见 [DescribeImage](describe_image.html)
	ImageId *string `required:"true"`
}

// TerminateCustomImageResponse is response schema for TerminateCustomImage action
type TerminateCustomImageResponse struct {
	response.CommonBase

	// 自制镜像Id
	ImageId string
}

// NewTerminateCustomImageRequest will create request of TerminateCustomImage action.
func (c *UHostClient) NewTerminateCustomImageRequest() *TerminateCustomImageRequest {
	cfg := c.client.GetConfig()

	return &TerminateCustomImageRequest{
		CommonBase: request.CommonBase{
			Region:    sdk.String(cfg.Region),
			ProjectId: sdk.String(cfg.ProjectId),
		},
	}
}

// TerminateCustomImage - 删除用户自定义镜像
func (c *UHostClient) TerminateCustomImage(req *TerminateCustomImageRequest) (*TerminateCustomImageResponse, error) {
	var err error
	var res TerminateCustomImageResponse

	err = c.client.InvokeAction("TerminateCustomImage", req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}