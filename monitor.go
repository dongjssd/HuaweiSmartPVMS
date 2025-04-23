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

// DevRealKpi里用到
type InverterRealData struct {
	InverterState float64 `json:"inverter_state"` // 逆变器状态，参见表5-1 (无)
	AbU           float64 `json:"ab_u"`           // 电网AB电压 (V)
	BcU           float64 `json:"bc_u"`           // 电网BC电压 (V)
	CaU           float64 `json:"ca_u"`           // 电网CA电压 (V)
	AU            float64 `json:"a_u"`            // A相电压 (V)
	BU            float64 `json:"b_u"`            // B相电压 (V)
	CU            float64 `json:"c_u"`            // C相电压 (V)
	AI            float64 `json:"a_i"`            // 电网A相电流 (A)
	BI            float64 `json:"b_i"`            // 电网B相电流 (A)
	CI            float64 `json:"c_i"`            // 电网C相电流 (A)
	Efficiency    float64 `json:"efficiency"`     // 逆变器转换效率(厂家) (%)
	Temperature   float64 `json:"temperature"`    // 机内温度 (℃)
	PowerFactor   float64 `json:"power_factor"`   // 功率因数 (无)
	ElecFreq      float64 `json:"elec_freq"`      // 电网频率 (Hz)
	ActivePower   float64 `json:"active_power"`   // 有功功率 (kW)
	ReactivePower float64 `json:"reactive_power"` // 输出无功功率 (kVar)
	DayCap        float64 `json:"day_cap"`        // 当日发电量 (kWh)
	MpptPower     float64 `json:"mppt_power"`     // MPPT输入总功率 (kW)
	Pv1U          float64 `json:"pv1_u"`          // PV1输入电压 (V)
	Pv2U          float64 `json:"pv2_u"`          // PV2输入电压 (V)
	Pv3U          float64 `json:"pv3_u"`          // PV3输入电压 (V)
	Pv4U          float64 `json:"pv4_u"`          // PV4输入电压 (V)
	Pv5U          float64 `json:"pv5_u"`          // PV5输入电压 (V)
	Pv6U          float64 `json:"pv6_u"`          // PV6输入电压 (V)
	Pv7U          float64 `json:"pv7_u"`          // PV7输入电压 (V)
	Pv8U          float64 `json:"pv8_u"`          // PV8输入电压 (V)
	Pv9U          float64 `json:"pv9_u"`          // PV9输入电压 (V)
	Pv10U         float64 `json:"pv10_u"`         // PV10输入电压 (V)
	Pv11U         float64 `json:"pv11_u"`         // PV11输入电压 (V)
	Pv12U         float64 `json:"pv12_u"`         // PV12输入电压 (V)
	Pv13U         float64 `json:"pv13_u"`         // PV13输入电压 (V)
	Pv14U         float64 `json:"pv14_u"`         // PV14输入电压 (V)
	Pv15U         float64 `json:"pv15_u"`         // PV15输入电压 (V)
	Pv16U         float64 `json:"pv16_u"`         // PV16输入电压 (V)
	Pv17U         float64 `json:"pv17_u"`         // PV17输入电压 (V)
	Pv18U         float64 `json:"pv18_u"`         // PV18输入电压 (V)
	Pv19U         float64 `json:"pv19_u"`         // PV19输入电压 (V)
	Pv20U         float64 `json:"pv20_u"`         // PV20输入电压 (V)
	Pv21U         float64 `json:"pv21_u"`         // PV21输入电压 (V)
	Pv22U         float64 `json:"pv22_u"`         // PV22输入电压 (V)
	Pv23U         float64 `json:"pv23_u"`         // PV23输入电压 (V)
	Pv24U         float64 `json:"pv24_u"`         // PV24输入电压 (V)
	Pv25U         float64 `json:"pv25_u"`         // PV25输入电压 (V)
	Pv26U         float64 `json:"pv26_u"`         // PV26输入电压 (V)
	Pv27U         float64 `json:"pv27_u"`         // PV27输入电压 (V)
	Pv28U         float64 `json:"pv28_u"`         // PV28输入电压 (V)
	Pv1I          float64 `json:"pv1_i"`          // PV1输入电流 (A)
	Pv2I          float64 `json:"pv2_i"`          // PV2输入电流 (A)
	Pv3I          float64 `json:"pv3_i"`          // PV3输入电流 (A)
	Pv4I          float64 `json:"pv4_i"`          // PV4输入电流 (A)
	Pv5I          float64 `json:"pv5_i"`          // PV5输入电流 (A)
	Pv6I          float64 `json:"pv6_i"`          // PV6输入电流 (A)
	Pv7I          float64 `json:"pv7_i"`          // PV7输入电流 (A)
	Pv8I          float64 `json:"pv8_i"`          // PV8输入电流 (A)
	Pv9I          float64 `json:"pv9_i"`          // PV9输入电流 (A)
	Pv10I         float64 `json:"pv10_i"`         // PV10输入电流 (A)
	Pv11I         float64 `json:"pv11_i"`         // PV11输入电流 (A)
	Pv12I         float64 `json:"pv12_i"`         // PV12输入电流 (A)
	Pv13I         float64 `json:"pv13_i"`         // PV13输入电流 (A)
	Pv14I         float64 `json:"pv14_i"`         // PV14输入电流 (A)
	Pv15I         float64 `json:"pv15_i"`         // PV15输入电流 (A)
	Pv16I         float64 `json:"pv16_i"`         // PV16输入电流 (A)
	Pv17I         float64 `json:"pv17_i"`         // PV17输入电流 (A)
	Pv18I         float64 `json:"pv18_i"`         // PV18输入电流 (A)
	Pv19I         float64 `json:"pv19_i"`         // PV19输入电流 (A)
	Pv20I         float64 `json:"pv20_i"`         // PV20输入电流 (A)
	Pv21I         float64 `json:"pv21_i"`         // PV21输入电流 (A)
	Pv22I         float64 `json:"pv22_i"`         // PV22输入电流 (A)
	Pv23I         float64 `json:"pv23_i"`         // PV23输入电流 (A)
	Pv24I         float64 `json:"pv24_i"`         // PV24输入电流 (A)
	Pv25I         float64 `json:"pv25_i"`         // PV25输入电流 (A)
	Pv26I         float64 `json:"pv26_i"`         // PV26输入电流 (A)
	Pv27I         float64 `json:"pv27_i"`         // PV27输入电流 (A)
	Pv28I         float64 `json:"pv28_i"`         // PV28输入电流 (A)
	TotalCap      float64 `json:"total_cap"`      // 累计发电量 (kWh)
	OpenTime      float64 `json:"open_time"`      // 逆变器开机时间 (ms)
	CloseTime     float64 `json:"close_time"`     // 逆变器关机时间 (ms)
	MpptTotalCap  float64 `json:"mppt_total_cap"` // 直流输入总电量 (kWh)
	Mppt1Cap      float64 `json:"mppt_1_cap"`     // MPPT1直流累计发电量 (kWh)
	Mppt2Cap      float64 `json:"mppt_2_cap"`     // MPPT2直流累计发电量 (kWh)
	Mppt3Cap      float64 `json:"mppt_3_cap"`     // MPPT3直流累计发电量 (kWh)
	Mppt4Cap      float64 `json:"mppt_4_cap"`     // MPPT4直流累计发电量 (kWh)
	Mppt5Cap      float64 `json:"mppt_5_cap"`     // MPPT5直流累计发电量 (kWh)
	Mppt6Cap      float64 `json:"mppt_6_cap"`     // MPPT6直流累计发电量 (kWh)
	Mppt7Cap      float64 `json:"mppt_7_cap"`     // MPPT7直流累计发电量 (kWh)
	Mppt8Cap      float64 `json:"mppt_8_cap"`     // MPPT8直流累计发电量 (kWh)
	Mppt9Cap      float64 `json:"mppt_9_cap"`     // MPPT9直流累计发电量 (kWh)
	Mppt10Cap     float64 `json:"mppt_10_cap"`    // MPPT10直流累计发电量 (kWh)
	RunState      int64   `json:"run_state"`      // 状态(0：断连，1：连接) (无)
}

/*

 */

// 设备实时数据接口 返回结构体
type DevRealKpiDataResponse struct {
	DevId       int64           `json:"devId"`       //设备编号
	Sn          string          `json:"sn"`          //设备SN号
	DataItemMap json.RawMessage `json:"dataItemMap"` //每个数据项的内容，用key-value形式返回，不同设备类型的数据项内容不一样，数据项列表参见下方设备实时数据列表。
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
	DevId       int             `json:"devId"`
	Sn          string          `json:"sn"`
	CollectTime int64           `json:"collectTime"`
	DataItemMap json.RawMessage `json:"dataItemMap"`
}
