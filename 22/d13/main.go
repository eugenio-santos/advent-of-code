package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type item struct {
	i      int
	l      []item
	isList bool
}

func main() {
	dat, _ := os.ReadFile("in")

	datA := strings.Split(string(dat), "\n\n")

	ss := []string{}
	for _, v := range datA {
		strs := strings.Split(v, "\n")
		ss = append(ss, strs[0])
		ss = append(ss, strs[1])
	}

	ordered := []int{}
	for i := 0; i < len(ss); i += 2 {
		ts := ss[i]
		bs := ss[i+1]

		fmt.Println(ts)
		fmt.Println(bs)

		ind := 1
		top := item{i: 0, l: *getItems(&ts, &ind, &[]item{}), isList: true}
		ind = 1
		bot := item{i: 0, l: *getItems(&bs, &ind, &[]item{}), isList: true}

		isOrder, _ := comparePairs(top.l, bot.l)

		if isOrder {
			ordered = append(ordered, (i/2)+1)
		}
	}

	fmt.Println(ordered)
	sum := 0
	for _, v := range ordered {
		sum += v
	}
	fmt.Println(sum)
}

func comparePairs(top, bot []item) (bool, bool) {
	biggerL := 0
	if len(top) > len(bot) {
		biggerL = len(top)
	} else {
		biggerL = len(bot)
	}

	for i := 0; i < biggerL; i++ {
		//check both have intem in pos
		if hasNext(i, top) && !hasNext(i, bot) {
			return false, true
		} else if !hasNext(i, top) && hasNext(i, bot) {
			return true, true
		}

		// check what is the type of each element
		if !top[i].isList && !bot[i].isList {
			if top[i].i > bot[i].i {
				return false, true
			} else if top[i].i < bot[i].i {
				return true, true
			}
		} else if top[i].isList && bot[i].isList {
			isOrdered, isFinished := comparePairs(top[i].l, bot[i].l)
			if isFinished && isOrdered {
				return true, true
			} else if isFinished && !isOrdered {
				return false, true
			}
		} else {
			if !top[i].isList {
				top[i].l = []item{{i: top[i].i, l: nil, isList: false}}
			} else if !bot[i].isList {
				bot[i].l = []item{{i: bot[i].i, l: nil, isList: false}}
			}

			isOrdered, isFinished := comparePairs(top[i].l, bot[i].l)
			if isFinished && isOrdered {
				return true, true
			} else if isFinished && !isOrdered {
				return false, true
			}
		}
	}
	return false, false
}

func getItems(str *string, ind *int, items *[]item) *[]item {
	for *ind < len(*str) {
		switch (*str)[*ind] {
		case ',':
			*ind++
			continue
		case '[':
			// getItmes
			*ind++
			res := item{i: 0, l: *getItems(str, ind, &[]item{}), isList: true}
			// append
			*items = append(*items, res)

		case ']':
			*ind++
			// return
			return items
		default:
			// parse number and append
			fstd, _ := strconv.Atoi(string((*str)[*ind]))
			snd, err := strconv.Atoi(string((*str)[*ind+1]))
			num := 0
			if err != nil {
				num = fstd
			} else {
				num, _ = strconv.Atoi(strconv.Itoa(fstd) + strconv.Itoa(snd))
			}

			*items = append(*items, item{i: num, l: nil, isList: false})
			*ind++
		}
	}
	return items
}

func hasNext(i int, arr []item) bool {
	if i < len(arr) {
		return true
	} else {
		return false
	}
}
