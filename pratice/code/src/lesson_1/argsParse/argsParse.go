package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("os args is:", os.Args)

	name := flag.String("name", "xxj", "eg: --name xqx")
	flag.Parse()
	fmt.Printf("name:%v\n", name)

}
