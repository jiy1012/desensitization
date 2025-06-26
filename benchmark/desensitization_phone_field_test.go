package test

import (
	"github.com/jiy1012/desensitization"
	"testing"
)

func Benchmark_Phone_Field(b *testing.B) {
	type TestCommonFields struct {
		Phone string `json:"phone" desensitization:"PHONE"`
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if err := desensitization.Desensitization(&TestCommonFields{Phone: "12345678911"}); err != nil {
			b.Fatal(err)
		}
	}
}
