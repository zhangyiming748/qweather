# QWeather Go SDK

这是一个用于访问和使用和风天气 API 的 Go 语言 SDK，提供了简单易用的接口来获取天气信息。

## 安装

使用 Go modules 安装：

```bash
go get github.com/zhangyiming748/qweather
```

## 功能特性

- 实时天气查询
- 地理位置信息查询
- 支持多语言（默认中文）
- 错误处理机制

## 使用示例

### 初始化

```go
package main

import "github.com/zhangyiming748/qweather"

const (
    host = "api.qweather.com"  // API 域名
    key  = "your-api-key"      // 你的 API 密钥
)
```

### 查询地理位置信息

```go
// 查询城市信息
response, err := function.GeoApi("北京", host, key)
if err != nil {
    log.Fatal(err)
}
// 使用返回的地理位置信息
```

### 查询实时天气

```go
// 使用位置 ID 查询实时天气
weather, err := function.NowApi(locationId, host, key)
if err != nil {
    log.Fatal(err)
}
// 处理天气信息
```

## 注意事项

- 使用前需要先在和风天气官网注册并获取 API 密钥
- API 调用需要遵循和风天气的使用限制和规范
- 建议在生产环境中做好错误处理和重试机制

## 许可证

本项目采用 MIT 许可证，详见 LICENSE 文件。