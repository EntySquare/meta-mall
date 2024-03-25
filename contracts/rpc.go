package contract

import (
	"bytes"
	"net/http"
	"unsafe"
)

//	func chainGetBlockByHashRpc(ip string, hash string) (jsonstr string, err error) {
//		json := `
//				{ "jsonrpc": "2.0", "method":"eth_getBlockByHash", "params": [` + hash + `, false], "id": 0}
//			`
//		reader := bytes.NewReader([]byte(json))
//		resp, err := doRequest(ip, reader)
//		if err != nil {
//			return "", err
//		}
//		buf := new(bytes.Buffer)
//		buf.ReadFrom(resp.Body)
//		respBytes := buf.Bytes()
//		//	respBytes, err := ioutil.ReadAll(resp.Body)
//		if err != nil {
//			return "", err
//		}
//		str := (*string)(unsafe.Pointer(&respBytes))
//		return *str, nil
//	}
func chainGetTransactionByHashRpc(apiUrl string, hash string) (jsonstr string, err error) {
	json := `
			{ "jsonrpc": "2.0", "method":"eth_getTransactionByHash", "params": ["` + hash + `"], "id": 0}
		`
	reader := bytes.NewReader([]byte(json))
	resp, err := doRequest(apiUrl, reader)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respBytes := buf.Bytes()
	//	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	str := (*string)(unsafe.Pointer(&respBytes))
	return *str, nil
}
func doRequest(ip string, reader *bytes.Reader) (resp *http.Response, err error) {
	request, err := http.NewRequest("POST", ip, reader)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	return client.Do(request)
}
