package gotester

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/skeptycal/defaults"
	"github.com/skeptycal/types"
)

func NewParamMap(name string, args ...Param) ParamMap {
	m := paramMap{
		m:        make(map[string]Param, 2),
		id:       generateID(), // TODO not yet implemented
		testName: name,
	}

	return &m
}

type (
	ParamMap interface {
		Args() []Param
		LoadParams(args []Param) error

		Name() string
		Id() int64
		types.GetSetter

		KeyGuard(key Any) bool
		String() string
	}
	paramMap struct {
		id       int64 // TODO not yet implemented
		testName string
		m        pmap
	}

	pmap map[string]Param
)

func (p pmap) Get(key string) (Param, error) {
	if key == "" {
		return nil, fmt.Errorf("key not provided: %q", key)
	}

	if v, ok := p[key]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("key not found: %q", key)
}

// TODO not yet implemented
func generateID() int64 {
	return 12345
}

func (d paramMap) Id() int64 {
	return d.id
}

func (d paramMap) Name() string {
	return d.testName
}

func (d paramMap) LoadParams(args []Param) error {
	for _, arg := range args {
		err := d.Set(arg.Name(), arg.Value())
		if err != nil {
			log.Error(err)
		}
	}
	return nil
}

func (d paramMap) Args() []Param {
	return d.toSlice()
}

func (d paramMap) toSlice() []Param {
	args := make([]Param, 0, len(d.m))
	for _, v := range d.m {
		args = append(args, v)
	}
	return args
}

// KeyGuard returns true if the key is valid for the
// underlying map type.
//
// examples:
//
//  map[time.Time]struct{temp,pressure,pH}
// the underlying key must be a time.Time
//
// 	 map[string]Any
// the underlying key must be a string
//
// 	 map[int]*os.File
// the underlying key must be an integer
//
func (d paramMap) KeyGuard(key Any) bool {
	return isString(key)
}

func (d paramMap) Set(key Any, value Any) error {
	switch v := key.(type) {
	case string:
		d.m[v] = NewParam(v, value)
		return nil
	default:
		return fmt.Errorf("key type not string: %v", defaults.GetType(key))
	}
}

// isString returns true if the underlying value is a string.
// It guards against non-string keys causing panics.
func isString(key Any) bool {
	return reflect.TypeOf(key).Kind() == reflect.String
}

// an alternate version for performace profiling ...
func isStringNoReflect(key Any) bool {
	switch key.(type) {
	case string:
		return true
	default:
		return false
	}
}

// any2str converts an Any interface to a string if the underlying
// Kind is string. If the underlying Kind is not string, it returns
// the empty string.
func any2str(key Any) string {
	switch t := key.(type) {
	case string:
		return t
	default:
		return ""
	}
}

func (d paramMap) Get(key Any) (Any, error) {
	if k := any2str(key); k != "" {
		return d.m.Get(k)
	}

	switch t := key.(type) {
	case string:
		return d.m.Get(t)
	default:
		return nil, fmt.Errorf("key type not string: %v", defaults.GetType(key))
	}
}

func (d paramMap) String() string {

	const format = "%-20s = %-20v\n"
	sb := &strings.Builder{}
	defer sb.Reset()

	fmt.Fprintf(sb, "Test Item Parameters for %s (id: %d):\n", d.testName, d.id)
	fmt.Fprintf(sb, format, "Key", "Value")

	for key, value := range d.m {
		fmt.Fprintf(sb, format, key, value)
	}

	return sb.String()
}
