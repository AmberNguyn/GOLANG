package main

import "fmt"

type mystack struct {
	data   []int
	height int
}

func stackpush(a int, b *mystack) {
	b.data = append(b.data, a)
	b.height = b.height + 1
}

func stackpop(b *mystack) (a int) {

	a = b.data[b.height-1]
	b.data[b.height-1] = 0
	b.height = b.height - 1
	return a

}

func main() {
	var e mystack

	stackpush(2, &e)
	fmt.Println(e)

	stackpush(3, &e)
	fmt.Println(e)

	stackpush(4, &e)
	fmt.Println(e)

	stackpush(5, &e)
	fmt.Println(e)

	stackpush(6, &e)
	fmt.Println(e)

	stackpush(7, &e)
	fmt.Println(e)

	f := stackpop(&e)
	fmt.Print("\n Take out no.2: ", f)
	fmt.Print(e)

	h := stackpop(&e)
	fmt.Print("\n Take out no.3: ", h)
	fmt.Print(e)

	j := stackpop(&e)
	fmt.Print("\n Take out no.4: ", j)
	fmt.Print(e)

	x := stackpop(&e)
	fmt.Print("\n Take out no.5: ", x)
	fmt.Print(e)

	y := stackpop(&e)
	fmt.Print("\n Take out no.6: ", y)
	fmt.Print(e)

	z := stackpop(&e)
	fmt.Print("\n Take out no.7: ", z)
	fmt.Print(e)

	t := stackpop(&e)
	fmt.Print("\n Take out no.8: ", t)
	fmt.Print(e)

	u := stackpop(&e)
	fmt.Print("\n Take out no.9: ", u)
	fmt.Print(e)
}
