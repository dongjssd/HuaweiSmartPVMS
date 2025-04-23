/**
    @author: dongjs
    @date: 2025/1/16
    @description: 控制类接口
**/

package HuaweiSmartPVMS

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// 储能工作模式设置任务下发接口
// 支持电站级储能工作模式设置，根据电站DN进行储能工作模式设置任务下发，一次任务最多支持10个电站，支持对每个电站下发不同的参数。
// 支持电站单控制器(Dongle、EMMA、分布式数采、逆变器直连)组网场景，储能工作模式支持"最大自发自用"和"TOU"两种模式。
func (c *Client) BatteryModeAsyncTask(request BatteryModeAsyncTaskRequest) (*BatteryModeAsyncTaskResponse, error) {
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("/rest/openapi/pvms/nbi/v1/control/battery/mode/async-task", buf)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("body:%+v", string(body))
	response := BatteryModeAsyncTaskResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 储能工作模式设置任务下发接口 请求参数
type BatteryModeAsyncTaskRequest struct {
	Tasks []BatteryModeAsyncTaskTaskRequest `json:"tasks"`
}

// 储能工作模式设置任务下发接口 请求参数
type BatteryModeAsyncTaskTaskRequest struct {
	PlantCode                        string                                  `json:"plantCode"`                        //电站DN
	OperationMode                    string                                  `json:"operationMode"`                    //工作模式 TOU：TOU模式maximumSelfConsumption：最大自发自用模式设置为“TOU模式”时，可以手动设置充放电时间段。适用于光储系统及纯储系统，且存在电峰谷价差和有电表场景。设置“最大自发自用模式”时，光伏发电优先供负载，将多余的发电功率给储能充电。若储能已充满或满功率在充电，将多余发电量输出到电网。在发电不足或夜间无光伏发电时，储能放电供负载用电，提高系统的自发自用率和能源自给自足率，节省电费支出。电网不能给储能充电，但可给负载供电。适用于光储系统，光伏配比高，光伏发电足够负载使用，多余的发电功率给储能充电（若光伏电量小于负载功率，建议使用TOU），且电价高，上网电价补贴低或无上网电价补贴的场景。
	RedundantPVEnergyPriority        string                                  `json:"redundantPVEnergyPriority"`        //多余PV能量优先级 fedToGridPreference：上网优先chargePreference：充电优先设置为“上网优先”时，是指光伏发电功率大于负载时，多余的光伏能量优先上网，设备输出功率达到最大值后，多余能量给储能充电。该设置一般适用于FIT电价高于用电电价场景，电网不能给储能充电。设置为“充电优先”时，是指光伏发电比负载多时，多余的光伏能量给储能充电，充电功率达到最大或储能充满后，多余光伏能量上网。
	AllowedAcChargePower             float64                                 `json:"allowedAcChargePower"`             //电网充电最大功率，单位kW 该参数的范围和电站中的控制器相关，具体范围如下： Dongle：[0.000, 30.000]EMMA：[0.000, 50.000]数采：[0.000, 50000.000]逆变器：[0.000, 电网充电最大功率上限]电网充电最大功率用于设置电网反向给储能设备的充电最大功率。
	ChargingAndDischargingTimeWindow []BatteryModeAsyncTaskTaskWindowRequest `json:"chargingAndDischargingTimeWindow"` //充放电时间窗口设置充电/放电开始和结束时间。最多可以设置14个时间窗。 对于同一个电站，同一天不能设置多个重叠的时间段。
}

// 储能工作模式设置任务下发接口 请求参数
type BatteryModeAsyncTaskTaskWindowRequest struct {
	StartTime         string `json:"startTime"`         //开始时间，格式：HH:MM范围：00:00~23:59
	EndTime           string `json:"endTime"`           //结束时间，格式：HH:MM 范围：00:00~23:59
	ChargeOrDischarge string `json:"chargeOrDischarge"` //充电/放电 charge：充电discharge：放电
	Repeat            []int  `json:"repeat"`            //重复日期，每个电站最多设置7个日期 1：周一 2：周二3：周三4：周四5：周五6：周六7：周日
}

// 储能工作模式设置任务下发接口 返回参数
type BatteryModeAsyncTaskResponse struct {
	Success  bool                             `json:"success"` //请求成功或失败标识 true：请求成功 false：请求失败
	FailCode int                              `json:"failCode"`
	Message  string                           `json:"message"`
	Data     BatteryModeAsyncTaskDataResponse `json:"data"`
}

// 储能工作模式设置任务下发接口 返回参数
type BatteryModeAsyncTaskDataResponse struct {
	TaskId string                                   `json:"taskId"` //请求的任务唯一ID，可用于查询任务下发结果
	Result []BatteryModeAsyncTaskDataResultResponse `json:"result"` //任务下发结果
}

// 储能工作模式设置任务下发接口 返回参数
type BatteryModeAsyncTaskDataResultResponse struct {
	PlantCode string `json:"plantCode"` //电站DN
	Status    string `json:"status"`    //电站下发任务的当前状态 RUNNING：电站下发任务运行中FAIL：电站下发任务失败
	Message   string `json:"message"`   //下发结果描述
}

// 储能工作模式设置任务查询接口
// 根据taskId查询储能工作模式设置任务执行情况，一次查询支持一个任务信息。
func (c *Client) BatteryModeTaskInfo(taskId string) (*BatteryModeTaskInfoResponse, error) {
	request := map[string]interface{}{
		"taskId": taskId,
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("/rest/openapi/pvms/nbi/v1/control/battery/mode/task-info", buf)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("body:%+v", string(body))
	response := BatteryModeTaskInfoResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 储能工作模式设置任务查询接口 返回参数
type BatteryModeTaskInfoResponse struct {
	Success  bool                            `json:"success"` //请求成功或失败标识 true：请求成功 false：请求失败
	FailCode int                             `json:"failCode"`
	Message  string                          `json:"message"`
	Data     BatteryModeTaskInfoDataResponse `json:"data"`
}

// 储能工作模式设置任务查询接口 返回参数
type BatteryModeTaskInfoDataResponse struct {
	DispatchResult []BatteryModeTaskInfoDataResultResponse `json:"dispatchResult"` //dispatchResult里面是请求执行的返回信息List，包含如下信息：
	StartTime      string                                  `json:"startTime"`      //收到任务的时间，包含时区
	EndTime        string                                  `json:"endTime"`        //任务完成的时间（任务未完成时返回null），包含时区
}

// 储能工作模式设置任务查询接口 返回参数
type BatteryModeTaskInfoDataResultResponse struct {
	PlantCode                        string                                        `json:"plantCode"`                        //电站DN
	Status                           string                                        `json:"status"`                           //电站下发任务的当前状态。 RUNNING：电站下发任务运行中 SUCCESS：电站下发任务成功FAIL：电站下发任务失败
	Message                          string                                        `json:"message"`                          //请求描述。当status为FAIL时，返回电站下发失败的描述，其他情况不返回。 FAILURE：下发失败 TIMEOUT：超时BUSY：设备忙INVALID：非法设备EXCEPTION：异常
	OperationMode                    string                                        `json:"operationMode"`                    //工作模式 TOU：TOU模式 maximumSelfConsumption：最大自发自用模式
	RedundantPVEnergyPriority        string                                        `json:"redundantPVEnergyPriority"`        //多余PV能量优先级 fedToGridPreference：上网优先 chargePreference：充电优先
	AllowedAcChargePower             float64                                       `json:"allowedAcChargePower"`             //电网充电最大功率，单位kW
	ChargingAndDischargingTimeWindow []BatteryModeTaskInfoDataResultWindowResponse `json:"chargingAndDischargingTimeWindow"` //充放电时间窗口
}

// 储能工作模式设置任务查询接口 返回参数
type BatteryModeTaskInfoDataResultWindowResponse struct {
	StartTime         string `json:"startTime"`         //开始时间 格式：HH:MM
	EndTime           string `json:"endTime"`           //结束时间 格式：HH:MM
	ChargeOrDischarge string `json:"chargeOrDischarge"` //充电/放电 charge：充电 discharge：放电
	Repeat            []int  `json:"repeat"`            //重复日期 1：周一 2：周二3：周三4：周四5：周五6：周六7：周日
}

// 储能参数设置任务下发接口
// 支持电站级储能参数（充电截止SOC、放电截止SOC、最大充电功率、最大放电功率）设置，
// 根据电站DN进行储能参数设置任务下发，一次任务最多支持10个电站，支持对每个电站下发不同的参数。
func (c *Client) BatteryConfigurationAsyncTask(request BatteryConfigurationAsyncTaskRequest) (*BatteryConfigurationAsyncTaskResponse, error) {
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("/rest/openapi/pvms/nbi/v1/control/battery/configuration/async-task", buf)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("body:%+v", string(body))
	response := BatteryConfigurationAsyncTaskResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 储能参数设置任务下发接口 请求参数
type BatteryConfigurationAsyncTaskRequest struct {
	Tasks []BatteryConfigurationAsyncTaskTaskRequest `json:"tasks"` //储能参数设置任务列表，一次任务下发最多10电站
}

// 储能参数设置任务下发接口 请求参数
type BatteryConfigurationAsyncTaskTaskRequest struct {
	PlantCode                string                                       `json:"plantCode"`                //电站DN
	BatteryConfigurationInfo BatteryConfigurationAsyncTaskTaskInfoRequest `json:"batteryConfigurationInfo"` //储能参数设置有效参数
}

// 储能参数设置任务下发接口 请求参数
type BatteryConfigurationAsyncTaskTaskInfoRequest struct {
	EndOfChargeSoc        float64 `json:"endOfChargeSoc"`        //充电截止SOC，单位% 范围：[90.0, 100.0]
	EndOfDischargeSoc     float64 `json:"endOfDischargeSoc"`     //放电截止SOC，单位% 范围：[0.0, 20.0]
	MaximumChargePower    int     `json:"maximumChargePower"`    //最大充电功率，单位W 范围：[0, 最大充电功率上限]如果下发的最大充电功率大于最大充电功率上限，则默认取上限值。
	MaximumDischargePower int     `json:"maximumDischargePower"` //最大放电功率，单位W 范围：[0, 最大放电功率上限]如果下发的最大放电功率大于最大放电功率上限，则默认取上限值。
}

// 储能参数设置任务下发接口 返回参数
type BatteryConfigurationAsyncTaskResponse struct {
	Success  bool                                      `json:"success"` //请求成功或失败标识 true：请求成功 false：请求失败
	FailCode int                                       `json:"failCode"`
	Message  string                                    `json:"message"`
	Data     BatteryConfigurationAsyncTaskDataResponse `json:"data"`
}

// 储能参数设置任务下发接口 返回参数
type BatteryConfigurationAsyncTaskDataResponse struct {
	TaskId string                                            `json:"taskId"`
	Result []BatteryConfigurationAsyncTaskDataResultResponse `json:"result"`
}

// 储能参数设置任务下发接口 返回参数
type BatteryConfigurationAsyncTaskDataResultResponse struct {
	PlantCode string `json:"plantCode"`
	Status    string `json:"status"`
	Message   string `json:"message"`
}

// 储能参数设置任务查询接口
// 根据taskId查询储能参数设置任务执行情况，一次查询支持一个任务信息。
func (c *Client) BatteryConfigurationTaskInfo(taskId string) (*BatteryConfigurationTaskInfoResponse, error) {
	request := map[string]interface{}{
		"taskId": taskId,
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("/rest/openapi/pvms/nbi/v1/control/battery/configuration/task-info", buf)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("body:%+v", string(body))
	response := BatteryConfigurationTaskInfoResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 储能参数设置任务查询接口 返回参数
type BatteryConfigurationTaskInfoResponse struct {
	Success  bool                                     `json:"success"` //请求成功或失败标识 true：请求成功 false：请求失败
	FailCode int                                      `json:"failCode"`
	Message  string                                   `json:"message"`
	Data     BatteryConfigurationTaskInfoDataResponse `json:"data"`
}

// 储能参数设置任务查询接口 返回参数
type BatteryConfigurationTaskInfoDataResponse struct {
	DispatchResult []BatteryConfigurationTaskInfoDataResultResponse `json:"dispatchResult"`
	StartTime      string                                           `json:"start_time"`
	EndTime        string                                           `json:"end_time"`
}

// 储能参数设置任务查询接口 返回参数
type BatteryConfigurationTaskInfoDataResultResponse struct {
	PlantCode                string `json:"plant_code"`
	Status                   string `json:"status"`
	Message                  string `json:"message"`
	BatteryConfigurationInfo struct {
		EndOfChargeSoc        float64 `json:"end_of_charge_soc"`
		EndOfDischargeSoc     float64 `json:"end_of_discharge_soc"`
		MaximumChargePower    int     `json:"maximum_charge_power"`
		MaximumDischargePower int     `json:"maximum_discharge_power"`
	} `json:"battery_configuration_info"`
}

// 逆变器有功功率设置任务下发接口
// 支持电站级有功功率设置，根据电站DN进行逆变器有功功率设置任务下发，一次任务最多支持10个电站，支持对每个电站下发不同的参数。
// 支持电站单控制器(Dongle、EMMA、数采、逆变器直连)组网场景，有功功率控制方式支持"无限制"和"限功率并网(kW)"两种模式。
func (c *Client) ActivePowerControlAsyncTask(request ActivePowerControlAsyncTaskRequest) (*ActivePowerControlAsyncTaskResponse, error) {
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("/rest/openapi/pvms/nbi/v2/control/active-power-control/async-task", buf)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("body:%+v", string(body))
	response := ActivePowerControlAsyncTaskResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 逆变器有功功率设置任务下发接口 请求参数
type ActivePowerControlAsyncTaskRequest struct {
	Tasks []ActivePowerControlAsyncTaskTaskRequest `json:"tasks"` //储能工作模式设置任务列表，一次任务下发最多10电站
}

// 逆变器有功功率设置任务下发接口 请求参数
type ActivePowerControlAsyncTaskTaskRequest struct {
	PlantCode   string `json:"plantCode"`   //电站DN
	ControlMode string `json:"controlMode"` //有功功率控制方式 0：无限制。6：限功率并网(kW)设置为“无限制”时，逆变器输出功率不受限制，逆变器能够以额定功率输出并网。
	ControlInfo struct {
		MaxGridFeedInPower float64 `json:"maxGridFeedInPower"` //最大馈送电网功率，单位kW 该参数的范围和电站中的控制器相关，具体范围如下： Dongle：[-1000.000, 5000.000]EMMA：[-1000.000, 逆变器额定功率]数采：[-1000.000, 5000.000]分布式数采：[-1000.000, 50000.000]逆变器：[-1000.000, 5000.000]设置并网点输送到电网的最大有功功率。
		LimitationMode     string  `json:"limitationMode"`     //限制方式 0：总功率 1：单相功率设置为“总功率”时，控制并网点总功率不逆流。设置为“单相功率”时，控制并网点各相功率均不逆流。
	} `json:"controlInfo"` //有功功率设置参数 当controlMode="6"时，可以通过该参数进行有功功率设置。 当controlMode="0"时，忽略该参数。
}

// 逆变器有功功率设置任务下发接口 返回参数
type ActivePowerControlAsyncTaskResponse struct {
	Success  bool                                    `json:"success"` //请求成功或失败标识 true：请求成功 false：请求失败
	FailCode int                                     `json:"failCode"`
	Message  string                                  `json:"message"`
	Data     ActivePowerControlAsyncTaskDataResponse `json:"data"`
}

// 逆变器有功功率设置任务下发接口 返回参数
type ActivePowerControlAsyncTaskDataResponse struct {
	TaskId string                                          `json:"taskId"` //请求的任务唯一ID，可用于查询任务下发结果
	Result []ActivePowerControlAsyncTaskDataResultResponse `json:"result"` //任务下发结果
}

// 逆变器有功功率设置任务下发接口 返回参数
type ActivePowerControlAsyncTaskDataResultResponse struct {
	PlantCode string `json:"plantCode"` //电站DN
	Status    string `json:"status"`    //电站下发任务的当前状态 RUNNING：电站下发任务运行中 FAIL：电站下发任务失败
	Message   string `json:"message"`   //下发结果描述
}

// 逆变器有功功率设置任务查询接口
// 根据taskId查询逆变器有功功率设置任务执行情况，一次查询支持一个任务信息。
func (c *Client) ActivePowerControlTaskInfo(taskId string) (*ActivePowerControlTaskInfoResponse, error) {
	request := map[string]interface{}{
		"taskId": taskId,
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("/rest/openapi/pvms/nbi/v2/control/active-power-control/task-info", buf)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("body:%+v", string(body))
	response := ActivePowerControlTaskInfoResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 逆变器有功功率设置任务查询接口 返回参数
type ActivePowerControlTaskInfoResponse struct {
	Success  bool                                   `json:"success"` //请求成功或失败标识 true：请求成功 false：请求失败
	FailCode int                                    `json:"failCode"`
	Message  string                                 `json:"message"`
	Data     ActivePowerControlTaskInfoDataResponse `json:"data"`
}

// 逆变器有功功率设置任务查询接口 返回参数
type ActivePowerControlTaskInfoDataResponse struct {
	DispatchResult []ActivePowerControlTaskInfoDataResultResponse `json:"dispatchResult"`
	StartTime      string                                         `json:"startTime"`
	EndTime        string                                         `json:"endTime"`
}

// 逆变器有功功率设置任务查询接口 返回参数
type ActivePowerControlTaskInfoDataResultResponse struct {
	PlantCode   string `json:"plantCode"`
	Status      string `json:"status"`
	Message     string `json:"message"`
	ControlMode string `json:"controlMode"`
	ControlInfo struct {
		MaxGridFeedInPower float64 `json:"maxGridFeedInPower"`
		LimitationMode     string  `json:"limitationMode"`
	} `json:"controlInfo"`
}
