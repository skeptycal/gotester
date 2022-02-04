package gotester

import "fmt"

type (
	GotFunc func(args ...Param) Any

	Test interface {
		Args() []Param
		Want() Any
		Got() Any
		WantError() bool
	}

	test struct {
		args      paramMap
		want      Any
		got       GotFunc
		wantError bool
	}
)

func NewTest(
	name string,
	want Any,
	got GotFunc,
	wantErr bool,
	args paramMap) Test {

	if name == "" {
		name = argString(args.Args())
	}

	return &test{}
}

func argString(args ...Any) string {
	return fmt.Sprint(args)
}

func (t *test) Args() []Param {
	return t.args.Args()
}

func (t *test) Want() Any {
	return t.want
}

func (t *test) Got() Any {
	return t.got(t.Args())
}

func (t *test) WantError() bool {
	return t.wantError
}
