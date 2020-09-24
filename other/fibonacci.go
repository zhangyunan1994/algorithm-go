package other

func Fib(N int) int {
	if N < 2 {
		return N
	}
	if N == 2 {
		return 1
	}
	var a = 0
	var b = 1
	for i := 2; i <= N; i++ {
		a, b = b, a+b
	}
	return b
}
