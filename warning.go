package function

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/zhangyiming748/qweather/constant"
	"github.com/zhangyiming748/qweather/util"
)

/*
城市搜索API提供全球地理位位置、全球城市搜索服务，支持经纬度坐标反查、多语言、模糊搜索等功能。
天气数据是基于地理位置的数据，因此获取天气之前需要先知道具体的位置信息。使用城市搜索，可获取到该城市的基本信息，包括城市的Location ID（你需要这个ID去查询天气），多语言名称、经纬度、时区、海拔、Rank值、归属上级行政区域、所在行政区域等。
另外，城市搜索也可以帮助你在你的APP中实现模糊搜索，用户只需要输入1-2个字即可获得结果。
*/
const (
	WARANING = "/v7/warning/now"
)

// WarningResponse 天气预警信息返回结构
type WarningResponse struct {
	Code       string    `json:"code"`
	UpdateTime string    `json:"updateTime"`
	FxLink     string    `json:"fxLink"`
	Warning    []Warning `json:"warning"`
	Refer      constant.Refer    `json:"refer"`
}

// Warning 预警详细信息
type Warning struct {
	ID            string `json:"id"`
	Sender        string `json:"sender"`
	PubTime       string `json:"pubTime"`
	Title         string `json:"title"`
	StartTime     string `json:"startTime"`
	EndTime       string `json:"endTime"`
	Status        string `json:"status"`
	Level         string `json:"level"`
	Severity      string `json:"severity"`
	SeverityColor string `json:"severityColor"`
	Type          string `json:"type"`
	TypeName      string `json:"typeName"`
	Urgency      string `json:"urgency"`
	Certainty    string `json:"certainty"`
	Text         string `json:"text"`
	Related      string `json:"related"`
}

func WarningApi(location, host, key string) (WarningResponse, error) {
	params := map[string]string{
		"location": location,
		"lang":     "zh-hans",
	}
	header := map[string]string{
		"Content-Type":  "application/json",
		"X-QW-Api-Key": key,
	}
	url := strings.Join([]string{constant.SCHEME, host, WARANING}, "")
	b, err := util.HttpGet(header, params, url)
	if err != nil {
		return WarningResponse{}, err
	} else {
		log.Printf("请求直接返回的内容:%v", string(b))
	}

	// 先尝试解析是否为错误响应
	var errRes constant.ErrorResponse
	if e := json.Unmarshal(b, &errRes); e == nil && errRes.Error.Status != 0 {
		return WarningResponse{}, fmt.Errorf("API错误: %s - %s", errRes.Error.Title, errRes.Error.Detail)
	}

	var res WarningResponse
	err = json.Unmarshal(b, &res)
	if err != nil {
		return res, err
	}
	log.Printf("res:%v", res)
	return res, nil
}
