# simpleConf
Simple configuration file reading <br>
##1.config.conf 描述<br>

浅蓝色文字：<font color="#0000dd">浅蓝色文字</font><br /> 


<b>除非是纯数字,否则一律按字符串处理</b><br>
`#`表示注释<br>
`##` 表示第一级 key<br>
key 值建议首字母小写,`key-value`必须一一对应<br>
key 值必须与 struct{} 的 `json tag` 一致

##2.使用
```go
# server config
## server
port = 8010
log_mode = true

# mysql config
## mysql
host = localhost
port = 3306
user = root
pw   = 123456qw

# image config
## image
```
对应着下面这段 json
```json

{
    "server":{
        "port":8010,
        "log_mode":true
    },
    "mysql":{
        "host":"localhost",
        "port":3306,
        "user":"root",
        "pw":"123456pw"
    },
    "image":{

    }
}
```
对应 GoLang struct{} 为
```go
// 配置文件结构体
type Table struct {
	Server			Server			`json:"server"`
	Database		Database		`json:"mysql"`
	Image			Image			`json:"image"`
}

// 服务
type Server struct {
	Port			int				`json:"port"`
	LogModel		bool			`json:"log_model"`
}

// database
type Database struct{
	Host			string			`json:"host"`
	Port			int				`json:"port"`
	User			string			`json:"user"`
	Pw		        string			`json:"pw"`
}

type Image struct {
	
}

```