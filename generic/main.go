package main

import "fmt"

type MyGenericConstraint interface {
	string | int | float64
}

func main() {
	d1 := cetak(1)
	fmt.Println(d1)

	d2 := cetak("Hello")
	fmt.Println(d2)

	cetak(1.4)

	// d := cetak2(1)
	// d.(int64)

	fmt.Println(isEqual(1, 1))         // true
	fmt.Println(isEqual("abc", "abc")) // true

	s1 := []string{"satu", "dua"}
	s2 := []string{"tiga", "empat"}
	sMerged := mergeSlice(s1, s2)
	fmt.Println(sMerged)

	m := map[string]int{
		"satu": 1,
		"dua":  2,
	}
	v := findValueFromMap(m, "satu")
	fmt.Println(v)
}

func cetak[T MyGenericConstraint](v T) T {
	return v
}

func cetak2(v any) any {
	return v
}

func isEqual[T comparable](arg1 T, arg2 T) bool {
	return arg1 == arg2
}

func mergeSlice[T comparable](s1 []T, s2 []T) []T {
	return append(s1, s2...)
}

func findValueFromMap[K comparable, V comparable](m map[K]V, key K) V {
	return m[key]
}
