package test

import (
	"desensitization"
	"testing"
)

func Benchmark_Array_Phone_And_Email_Field(b *testing.B) {
	type TestCommonFields struct {
		Username string `json:"username"`
		Phone    string `json:"phone" desensitization:"PHONE"`
		Email    string `json:"email" desensitization:"EMAIL"`
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if err := desensitization.Desensitization(&[]TestCommonFields{
			{Phone: "12345678910", Email: "example0@example.com", Username: "A"},
			{Phone: "12345678911", Email: "example1@example.com", Username: "B"},
			{Phone: "12345678912", Email: "example2@example.com", Username: "C"},
			{Phone: "12345678913", Email: "example3@example.com", Username: "D"},
			{Phone: "12345678914", Email: "example4@example.com", Username: "E"},
			{Phone: "12345678915", Email: "example5@example.com", Username: "F"},
			{Phone: "12345678916", Email: "example6@example.com", Username: "G"},
			{Phone: "12345678917", Email: "example7@example.com", Username: "H"},
			{Phone: "12345678918", Email: "example8@example.com", Username: "I"},
			{Phone: "12345678919", Email: "example9@example.com", Username: "J"}}); err != nil {
			b.Fatal(err)
		}
	}
}
