package goconfig

import (
	"testing"
)

type getvalue struct {
	title       string
	key         string
	value       interface{}
	shouldMatch interface{}
}

var x = `
uint64 =9223372036854775808
int64=123456223
int=1234556223

[one]
float=0.25
string=goconfig

[two]
bool=true
password= "adlfjlskdf "
`

func TestGet(t *testing.T) {
	// 方便测试，
	InitFromBytes([]byte(x))
	tests := []getvalue{
		{
			title: "float64",
			key:   "one.float",
			value: 0.25,
		},
		{
			title: "string",
			key:   "one.string",
			value: "goconfig",
		},
		{
			title: "uint64",
			key:   "uint64",
			value: uint64(1 << 63),
		},
		{
			title: "int64",
			key:   "int64",
			value: int64(123456223),
		},
		{
			title: "int",
			key:   "int",
			value: 1234556223,
		},
		{
			title: "bool",
			key:   "two.bool",
			value: true,
		},
		{
			title: "password",
			key:   "two.password",
			value: "adlfjlskdf ",
		},
	}
	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			testGet(t, test)
		})
	}
}

func testGet(t *testing.T, test getvalue) {
	switch test.value.(type) {
	case string:
		if test.key == "string" {
			if test.value != ReadString(test.key) {
				t.Error("fail")
			}
		}
		if test.key == "file" {
			if test.value != ReadFile(test.key) {
				t.Error("fail")
			}
		}
		if test.key == "password" {
			if test.value != ReadPassword(test.key) {
				t.Logf("--%s---", ReadPassword(test.key))
				t.Error("fail")
			}
		}
	case int:
		if test.value != ReadInt(test.key) {
			t.Error("fail")
		}
	case uint64:
		if test.value != ReadUint64(test.key) {
			t.Error("fail")
		}
	case int64:
		if test.value != ReadInt64(test.key) {
			t.Error("fail")
		}
	case float64:
		if test.value != ReadFloat64(test.key) {
			t.Error("fail")
		}
	default:

	}
}
