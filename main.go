package main

import (
	"log"

	"github.com/CatchZeng/feishu/cmd/feishu"
)

func main() {
	log.SetFlags(0)
	feishu.Execute()
}
