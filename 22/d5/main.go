package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"container/list"

	"github.com/golang-collections/collections/stack"
)

func main() {

	l1 := list.New()
	l2 := list.New()

	l1.PushFront(1)
	l1.PushFront(2)
	l1.PushFront(3)
	l1.PushFront(4)
	l1.PushFront(5)

	l2.PushFront(1)
	l2.PushFront(2)
	l2.PushFront(3)

	fmt.Println(l1.Front().Value, l2.Front().Value)

	e := l1.Front()

	for i := 0; i < 2; i++ {
		e = e.Prev()
	}

	l2.PushFront(e)
	// l1.Remove(e)

	fmt.Println(l1.Front().Value, l2.Front().Value)

	return

	dat, err := os.ReadFile("in")
	if err != nil {
		panic(err)
	}
	datA := strings.Split(string(dat), "\n")
	var bPoint int
	var numStacks int
	for i, v := range datA {
		if strings.HasPrefix(v, " 1") {
			bPoint = i
			numStacks = int(v[len(v)-2] - '0')
			break
		}
	}

	fmt.Println("num stacks", numStacks)
	fmt.Println(bPoint)
	stacks := make([]stack.Stack, numStacks)

	for i := bPoint - 1; i >= 0; i-- {
		v := datA[i]
		for i := 0; i < numStacks; i++ {
			char := string(v[(i*4)+1])
			if char != " " {
				stacks[i].Push(v[(i*4)+1])
			}
		}
	}

	fmt.Println("Stacks")
	for _, v := range datA[bPoint+2:] {
		s := strings.Split(v, " ")

		move, _ := strconv.Atoi(s[1])
		from, _ := strconv.Atoi(s[3])
		to, _ := strconv.Atoi(s[5])

		tmp := stack.New()

		for i := 0; i < move; i++ {
			tmp.Push(stacks[from-1].Pop())
		}

		for tmp.Len() > 0 {
			stacks[to-1].Push(tmp.Pop())
		}
	}

	for _, v := range stacks {
		str := fmt.Sprintf("%v", v.Peek())
		i, _ := strconv.Atoi(str)
		fmt.Printf("%s", string(i))
	}
	fmt.Println()

	// for _, v := range datA[bPoint:] {
	// 	fmt.Println(v)
	// }
}
