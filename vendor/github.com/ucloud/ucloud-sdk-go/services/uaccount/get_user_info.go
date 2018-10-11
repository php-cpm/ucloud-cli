//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UAccount GetUserInfo

package uaccount

import (
	"github.com/ucloud/ucloud-sdk-go/sdk"
	"github.com/ucloud/ucloud-sdk-go/sdk/request"
	"github.com/ucloud/ucloud-sdk-go/sdk/response"
)

// GetUserInfoRequest is request schema for GetUserInfo action
type GetUserInfoRequest struct {
	request.CommonBase
}

// GetUserInfoResponse is response schema for GetUserInfo action
type GetUserInfoResponse struct {
	response.CommonBase

	// 用户信息返回数组
	DataSet []UserInfo
}

// NewGetUserInfoRequest will create request of GetUserInfo action.
func (c *UAccountClient) NewGetUserInfoRequest() *GetUserInfoRequest {
	cfg := c.client.GetConfig()

	return &GetUserInfoRequest{
		CommonBase: request.CommonBase{
			Region:    sdk.String(cfg.Region),
			ProjectId: sdk.String(cfg.ProjectId),
		},
	}
}

// GetUserInfo - 获取用户信息
func (c *UAccountClient) GetUserInfo(req *GetUserInfoRequest) (*GetUserInfoResponse, error) {
	var err error
	var res GetUserInfoResponse

	err = c.client.InvokeAction("GetUserInfo", req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}