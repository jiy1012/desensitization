package test

import (
	"desensitization"
	"testing"
)

func Benchmark_Email_Field(b *testing.B) {
	type TestCommonFields struct {
		Email string `json:"email" desensitization:"EMAIL"`
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if err := desensitization.Desensitization(&TestCommonFields{Email: "example@example.com"}); err != nil {
			b.Fatal(err)
		}
	}
}
