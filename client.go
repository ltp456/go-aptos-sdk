package go_apots_sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-apots-sdk/types"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

type ApotsClient struct {
	imp      *http.Client
	endpoint string
	debug    bool
}

func NewApotsClient(endpoint string) (*ApotsClient, error) {
	fbClient := &ApotsClient{
		endpoint: endpoint,
		imp:      http.DefaultClient,
	}
	return fbClient, nil
}

func (ap *ApotsClient) GetAccount(address string) (*types.Account, error) {
	result := &types.Account{}
	err := ap.GET(fmt.Sprintf("/accounts/%s", address), Params{}, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ap *ApotsClient) POST(method string, param Params, value interface{}, options ...Value) error {
	for _, v := range options {
		if param == nil {
			break
		}
		param.SetValue(v.Key, v.Value)
	}
	return ap.httpReq(http.MethodPost, method, param, value)
}

func (ap *ApotsClient) PUT(method string, param Params, value interface{}, options ...Value) error {
	for _, v := range options {
		if param == nil {
			break
		}
		param.SetValue(v.Key, v.Value)
	}
	return ap.httpReq(http.MethodPut, method, param, value)
}

func (ap *ApotsClient) DELETE(method string, param Params, value interface{}, options ...Value) error {
	for _, v := range options {
		if param == nil {
			break
		}
		param.SetValue(v.Key, v.Value)
	}
	return ap.httpReq(http.MethodDelete, method, param, value)
}

func (ap *ApotsClient) GET(path string, params Params, value interface{}, options ...Value) error {
	for _, v := range options {
		if params == nil {
			break
		}
		params.SetValue(v.Key, v.Value)
	}
	return ap.httpReq(http.MethodGet, fmt.Sprintf("%v?%v", path, params.Encode()), nil, value)

}

func (ap *ApotsClient) newRequest(httpMethod, url string, reqData []byte) (*http.Request, error) {
	req, err := http.NewRequest(httpMethod, fmt.Sprintf("%s%s", ap.endpoint, url), bytes.NewReader(reqData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func (ap *ApotsClient) httpReq(httpMethod, path string, param interface{}, value interface{}) (err error) {
	vi := reflect.ValueOf(value)
	if vi.Kind() != reflect.Ptr {
		return fmt.Errorf("value must be pointer")
	}
	requestData, err := json.Marshal(param)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	if Debug {
		log.Printf("httpReq request: %v  %v \n", path, string(requestData))
	}
	req, err := ap.newRequest(httpMethod, path, requestData)
	if err != nil {
		return err
	}

	resp, err := ap.imp.Do(req)
	if err != nil {
		if resp != nil && resp.Body != nil {
			defer resp.Body.Close()
		}

		return err
	}
	if resp == nil || resp.StatusCode < http.StatusOK || resp.StatusCode > 300 {
		data, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("response err: %v %v %v", resp.StatusCode, resp.Status, string(data))
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if Debug {
		log.Printf("httpReq response: %v %v \n", path, string(data))
	}
	if len(data) != 0 {
		err = json.Unmarshal(data, value)
		if err != nil {
			return fmt.Errorf("%s%s", path, string(data))
		}
	}
	return nil

}
