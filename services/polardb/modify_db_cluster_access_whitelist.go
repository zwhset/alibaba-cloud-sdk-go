package polardb

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

// ModifyDBClusterAccessWhitelist invokes the polardb.ModifyDBClusterAccessWhitelist API synchronously
// api document: https://help.aliyun.com/api/polardb/modifydbclusteraccesswhitelist.html
func (client *Client) ModifyDBClusterAccessWhitelist(request *ModifyDBClusterAccessWhitelistRequest) (response *ModifyDBClusterAccessWhitelistResponse, err error) {
	response = CreateModifyDBClusterAccessWhitelistResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBClusterAccessWhitelistWithChan invokes the polardb.ModifyDBClusterAccessWhitelist API asynchronously
// api document: https://help.aliyun.com/api/polardb/modifydbclusteraccesswhitelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDBClusterAccessWhitelistWithChan(request *ModifyDBClusterAccessWhitelistRequest) (<-chan *ModifyDBClusterAccessWhitelistResponse, <-chan error) {
	responseChan := make(chan *ModifyDBClusterAccessWhitelistResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBClusterAccessWhitelist(request)
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

// ModifyDBClusterAccessWhitelistWithCallback invokes the polardb.ModifyDBClusterAccessWhitelist API asynchronously
// api document: https://help.aliyun.com/api/polardb/modifydbclusteraccesswhitelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDBClusterAccessWhitelistWithCallback(request *ModifyDBClusterAccessWhitelistRequest, callback func(response *ModifyDBClusterAccessWhitelistResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBClusterAccessWhitelistResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBClusterAccessWhitelist(request)
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

// ModifyDBClusterAccessWhitelistRequest is the request struct for api ModifyDBClusterAccessWhitelist
type ModifyDBClusterAccessWhitelistRequest struct {
	*requests.RpcRequest
	ResourceOwnerId           requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityIps               string           `position:"Query" name:"SecurityIps"`
	DBClusterIPArrayAttribute string           `position:"Query" name:"DBClusterIPArrayAttribute"`
	ResourceOwnerAccount      string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId               string           `position:"Query" name:"DBClusterId"`
	OwnerAccount              string           `position:"Query" name:"OwnerAccount"`
	OwnerId                   requests.Integer `position:"Query" name:"OwnerId"`
	DBClusterIPArrayName      string           `position:"Query" name:"DBClusterIPArrayName"`
}

// ModifyDBClusterAccessWhitelistResponse is the response struct for api ModifyDBClusterAccessWhitelist
type ModifyDBClusterAccessWhitelistResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDBClusterAccessWhitelistRequest creates a request to invoke ModifyDBClusterAccessWhitelist API
func CreateModifyDBClusterAccessWhitelistRequest() (request *ModifyDBClusterAccessWhitelistRequest) {
	request = &ModifyDBClusterAccessWhitelistRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "ModifyDBClusterAccessWhitelist", "polardb", "openAPI")
	return
}

// CreateModifyDBClusterAccessWhitelistResponse creates a response to parse from ModifyDBClusterAccessWhitelist response
func CreateModifyDBClusterAccessWhitelistResponse() (response *ModifyDBClusterAccessWhitelistResponse) {
	response = &ModifyDBClusterAccessWhitelistResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
