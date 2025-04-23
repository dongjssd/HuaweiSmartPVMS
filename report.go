/**
    @author: dongjs
    @date: 2025/1/15
    @description:
**/

package HuaweiSmartPVMS

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// 电站小时数据接口
// 获取电站小时数据接口，一次最多支持查询100个电站。
// 后台会根据请求参数collectTime（采集时间毫秒数）以及电站所在时区，计算出该采集时间属于哪一天，
// 然后通过电站编号查询出电站该天每小时的数据。如果该天n（0≤n≤24）小时有数据，那么会返回n（0≤n≤24）条数据。
// stationCodes 电站编号列表，多个电站用英文逗号分隔
// collectTime 采集时间毫秒数
func (c Client) KpiStationHour(stationCodes string, collectTime int64) (*KpiStationHourResponse, error) {
	request := map[string]interface{}{
		"stationCodes": stationCodes,
		"collectTime":  collectTime,
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("/thirdData/getKpiStationHour", buf)
	if err != nil {
		return nil, err
	}
	response := KpiStationHourResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 电站小时数据接口 返回参数结构体
type KpiStationHourResponse struct {
	Success  bool                         `json:"success"` //请求成功或失败标识 true：请求成功 false：请求失败
	FailCode int                          `json:"failCode"`
	Message  string                       `json:"message"`
	Params   KpiStationHourParamsResponse `json:"params"`
	Data     []KpiStationHourDataResponse `json:"data"`
}

// 电站小时数据接口 返回参数结构体
type KpiStationHourParamsResponse struct {
	StationCodes string `json:"stationCodes"` //请求参数中的电站编号列表
	CollectTime  int64  `json:"collectTime"`  //请求参数中的采集时间毫秒数
	CurrentTime  int64  `json:"currentTime"`  //系统当前时间毫秒数
}

// 电站小时数据接口 返回参数结构体
type KpiStationHourDataResponse struct {
	StationCode string                 `json:"stationCode"` //电站编号
	CollectTime int64                  `json:"collectTime"` //采集时间毫秒数
	DataItemMap map[string]interface{} `json:"dataItemMap"` //每个数据项的内容，用key-value形式返回，数据项列表参见下方电站小时数据列表。
}

// 电站日数据接口
// 获取电站日数据接口，stationCodes 一次最多支持查询100个电站。 电站编号列表，多个电站用英文逗号分隔
// 后台会根据请求参数collectTime（采集时间毫秒数）以及电站所在时区，计算出该采集时间属于哪一月，
// 然后通过电站编号查询出电站该月每天的数据。如果该月n（0≤n≤31）天有数据，那么会返回n（0≤n≤31）条数据。
func (c *Client) KpiStationDay(stationCodes string, collectTime int64) (*KpiStationDayResponse, error) {
	request := map[string]interface{}{
		"stationCodes": stationCodes,
		"collectTime":  collectTime,
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("/thirdData/getKpiStationDay", buf)
	if err != nil {
		return nil, err
	}
	response := KpiStationDayResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 电站日数据接口 返回值结构体
type KpiStationDayResponse struct {
	Success  bool                        `json:"success"` //请求成功或失败标识 true：请求成功 false：请求失败
	FailCode int                         `json:"failCode"`
	Message  string                      `json:"message"`
	Params   KpiStationDayParamsResponse `json:"params"`
	Data     []KpiStationDayDataResponse `json:"data"`
}

// 电站日数据接口 返回参数结构体
type KpiStationDayParamsResponse struct {
	StationCodes string `json:"stationCodes"` //请求参数中的电站编号列表
	CollectTime  int64  `json:"collectTime"`  //请求参数中的采集时间毫秒数
	CurrentTime  int64  `json:"currentTime"`  //系统当前时间毫秒数
}

// 电站日数据接口 返回参数结构体
type KpiStationDayDataResponse struct {
	StationCode string                 `json:"stationCode"` //电站编号
	CollectTime int64                  `json:"collectTime"` //采集时间毫秒数
	DataItemMap map[string]interface{} `json:"dataItemMap"` //每个数据项的内容，用key-value形式返回，数据项列表参见下方电站小时数据列表。
}

// 电站月数据接口
// 获取电站月数据接口，一次查询最多支持100个电站。
// 后台会根据请求参数collectTime（采集时间毫秒数）以及电站所在时区，计算出该采集时间属于哪一年，
// 然后通过电站编号查询出电站该年每月的数据。如果该年n（0≤n≤12）个月有数据，那么会返回n（0≤n≤12）条数据。
func (c *Client) KpiStationMonth(stationCodes string, collectTime int64) (*KpiStationMonthResponse, error) {
	request := map[string]interface{}{
		"stationCodes": stationCodes,
		"collectTime":  collectTime,
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("/thirdData/getKpiStationMonth", buf)
	if err != nil {
		return nil, err
	}
	response := KpiStationMonthResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 电站月数据接口 返回值结构体
type KpiStationMonthResponse struct {
	Success  bool                          `json:"success"` //请求成功或失败标识 true：请求成功 false：请求失败
	FailCode int                           `json:"failCode"`
	Message  string                        `json:"message"`
	Params   KpiStationMonthParamsResponse `json:"params"`
	Data     []KpiStationMonthDataResponse `json:"data"`
}

// 电站月数据接口 返回参数结构体
type KpiStationMonthParamsResponse struct {
	StationCodes string `json:"stationCodes"` //请求参数中的电站编号列表
	CollectTime  int64  `json:"collectTime"`  //请求参数中的采集时间毫秒数
	CurrentTime  int64  `json:"currentTime"`  //系统当前时间毫秒数
}

// 电站月数据接口 返回参数结构体
type KpiStationMonthDataResponse struct {
	StationCode string                 `json:"stationCode"` //电站编号
	CollectTime int64                  `json:"collectTime"` //采集时间毫秒数
	DataItemMap map[string]interface{} `json:"dataItemMap"` //每个数据项的内容，用key-value形式返回，数据项列表参见下方电站小时数据列表。
}

// 电站年数据接口
// 获取电站年数据接口，一次最多支持查询100个电站。
// 后台会根据电站编号查询出电站建站至今每一年的数据（包括当前时间所属年）。
func (c *Client) KpiStationYear(stationCodes string, collectTime int64) (*KpiStationYearResponse, error) {
	request := map[string]interface{}{
		"stationCodes": stationCodes,
		"collectTime":  collectTime,
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("/thirdData/getKpiStationYear", buf)
	if err != nil {
		return nil, err
	}
	response := KpiStationYearResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 电站年数据接口 返回值结构体
type KpiStationYearResponse struct {
	Success  bool                         `json:"success"` //请求成功或失败标识 true：请求成功 false：请求失败
	FailCode int                          `json:"failCode"`
	Message  string                       `json:"message"`
	Params   KpiStationYearParamsResponse `json:"params"`
	Data     []KpiStationYearDataResponse `json:"data"`
}

// 电站年数据接口 返回参数结构体
type KpiStationYearParamsResponse struct {
	StationCodes string `json:"stationCodes"` //请求参数中的电站编号列表
	CollectTime  int64  `json:"collectTime"`  //请求参数中的采集时间毫秒数
	CurrentTime  int64  `json:"currentTime"`  //系统当前时间毫秒数
}

// 电站年数据接口 返回参数结构体
type KpiStationYearDataResponse struct {
	StationCode string                 `json:"stationCode"` //电站编号
	CollectTime int64                  `json:"collectTime"` //采集时间毫秒数
	DataItemMap map[string]interface{} `json:"dataItemMap"` //每个数据项的内容，用key-value形式返回，数据项列表参见下方电站小时数据列表。
}

// 设备日数据接口
// 获取设备日维度数据接口，一次最多支持查询1种设备类型、100个设备。
// 后台会根据请求参数collectTime（采集时间毫秒数）以及设备所在时区，计算出该采集时间属于哪一月，
// 然后通过设备编号查询出设备该月每天的数据。如果该月n（0≤n≤31）天有数据，那么会返回n（0≤n≤31）条数据。
func (c *Client) DevKpiDay(devIds, sns string, devTypeId int, collectTime int64) (*DevKpiDayResponse, error) {
	request := map[string]interface{}{
		"devTypeId":   devTypeId,
		"collectTime": collectTime,
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
	body, err := c.doRequest("/thirdData/getDevKpiDay", buf)
	if err != nil {
		return nil, err
	}
	response := DevKpiDayResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 设备日数据接口 返回值对象
type DevKpiDayResponse struct {
	Success  bool                    `json:"success"` //请求成功或失败标识 true：请求成功 false：请求失败
	FailCode int                     `json:"failCode"`
	Message  string                  `json:"message"`
	Params   DevKpiDayParamsResponse `json:"params"`
	Data     []DevKpiDayDataResponse `json:"data"`
}

// 设备日数据接口 返回值对象
type DevKpiDayParamsResponse struct {
	DevIds      string `json:"devIds"`      //请求参数中的设备编号列表
	Sns         string `json:"sns"`         //请求参数中的设备sn号列表
	DevTypeId   int    `json:"devTypeId"`   //请求参数中的设备类型ID
	CollectTime int64  `json:"collectTime"` //请求参数中的采集时间毫秒数
	CurrentTime int64  `json:"currentTime"` //系统当前时间毫秒数
}

// 设备日数据接口 返回值对象
type DevKpiDayDataResponse struct {
	DevId       int64           `json:"devId"`       //设备编号
	Sn          string          `json:"sn"`          //设备SN号
	CollectTime int64           `json:"collectTime"` //采集时间毫秒数
	DataItemMap json.RawMessage `json:"dataItemMap"` //每个数据项的内容，用key-value形式返回，不同设备类型的数据项内容不一样，数据项列表参见下方设备日数据接口
}

// 设备月数据接口
// 获取设备月维度数据接口，一次最多支持查询1种设备类型、100个设备。
// 后台会根据请求参数collectTime（采集时间毫秒数）以及设备所在时区，计算出该采集时间属于哪一年，
// 然后通过设备编号查询出设备该年每月的数据。如果该年n（0≤n≤12）个月有数据，那么会返回n（0≤n≤12）条数据。
func (c *Client) DevKpiMonth(devIds, sns string, devTypeId int, collectTime int64) (*DevKpiMonthResponse, error) {
	request := map[string]interface{}{
		"devTypeId":   devTypeId,
		"collectTime": collectTime,
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
	body, err := c.doRequest("/thirdData/getDevKpiMonth", buf)
	if err != nil {
		return nil, err
	}
	response := DevKpiMonthResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 设备月数据接口 返回值对象
type DevKpiMonthResponse struct {
	Success  bool                      `json:"success"` //请求成功或失败标识 true：请求成功 false：请求失败
	FailCode int                       `json:"failCode"`
	Message  string                    `json:"message"`
	Params   DevKpiMonthParamsResponse `json:"params"`
	Data     []DevKpiMonthDataResponse `json:"data"`
}

// 设备月数据接口 返回值对象
type DevKpiMonthParamsResponse struct {
	DevIds      string `json:"devIds"`      //请求参数中的设备编号列表
	Sns         string `json:"sns"`         //请求参数中的设备sn号列表
	DevTypeId   int    `json:"devTypeId"`   //请求参数中的设备类型ID
	CollectTime int64  `json:"collectTime"` //请求参数中的采集时间毫秒数
	CurrentTime int64  `json:"currentTime"` //系统当前时间毫秒数
}

type InverterKpiMonthDataItem struct {
	ProductPower      float64 `json:"product_power"`
	PerpowerRatio     float64 `json:"perpower_ratio"`
	InstalledCapacity float64 `json:"installed_capacity"`
}

// 设备月数据接口 返回值对象
type DevKpiMonthDataResponse struct {
	DevId       int64           `json:"devId"`       //设备编号
	Sn          string          `json:"sn"`          //设备SN号
	CollectTime int64           `json:"collectTime"` //采集时间毫秒数
	DataItemMap json.RawMessage `json:"dataItemMap"` //每个数据项的内容，用key-value形式返回，不同设备类型的数据项内容不一样，数据项列表参见下方设备日数据接口
}

// 设备年数据接口
// 获取设备年维度数据接口，一次最多支持查询1种设备类型、100个设备。
// 后台会根据设备编号查询出设备接入至今每一年的数据。
func (c *Client) DevKpiYear(devIds, sns string, devTypeId int, collectTime int64) (*DevKpiYearResponse, error) {
	request := map[string]interface{}{
		"devTypeId":   devTypeId,
		"collectTime": collectTime,
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
	body, err := c.doRequest("/thirdData/getDevKpiYear", buf)
	if err != nil {
		return nil, err
	}
	response := DevKpiYearResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 设备年数据接口 返回值对象
type DevKpiYearResponse struct {
	Success  bool                     `json:"success"` //请求成功或失败标识 true：请求成功 false：请求失败
	FailCode int                      `json:"failCode"`
	Message  string                   `json:"message"`
	Params   DevKpiYearParamsResponse `json:"params"`
	Data     []DevKpiYearDataResponse `json:"data"`
}

// 设备年数据接口 返回值对象
type DevKpiYearParamsResponse struct {
	DevIds      string `json:"devIds"`      //请求参数中的设备编号列表
	Sns         string `json:"sns"`         //请求参数中的设备sn号列表
	DevTypeId   int    `json:"devTypeId"`   //请求参数中的设备类型ID
	CollectTime int64  `json:"collectTime"` //请求参数中的采集时间毫秒数
	CurrentTime int64  `json:"currentTime"` //系统当前时间毫秒数
}

// 设备年数据接口 返回值对象
type DevKpiYearDataResponse struct {
	DevId       int64                  `json:"devId"`       //设备编号
	Sn          string                 `json:"sn"`          //设备SN号
	CollectTime int64                  `json:"collectTime"` //采集时间毫秒数
	DataItemMap map[string]interface{} `json:"dataItemMap"` //每个数据项的内容，用key-value形式返回，不同设备类型的数据项内容不一样，数据项列表参见下方设备日数据接口
}
