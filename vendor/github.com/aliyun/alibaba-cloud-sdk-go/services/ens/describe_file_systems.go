package ens

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

// DescribeFileSystems invokes the ens.DescribeFileSystems API synchronously
func (client *Client) DescribeFileSystems(request *DescribeFileSystemsRequest) (response *DescribeFileSystemsResponse, err error) {
	response = CreateDescribeFileSystemsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeFileSystemsWithChan invokes the ens.DescribeFileSystems API asynchronously
func (client *Client) DescribeFileSystemsWithChan(request *DescribeFileSystemsRequest) (<-chan *DescribeFileSystemsResponse, <-chan error) {
	responseChan := make(chan *DescribeFileSystemsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeFileSystems(request)
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

// DescribeFileSystemsWithCallback invokes the ens.DescribeFileSystems API asynchronously
func (client *Client) DescribeFileSystemsWithCallback(request *DescribeFileSystemsRequest, callback func(response *DescribeFileSystemsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeFileSystemsResponse
		var err error
		defer close(result)
		response, err = client.DescribeFileSystems(request)
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

// DescribeFileSystemsRequest is the request struct for api DescribeFileSystems
type DescribeFileSystemsRequest struct {
	*requests.RpcRequest
	PageNumber     requests.Integer `position:"Query" name:"PageNumber"`
	FileSystemName string           `position:"Query" name:"FileSystemName"`
	EnsRegionId    string           `position:"Query" name:"EnsRegionId"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	FileSystemId   string           `position:"Query" name:"FileSystemId"`
}

// DescribeFileSystemsResponse is the response struct for api DescribeFileSystems
type DescribeFileSystemsResponse struct {
	*responses.BaseResponse
	PageNumber  int               `json:"PageNumber" xml:"PageNumber"`
	PageSize    int               `json:"PageSize" xml:"PageSize"`
	TotalCount  int               `json:"TotalCount" xml:"TotalCount"`
	RequestId   string            `json:"RequestId" xml:"RequestId"`
	FileSystems []FileSystemsItem `json:"FileSystems" xml:"FileSystems"`
}

// CreateDescribeFileSystemsRequest creates a request to invoke DescribeFileSystems API
func CreateDescribeFileSystemsRequest() (request *DescribeFileSystemsRequest) {
	request = &DescribeFileSystemsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeFileSystems", "ens", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeFileSystemsResponse creates a response to parse from DescribeFileSystems response
func CreateDescribeFileSystemsResponse() (response *DescribeFileSystemsResponse) {
	response = &DescribeFileSystemsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}