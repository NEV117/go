package main

func Sum(x, y int) int {
	return x + y
}

func GetMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

//--------comando en consola para el coverage---------
//go test -cover
//go test -coverprofile=coverage.out
//go tool cover -func=coverage.out
//go tool cover -html=coverage.out

//---------tiempo de ejecucion Profiling--------------
//go test -cpuprofile=cou.out
//go tool pprof cou.out
//	top
//	list Fibonacci
//  web
//  pdf

//crear modulo en go
//go mod init github.com/NEV117/mockTest
