package main

import "fmt"

type Show interface {
	show() string
}

type node[T any] struct {
	val  T
	next *node[T]
}

type Integer int

func (i Integer) show() string {
	return fmt.Sprintf("%d", i)
}

type String string

func (s String) show() string {
	return string(s)
}

func showNode[T Show](n *node[T]) string {
	var s string

	for n != nil {
		s = s + n.val.show() + " -> "
		n = n.next
	}

	s = s + " nil"

	return s
}

func Task4() {
	n1 := &node[Integer]{Integer(1), nil}
	n2 := &node[Integer]{Integer(2), n1}
	n3 := &node[Integer]{Integer(3), n2}

	fmt.Println(showNode[Integer](n3))

	s1 := &node[String]{String("A"), nil}
	s2 := &node[String]{String("B"), s1}
	s3 := &node[String]{String("C"), s2}

	fmt.Println(showNode[String](s3))
}
