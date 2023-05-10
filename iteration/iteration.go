package iteration

func Repeat(input string, times int) (repeated string) {
	for i := 0; i < times; i++ {
		repeated += input
	}
	return repeated
}
