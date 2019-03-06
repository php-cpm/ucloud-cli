//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UDisk DescribeUDiskPrice

package udisk

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeUDiskPriceRequest is request schema for DescribeUDiskPrice action
type DescribeUDiskPriceRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"true"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 购买UDisk大小,单位:GB,范围[1~1000]
	Size *int `required:"true"`

	// Year， Month， Dynamic，Trial，默认: Dynamic 如果不指定，则一次性获取三种计费
	ChargeType *string `required:"false"`

	// 购买UDisk的时长，默认值为1
	Quantity *int `required:"false"`

	// 是否打开数据方舟, 打开"Yes",关闭"No", 默认关闭
	UDataArkMode *string `required:"false"`

	// UDisk 类型: DataDisk（普通数据盘），SSDDataDisk（SSD数据盘），默认值（DataDisk）
	DiskType *string `required:"false"`
}

// DescribeUDiskPriceResponse is response schema for DescribeUDiskPrice action
type DescribeUDiskPriceResponse struct {
	response.CommonBase

	// 价格参数列表，具体说明见 UDiskPriceDataSet
	DataSet []UDiskPriceDataSet
}

// NewDescribeUDiskPriceRequest will create request of DescribeUDiskPrice action.
func (c *UDiskClient) NewDescribeUDiskPriceRequest() *DescribeUDiskPriceRequest {
	req := &DescribeUDiskPriceRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeUDiskPrice - 获取UDisk实例价格信息
func (c *UDiskClient) DescribeUDiskPrice(req *DescribeUDiskPriceRequest) (*DescribeUDiskPriceResponse, error) {
	var err error
	var res DescribeUDiskPriceResponse

	err = c.client.InvokeAction("DescribeUDiskPrice", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
