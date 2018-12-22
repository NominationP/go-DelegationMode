package v2

import (
	"errors"
	"fmt"
)

/**
UndoDelegationMode.go 中的数据容器平淡无奇,我们想给它加一个Undo功能
 */
type UndoableIntSet struct {
	IntSet // Embedding (delegation)
	functions []func()
}

func NewUndoableIntSet() UndoableIntSet {
	return UndoableIntSet{NewIntSet(), nil}
}

func (set *UndoableIntSet) Add(x int) { // Override
	if !set.Contains(x) {
		set.data[x] = true
		set.functions = append(set.functions, func() { set.Delete(x) })
	} else {
		set.functions = append(set.functions, nil)
	}
}

func (set *UndoableIntSet) Delete(x int) {
	//Overried
	if set.Contains(x) {
		delete(set.data, x)
		set.functions = append(set.functions, func() { set.Add(x) })
	} else {
		set.functions = append(set.functions, nil)
	}
}

func (set *UndoableIntSet) Undo() error {
	if len(set.functions) == 0 {
		return errors.New("No functions to undo")
	}

	index := len(set.functions) - 1
	if function := set.functions[index]; function != nil {
		function()
		set.functions[index] = nil // Free closure for garbage collection
	}
	set.functions = set.functions[:index]
	return nil
}

func main() {
	ints := NewUndoableIntSet()
	for _, i := range []int{1, 3, 5, 7} {
		ints.Add(i)
		fmt.Println(ints)
	}

	for _, i := range []int{1, 2, 3, 4, 5, 6, 7} {
		fmt.Println(i, ints.Contains(i), " ")
		ints.Delete(i)
		fmt.Println(ints)
	}

	fmt.Println()
	for {
		if err := ints.Undo(); err != nil {
			break
		}
		fmt.Println(ints)
	}
}
