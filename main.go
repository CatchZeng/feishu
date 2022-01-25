package main

import (
	"log"
	"os"

	"github.com/CatchZeng/feishu/cmd/feishu"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(0)
	feishu.Execute()
}
