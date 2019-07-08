package main

import (
	"strings"
	"fmt"
	"flag"
)


var n = flag.Bool("n",true,"omit trailing")
var sep = flag.String("s"," ","seperate")

func main(){
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(),*sep))
	if!*n {
		fmt.Println()
	}
}