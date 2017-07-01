# Introduction
A tool provides health check for HTTP address in timing.

# Configure

```conf
appname = url-checker
httpport = 8080
runmode = dev

# 请求参数
copyrequestbody = true

# database config
db.host = 127.0.0.1
db.user = root
db.pass = "" 
db.port = 3306
db.name = url-checker
db.timezone = Asia/Shanghai

# 检测参数
check.frequency = 5

# influxdb config
indb.host = "http://127.0.0.1:8086"
indb.user = ''
indb.pass = ''
indb.db   = 'test'

# 数据存储

exit.where = "influxdb"
```

# Build
```shell
go build  or  gox -osarch="linux/amd64"
```

# Start
```shell
./url-checker  &
```
# Get Items
### get all items
##### request: `/api/item/list`
##### respone: `json data`
```
{
code: 0,
message: "",
data: [
{
Id: 163,
InstanceName: "server-test",
Item: "http://10.11.12.14:10000/health",
UrlType: "buiness",
Timeout: 10,
Keyword: "",
Maintainer: "xxx@gmail.com"
},
{
Id: 162,
InstanceName: "server-test",
Item: "http://10.11.12.13:10000/health",
UrlType: "buiness",
Timeout: 10,
Keyword: "",
Maintainer: "xxx@gmail.com"
}]
}
```
### get specfiy item
##### request: `/api/item/list/9`
##### respone: `json data`
```
{
code: 0,
message: "",
data: {
Id: 9,
InstanceName: "server-online",
Item: "http://10.11.12.12:10013/buiness/health",
UrlType: "buiness",
Timeout: 10,
Keyword: "",
Maintainer: "xxx@gmail.com"
}
}
```

# Query History Resptime From Influxdb
### examples:
![resp1](https://github.com/yujianglei/url-checker/blob/master/snapshots/resp1.jpeg)
![resp1](https://github.com/yujianglei/url-checker/blob/master/snapshots/resp2.jpeg)

# To Do Lists
* add alarm 

# You Can See 
* [XiaoMi](https://github.com/XiaoMi)
* [URLooker](https://github.com/URLooker/)
* [OpenFalcon](https://github.com/XiaoMi/open-falcon)

