package qweather

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
	NOW = "/v7/weather/now"
)

func NowApi(location, host, key string) (WeatherNowResponse, error) {
	params := map[string]string{
		"location": location,
		"lang":     "zh-hans",
	}
	header := map[string]string{
		"Content-Type": "application/json",
		"X-QW-Api-Key": key,
	}
	url := strings.Join([]string{constant.SCHEME, host, NOW}, "")
	b, err := util.HttpGet(header, params, url)
	if err != nil {
		return WeatherNowResponse{}, err
	} else {
		log.Printf("请求直接返回的内容:%v", string(b))
	}
	// 先尝试解析是否为错误响应
	var errRes constant.ErrorResponse
	if e := json.Unmarshal(b, &errRes); e == nil && errRes.Error.Status != 0 {
		return WeatherNowResponse{}, fmt.Errorf("API错误: %s - %s", errRes.Error.Title, errRes.Error.Detail)
	}
	var res WeatherNowResponse
	err = json.Unmarshal(b, &res)
	if err != nil {
		return res, err
	}
	log.Printf("res:%v", res)
	return res, nil
}

// WeatherNowResponse 实时天气信息返回结构
type WeatherNowResponse struct {
	Code       string         `json:"code"`
	UpdateTime string         `json:"updateTime"`
	FxLink     string         `json:"fxLink"`
	Now        Now            `json:"now"`
	Refer      constant.Refer `json:"refer"`
}

// Now 实时天气详细信息
type Now struct {
	ObsTime   string `json:"obsTime"`   // 数据观测时间
	Temp      string `json:"temp"`      // 温度
	FeelsLike string `json:"feelsLike"` // 体感温度
	Icon      string `json:"icon"`      // 天气状况图标代码
	Text      string `json:"text"`      // 天气状况描述
	Wind360   string `json:"wind360"`   // 风向360角度
	WindDir   string `json:"windDir"`   // 风向
	WindScale string `json:"windScale"` // 风力等级
	WindSpeed string `json:"windSpeed"` // 风速
	Humidity  string `json:"humidity"`  // 相对湿度
	Precip    string `json:"precip"`    // 降水量
	Pressure  string `json:"pressure"`  // 大气压强
	Vis       string `json:"vis"`       // 能见度
	Cloud     string `json:"cloud"`     // 云量
	Dew       string `json:"dew"`       // 露点温度
}
