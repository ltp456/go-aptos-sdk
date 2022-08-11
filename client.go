package go_apots_sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ltp456/go-apots-sdk/form"
	"github.com/ltp456/go-apots-sdk/types"
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

// general

func (ap *ApotsClient) LedgerInformation() (*types.LedgerInformation, error) {
	result := &types.LedgerInformation{}
	err := ap.get("", Params{}, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ap *ApotsClient) apiDocument() (interface{}, error) {
	var result interface{}
	err := ap.get("spec.html", Params{}, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ap *ApotsClient) openApiDocument() (interface{}, error) {
	var result interface{}
	err := ap.get("openapi.yaml", Params{}, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// transaction

func (ap *ApotsClient) Transactions(start, limit int64) ([]types.Transaction, error) {
	var result []types.Transaction
	params := Params{}
	params.SetValue("start", start)
	params.SetValue("limit", limit)
	err := ap.get("transactions", params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ap *ApotsClient) SubmitTransaction(sender, sequenceNumber, maxGasAmount, gasUnitPrice, gasCurrencyCode,
	expirationTimestampSec string, payload form.Payload, signature form.Signature) (*types.SubmitTransactionResp, error) {
	params := Params{}
	params.SetValue("sender", sender)
	params.SetValue("sequenceNumber", sequenceNumber)
	params.SetValue("maxGasAmount", maxGasAmount)
	params.SetValue("gasUnitPrice", gasUnitPrice)
	params.SetValue("gasCurrencyCode", gasCurrencyCode)
	params.SetValue("expirationTimestampSec", expirationTimestampSec)
	params.SetValue("payload", payload)
	params.SetValue("signature", signature)
	result := &types.SubmitTransactionResp{}
	err := ap.post("transactions", params, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ap *ApotsClient) SimulateTransaction(sender, sequenceNumber, maxGasAmount, gasUnitPrice, gasCurrencyCode,
	expirationTimestampSec string, payload form.Payload, signature form.Signature) (*types.Transaction, error) {
	params := Params{}
	params.SetValue("sender", sender)
	params.SetValue("sequenceNumber", sequenceNumber)
	params.SetValue("maxGasAmount", maxGasAmount)
	params.SetValue("gasUnitPrice", gasUnitPrice)
	params.SetValue("gasCurrencyCode", gasCurrencyCode)
	params.SetValue("expirationTimestampSec", expirationTimestampSec)
	params.SetValue("payload", payload)
	params.SetValue("signature", signature)
	result := &types.Transaction{}
	err := ap.post("transactions/simulate", params, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ap *ApotsClient) AccountTransactions(address string, start, limit int64) ([]types.Transaction, error) {
	var result []types.Transaction
	params := Params{}
	params.SetValue("start", start)
	params.SetValue("limit", limit)
	err := ap.get(fmt.Sprintf("accounts/%s/transactions", address), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ap *ApotsClient) Transaction(txHash string) (*types.Transaction, error) {
	result := &types.Transaction{}
	err := ap.get(fmt.Sprintf("transactions/%s", txHash), Params{}, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ap *ApotsClient) CreateTxSignMessage(sender, sequenceNumber, maxGasAmount, gasUnitPrice, gasCurrencyCode,
	expirationTimestampSec string, payload form.Payload, signature form.Signature) (*types.SignMessage, error) {
	params := Params{}
	params.SetValue("sender", sender)
	params.SetValue("sequenceNumber", sequenceNumber)
	params.SetValue("maxGasAmount", maxGasAmount)
	params.SetValue("gasUnitPrice", gasUnitPrice)
	params.SetValue("gasCurrencyCode", gasCurrencyCode)
	params.SetValue("expirationTimestampSec", expirationTimestampSec)
	params.SetValue("payload", payload)
	params.SetValue("signature", signature)
	result := &types.SignMessage{}
	err := ap.post("transactions/signing_message", params, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// account

func (ap *ApotsClient) GetAccount(address string) (*types.Account, error) {
	result := &types.Account{}
	err := ap.get(fmt.Sprintf("accounts/%s", address), Params{}, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ap *ApotsClient) AccountResource(address, version string) ([]types.AccountResource, error) {
	params := Params{}
	params.SetValue("version", version)
	var result []types.AccountResource
	err := ap.get(fmt.Sprintf("accounts/%s/resources", address), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ap *ApotsClient) AccountResourceByResType(address, resType, version string) (*types.AccountResource, error) {
	params := Params{}
	params.SetValue("version", version)
	result := &types.AccountResource{}
	err := ap.get(fmt.Sprintf("/accounts/%s/resource/%s", address, resType), params, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ap *ApotsClient) AccountModules(address, version string) ([]types.AccountModules, error) {
	params := Params{}
	params.SetValue("version", version)
	var result []types.AccountModules
	err := ap.get(fmt.Sprintf("/accounts/%s/modules", address), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ap *ApotsClient) AccountModuleByID(address, moduleName, version string) (*types.AccountModules, error) {
	params := Params{}
	params.SetValue("version", version)
	result := &types.AccountModules{}
	err := ap.get(fmt.Sprintf("/accounts/%s/module/%s", address, moduleName), params, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// events

func (ap *ApotsClient) EventsByKey(eventKey string) ([]types.Event, error) {

	var result []types.Event
	err := ap.get(fmt.Sprintf("events/%s", eventKey), Params{}, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ap *ApotsClient) EventsByEventHandle(address, handle, fieldName string, start, limit int) ([]types.Event, error) {

	params := Params{}
	params.SetValue("start", start)
	params.SetValue("limit", limit)
	var result []types.Event
	err := ap.get(fmt.Sprintf("accounts/%s/events/%s/%s", address, handle, fieldName), params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// state

func (ap *ApotsClient) TableItemByHandleAndKey(tableHandle, keyType, valueType, key string) error {
	params := Params{}
	params.SetValue("key_type", keyType)
	params.SetValue("value_type", valueType)
	params.SetValue("key", key)
	var result string
	err := ap.post(fmt.Sprintf("tables/%s/item", tableHandle), params, &result)
	if err != nil {
		return err
	}
	return nil
}

func (ap *ApotsClient) post(method string, param Params, value interface{}, options ...Value) error {

	return ap.httpReq(http.MethodPost, method, param, value, options...)
}

func (ap *ApotsClient) put(method string, param Params, value interface{}, options ...Value) error {

	return ap.httpReq(http.MethodPut, method, param, value, options...)
}

func (ap *ApotsClient) delete(method string, param Params, value interface{}, options ...Value) error {
	return ap.httpReq(http.MethodDelete, method, param, value, options...)
}

func (ap *ApotsClient) get(path string, params Params, value interface{}, options ...Value) error {

	return ap.httpReq(http.MethodGet, fmt.Sprintf("%v?%v", path, params.Encode()), params, value, options...)

}

func (ap *ApotsClient) newRequest(httpMethod, url string, reqData []byte) (*http.Request, error) {
	req, err := http.NewRequest(httpMethod, fmt.Sprintf("%s/%s", ap.endpoint, url), bytes.NewReader(reqData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func (ap *ApotsClient) httpReq(httpMethod, path string, param Params, value interface{}, options ...Value) (err error) {
	vi := reflect.ValueOf(value)
	if vi.Kind() != reflect.Ptr {
		return fmt.Errorf("value must be pointer")
	}

	if param != nil && len(options) > 0 {
		for _, opt := range options {
			param.SetValue(opt.Key, opt.Value)
		}
	}

	var requestData []byte
	if param != nil {
		requestData, err = json.Marshal(param)
		if err != nil {
			return fmt.Errorf("%v", err)
		}
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
