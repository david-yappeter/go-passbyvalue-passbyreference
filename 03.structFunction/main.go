package main

import "fmt"

type StructVal struct {
	IntVal int
}

func (s StructVal) Add() {
	fmt.Printf("====================\n")
	fmt.Printf("Func Add\n")
	fmt.Printf("====================\n")
	fmt.Printf("Memory Location %p\n", &s)
	fmt.Printf("Value Before: %+v\n", s)
	s.IntVal++
	fmt.Printf("Value After: %+v\n\n", s)
}

func (s *StructVal) AddPtr() {
	fmt.Printf("====================\n")
	fmt.Printf("Func AddPtr\n")
	fmt.Printf("====================\n")
	fmt.Printf("Memory Location %p\n", s)
	fmt.Printf("Value Before: %+v\n", s)
	s.IntVal++
	fmt.Printf("Value After: %+v\n\n", s)
}

func main() {
	init := StructVal{
		IntVal: 0,
	}

	fmt.Printf("================\n")
	fmt.Printf("MAIN\n")
	fmt.Printf("================\n")
	fmt.Printf("Memory Location: %p\n", &init)
	fmt.Printf("Value: %+v\n\n", init)

	init.Add()

	fmt.Printf("================\n")
	fmt.Printf("AFTER Func Add()\n")
	fmt.Printf("================\n")
	fmt.Printf("Value: %+v\n\n", init)

	init.AddPtr()
	fmt.Printf("================\n")
	fmt.Printf("AFTER Func AddPtr()\n")
	fmt.Printf("================\n")
	fmt.Printf("Value: %+v\n\n", init)

	fmt.Printf("================\n")
	fmt.Printf("FINAL\n")
	fmt.Printf("================\n")
	fmt.Printf("Value: %+v\n", init)
}
