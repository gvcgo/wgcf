package main

import (
	"log"

	"github.com/moqsien/wgcf/cmd"
	"github.com/moqsien/wgcf/util"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(util.GetErrorMessage(err))
	}
}
