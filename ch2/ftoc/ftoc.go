package main

import (
	"fmt"
)

func main(){
	const freezingF, boilingF = 32, 212.0
	fmt.Printf("freezingF  = %gC boilingF  %gC\n",fToC(freezingF),fToC(boilingF))
}

func fToC(f float64) float64{
	return(f - 32)*5 /9
}