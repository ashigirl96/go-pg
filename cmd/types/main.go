package main

import (
	"fmt"
	"github.com/ashigirl96/go-pg/pkg"
	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
	"strings"
)

func pointers() {
	pkg.FuncName()
	i := 42
	p := &i
	fmt.Println(p)
	fmt.Println(*p + 1)
}

func structs() {
	pkg.FuncName()
	type Vertex struct {
		X int
		Y int
	}
	fmt.Println(Vertex{})
	fmt.Println(Vertex{1, 2})
	fmt.Println(Vertex{X: 3, Y: 4})

	x := struct {
		X int
		Y int
	}{
		X: 1,
		Y: 2,
	}
	fmt.Println(x)
	p := &x
	p.X = 100
	fmt.Println(x)
}

func slices() {
	pkg.FuncName()
	// def array
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// スライスは配列への参照みたいなもの
	s := primes[1:4]
	s[0] = 100
	fmt.Println(s, primes)
}

func sliceLiterals() {
	pkg.FuncName()
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
	}
	fmt.Println(s)
}

func sliceLengthAndCapacity() {
	pkg.FuncName()
	printSlice := func(s []int) {
		fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	}
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)
	s = s[:0]
	printSlice(s)
	s = s[:4]
	printSlice(s)
	s = s[2:]
	printSlice(s)
	s = s[:4]
	printSlice(s)
}

func appendSlice() {
	pkg.FuncName()
	printSlice := func(s []int) {
		fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	}
	var s []int
	printSlice(s)
	s = append(s, 0)
	printSlice(s)
	s = append(s, 1)
	printSlice(s)
	s = append(s, 1, 2, 3)
	printSlice(s)
}

func ranges() {
	pkg.FuncName()
	pow := []int{1, 2, 4, 8}
	for i, v := range pow {
		fmt.Println(i, v)
	}
}

func exerciseSlices() {
	pkg.FuncName()
	Pic := func(dx, dy int) [][]uint8 {
		result := make([][]uint8, dx, dy)
		for i := range result {
			for j := range result {
				result[i] = append(result[i], uint8(i*j))
			}
		}
		return result
	}
	pic.Show(Pic)
}

func maps() {
	pkg.FuncName()
	type Vertex struct {
		Lat, Long float64
	}
	// map はキーと値とを関連付けます(マップします)。
	var m map[string]Vertex

	m = make(map[string]Vertex)
	n := make(map[string]Vertex)
	m["Labs"] = Vertex{40.1, -784.3}

	// ↓エラーになる
	//v := &m["Labs"]
	v := m["Labs"]
	v.Long = 10
	fmt.Println(m, n, v)
}

func mapLiterals() {
	pkg.FuncName()
	type Vertex struct {
		Lat, Long float64
	}
	m := map[string]Vertex{
		"Bell": {
			Lat:  1,
			Long: 2,
		},
		"Google": {
			Lat:  3,
			Long: 4,
		},
	}
	fmt.Println(m)
}

func mutateMaps() {
	pkg.FuncName()
	m := map[string]int{"x": 3, "y": 4}
	n := make(map[string]int)
	n["x"] = 5
	fmt.Println(m, n)
	delete(m, "x")
	v, ok := m["x"]
	fmt.Println(m, n, v, ok)
}

func exerciseMaps() {
	wordCount := func(s string) map[string]int {
		result := make(map[string]int)
		for _, s := range strings.Split(s, " ") {
			result[s] += 1
		}
		return result
	}
	wc.Test(wordCount)
}

func exerciseFib() {
	pkg.FuncName()
	fib := func() func() int {
		i, j := 1, 0
		return func() int {
			tmp := j
			j += i
			i = tmp
			return i
		}
	}

	f := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func main() {
	pointers()
	structs()
	slices()
	sliceLiterals()
	sliceLengthAndCapacity()
	appendSlice()
	ranges()
	//exerciseSlices()
	maps()
	mapLiterals()
	mutateMaps()
	//exerciseMaps()
	exerciseFib()
}
