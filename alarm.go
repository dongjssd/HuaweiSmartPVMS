/**
    @author: dongjs
    @date: 2025/1/15
    @description: 告警
**/

package HuaweiSmartPVMS

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// 查询活动告警接口
// 查询设备当前（活动）的告警信息，如果以电站为参数查询，一次支持查询100个电站；
// 如果以设备SN为参数查询，一次支持查询100个设备。
// 如果传入的电站编号列表不为空，则基于电站编号列表查询设备告警信息，
// 如果电站编号列表为空且设备SN列表不为空，则基于设备SN列表查询设备告警信息。
func (c *Client) AlarmList(stationCodes, sns string, beginTime, endTime int64,
	language string, levels string, devTypes string) (*AlarmListResponse, error) {
	request := map[string]interface{}{
		"beginTime": beginTime,
		"endTime":   endTime,
		"language":  language,
		"levels":    levels,
		"devTypes":  devTypes,
	}
	if stationCodes != "" {
		request["stationCodes"] = stationCodes
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
	body, err := c.doRequest("/thirdData/getAlarmList", buf)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("body:%+v", string(body))
	response := AlarmListResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 查询活动告警接口 返回对象
type AlarmListResponse struct {
	Success  bool                    `json:"success"` //请求成功或失败标识 true：请求成功 false：请求失败
	FailCode int                     `json:"failCode"`
	Message  string                  `json:"message"`
	Params   AlarmListParamsResponse `json:"params"`
	Data     []AlarmListDataResponse `json:"data"`
}

// 查询活动告警接口 返回对象
type AlarmListParamsResponse struct {
	StationCodes string `json:"stationCodes"` //请求参数中的电站编号列表
	Sns          string `json:"sns"`          //请求参数中的设备SN列表
	BeginTime    int64  `json:"beginTime"`    //请求参数中的查询活动告警的开始时间戳（毫秒）
	EndTime      int64  `json:"endTime"`      //请求参数中的查询活动告警的结束时间戳（毫秒）
	Language     string `json:"language"`     //请求参数中的语言
	Levels       string `json:"levels"`       //请求参数中的告警级别
	DevTypes     string `json:"devTypes"`     //请求参数中的设备类型
	CurrentTime  int64  `json:"currentTime"`  //系统当前时间戳（毫秒）
}

// 查询活动告警接口 返回对象
type AlarmListDataResponse struct {
	StationCode      string `json:"stationCode"`      //电站编号，电站唯一标识
	AlarmName        string `json:"alarmName"`        //告警名称
	DevName          string `json:"devName"`          //设备名称
	RepairSuggestion string `json:"repairSuggestion"` //修复建议
	EsnCode          string `json:"esnCode"`          //设备sn
	DevTypeId        int    `json:"devTypeId"`        //设备类型ID
	CauseId          int    `json:"causeId"`          //原因ID
	AlarmCause       string `json:"alarmCause"`       //告警原因
	AlarmType        int    `json:"alarmType"`        //告警类型 0：其它告警 1：变位信号2：异常告警3：保护事件4：通知状态5：告警信息
	RaiseTime        int64  `json:"raiseTime"`        //告警产生时间戳（毫秒）
	AlarmId          int    `json:"alarmId"`          //告警ID
	StationName      string `json:"stationName"`      //电站名称
	Lev              int    `json:"lev"`              //告警级别 //1：严重 2：重要3：次要4：提示
	Status           int    `json:"status"`           //告警状态 //1：未处理（活动）
}
