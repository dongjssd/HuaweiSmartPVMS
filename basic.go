/**
    @author: dongjs
    @date: 2025/1/15
    @description:基础类接口
**/

package HuaweiSmartPVMS

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// 接口描述
// 查询电站列表，传入分页参数（其中每页最多展示100条数据）和并网时间（如果只传入开始并网时间，
// 则结束并网时间默认为当前时间，如果只传入结束并网时间，则开始并网时间默认为时间戳起点），
// 则基于并网时间分页查询电站列表；仅传入分页参数，则直接分页查询电站列表。
func (c *Client) Stations(pageNo int, gridConnectedStartTime, gridConnectedEndTime int64) (*StationResponse, error) {
	request := map[string]interface{}{
		"pageNo": pageNo,
	}
	if gridConnectedStartTime > 0 {
		request["gridConnectedStartTime"] = gridConnectedStartTime
	}
	if gridConnectedEndTime > 0 {
		request["gridConnectedEndTime"] = gridConnectedEndTime
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("/thirdData/stations", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := StationResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 查询电站列表 请求参数
type stationRequest struct {
	PageNo                 int   `json:"pageNo"`
	GridConnectedStartTime int64 `json:"gridConnectedStartTime"`
	GridConnectedEndTime   int64 `json:"gridConnectedEndTime"`
}

// 查询电站列表 返回参数
type StationResponse struct {
	Success  bool                `json:"success"` //请求成功或失败标识 true：请求成功 false：请求失败
	FailCode int                 `json:"failCode"`
	Message  string              `json:"message"`
	Data     StationDataResponse `json:"data"`
}

type StationDataResponse struct {
	Total     int64                     `json:"total"`     //总条数
	PageCount int64                     `json:"pageCount"` //总页数
	PageNo    int                       `json:"pageNo"`    //分页查询，第n页
	PageSize  int                       `json:"pageSize"`  //分页查询，每页条数
	List      []StationDataListResponse `json:"list"`      //电站相关数据对象列表
}

type StationDataListResponse struct {
	PlantCode          string  `json:"plantCode"`          //电站编号，系统中电站唯一标识
	PlantName          string  `json:"plantName"`          //名称
	PlantAddress       string  `json:"plantAddress"`       //电站详细地址
	Longitude          string  `json:"longitude"`          //电站经度
	Latitude           string  `json:"latitude"`           //电站纬度
	Capacity           float64 `json:"capacity"`           //组串总容量 kWp
	ContactPerson      string  `json:"contactPerson"`      //电站联系人
	ContactMethod      string  `json:"contactMethod"`      //电站联系人联系方式，手机或邮箱
	GridConnectionDate string  `json:"gridConnectionDate"` //电站的并网时间，包含时区 2020-02-06T00:00:00+08:00
}

// 设备列表接口
// 获取设备基本信息，在调用其余接口获取设备数据前需先调用此接口获取设备编号。
// 通过电站编号集合查询，一次最多支持查询100个电站。
// 电站编号列表，多个电站用英文逗号分隔，电站编号由电站列表接口中plantCode获取。
func (c *Client) DevList(stationCodes string) (*DevListResponse, error) {
	request := map[string]interface{}{
		"stationCodes": stationCodes,
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("/thirdData/getDevList", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := DevListResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 设备列表接口 结果结构提
type DevListResponse struct {
	Success  bool                  `json:"success"` //请求成功或失败标识 true：请求成功 false：请求失败
	FailCode int                   `json:"failCode"`
	Message  string                `json:"message"`
	Data     []DevListDataResponse `json:"data"`
}

type DevListDataResponse struct {
	Id              int64  `json:"id"`              //设备ID（设备编号）
	DevDn           string `json:"devDn"`           //设备唯一ID（系统中设备唯一编号）
	DevName         string `json:"devName"`         //设备名称
	StationCode     string `json:"stationCode"`     //电站编号
	EsnCode         string `json:"esnCode"`         //设备SN号
	DevTypeId       int    `json:"devTypeId"`       //设备类型ID
	SoftwareVersion string `json:"softwareVersion"` //软件版本号
	OptimizerNumber int    `json:"optimizerNumber"` //优化器数量
	InvType         string `json:"invType"`         //机型（只有逆变器有机型）
	Longitude       string `json:"longitude"`       //经度
	Latitude        string `json:"latitude"`        //纬度
}
