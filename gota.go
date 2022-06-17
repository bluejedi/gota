package main

import (
	scanner "github.com/bluejedi/gota/go-tvscanner"

	"fmt"
)


func main() {
	fmt.Println("Starting")
	cl := scanner.New()
	data, _ := cl.GetAnalysis("indonesia", "IDX", "GOTO", "1d")
	//data, _ := cl.GetAnalysis("crypto", "BITFINEX", "BTCUSD", "1d")
	//fmt.Println("result: ", data)
	fmt.Printf("%+v\n", data)
}