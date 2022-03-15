package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func add2(a, b int) int {
	return a + b
}

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	v4 = &Vertex{1, 2}
)

func main() {
	fmt.Println(add(37, 25))
	fmt.Println(add2(37, 25))

	//struct
	v := Vertex{1, 2}
	p := &v
	p.X = 4

	fmt.Println(v)

	fmt.Println(v1, v2, v3, v4)

	//arrays
	var a [2]string
	a[0] = "bla"
	a[1] = "ble"

	fmt.Println(a[0], a[1])

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	//slices
	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"John",
		"paul",
		"george",
		"ringo",
	}

	fmt.Println(names)

	c := names[0:2]
	d := names[1:3]
	fmt.Println(c, d)

	d[0] = "XXX"
	fmt.Println(c, d)
	fmt.Println(names)

	//slice literals
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}

	sl := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
	}

	fmt.Println(r, sl)
	// nil slice
	var s1 []int
	fmt.Println(s1, s1 == nil)
	fmt.Println(s1, len(s1), cap(s1))

	// make function
	a1 := make([]int, 5)
	fmt.Println(a1, len(a1), cap(a1))

	a1 = make([]int, 0, 5)

	//extends slice
	a1 = a1[:2]
	fmt.Println(a1, len(a1), cap(a1))

	var pow = []int{1, 2, 3}

	for i, v := range pow {
		fmt.Println(i, v)
	}

	//maps
	var map1 = make(map[string]Vertex)

	//check if exists
	if _, ok := map1["bla"]; ok {

	}

	map1["bla"] = Vertex{
		1,
		2,
	}

	fmt.Println(map1)

	var map2 = map[string]Vertex{
		"ble": Vertex{
			1, 2,
		},
	}
	fmt.Println(map2)
	var map3 = map[string]Vertex{
		"blo": {
			1, 2,
		},
	}
	fmt.Println(map3)
}
