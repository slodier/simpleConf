package conf

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	Tab = NewTab()
)

// 读取配置文件
func Init() {
	// read config.conf
	f, err := os.Open("conf/config.conf")
	if err != nil {
		log.Fatal("open config.conf err:", err.Error())
		return
	}

	format(bufio.NewReader(f))
}

// 转换为结构体
func format(br *bufio.Reader)  {
	var (
		first_key string
		Dict = make(map[string]map[string]interface{})
	)

	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		// 1. []byte 转为 string
		lineStr := string(a)
		lineStr = strings.Replace(lineStr, " ", "", -1)

		// get first level key
		if strings.Contains(lineStr, "##") {
			first_key = strings.Replace(lineStr, "##", "", -1)
			if Dict[first_key] == nil {		// 如果第一级为空,初始化 map
				Dict[first_key] = make(map[string]interface{})
			}
		}

		// 2. # 表示注释,跳过 # 或者 为空
		if strings.Contains(lineStr, "#") || isEmpty(lineStr){
			continue
		}

		// 3.
		arr := strings.Split(lineStr, "=")		// 根据 = 分割数组
		if len(arr) == 2 && len(arr[0]) > 0 && len(arr[1]) > 0 {
			// 如果是正整数
			if isNum(arr[1]) {
				n, e := strconv.Atoi(arr[1])
				if e != nil{
					log.Fatal("conf fail convert int to string:",e)
				}
				Dict[first_key][arr[0]] = n
			}else {		// 字符串
				Dict[first_key][arr[0]] = arr[1]
			}
		}else {
			log.Fatal("格式填写有误,注意等于号前后不能为空!")
		}
	}
	data, _ := json.Marshal(Dict)		// covert map to bytes
	if err := json.Unmarshal(data, Tab); err != nil{		// covert bytes to struct
		log.Fatal("conf unmarshal json err:", err)
	}
}

// 是否为空
func isEmpty(str string) bool {
	if len(str) == 0 {
		return true
	}

	temp := strings.Replace(str, " ", "", -1)
	if len(temp) == 0 {
		return true
	}
	return false
}

// 是否是正整数
func isNum(n string) bool {
	return reg("^[1-9]\\d*$", n)
}

// 验证通用函数
func reg(reg, para string) bool {
	matched, err := regexp.MatchString(reg, para)
	if err != nil {
		return false
	}
	return matched
}