package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/hacker65536/hello/morestrings"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
