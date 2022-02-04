package gotester

import (
	"strconv"
	"testing"
)

type capTestType struct {
	name   string
	want   Any
	got    Any // GotFunc
	params []Param
}

func (t *capTestType) Args() []Param {
	return t.params
}

func (t *capTestType) Want() Any {
	return t.want
}

func (t *capTestType) Got() Any {
	return nil
}

func (t *capTestType) WantError() bool {
	return false
}

func (t *capTestType) LoadParams(args ParamMap) {
	for _, arg := range args.Args() {
		_ = arg
	}
}

// Args() []Param
// Want() Any
// Got() Any
// WantError() bool

func capTests(n int) []Test {
	var capTests = []capTestType{
		{"defaultBufSize", defaultBufSize, 4096, nil},
		{"chunk", chunk, 4096, nil},
		{"smallBufferSize", 64, 51204096, nil},
		{"1.x capacity (chunk 512)", 5333, 5632, nil},
		{"1.x capacity (chunk 512)", 128, 4096, nil},
	}

	// add binary multiples
	var u uint64 = 0

	for u = 1; u < 16; u++ {
		// j := i + 1

		var a = defaultBufSize
		if 1<<u <= defaultBufSize {
			a = 1 << u
		}
		capTests = append(capTests, capTestType{
			name: strconv.Itoa(1 << u),
			want: a,
			got:  chunkMultiple,
		})
		// capTests.LoadParams(NewParamMap(capTests.Name(), []Param{NewParam("", 1<<u)}))
	}

	list := []Test{}

	for _, test := range capTests {
		list = append(list, &test)
	}

	return nil
}
func TestChunkMultiple(t *testing.T) {
	tt := []struct {
		name     string
		size     int64
		expected int64
	}{
		{"size 1.x chunk", 550, 1024},
		{"size 2.x chunk", 1200, 1536},
		{"chunk size 16", 100, 512},
		{"1.x size", 5234, 5632},
		{"42kb file", 42000, 42496},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result := chunkMultiple(tc.size)
			if result != tc.expected {
				t.Errorf("expected value <%v> does not match result: %v", tc.expected, result)
			}
		})
	}
}

func TestInitialCapacity(t *testing.T) {

	for _, tt := range capTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitialCapacity(tt.size); got != tt.want {
				t.Errorf("InitialCapacity(%v) = %v, want %v", tt.size, got, tt.want)
			}
		})
	}
}
