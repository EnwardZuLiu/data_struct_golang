package main

import (
	"flag"
	"fmt"
)

func main1() {
	var name = *flag.String("name", "everyone", "The greeting object.")
	flag.Parse()

	fmt.Printf("Hello, %s! \n", name)
}
