package main

import (
	"fmt"
	"golang.org/x/example/stringutil"
)

func main() {
	sourceStr := "Hello, OTUS!";
	fmt.Println(stringutil.Reverse(sourceStr))
}
