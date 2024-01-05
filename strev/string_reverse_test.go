package strev

import "testing"

type ReverseStruct struct {
	in, re string
}

var reverseStructs = []ReverseStruct{
	{"ABC", "CBA"},
	{"123456789", "987654321"},
	// {"dnsakj", "123"},
}

func TestReverse(t *testing.T) {
	for _, v := range reverseStructs {
		if r := Reverse(v.in); r != v.re {
			t.Log(v.in + "不等于" + v.re)
			t.Fail()
		}
	}
}

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse("ABCDEF")
	}
}
