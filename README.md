# IP地理位置信息查询服务


* 定位信息来源：[ip2region](https://github.com/lionsoul2014/ip2region)
* 地理位置来源：国内地理位置收集自**百度地图API**，国外地理位置收集自**谷歌地图API**



## Run

### Install Kratos
```
go get -u github.com/go-kratos/kratos/cmd/kratos/v2@latest
```

### Run service
```
kratos run
```



## Example

```json
GET localhost:18080/geoip?ip=123.123.123.123

{
	"city": "北京",
	"nation": "中国",
	"province": "北京",
	"location": {
		"latitude": 39.910924547299565,
		"longitude": 116.4133836971231
	},
	"isp": "联通"
}
```