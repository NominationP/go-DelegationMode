package v1

import (
	"fmt"
	. "sort"
	"strings"
)

/**
声明一个数据容器,其中有Add(), Delete() Contains(), 还有一个转字符串的方法
 */

type IntSet struct {
	data map[int]bool
}

func NewIntSet() IntSet {
	return IntSet{make(map[int]bool)}
}

func (set *IntSet) Add(x int) {
	set.data[x] = true
}

func (set *IntSet) Delete(x int) {
	delete(set.data, x)
}

func (set *IntSet) Contains(x int) bool {
	return set.data[x]
}

func (set *IntSet) String() string { // Staisfies fmt.Stringer interface
	if len(set.data) == 0 {
		return "{}"
	}
	ints := make([]int, 0, len(set.data))
	for i := range set.data {
		ints = append(ints, i)
	}
	Ints(ints)
	parts := make([]string, 0, len(ints))
	for _, i := range ints {
		parts = append(parts, fmt.Sprint(i))
	}
	return "{" + strings.Join(parts, ",") + "}"
}

//func main() {
//	ints := NewIntSet()
//	for _, i := range []int{1, 3, 5, 7} {
//		ints.Add(i)
//		fmt.Println(ints)
//	}
//
//	for _, i := range []int{1, 2, 3, 4, 5, 6, 7} {
//		fmt.Print(i, ints.Contains(i), " ")
//		ints.Delete(i)
//		fmt.Println(ints)
//	}
//}
