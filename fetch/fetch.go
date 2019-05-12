package main

import (
	"strings"
	"io/ioutil"
	"fmt"
	"net/http"
	"os"
	
)
func  main()  {
	for _, url := range os.Args[1:]{
		if !strings.HasPrefix("https://",url) {
			url = "https://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprint(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		code := resp.Status
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr,"fetch %s: %v\n",url,err)
		    os.Exit(1)
		}
		fmt.Printf("%s",b)
		fmt.Printf("%s",code)
	}
}