package main

import (
	"fmt"
	"sort"
)

type Hero struct {
	Name string
	Age  int
}

type HeroSlice []Hero

func (h HeroSlice) Len() int           { return len(h) }
func (h HeroSlice) Less(i, j int) bool { return h[i].Age < h[j].Age }
func (h HeroSlice) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func main() {
	fmt.Println("hello world")

	var slice = []int{0, -110, 10, 7}
	sort.Ints(slice)
	fmt.Println(slice)

	// 结构体切片排序
	heroSlice := HeroSlice{{Name: "zxc", Age: 12}, {Name: "zx", Age: 13}}
	sort.Sort(heroSlice)
	fmt.Println(heroSlice)
	sort.Sort(sort.Reverse(heroSlice))
	fmt.Println(heroSlice)

}
