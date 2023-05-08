package test7

var fibRes []int

func fib(n int) int {
	if n < 3 {
		return 1
	}
	if fibRes[n] == 0 {
		fibRes[n] = fib(n-1) + fib(n-2)
	}
	return fibRes[n]
}

func Recursion() {
	n := 45
	fibRes = make([]int, n+1)
	println(fib(n))
}
