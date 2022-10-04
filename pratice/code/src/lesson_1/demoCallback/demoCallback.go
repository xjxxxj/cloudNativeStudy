package main

import (
	"fmt"
	"strings"
)

func doPrint(context string) {
	fmt.Println(context)
}

func parseAndPrint(context string) {
	parseText := strings.Fields(context)
	for _, item := range parseText {
		fmt.Println(item)
	}
}

func doTextResolve(target string, resolvers []func(string)) {
	for _, resolve := range resolvers {
		resolve(target)
	}
}

func main() {
	text := "Hello World!"
	trimAndPrint := func(context string) {
		fmt.Println(strings.TrimSpace(context))
	}
	resolves := []func(string){doPrint, parseAndPrint, trimAndPrint}
	doTextResolve(text, resolves)
}
