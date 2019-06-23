package main

import (
	"fmt"
	"os"
)

func main() {
	resp, err := StartCli(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, s := range resp {
		fmt.Println(s)
	}
}
