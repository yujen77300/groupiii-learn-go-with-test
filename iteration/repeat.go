package iteration

func Repeat(character string, times int) string {
	result := ""
	// no syntax like 'for i in range(5)' in Go, as in Python.
	for i := 0; i < times; i++ {
		result += character
	}
	return result
}
