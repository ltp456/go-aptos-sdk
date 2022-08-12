package go_apots_sdk

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ltp456/go-apots-sdk/crypto"
	"github.com/ltp456/go-apots-sdk/form"
	"github.com/ltp456/go-apots-sdk/types"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"reflect"
	"strings"
	"time"
)

type ApotsClient struct {
	imp            *http.Client
	endpoint       string
	debug          bool
	faucetEndpoint string
}

func NewApotsClient(endpoint string) (*ApotsClient, error) {
	fbClient := &ApotsClient{
		endpoint: endpoint,
		imp:      http.DefaultClient,
	}
	return fbClient, nil
}

// faucet

func (ap *ApotsClient) SetFaucetEndpoint(url string) {
	ap.faucetEndpoint = url
}

func (ap *ApotsClient) FaucetFundAccount(address string, amount uint64) (string, error) {
	if ap.faucetEndpoint == "" {
		return "", fmt.Errorf("please SetFaucetEndpoint")
	}
	data, err := ap.http(http.MethodPost, fmt.Sprintf("%s/mint?amount=%v&auth_key=%s", ap.faucetEndpoint, amount, address), nil)
	if err != nil {
		return "", err
	}
	return string(data), nil

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

func (ap *ApotsClient) apiDocument() (string, error) {

	data, err := ap.http(http.MethodPost, fmt.Sprintf("%s/spec.html", ap.endpoint), nil)
	if err != nil {
		return "", err
	}
	return string(data), nil

}

func (ap *ApotsClient) openApiDocument() (string, error) {

	data, err := ap.http(http.MethodPost, fmt.Sprintf("%s/openapi.yaml", ap.endpoint), nil)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// transaction

func (ap *ApotsClient) Transfer(seed []byte, sender, recipient, amount, maxGasAmount, gasUnitPrice string, options ...Value) (*types.SubmitTransactionResp, error) {

	account, err := ap.GetAccount(sender)
	if err != nil {
		return nil, err
	}
	keyPair, err := crypto.NewKeyPairFromSeed(seed)
	if err != nil {
		return nil, err
	}
	payload := form.Payload{
		Type:          string(ScriptFunctionPayload),
		Function:      string(CoinTransfer),
		TypeArguments: []string{string(ApotsCoin)},
		Arguments:     []string{recipient, amount},
	}

	expirationTimestampSec := fmt.Sprintf("%v", time.Now().Unix()+600)

	message, err := ap.CreateTxSignMessage(sender, account.SequenceNumber, maxGasAmount, gasUnitPrice, expirationTimestampSec, payload, options...)
	if err != nil {
		return nil, err
	}
	toSign, err := hex.DecodeString(strings.TrimPrefix(message.Message, "0x"))
	if err != nil {
		return nil, err
	}
	signature, err := keyPair.Sign(toSign)
	if err != nil {
		return nil, err
	}
	signatureInfo := form.Signature{
		Type:      string(Ed25519Signature),
		PublicKey: fmt.Sprintf("0x%x", keyPair.Public()),
		Signature: fmt.Sprintf("0x%x", signature),
	}
	transaction, err := ap.SubmitTransaction(sender, account.SequenceNumber, maxGasAmount, gasUnitPrice, expirationTimestampSec,
		payload, signatureInfo, options...)
	if err != nil {
		return nil, err
	}
	return transaction, nil

}

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

func (ap *ApotsClient) SubmitTransaction(sender, sequenceNumber, maxGasAmount, gasUnitPrice,
	expirationTimestampSec string, payload form.Payload, signature form.Signature, options ...Value) (*types.SubmitTransactionResp, error) {
	params := Params{}
	params.SetValue("sender", sender)
	params.SetValue("sequence_number", sequenceNumber)
	params.SetValue("max_gas_amount", maxGasAmount)
	params.SetValue("gas_unit_price", gasUnitPrice)
	params.SetValue("expiration_timestamp_secs", expirationTimestampSec)
	params.SetValue("payload", payload)
	params.SetValue("signature", signature)
	result := &types.SubmitTransactionResp{}
	err := ap.post("transactions", params, result, options...)
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

func (ap *ApotsClient) CreateTxSignMessage(sender, sequenceNumber, maxGasAmount, gasUnitPrice,
	expirationTimestampSec string, payload form.Payload, options ...Value) (*types.SignMessage, error) {
	params := Params{}
	params.SetValue("sender", sender)
	params.SetValue("sequence_number", sequenceNumber)
	params.SetValue("max_gas_amount", maxGasAmount)
	params.SetValue("gas_unit_price", gasUnitPrice)
	params.SetValue("expiration_timestamp_secs", expirationTimestampSec)
	params.SetValue("payload", payload)
	result := &types.SignMessage{}
	err := ap.post("transactions/signing_message", params, result, options...)
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

func (ap *ApotsClient) ApotsBalance(address string) (*big.Int, error) {
	accountResources, err := ap.AccountResource(address)
	if err != nil {
		return nil, err
	}
	for _, item := range accountResources {
		if item.Type == string(ApotsCoinRes) {
			balanceBig, ok := big.NewInt(0).SetString(item.Data.Coin.Value, 10)
			if !ok {
				return nil, fmt.Errorf("parse big error: %v", item.Data)
			}
			return balanceBig, nil
		}
	}
	return big.NewInt(0), nil
}

func (ap *ApotsClient) AccountResource(address string, options ...Value) ([]types.AccountResource, error) {
	params := Params{}
	var result []types.AccountResource
	err := ap.get(fmt.Sprintf("accounts/%s/resources", address), params, &result, options...)
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
	for _, opt := range options {
		if params == nil {
			break
		}
		params.SetValue(opt.Key, opt.Value)
	}
	return ap.httpReq(http.MethodGet, fmt.Sprintf("%v?%v", path, params.Encode()), nil, value, []Value{}...)

}

func (ap *ApotsClient) newRequest(httpMethod, url string, reqData []byte) (*http.Request, error) {
	req, err := http.NewRequest(httpMethod, url, bytes.NewReader(reqData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func (ap *ApotsClient) http(httpMethod, url string, reqData []byte) ([]byte, error) {
	request, err := ap.newRequest(httpMethod, url, reqData)
	if err != nil {
		return nil, err
	}
	response, err := ap.imp.Do(request)
	if err != nil {
		panic(err)
	}
	if response != nil && response.Body != nil {
		defer response.Body.Close()
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("%v %v %v", response.StatusCode, response.Status, string(data))
	}
	return data, nil
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
	req, err := ap.newRequest(httpMethod, fmt.Sprintf("%s/%s", ap.endpoint, path), requestData)
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
