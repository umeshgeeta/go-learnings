package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

/*

Both comments match the program. The type name is the part to tighten: `pos` and `neg` have type `func(int) int`, the function `adder` returns. They are not of type `adder`.

`adder` itself has type `func() func(int) int`: it takes no arguments and returns a function. The short declaration

```go
pos, neg := adder(), adder()
```

is the same inference as `i := 0`. The compiler takes the type from the right-hand side, so this is the explicit form:

```go
var pos func(int) int = adder()
var neg func(int) int = adder()
```

Each call `adder()` runs `sum := 0` again and returns a new function value that closes over that call's `sum`. `pos` and `neg` are two of those returned functions, so they hold two different `sum` variables. Every later call `pos(i)` updates and reads `pos`'s `sum`. Every `neg(-2*i)` does the same with `neg`'s `sum`. The two running totals never meet.

*/
