## GET 查询实时天气

GET /v7/weather/now

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|location|query|string| 是 |需要查询地区的LocationID或以英文逗号分隔的经度,纬度坐标（十进制，最多支持小数点后两位），LocationID可通过GeoAPI获取。例如 location=101010100 或 location=116.41,39.92|
|lang|query|string| 否 |多语言设置，请阅读多语言文档，了解我们的多语言是如何工作、如何设置以及数据是否支持多语言。数据单位设置，可选值包括unit=m（公制单位，默认）和unit=i（英制单位）。更多选项和说明参考度量衡单位。|
|key|query|string| 否 |你自己的apikey|
|X-QW-Api-Key|header|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "code": "200",
  "updateTime": "2025-04-13T14:43+08:00",
  "fxLink": "https://www.qweather.com/weather/shijingshan-101011000.html",
  "now": {
    "obsTime": "2025-04-13T14:40+08:00",
    "temp": "13",
    "feelsLike": "6",
    "icon": "104",
    "text": "阴",
    "wind360": "0",
    "windDir": "北风",
    "windScale": "5",
    "windSpeed": "31",
    "humidity": "19",
    "precip": "0.0",
    "pressure": "994",
    "vis": "22",
    "cloud": "91",
    "dew": "-12"
  },
  "refer": {
    "sources": [
      "QWeather"
    ],
    "license": [
      "QWeather Developers License"
    ]
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|string|true|none||请参考状态码|
|» updateTime|string|true|none||API的最近更新时间|
|» fxLink|string|true|none||当前数据的响应式页面，便于嵌入网站或应用|
|» now|object|true|none||none|
|»» obsTime|string|true|none||数据观测时间|
|»» temp|string|true|none||温度，默认单位：摄氏度|
|»» feelsLike|string|true|none||体感温度，默认单位：摄氏度|
|»» icon|string|true|none||天气状况的图标代码，另请参考天气图标项目|
|»» text|string|true|none||天气状况的文字描述，包括阴晴雨雪等天气状态的描述|
|»» wind360|string|true|none||风向360角度|
|»» windDir|string|true|none||风向|
|»» windScale|string|true|none||风力等级|
|»» windSpeed|string|true|none||风速，公里/小时|
|»» humidity|string|true|none||相对湿度，百分比数值|
|»» precip|string|true|none||过去1小时降水量，默认单位：毫米|
|»» pressure|string|true|none||大气压强，默认单位：百帕|
|»» vis|string|true|none||能见度，默认单位：公里|
|»» cloud|string|true|none||云量，百分比数值。可能为空|
|»» dew|string|true|none||露点温度。可能为空|
|» refer|object|true|none||none|
|»» sources|[string]|true|none||原始数据来源，或数据源说明，可能为空|
|»» license|[string]|true|none||数据许可或版权声明，可能为空|

# 数据模型

