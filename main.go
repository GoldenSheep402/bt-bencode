package main

import (
	"bt-bencode/bencode"
	"errors"
	"fmt"
)

func main() {
	mode, err := modeSelect()
	if err != nil {
		fmt.Println("[Error selecting mode]: %v\n", err)
		return
	}

	if mode == 1 {
		fmt.Println("[You have selected Decode mode]")
		bencode.DecodeStart()
	} else if mode == 2 {
		fmt.Println("[You have selected Encode mode]")
		// 执行Encode的代码
	}
}

func modeSelect() (int, error) {
	fmt.Println("Please choose a mode:")
	fmt.Println("1. Decode")
	fmt.Println("2. Encode")
	var mode int

	_, err := fmt.Scanln(&mode)

	if err != nil {
		return 0, err
	}

	if mode == 1 {
		return 1, nil
	} else if mode == 2 {
		return 2, nil
	} else {
		return 0, errors.New("[Invalid mode]")
	}
}
