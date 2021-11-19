package main

import (
	"github.com/rodge0/sameas"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(sameas.Analyzer)
}
