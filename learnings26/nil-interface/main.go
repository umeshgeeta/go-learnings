package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

/*

An interface value is a pair: the concrete type, and the concrete value. The interface itself is nil only when both halves are empty. After line 25 the type half is `*T`, so `i` is not nil, even though the pointer stored in it is nil.

```22:27:learnings26/nil-interface/main.go
	var i I

	var t *T
	i = t
	describe(i)
	i.M()
```

`var i I` starts as that empty pair. Nothing has been stored, so `i == nil` is true, and `i.M()` would panic because there is no type to look the method up on.

`var t *T` is a nil pointer of type `*T`. The assignment `i = t` copies both pieces into the interface: type `*T`, value `nil`. The box is now occupied. `i == nil` is false. `describe` prints that pair:

```text
(<nil>, *main.T)
```

`<nil>` is the pointer inside. `*main.T` is the type, and that type is enough to make the interface non-nil.

`i.M()` therefore works. Go finds `M` on `*T` and passes the nil pointer as the receiver. Line 14 sees `t == nil` and prints `<nil>`. The method is running on a nil pointer; the interface that held it was never nil.

Line 29 replaces the pair with a non-nil pointer, `(&{hello}, *main.T)`, and `M` prints `hello`.

*/
