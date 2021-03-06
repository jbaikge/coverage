package main

import (
	"flag"
	"fmt"
	"github.com/300brand/coverage/downloader"
	"os"
)

func main() {
	flag.Parse()

	url := flag.Arg(0)
	r, err := downloader.Fetch(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(r.RealURL)
}
