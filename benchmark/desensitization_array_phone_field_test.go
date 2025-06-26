package test

import (
	"github.com/jiy1012/desensitization"
	"testing"
)

func Benchmark_Array_Phone_Field(b *testing.B) {
	type TestCommonFields struct {
		Phone string `json:"phone" desensitization:"PHONE"`
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if err := desensitization.Desensitization(
			&[]TestCommonFields{
				{Phone: "12345678910"},
				{Phone: "12345678911"},
				{Phone: "12345678912"},
				{Phone: "12345678913"},
				{Phone: "12345678914"},
				{Phone: "12345678915"},
				{Phone: "12345678916"},
				{Phone: "12345678917"},
				{Phone: "12345678918"},
				{Phone: "12345678919"}}); err != nil {
			b.Fatal(err)
		}
	}
}
