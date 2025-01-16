/**
    @author: dongjs
    @date: 2025/1/15
    @description:基础类接口 监控
**/

package HuaweiSmartPVMS

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// 电站实时数据接口
// 获取电站实时数据，最小采集周期为5分钟；通过电站编号集合查询，一次最多查询100个电站。
// 电站编号列表，多个电站用英文逗号分隔
func (c *Client) StationRealKpi(stationCodes string) (*StationRealKpiResponse, error) {
	request := map[string]interface{}{
		"stationCodes": stationCodes,
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("/thirdData/getStationRealKpi", buf)
	if err != nil {
		return nil, err
	}
	response := StationRealKpiResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 电站实时数据接口 返回数据结构体
type StationRealKpiResponse struct {
	Success  bool                         `json:"success"` //请求成功或失败标识 true：请求成功 false：请求失败
	FailCode int                          `json:"failCode"`
	Message  string                       `json:"message"`
	Data     []StationRealKpiDataResponse `json:"data"`
}

// 电站实时数据接口 返回数据结构体
type StationRealKpiDataResponse struct {
	StationCode string                        `json:"stationCode"`
	DataItemMap StationRealKpiDataMapResponse `json:"dataItemMap"`
}

// 电站实时数据接口 返回数据结构体
type StationRealKpiDataMapResponse struct {
	DayPower        float64 `json:"day_power"`
	MonthPower      float64 `json:"month_power"`
	TotalPower      float64 `json:"total_power"`
	DayIncome       float64 `json:"day_income"`
	TotalIncome     float64 `json:"total_income"`
	DayOnGridEnergy float64 `json:"day_on_grid_energy"`
	DayUseEnergy    float64 `json:"day_use_energy"`
	RealHealthState int     `json:"real_health_state"`
}

// 设备实时数据接口
// 获取设备实时数据，最小采集周期为5分钟；不同设备类型的实时数据不同；
// 通过设备类型、设备编号集合查询，一次查询最多支持1种设备类型、100个设备。
func (c *Client) DevRealKpi(devIds, sns string, devTypeId int) (*DevRealKpiResponse, error) {
	request := map[string]interface{}{
		"devTypeId": devTypeId,
	}
	if devIds != "" {
		request["devIds"] = devIds
	}
	if sns != "" {
		request["sns"] = sns
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("/thirdData/getDevRealKpi", buf)
	if err != nil {
		return nil, err
	}
	response := DevRealKpiResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 设备实时数据接口 返回结构体
type DevRealKpiResponse struct {
	Success  bool                     `json:"success"` //请求成功或失败标识 true：请求成功 false：请求失败
	FailCode int                      `json:"failCode"`
	Message  string                   `json:"message"`
	Params   DevRealKpiParamsResponse `json:"params"`
	Data     []DevRealKpiDataResponse `json:"data"`
}

// 设备实时数据接口 返回结构体
type DevRealKpiParamsResponse struct {
	DevIds      string `json:"devIds"`
	Sns         string `json:"sns"`
	DevTypeId   int    `json:"devTypeId"`
	CurrentTime int64  `json:"currentTime"`
}

// 设备实时数据接口 返回结构体
type DevRealKpiDataResponse struct {
	DevId       int64                  `json:"devId"`       //设备编号
	Sn          string                 `json:"sn"`          //设备SN号
	DataItemMap map[string]interface{} `json:"dataItemMap"` //每个数据项的内容，用key-value形式返回，不同设备类型的数据项内容不一样，数据项列表参见下方设备实时数据列表。
}

// 设备历史数据接口
// 获取设备指定时间段内的5分钟维度数据，一次最多支持查询1种设备类型、10个设备、3天数据。
func (c *Client) DevHistoryKpi(devIds, sns string, devTypeId int, startTime, endTime int64) (*DevHistoryKpiResponse, error) {
	request := map[string]interface{}{
		"devTypeId": devTypeId,
		"startTime": startTime,
		"endTime":   endTime,
	}
	if devIds != "" {
		request["devIds"] = devIds
	}
	if sns != "" {
		request["sns"] = sns
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("/thirdData/getDevHistoryKpi", buf)
	if err != nil {
		return nil, err
	}
	response := DevHistoryKpiResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 设备历史数据接口 返回参数
type DevHistoryKpiResponse struct {
	Success  bool                      `json:"success"` //请求成功或失败标识 true：请求成功 false：请求失败
	FailCode int                       `json:"failCode"`
	Message  string                    `json:"message"`
	Data     DevHistoryKpiDataResponse `json:"data"`
}

// 设备历史数据接口 返回参数
type DevHistoryKpiDataResponse struct {
	DevId       int                    `json:"devId"`
	Sn          string                 `json:"sn"`
	CollectTime int64                  `json:"collectTime"`
	DataItemMap map[string]interface{} `json:"dataItemMap"`
}
