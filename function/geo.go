package function

import (
	"encoding/json"
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
	GEO = "/geo/v2/city/lookup"
)

func GeoApi(name, host, api string) (GeoResponse, error) {
	params := map[string]string{
		"location": name,
	}
	header := map[string]string{
		"Content-Type": "application/json",
		"X-QW-Api-Key": api,
	}
	url := strings.Join([]string{constant.SCHEME, host, GEO}, "")
	b, err := util.HttpGet(header, params, url)
	if err != nil {
		return GeoResponse{}, err
	}
	var res GeoResponse
	err = json.Unmarshal(b, &res)
	if err != nil {
		return res, err
	}
	log.Printf("res:%v", res)
	return res, nil
}

// GeoResponse 地理位置信息返回结构
type GeoResponse struct {
	Code     string     `json:"code"`
	Location []Location `json:"location"`
	Refer    Refer      `json:"refer"`
}

// Location 位置详细信息
type Location struct {
	Name      string `json:"name"`
	ID        string `json:"id"`
	Lat       string `json:"lat"`
	Lon       string `json:"lon"`
	Adm2      string `json:"adm2"`
	Adm1      string `json:"adm1"`
	Country   string `json:"country"`
	Tz        string `json:"tz"`
	UtcOffset string `json:"utcOffset"`
	IsDst     string `json:"isDst"`
	Type      string `json:"type"`
	Rank      string `json:"rank"`
	FxLink    string `json:"fxLink"`
}

// Refer 引用信息
type Refer struct {
	Sources []string `json:"sources"`
	License []string `json:"license"`
}
