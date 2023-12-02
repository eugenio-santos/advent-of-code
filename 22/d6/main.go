package main

import (
	"fmt"
	"os"
)

func main() {
	dat, err := os.ReadFile("in")
	if err != nil {
		panic(err)
	}

	str := string(dat)
	letters := make(map[byte]int, 26)
	gap := 14

	for i := 0; i < gap; i++ {
		letters[str[i]]++
	}

	for i := 14; i < len(str); i++ {
		letters[str[i-gap]]--
		letters[str[i]]++
		flag := true

		for _, v := range letters {
			if v > 1 {
				flag = false
				break
			}
		}

		if flag {
			fmt.Println(i + 1)
			break
		}
	}
}
