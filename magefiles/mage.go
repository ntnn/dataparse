package main

import (
	//mage:import
	"github.com/ntnn/magefiles/base"
)

func init() {
	base.PreGenerateDeletePatterns = []string{
		"*_gen.go",
		"*_gen_test.go",
	}
}
