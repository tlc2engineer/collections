package main

import "fmt"

func main() {

	fmt.Println("go")
	d := []int{1, 4, 35, -1}
	f := func(f float64, i int) float64 {
		return f + float64(i)
	}
	fmt.Println(reduce(d, f, 5.0))
	fmt.Println(filter(d, func(i int) bool {
		return i > 1
	}))
	list := MakeList[int]()
	list.add(1)

	lint := last[int]
	print(lint(d))

}

func last[T any](data []T) *T {
	if len(data) == 0 {
		return nil
	}
	return &data[len(data)-1]

}

func reduce[T any, M any](data []T, f func(m M, t2 T) M, beg M) M {
	if len(data) == 0 {
		return beg
	}
	m := beg
	for _, t := range data {
		m = f(m, t)
	}
	return m
}

func filter[T any](data []T, f func(T) bool) []T {
	out := make([]T, 0)
	for _, t := range data {
		if f(t) {
			out = append(out, t)
		}
	}
	return out
}

func map_[T, M any](data []T, f func(T) M) []M {
	out := make([]M, 0, len(data))
	for _, t := range data {
		out = append(out, f(t))
	}
	return out
}

type List[T any] struct {
	data []T
}

func (l List[T]) add(el T) {
	l.data = append(l.data, el)
}

func (l List[T]) get(i int) T {
	return l.data[i]
}

func MakeList[T any]() List[T] {
	list := make([]T, 0)
	return List[T]{
		data: list,
	}
}
