package validator

import (
	"reflect"
	"testing"
)

func TestPattern_Struct(t *testing.T) {
	p := Pattern{
		Regex: "[a-z]+",
		Match: true,
	}
	if reflect.TypeOf(p.Regex).String() != "string" {
		t.Fatal("expected Regex of type string")
	}
	if reflect.TypeOf(p.Match).String() != "bool" {
		t.Fatal("expected Match of type bool")
	}
}
