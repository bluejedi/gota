# go-tvscanner
TradingView scanner api client

## Thx to :
https://github.com/elliottcarlson/go-tvscanner/tree/use_screener
https://github.com/Canibus/go-tvscanner


## Usage
~~~ go
package main

import (
	scanner "github.com/bluejedi/go-tvscanner"

	"fmt"
)


func main() {
	fmt.Println("Starting")
	cl := scanner.New()
	cl.GetAnalysis("crypto", "BITTREX", "BTCUSD", "1h")
}
~~~