package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}

/*

Yes. `Scale` receives either the address of `main`'s `Vertex` or a copy of it, and that is the whole difference.

`v` in `main` is the struct `{X: 3, Y: 4}`. The receiver is an ordinary parameter of the method, filled in by the call `v.Scale(10)`.

With `func (v *Vertex) Scale(f float64)`, that parameter holds the address of `main`'s `v`. Go inserts the `&` for you, so the call behaves as `(&v).Scale(10)`. Lines 17 and 18 write through that address:

```text
main's v          Scale's receiver
{X: 3, Y: 4}  <---  pointer to that same struct
       |
       v
{X: 30, Y: 40}
```

`Abs` then reads those fields: `sqrt(30² + 40²) = 50`.

With `func (v Vertex) Scale(f float64)`, the parameter is a second struct. The call copies `{3, 4}` into it. Lines 17 and 18 change only the copy:

```text
main's v              Scale's receiver
{X: 3, Y: 4}          {X: 3, Y: 4}  --copy-->  {X: 30, Y: 40}
unchanged             discarded when Scale returns
```

`Abs` reads `main`'s struct, still `{3, 4}`, so `sqrt(9 + 16) = 5`. The scaled numbers existed only inside `Scale`.

*/
