/**
    @author: dongjs
    @date: 2025/1/11
    @description:
**/

package HuaweiSmartPVMS

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type CommonResponse struct {
	Data any `json:"data"`
	//FailCode int         `json:"failCode"`
	//Params   interface{} `json:"params"`
	Success bool `json:"success"`
}

// 请求方法
func (c *Client) doRequest(url string, buf io.Reader) ([]byte, error) {
	fmt.Println("\033[97;34m", "call:", url, "\033[0m")
	req, _ := http.NewRequest("POST", ManagementSystemDomainName+url, buf)
	req.Header.Set("Content-Type", "application/json")
	token, err := c.getTokenFromRedis() //获取redis中的token
	if err != nil {
		return nil, err
	}
	req.Header.Set("XSRF-TOKEN", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%s%s%s\n", "\033[97;35m", string(body), "\033[0m")
	var check CommonResponse
	err = json.Unmarshal(body, &check)
	if err != nil {
		return nil, err
	}
	if !check.Success {
		return nil, errors.New(fmt.Sprintf("%s:%v", url, check.Data))
	}

	return body, nil
}
