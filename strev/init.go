package strev

func Reverse(str string) string {
	var res string
	for i := len(str) - 1; i >= 0; i-- {
		res += string(str[i])
	}
	return res
}
