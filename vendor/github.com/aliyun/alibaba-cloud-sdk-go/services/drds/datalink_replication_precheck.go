package drds

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DatalinkReplicationPrecheck invokes the drds.DatalinkReplicationPrecheck API synchronously
// api document: https://help.aliyun.com/api/drds/datalinkreplicationprecheck.html
func (client *Client) DatalinkReplicationPrecheck(request *DatalinkReplicationPrecheckRequest) (response *DatalinkReplicationPrecheckResponse, err error) {
	response = CreateDatalinkReplicationPrecheckResponse()
	err = client.DoAction(request, response)
	return
}

// DatalinkReplicationPrecheckWithChan invokes the drds.DatalinkReplicationPrecheck API asynchronously
// api document: https://help.aliyun.com/api/drds/datalinkreplicationprecheck.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DatalinkReplicationPrecheckWithChan(request *DatalinkReplicationPrecheckRequest) (<-chan *DatalinkReplicationPrecheckResponse, <-chan error) {
	responseChan := make(chan *DatalinkReplicationPrecheckResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DatalinkReplicationPrecheck(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DatalinkReplicationPrecheckWithCallback invokes the drds.DatalinkReplicationPrecheck API asynchronously
// api document: https://help.aliyun.com/api/drds/datalinkreplicationprecheck.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DatalinkReplicationPrecheckWithCallback(request *DatalinkReplicationPrecheckRequest, callback func(response *DatalinkReplicationPrecheckResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DatalinkReplicationPrecheckResponse
		var err error
		defer close(result)
		response, err = client.DatalinkReplicationPrecheck(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DatalinkReplicationPrecheckRequest is the request struct for api DatalinkReplicationPrecheck
type DatalinkReplicationPrecheckRequest struct {
	*requests.RpcRequest
	SrcTableName   string `position:"Query" name:"SrcTableName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	DbName         string `position:"Query" name:"DbName"`
	DstTableName   string `position:"Query" name:"DstTableName"`
}

// DatalinkReplicationPrecheckResponse is the response struct for api DatalinkReplicationPrecheck
type DatalinkReplicationPrecheckResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      int64  `json:"Data" xml:"Data"`
}

// CreateDatalinkReplicationPrecheckRequest creates a request to invoke DatalinkReplicationPrecheck API
func CreateDatalinkReplicationPrecheckRequest() (request *DatalinkReplicationPrecheckRequest) {
	request = &DatalinkReplicationPrecheckRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "DatalinkReplicationPrecheck", "Drds", "openAPI")
	return
}

// CreateDatalinkReplicationPrecheckResponse creates a response to parse from DatalinkReplicationPrecheck response
func CreateDatalinkReplicationPrecheckResponse() (response *DatalinkReplicationPrecheckResponse) {
	response = &DatalinkReplicationPrecheckResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}