package utils

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func OutputValue(value interface{}) {
	switch v := value.(type) {
	case string:
		// 将UTF-8编码的字符串转换为GBK编码的字节流
		encoder := simplifiedchinese.GBK.NewEncoder()
		gbkBytes, _, err := transform.Bytes(encoder, []byte(v))
		if err != nil {
			fmt.Printf("Error encoding string: %v\n", err)
			return
		}
		// 将GBK编码的字节流转换为GBK编码的字符串
		gbkString := string(gbkBytes)
		// 输出GBK编码的字符串
		fmt.Println(gbkString)
	case []interface{}:
		// 对于列表，递归输出每个元素
		for _, item := range v {
			OutputValue(item)
		}
	case map[string]interface{}:
		// 对于字典，递归输出每个键值对
		for k, v := range v {
			fmt.Printf("%s: ", k)
			OutputValue(v)
		}
	default:
		// 其他类型直接输出
		fmt.Printf("%v\n", v)
	}
}
