## GET 天气灾害预警

GET /v7/warning/now

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|location|query|string| 否 |需要查询地区的LocationID或以英文逗号分隔的经度,纬度坐标（十进制，最多支持小数点后两位），LocationID可通过GeoAPI获取。例如 location=101010100 或 location=116.41,39.92|
|lang|query|string| 否 |多语言设置，请阅读多语言文档，了解我们的多语言是如何工作、如何设置以及数据是否支持多语言。|
|key|query|string| 否 |none|
|X-QW-Api-Key|header|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "code": "200",
  "updateTime": "2025-04-13T15:04+08:00",
  "fxLink": "https://www.qweather.com/severe-weather/shijingshan-101011000.html",
  "warning": [
    {
      "id": "10101100020250412211000454915870",
      "sender": "石景山区气象台",
      "pubTime": "2025-04-12T21:10+08:00",
      "title": "石景山区气象台继续发布大风橙色预警",
      "startTime": "2025-04-12T21:10+08:00",
      "endTime": "2025-04-13T21:00+08:00",
      "status": "active",
      "level": "橙色",
      "severity": "Severe",
      "severityColor": "Orange",
      "type": "1006",
      "typeName": "大风",
      "urgency": "",
      "certainty": "",
      "text": "石景山区气象台12日21时10分继续发布大风橙色预警,石景山区气象台发布分区域大风橙色预警信号，我区大风持续，预计当前至13日20时石景山区五里坨街道、苹果园街道、金顶街街道、广宁街道，古城街道（首钢园区域）将出现9～11级阵风。其他区域阵风8～9级，降级发布大风黄色预警信号。我区仍然处于大风较强时段，严禁一切室外用火行为，大风橙色预警区域非必要不出行，大风黄色预警区域老、弱、病、幼减少户外活动，请继续做好防范。",
      "related": ""
    },
    {
      "id": "10101100020250410193940080456885",
      "sender": "北京应急局（森防办）",
      "pubTime": "2025-04-10T19:39+08:00",
      "title": "北京市森防办2025年04月10日19时12分发布森林火险橙色预警",
      "startTime": "2025-04-10T19:12+08:00",
      "endTime": "2025-04-14T00:00+08:00",
      "status": "active",
      "level": "橙色",
      "severity": "Severe",
      "severityColor": "Orange",
      "type": "1025",
      "typeName": "森林火险",
      "urgency": "",
      "certainty": "",
      "text": "北京市森防办发布森林火险橙色预警：预计4月11日至13日本市将出现大风天气，气象条件不利，林区植被干燥，森林火险等级高，森林防火形势严峻。严禁携带火种进山入林，严禁野外用火。一旦发现森林火情，请及时拨打12119或119报警。",
      "related": ""
    }
  ],
  "refer": {
    "sources": [
      "国家预警信息发布中心 (12379)"
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
|» updateTime|string|true|none||当前API的最近更新时间|
|» fxLink|string|true|none||当前数据的响应式页面，便于嵌入网站或应用|
|» warning|[object]|true|none||none|
|»» id|string|true|none||本条预警的唯一标识，可判断本条预警是否已经存在|
|»» sender|string|false|none||预警发布单位，可能为空|
|»» pubTime|string|true|none||预警发布时间|
|»» title|string|true|none||预警信息标题|
|»» startTime|string|false|none||预警开始时间，可能为空|
|»» endTime|string|false|none||预警结束时间，可能为空|
|»» status|string|true|none||预警信息的发布状态|
|»» level|string|false|none||预警等级（已弃用），不要再使用这个字段，该字段已弃用，目前返回为空或未更新的值。请使用severity和severityColor代替|
|»» severity|string|true|none||预警严重等级|
|»» severityColor|string|false|none||预警严重等级颜色，可能为空|
|»» type|string|true|none||预警类型ID|
|»» typeName|string|true|none||预警类型名称|
|»» urgency|string|false|none||预警信息的紧迫程度，可能为空|
|»» certainty|string|false|none||预警信息的确定性，可能为空|
|»» text|string|true|none||预警详细文字描述|
|»» related|string|false|none||与本条预警相关联的预警ID，当预警状态为cancel或update时返回。可能为空|
|» refer|object|true|none||none|
|»» sources|[string]|true|none||none|
|»» license|[string]|false|none||none|

# 数据模型

