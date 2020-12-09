package main

import (
	"github.com/ashtyn3/fondue/parent"
)

func main() {
	p := parent.MakeParent("http://localhost:8080", "")
	p.Get()
}
