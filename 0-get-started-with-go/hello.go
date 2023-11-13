package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
}

// Notes
// go mod init module/desc to create dependency tracking in go
// go mod tidy to add packages as requirement
