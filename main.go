package main

import "fmt"

func eiei() {
	//  declare variables
	gopher := "Gopher"
	fmt.Printf("eioei, %s.\n", gopher)
}

func sw() {
	var score int
	fmt.Scanln(&score)

	switch {
	case score < 50:
		fmt.Print("F")
	case score < 60:
		fmt.Printf("D")
	case score < 70:
		fmt.Printf("C")
	case score < 80:
		fmt.Printf("B")
	case score > 80:
		fmt.Printf("A")
	}
}

func main() {
	// Normal Arrays
	var a [5]int
	a[0] = 10
	a[2] = 20
	a[3] = 14

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	for i := range a {
		fmt.Println(a[i])
	}

	for _, v := range a {
		fmt.Println(v)
	}
	// 2D Arrays
	var b [2][3]int
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			b[i][j] = j
		}
	}
	fmt.Println(b)

	// Slice
	c := []int{1, 2, 3, 4, 5, 6, 7}
	c = append(c, 60)
	fmt.Println(c)
	fmt.Println(c[2:4]) // print index[2] to index[4]
	fmt.Println(c[:4])  // print index[0] to index[4]
	fmt.Println(c[4:])  // print index[4] plus

	// slice in slice array

	x := c[2:]
	fmt.Println(x[0])

	// Map, hash, dict
	g := make(map[string]string)
	g["hello"] = "gopher"
	g["name"] = "access"
	g["date"] = "Today"

	d := map[string]string{
		"id":   "afag55agag",
		"name": "Gopher",
	}

	fmt.Println(d)

	date, present := g["date"] // check date in hash["date"] present ?, if present return true

	if date, present := g["date"]; present {
		fmt.Println(date)
	}

	for key, value := range g {
		fmt.Println(key, ":", value)
	}

	fmt.Println(present)
	fmt.Println(date)
	fmt.Println(g["hello"])
	fmt.Println(g["hello"])

	fmt.Println("*******POINTER***********")
	// pointer

	aa := 10

	ptrA := &aa
	fmt.Println(ptrA)
	fmt.Println(*ptrA)

	zz, yy := 1, 2
	r := add(zz, yy)
	fmt.Println(r)

	zaa := [5]int{}
	fmt.Println(zaa)
	mutateArray(zaa[0:len(zaa)])
	mutateArray(zaa[:])
	fmt.Print(zaa)

	p1 := person{
		Id:   "oisjkakaf",
		Name: "Panuwat",
	}
	fmt.Println(p1)

}

// Struct

type person struct {
	Id   string
	Name string
}

// pass by value
func mutateArray(a []int) {
	a[0] = 10
}

func add(val1, val2 int) int {
	return val1 + val2
}
