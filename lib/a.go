package lib

import (
	"fmt"

	"github.com/IlyaFloppy/observable"
)

type ThisTypeDecalarationCausesInternalCompilerError struct {
	obj *observable.Object[string]
}

func Main() {
	_ = observable.New[string]("hello")
	fmt.Println("hello")
}
