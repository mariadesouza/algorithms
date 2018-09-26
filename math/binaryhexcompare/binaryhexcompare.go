package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	binStr := "1011"
	hexStr := "0xB"

	a, err := convertFromBase(binStr, 2)
	if err != nil {
		log.Fatal(err)
	}
	b, err := convertFromBase(hexStr, 16)
	if err != nil {
		log.Fatal(err)
	}
	if a == b {
		fmt.Println("equal", a, b)
	} else {
		fmt.Println(" not equal")
	}
}

func convertFromBase(numStr string, base int) (int64, error) {
	//var result int64
	cleaned := strings.Replace(numStr, "0x", "", -1)

	result, err := strconv.ParseInt(cleaned, base, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
