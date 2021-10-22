// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0
package zedcloudapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/zededa/zedcloud-api/swagger_models"
	"golang.org/x/net/publicsuffix"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
)

type Client struct {
	Bearer     runtime.ClientAuthInfoWriter
	HttpClient *http.Client
	// If bearerToken is not empty, it is used for Auth. This is the token
	// obtained from zedcontrol UI ( user profile page ).
	// Else, Login is done using username and password. bearerToken is set from
	// response of Login.
	BearerToken      string
	BaseUrl          string
	XRequestIdPrefix string
	CSRFToken        string

	// Set this flag for debugging information to be printed
	Debug bool
}

func (client *Client) logDebug(fmtStr string, args ...interface{}) {
	log.Printf(fmtStr, args...)
}

// UrlForNameOrId
//  ID is prefferred if specified. If not, use name.
func UrlForNameOrId(name, id string) (string, error) {
	if id != "" {
		return "id/" + id, nil
	}
	if name != "" {
		return "name/" + name, nil
	}
	return "", fmt.Errorf("Neither Name nor ID specified")
}

// UrlForObjectRequest
func UrlForObjectRequest(urlExtension, name, id, reqType string) string {
	url, err := UrlForNameOrId(name, id)
	if err != nil {
		return ""
	}
	url = fmt.Sprintf("%s/%s", urlExtension, url)
	switch reqType {
	case "config", "update", "delete":
		return url
	case "status":
		return url + "/status"
	case "publish":
		return url + "/publish"
	case "apply":
		return url + "/apply"
	case "uplink":
		return url + "/uplink"
	default:
		panic("UrlForObjectRequest - Invalid reqType: " + reqType)
	}
}

func reqBodyReaderForData(data interface{}) *bytes.Buffer {
	if data == nil {
		return nil
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("***Failed to Marshall Params. err: %v", err)
		return nil
	}
	return bytes.NewBuffer(jsonData)
}

func isHttpStatusCodeSuccess(statusCode int32) bool {
	switch statusCode {
	case http.StatusOK, http.StatusAccepted:
		return true
	}
	return false
}

func (client *Client) createRequest(method, urlExtension string,
	data interface{}) (*http.Request, error) {
	url := client.BaseUrl + urlExtension
	var err error
	var req *http.Request
	if data == nil {
		req, err = http.NewRequest(method, url, nil)
	} else {
		reqBodyReader := reqBodyReaderForData(data)
		req, err = http.NewRequest(method, url, reqBodyReader)
	}
	if err != nil {
		log.Printf("***Failed to get Request. err: %v", err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Request-Id", client.XRequestIdPrefix+"-"+time.Now().String())
	if client.BearerToken != "" {
		req.Header.Add("Authorization", "Bearer "+client.BearerToken)
	}
	if client.CSRFToken != "" {
		req.Header.Add("X-CSRF-Token", client.CSRFToken)
	}
	return req, nil
}

func isZsrvErrorSuccess(err *swagger_models.ZsrvError) bool {
    switch *err.Ec {
        case swagger_models.ZsrvErrorCodeZMsgSucess,
             swagger_models.ZsrvErrorCodeZMsgAccepted:
             return true
    }
    return false
}

func processZsrvResponseError(rspData interface{}) error {
	rsp, ok := rspData.(*swagger_models.ZsrvResponse)
	if !ok {
		// get calls don't use ZsrvResponse. Ignore.
		return nil
	}
	if rsp == nil {
		panic("unexpected nil rsp in processZsrvResponseError")
	}
	errStr := ""
	if !isHttpStatusCodeSuccess(rsp.HTTPStatusCode) {
		errStr = fmt.Sprintf("Request Failed. HTTPStatusCode: %d, HTTPStatusMsg: %s",
			rsp.HTTPStatusCode, rsp.HTTPStatusMsg)
	}
	for _, zerr := range rsp.Error {
		errCode := ""
		if zerr.Ec != nil && !isZsrvErrorSuccess(zerr) {
			errCode = string(*zerr.Ec)
			errStr += " " + fmt.Sprintf("ErrorCode: %s, ErrorDetails: %s\n",
				errCode, zerr.Details)
		}
	}
	if errStr != "" {
		return fmt.Errorf("%s", errStr)
	}
	return nil
}

// processResponse
//  Returns Nil if Response is a success.
//  If response failed, Returns an error
//      - include Response Status code in Error.
func (client *Client) processResponse(resp *http.Response,
	rspData interface{}) error {
	client.SetCsrfToken(resp)
	var err error
	reqSuccess := false
	if isHttpStatusCodeSuccess(int32(resp.StatusCode)) {
		client.logDebug("[INFO]Request SUCCESS")
		reqSuccess = true
	} else {
		err = fmt.Errorf("Request FAILED. StatusCode: %s", resp.Status)
		client.logDebug("***[ERROR]Request FAILED: %s, resp: %+v", resp.Status, resp)
	}
	if rspData == nil || resp.Body == nil {
		client.logDebug("[INFO] not processing response body - rspData: %+v, resp.Body: %+v",
			rspData, resp.Body)
		// caller  not interested in rsp body Or NO response body
		return err
	}
	defer resp.Body.Close()
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		if !reqSuccess {
			err = fmt.Errorf("%s\n Failed to read response Body. %s",
				err.Error(), readErr.Error())
		}
	}
	if body == nil {
		client.logDebug("***[ERROR]nil body")
		return err
	}
	unmarshalErr := json.Unmarshal(body, rspData)
	if unmarshalErr == nil {
		client.logDebug("[INFO] Successfully Unmarshalled Response. rspData: %++v", rspData)
		zerr := processZsrvResponseError(rspData)
		if zerr != nil {
			if err != nil {
				err = fmt.Errorf("%s\n%s", err, zerr)
			} else {
				err = zerr
			}
		}
		return err
	}
	err = fmt.Errorf("%v\n***[ERROR]Failed to unmarshal response body. err: %v\n"+
		"rsp body: %s", err, unmarshalErr, body)
	if reqSuccess {
		// Ignore errors if the request has been a success. Log them if debug
		// is set.
		client.logDebug("***[ERROR}Errors when unmarshalling Response: %+v", err)
		return nil
	}
	client.logDebug("%s", err.Error())
	return err
}

// SendReq
//  Return Values:
//   Returns the http.Response received and any errors encountered in the
//   process.
//   If Response is Successful ( isHttpStatusCodeSuccess(resp.StatusCode)), returns nil as error.
//   Caller can use err == nil to check for a successful response.
func (client *Client) SendReq(method, urlExtension string,
	data interface{},
	rspData interface{}) (*http.Response, error) {
	req, err := client.createRequest(method, urlExtension, data)
	if err != nil {
		return nil, err
	}
	client.logDebug("%s %s", method, req.URL)
	resp, err := client.HttpClient.Do(req)
	if err != nil {
		return resp, err
	}
	err = client.processResponse(resp, rspData)
	if err != nil {
		client.logDebug("SEND REQ - ERROR: %v", err)
	}
	return resp, err
}

func (client *Client) SetCsrfToken(resp *http.Response) {
	for k, v := range resp.Header {
		if k == "X-CSRF-Token" || k == "X-Csrf-Token" {
			client.CSRFToken = v[0]
			break
		}
	}
}

func (client *Client) login(username, password string) error {
	// Get CSRF token for Login by doing a get.
	client.SendReq("GET", "login", nil, nil)

	parm := swagger_models.AAAFrontendLoginWithPasswordRequest{
		UsernameAtRealm: username,
		Password:        password,
	}
	loginRsp := &swagger_models.AAAFrontendLoginResponse{}
	_, err := client.SendReq("POST", "login", &parm, loginRsp)
	if err != nil {
		return err
	}
	client.BearerToken = loginRsp.Token.Base64
	client.Bearer = httptransport.BearerToken(client.BearerToken)
	return nil
}

// CreateObj
//  typeUrl - Ex: "devices", "apps", "instances" etc
//  id - ID of the object to delete
func (client *Client) CreateObj(typeUrl string, reqBody interface{}) (
	*http.Response, *swagger_models.ZsrvResponse, error) {
	rspData := &swagger_models.ZsrvResponse{}
	resp, err := client.SendReq("POST", typeUrl, reqBody, rspData)
	return resp, rspData, err
}

// DeleteObj
//  typeUrl - Ex: "devices", "apps", "instances" etc
//  id - ID of the object to delete
func (client *Client) DeleteObj(typeUrl, id string) (
	*http.Response, *swagger_models.ZsrvResponse, error) {
	urlExtension := fmt.Sprintf("%s/id/%s", typeUrl, id)
	rspData := &swagger_models.ZsrvResponse{}
	resp, err := client.SendReq("DELETE", urlExtension, nil, rspData)
	return resp, rspData, err
}

// GetObj
//  typeUrl - Ex: "devices", "apps", "instances" etc
// rspBody -> pointer to the expected response body. Response body would be
// unmarshalled into this.
//  Ex:
//      rspData := &swagger_models.DeviceConfig{}
//      client.GetObj("devices", "testDevice", "", false, rspData)
func (client *Client) GetObj(typeUrl, name, id string,
	status bool,
	rspBody interface{}) error {

	if client == nil {
		return fmt.Errorf("client.GetObj - nil client")
	}

	// Get CSRF token for Login by doing a get.
	url, err := UrlForNameOrId(name, id)
	if err != nil {
		return err
	}
	url = typeUrl + "/" + url
	if status {
		url += "/status"
	}
	client.XRequestIdPrefix = "TF-GetObj"
	_, err = client.SendReq("GET", url, nil, rspBody)
	return err
}

func NewClient(zedcloudUrl string, token, username, password string) (*Client, error) {
	log.Println("[INFO] NewClient")
	// create the transport
	schemes := []string{"https"}
	client := &Client{
		HttpClient:       &http.Client{},
		BaseUrl:          zedcloudUrl + "/api/v1/",
		XRequestIdPrefix: "zedcloudapi-client",
		Debug:            true,
	}
	transport := httptransport.New(zedcloudUrl, "api", schemes)
	transport.SetDebug(true)

	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return nil, err
	}
	transport.Jar = jar
	client.HttpClient.Jar = jar

	// Set BearerToken
	if token != "" {
		client.Bearer = httptransport.BearerToken(token)
		client.BearerToken = token
		log.Println("[INFO] Created Bearer")
	} else {
		if username == "" || password == "" {
			err = fmt.Errorf("Token or Username/Password must be specified")
			log.Printf("***[ERROR] %s", err.Error())
			return nil, err
		}
		// Get BearerToken by doing a Login
		err = client.login(username, password)
		if err != nil {
			err = fmt.Errorf("LOGIN FAILED: %s", err.Error())
			return nil, err
		}
	}
	return client, nil
}
