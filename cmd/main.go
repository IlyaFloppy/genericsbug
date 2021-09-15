package main

import (
	"genericsbug/lib"

	"github.com/IlyaFloppy/observable"
)

func main() {
	var _ *observable.Object[string]
	lib.Main()
}
