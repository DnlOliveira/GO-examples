package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {

	// Struct - Used for holding values

	p1 := Person{} //initializes with default values "" for string, 0 for int
	p2 := Person{"Daniel", "Oliveira", 26}
	p3 := Person{LastName: "Oliveira", Age: 26} //explicitly specifying in case one of the values are to be skipped

	fmt.Println("Structs ------")
	fmt.Println("Person 1 :", p1)
	fmt.Println("Person 2 :", p2)
	fmt.Println("Person 3 :", p3)
	fmt.Println("--------------")

	// Map - The only collection with key:value pairs supported by Go

	m := make(map[int]string) // int as key, string as value
	// m := make(map[int]string, 10) // can also pass an initial capacity

	m[0] = "ABC"
	m[1] = "DEF"
	m[2] = "GHI"

	fmt.Println("Maps ------")
	fmt.Println("Length of map: ", len(m))
	fmt.Println(m)

	m[0] = "YZX" // updating value

	for k, v := range m {
		fmt.Println("Key:", k, " Value:", v)
	}
	fmt.Println("--------------")

	// Arrays - Holding elements of the same type
	a := [5]int{}
	fmt.Println("Arrays  ------")
	fmt.Println("Array a: ", a)
	fmt.Println("Array length: ", len(a))

	for i, _ := range a {
		a[i] = i
	}
	fmt.Println("Array after inserting data: ", a)

	b := [...]string{"A", "B", "C"}
	fmt.Println("Array b: ", b)
	fmt.Println("--------------")

	// Slice - Special data type that wraps on top of array

	// s := make([]int, length, capacity)
	// length: define the length of the slice
	// capacity: define the capacity of the underlying array - optional

	// -- Declaration
	//s := make(int[], 5, 10)
	// s := []int{0,1,2,3} another way to declare slices similar to arrays

	a2 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := a2[0:5]

	fmt.Println("Slice  ------")
	fmt.Println("Underlying Array: ", a2)
	fmt.Println("slice built on top of array: ", s)
	fmt.Println("--------------")

	x := a2[:]
	fmt.Println("Array: ", a2)
	fmt.Println("Slice: ", x)

	x = append(x, 11)

	fmt.Println("Array - updated: ", a2)
	fmt.Println("Slice - updated: ", x)
	fmt.Println("Array length- updated: ", len(a2))
	fmt.Println("Slice length - updated: ", len(x))
	fmt.Println("Array capacity - updated: ", cap(a2))
	fmt.Println("Slice capacity - updated: ", cap(x))

}
