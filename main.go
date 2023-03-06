package main

import (
	"de-bt-bencode/bencode"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: bdecode FILENAME")
		return
	}
	filename := os.Args[1]
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		return
	}

	// 解码Bencode字符串
	var value interface{}
	err = bencode.Decode(data, &value)
	if err != nil {
		fmt.Printf("Error decoding Bencode data: %v\n", err)
		return
	}

	// 打印解码结果
	fmt.Printf("Decoded value: %+v\n", value)
}
