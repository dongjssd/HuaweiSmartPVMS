package HuaweiSmartPVMS

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-redsync/redsync/v4"
)

// 此接口在10分钟内连续输入5次错误密码，API账户将被锁定，锁定时长为30分钟。
// 每个API账户的限流次数：每10分钟5次。
// 在使用过程中如果超过访问频度的限制，接口会返回错误码407。
// 登录接口，获取数据前必须调用此接口获取XSRF-TOKEN，XSRF-TOKEN过期时间为30分钟。
// 如果XSRF-TOKEN未过期，则可重复使用；如果XSRF-TOKEN已过期，则需再次调用登录接口获取新的XSRF-TOKEN。
// 调用此接口登录成功后，响应头里会返回XSRF-TOKEN。
// 每一次接口调用都会生成一个新的Token，先前获取的Token会失效
// 调用登录接口获取token，用于访问其他数据查询接口和控制接口的身份认证
func (c *Client) login() error {
	var request = loginRequest{
		UserName:   c.userName,
		SystemCode: c.password,
	}

	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(requestBytes)
	fmt.Println(string(requestBytes))
	token, err := c.loginDoRequest("/thirdData/login", buf)
	if err != nil {
		return err
	}
	return c.setTokenToRedis(token)
}

// 登录接口访问
func (c *Client) loginDoRequest(url string, buf io.Reader) (string, error) {
	req, _ := http.NewRequest("POST", ManagementSystemDomainName+url, buf)
	req.Header.Set("Content-Type", "application/json")
	var response Response
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return "", err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	defer resp.Body.Close()
	if err = json.Unmarshal(body, &response); err != nil {
		return "", err
	}
	if response.Success == false {
		//请求失败
		if response.FailCode != 0 {
			return "", errors.New(ErrorCodeMap[fmt.Sprintf("%d", response.FailCode)])
		}
	}
	token := resp.Header.Get("Xsrf-Token")
	return token, nil
}

// 登录接口请求参数结构体
type loginRequest struct {
	UserName   string `json:"userName"`
	SystemCode string `json:"systemCode"`
}

// 返回参数结构体
type Response struct {
	Success  bool `json:"success"`  //请求成功或失败标识 true：请求成功 false：请求失败
	FailCode int  `json:"failCode"` //错误码 0 表示正常，其他错误码参见错误码列表
	Params   struct {
		CurrentTime int64 `json:"currentTime"` //系统当前时间戳（毫秒）
	} `json:"params"`
	Message string      `json:"message"` //可选消息
	Data    interface{} `json:"data"`
}

// 注销接口请求参数
type logoutRequest struct {
	XsrfToken string `json:"xsrfToken"`
}

// 注销接口 如果希望XSRF-TOKEN立即失效，则可以调用此注销接口。
func (c *Client) Logout(token string) error {
	var request = logoutRequest{
		XsrfToken: token,
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(requestBytes)
	response, err := c.LogoutDoRequest("/thirdData/logout", buf)
	if err != nil {
		return err
	}
	if response.Success == false {
		if response.FailCode != 0 {
			return errors.New(ErrorCodeMap[fmt.Sprintf("%d", response.FailCode)])
		}
	}
	return nil
}

func (c *Client) LogoutDoRequest(url string, buf io.Reader) (Response, error) {
	req, _ := http.NewRequest("POST", url, buf)
	req.Header.Set("Content-Type", "application/json")
	var response Response
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return response, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err = json.Unmarshal(body, &response); err != nil {
		return response, err
	}
	return response, nil
}

// 将token存储到redis中
func (c *Client) setTokenToRedis(token string) error {
	lockKey := fmt.Sprintf("setTokenToRedis:userName:%s", c.userName)
	mutex := c.redisSync.NewMutex(
		lockKey,
		redisSyncCustomOptions...)
	if err := mutex.Lock(); err != nil {
		return err
	}

	defer func(mutex *redsync.Mutex) {
		_, _ = mutex.Unlock()
	}(mutex)
	return c.redisClient.Set(context.Background(), c.getTokenKey(), token, 20*time.Minute).Err()
}

// 读取redis中的token
func (c *Client) getTokenFromRedis() (string, error) {
	lockKey := fmt.Sprintf("getTokenFromRedis:userName:%s", c.userName)
	mutex := c.redisSync.NewMutex(
		lockKey,
		redisSyncCustomOptions...)
	if err := mutex.Lock(); err != nil {
		return "", err
	}

	defer func(mutex *redsync.Mutex) {
		_, _ = mutex.Unlock()
	}(mutex)

	token, err := c.redisClient.Get(context.Background(), c.getTokenKey()).Result()
	if err != nil && err.Error() != "redis: nil" {
		return "", err
	}
	if token == "" {
		err = c.login()
		if err != nil {
			return "", err
		}
		token, err = c.redisClient.Get(context.Background(), c.getTokenKey()).Result()
		if err != nil && err.Error() != "redis: nil" {
			return "", err
		}
		//fmt.Println(token)
		return token, nil
	}
	//fmt.Println(token)
	return token, nil
}
func (c *Client) getTokenKey() string {
	return "huawei_smart_pvms:token:" + c.userName
}
