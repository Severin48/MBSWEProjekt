package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Show interface {
	show() string
}

type ShowFunc func(interface{}) string

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

type dictionary[T any] struct {
	show ShowFunc
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

func showNodeDict[T Show](n *node[T], dict dictionary[T]) string {
	var s string

	for n != nil {
		s = s + dict.show(n.val) + " -> "
		n = n.next
	}

	s = s + " nil"

	return s
}

func generateNodeList[T Show](length int, generator func() T) *node[T] {
	if length <= 0 {
		return nil
	}

	val := generator()
	return &node[T]{val, generateNodeList(length-1, generator)}
}

func Task4() {
	nodeCount := 10000
	intGenerator := func() Integer { return Integer(rand.Int()) }
	strGenerator := func() String { return String(fmt.Sprintf("%d", rand.Int())) }

	intDict := dictionary[Integer]{
		show: func(v interface{}) string {
			if val, ok := v.(Integer); ok {
				return val.show()
			}
			return ""
		},
	}

	strDict := dictionary[String]{
		show: func(v interface{}) string {
			if val, ok := v.(String); ok {
				return val.show()
			}
			return ""
		},
	}

	rand.Seed(time.Now().UnixNano())

	start := time.Now()
	intNodeList := generateNodeList(nodeCount, intGenerator)
	showNode(intNodeList)
	// fmt.Println(showNode(intNodeList))
	strNodeList := generateNodeList(nodeCount, strGenerator)
	showNode(strNodeList)
	// fmt.Println(showNode(strNodeList))
	elapsed := time.Since(start)

	fmt.Printf("Run-time method lookup took %d ms\n", elapsed.Milliseconds())

	start = time.Now()
	intNodeList = generateNodeList(nodeCount, intGenerator)
	showNodeDict(intNodeList, intDict)
	// fmt.Println(showNodeDict(intNodeList, intDict))
	strNodeList = generateNodeList(nodeCount, strGenerator)
	showNodeDict(strNodeList, strDict)
	// fmt.Println(showNodeDict(strNodeList, strDict))
	elapsed = time.Since(start)

	fmt.Printf("Dictionary translation took %d ms\n", elapsed.Milliseconds())
}
