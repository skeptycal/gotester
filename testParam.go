package gotester

import (
	"fmt"
	"reflect"
)

type (
	Param interface {
		Name() string
		Value() Any
		String() string
	}

	param struct {
		name  string
		value Any
	}

	params struct {
		m map[string]Param
	}
)

var p = NewParam

// NewArg
func NewParam(name string, arg Any) Param {
	if name == "" {
		name = fmt.Sprintf("%v", arg)
	}
	return &param{name, arg}
}

func (a *param) Name() string {
	return a.name
}

func (a *param) Value() Any {
	return a.value
}

func (a *param) Kind() reflect.Kind {
	return reflect.ValueOf(a.value).Kind()
}
func (a *param) String() string {
	return fmt.Sprint("%v", a.value)
}

func (a *param) GoString() string {
	return fmt.Sprintf("parameter %q = %q", a.name, a.value)
}
