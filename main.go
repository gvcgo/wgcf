package main

import (
	"log"

	"github.com/gvcgo/wgcf/cmd"
	"github.com/gvcgo/wgcf/util"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(util.GetErrorMessage(err))
	}
}
