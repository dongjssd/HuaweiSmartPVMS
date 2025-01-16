/**
    @author: dongjs
    @date: 2025/1/11
    @description:
**/

package HuaweiSmartPVMS

import (
	"io"
	"io/ioutil"
	"net/http"
)

// 请求方法
func (c *Client) doRequest(url string, buf io.Reader) ([]byte, error) {
	req, _ := http.NewRequest("POST", ManagementSystemDomainName+url, buf)
	req.Header.Set("Content-Type", "application/json")
	token, err := c.getTokenFromRedis() //获取redis中的token
	if err != nil {
		return nil, err
	}
	req.Header.Set("XSRF-TOKEN", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return body, nil
}
