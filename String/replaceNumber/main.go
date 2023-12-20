package main

import (
	"Leetcode/String/replaceNumber/Methods"
	"fmt"
)

func main() {
	number := Methods.ReplaceNumber("a1b2c3")

	fmt.Printf("old: %v, new: %v", "a1b2c3", number)

	//answer

	var strByte []byte

	_, err := fmt.Scanln(&strByte)
	if err != nil {
		return
	}

	for i := 0; i < len(strByte); i++ {
		if strByte[i] <= '9' && strByte[i] >= '0' {
			inserElement := []byte{'n', 'u', 'm', 'b', 'e', 'r'}
			strByte = append(strByte[:i], append(inserElement, strByte[i+1:]...)...)
			i = i + len(inserElement) - 1
		}
	}

	fmt.Printf(string(strByte))

}
