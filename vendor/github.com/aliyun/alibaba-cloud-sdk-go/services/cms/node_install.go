package cms

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

// NodeInstall invokes the cms.NodeInstall API synchronously
// api document: https://help.aliyun.com/api/cms/nodeinstall.html
func (client *Client) NodeInstall(request *NodeInstallRequest) (response *NodeInstallResponse, err error) {
	response = CreateNodeInstallResponse()
	err = client.DoAction(request, response)
	return
}

// NodeInstallWithChan invokes the cms.NodeInstall API asynchronously
// api document: https://help.aliyun.com/api/cms/nodeinstall.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) NodeInstallWithChan(request *NodeInstallRequest) (<-chan *NodeInstallResponse, <-chan error) {
	responseChan := make(chan *NodeInstallResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.NodeInstall(request)
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

// NodeInstallWithCallback invokes the cms.NodeInstall API asynchronously
// api document: https://help.aliyun.com/api/cms/nodeinstall.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) NodeInstallWithCallback(request *NodeInstallRequest, callback func(response *NodeInstallResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *NodeInstallResponse
		var err error
		defer close(result)
		response, err = client.NodeInstall(request)
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

// NodeInstallRequest is the request struct for api NodeInstall
type NodeInstallRequest struct {
	*requests.RpcRequest
	UserId     string           `position:"Query" name:"UserId"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	Force      requests.Boolean `position:"Query" name:"Force"`
}

// NodeInstallResponse is the response struct for api NodeInstall
type NodeInstallResponse struct {
	*responses.BaseResponse
	ErrorCode    int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
}

// CreateNodeInstallRequest creates a request to invoke NodeInstall API
func CreateNodeInstallRequest() (request *NodeInstallRequest) {
	request = &NodeInstallRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2018-03-08", "NodeInstall", "cms", "openAPI")
	return
}

// CreateNodeInstallResponse creates a response to parse from NodeInstall response
func CreateNodeInstallResponse() (response *NodeInstallResponse) {
	response = &NodeInstallResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}