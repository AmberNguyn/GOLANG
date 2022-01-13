package main

import "fmt"

type myqueue struct {
	data   []int
	height int
}

func enqueue(a int, b *myqueue) {
	b.data = append(b.data, a)
	b.height = b.height + 1
}

func dequeue(b *myqueue) (a int) {
	if b.height == 0 {
		fmt.Println("\n Error")
	}

	a = b.data[0]

	newarray := make([]int, len(b.data))
	for i := 0; i < len(b.data)-1; i++ {
		newarray[i] = b.data[i+1]
	}
	b.data = newarray
	b.height = b.height - 1
	return a

}

func main() {
	var e myqueue

	enqueue(2, &e)
	fmt.Println(e)

	enqueue(3, &e)
	fmt.Println(e)

	enqueue(4, &e)
	fmt.Println(e)

	enqueue(5, &e)
	fmt.Println(e)

	enqueue(6, &e)
	fmt.Println(e)

	enqueue(7, &e)
	fmt.Println(e)

	f := dequeue(&e)
	fmt.Print("\n Take out no.2: ", f)
	fmt.Print(e)

	h := dequeue(&e)
	fmt.Print("\n Take out no.3: ", h)
	fmt.Print(e)

	j := dequeue(&e)
	fmt.Print("\n Take out no.4: ", j)
	fmt.Print(e)

	x := dequeue(&e)
	fmt.Print("\n Take out no.5: ", x)
	fmt.Print(e)

	y := dequeue(&e)
	fmt.Print("\n Take out no.6: ", y)
	fmt.Print(e)

	z := dequeue(&e)
	fmt.Print("\n Take out no.7: ", z)
	fmt.Print(e)

	t := dequeue(&e)
	fmt.Print("\n Take out no.8: ", t)
	fmt.Print(e)

}

// fksdjfjsdkfjsdkjfklasdj
// abc
