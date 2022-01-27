package main

import (
	"fmt"
)

type Numeric interface {
	int | int64 | string
}

type Amount[T Numeric] struct {
	val T
}

func NewAmount[T Numeric](val T) Amount[T] {
	return Amount[T]{val: val}
}

func (a Amount[T]) Add(b Amount[T]) Amount[T] {
	return Amount[T]{val: a.val + b.val}
}

// invalid operation: operator / not defined on a.val (variable of type T constrained by Numeric)
//
// func (a Amount[T]) Div(b Amount[T]) Amount[T] {
// 	return Amount[T]{val: a.val / b.val}
// }

func main() {
	a := NewAmount(64)
	fmt.Printf("a amount is: %d\n", a.val)
	b := NewAmount(32)
	fmt.Printf("amount sum: %d\n", a.Add(b).val)

	// as expected throws an error, because different types
	// c := NewAmount(int64(16))
	// fmt.Printf("int amount + int64 amount: %d", a.Add(c).val)

	// Let's try strings
	s1 := NewAmount("one")
	s2 := NewAmount("two")
	fmt.Printf("two strings: %s", s1.Add(s2).val)
}
