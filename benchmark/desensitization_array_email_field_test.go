package test

import (
	"github.com/jiy1012/desensitization"
	"testing"
)

func Benchmark_Array_Email_Field(b *testing.B) {
	type TestCommonFields struct {
		Email string `json:"email" desensitization:"EMAIL"`
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if err := desensitization.Desensitization(
			&[]TestCommonFields{
				{Email: "example1@example1.com"},
				{Email: "example2@example2.com"},
				{Email: "example3@example3.com"},
				{Email: "example4@example4.com"},
				{Email: "example5@example5.com"},
				{Email: "example6@example6.com"},
				{Email: "example7@example7.com"},
				{Email: "example8@example8.com"},
				{Email: "example9@example9.com"},
				{Email: "example0@example0.com"},
			}); err != nil {
			b.Fatal(err)
		}
	}
}
