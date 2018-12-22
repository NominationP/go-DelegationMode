package v3

import "errors"

/**
Undov2 几乎重写了 UndoV1
利用泛型编程,函数式编程,IoC等范式
 */

/**
declare Undo[] (其实是一个栈)
 */
type Undo []func()

/**
通用的Add() 需要一个指针,并把这个函数指针存放到 Undo[] 函数数组中
 */
func (undo *Undo) Add(function func()) {
	*undo = append(*undo, function)
}

/**
在Undo()函数中,遍历Undo[]函数数组,并执行,执行完后弹栈
 */
func (undo *Undo) Undo() error {
	functions := *undo
	if len(functions) == 0 {
		return errors.New("No function to undo")
	}
	index := len(functions) - 1
	if function := functions[index]; function != nil {
		function()
		functions[index] = nil // free closure for garbage collection
	}

	*undo = functions[:index]
	return nil
}

/**
IntSet 就可以改写成如下形式
 */

type IntSet struct {
	data map[int]bool
	undo Undo
}

func NewIntSet() IntSet {
	return IntSet{data: make(map[int]bool)}
}

/**
然后再Add和Delete中实现Undo操作
 */
func (set *IntSet) Add(x int) {
	if !set.Contains(x) {
		set.data[x] = true
		set.undo.Add(func() { set.Delete(x) })
	} else {
		set.undo.Add(nil)
	}
}

func (set *IntSet) Delete(x int) {
	if set.Contains(x) {
		delete(set.data, x)
		set.undo.Add(func() { set.Add(x) })
	} else {
		set.undo.Add(nil)
	}
}

func (set *IntSet) Undo() error {
	return set.undo.Undo()
}

func (set *IntSet) Contains(x int) bool {
	return set.data[x]
}

/**
Go语言的Undo接口把Undo的流程抽象出来
而要怎么Undo的事交给了业务代码来维护(用过注册一个Undo的方法)
这样在Undo的时候,就可以回调这个方法来做业务相关的操作了
 */

/**
这是不是和最开始的C++泛型编程也和像
也和map,reduce,filter这样只关心控制流程,不关心业务逻辑的做法很想
而且,一开始用一个UndoableIntSet来包装IntSet类,到反过来在IntSet里依赖Undo类,这就是控制反转IOC
 */
