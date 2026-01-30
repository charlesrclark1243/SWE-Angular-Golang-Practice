package main

import (
	"fmt"
	"math"
	"strings"
)

/*
 * Go has pointers which hold the memory address of a value.
 * The type *T is a pointer to a T value; its zero value is nil.
 *
 * The & operator generates a pointer to its operand.
 * The * operator denotes the pointer's underlying value. This is called dereferencing, or indirecting.
 *
 * Unlike in C, Go has no pointer arithmetic.
 */

/*
 * A struct is a collection of fields.
 *
 * Fields can be accessed using a dot.
 * Struct fields can be accessed through a struct pointer.
 * To access the field X of a struct with struct pointer p, we could write (*p).X, but we can also write p.X (as if p were a struct value itself).
 *
 * A struct literal denotes a newly allocated struct value by listing the values of its fields.
 * You can list just a subset of fields by using the Name: syntax (the order of the named fields is irrelevant).
 * The special prefix & returns a pointer to the struct value.
 */

/*
 * The type [n]T is an array of n values of type T.
 * The expression var a[10]int declares an array of ten integers.
 * An array's length is part of its type, so they can't be resized.
 *
 * A slice is a dynamically-sized, flexible view into the elements of an array.
 * The type []T is a slice with elements of type T.
 * A slice is formed by specifying two indices, a low and high bound, separated by a colon: a[low : high].
 * This selects a half-open range which includes the first element, but excludes the last one.
 *
 * Slices do not store any data, they just describe a section of an underlying array.
 * Changing the elements of a slice modifies the corresponding elements of its underlying array.
 * Other slices that share the same underlying array will see those changes.
 *
 * A slice literal is like an array literal without the length.
 *
 * When slicing, you can omit the high or low bounds to use their defaults instead.
 * The default low bound is 0, and the default high bound is the length of the slice.
 *
 * A slice has both a length and a capacity.
 * The length of a slice is the number of elements it contains.
 * The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
 * The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).
 * You can extend a slice's length by re-slicing it, provided it has sufficient capacity.
 *
 * The zero value of a slice is nil.
 * A nil slice has a length and capacity of 0 and has no underlying array.
 *
 * Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
 * The make function allocates a zeroed array and returns a slice that refers to that array.
 * The make function takes three arguments: the type of the slice to create, its length, and its capacity.
 *
 * Slices can contain any type, including other slices.
 *
 * It's common to append new elements to a slice, and so Go provides a built-in append function.
 * The first parameter is a slice of type T, and the rest are T values to append to the slice.
 * The resulting value of append is a slice containing all the elements of the original slice plus the provided values.
 * If the backing array of s is too small to fit all the given values a bigger array will be allocated; the returned
 * slice will point to the newly allocated array.
 *
 * More about slices here: https://go.dev/blog/go-slices-usage-and-internals
 */

/*
 * The range form of the for loop iterates over a slice or a map.
 * When ranging over a slice, two values are returned for each iteration: the first is the index, and the second
 * is a copy of the element at that index (like Python's enumerate).
 *
 * You can skip the index or value by assigning it to _.
 * If you only want the index, you can omit the second variable.
 */

func printSlice(slice []int) {
	fmt.Printf("\nlen = %d, cap = %d %v", len(slice), cap(slice), slice)
}

type Vertex struct {
	X int
	Y int
}

/**
 * Exercise - Implement the Pic function. It should return a slice of length dy, each element of which is a slice of dx
 * 8-bit unsigned integers. When you run the program, it will display your picture, interpreting the integers as grayscale
 * (well, bluescale) values.
 *
 * The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.
 */

func Pic(dx, dy int) [][]uint8 {
	var pic [][]uint8 = make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			pic[y] = append(pic[y], uint8((x*x + y*y)))
		}
	}

	return pic
}

/*
 * A map maps keys to values (like a Python dictionary).
 * The zero value of a map is nil.
 * A nil map has no keys, nor can keys be added.
 * To create a map, use the built-in make function: m := make(map[string]int)
 *
 * Map literals are like struct literals, but the keys are required.
 * If the top-level type is just a type name, you can omit it from the elements of the literal.
 *
 * Inserting or updating a key/value pair in a map is done with the syntax: m[key] = value
 * Retrieving a value for a key is done with the syntax: value = m[key]
 * Deleting a key/value pair is done with the built-in delete function: delete(m, key)
 *
 * Test that a key is present with a two-value assignment: value, ok = m[key]
 * If key is in m, ok is true. If not, ok is false.
 * If key is not in the map, then elem is the zero value for the map's eleemnt type.
 */

type Coord struct {
	Lat, Long float64
}

/*
 * Functions are values too which can be passed around like other values.
 * Function values may be used as function arguments and return values.
 *
 * Some functions may be closures, which are function values that reference variables from outside their body.
 * The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

 */

func adder() func(int) int {
	sum := 0 // captured variable
	return func(x int) int {
		sum += x
		return sum
	}
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func fibonnaci() func() int {
	a, b := 0, 1

	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	// pointers
	i, j := 42, 27

	p := &i // point to i
	*p = 21 // set i through the pointer
	fmt.Println("i:", i)

	p = &j      // point to j
	*p = *p * 2 // set j through the pointer
	fmt.Println("j:", j)

	// structs
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println("v:", v)

	// struct pointers
	vp := &v
	vp.Y = 5
	fmt.Println("v':", v)

	// struct literals
	v1 := Vertex{1, 2}  // has type Vertex
	v2 := Vertex{X: 1}  // Y:0 is implicit
	v3 := Vertex{}      // X:0 and Y:0
	p4 := &Vertex{1, 2} // has type *Vertex

	fmt.Println("v1:", v1, "v2:", v2, "v3:", v3, "p4:", p4)

	// arrays
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println("array:", a)
	fmt.Println("array length:", len(a))

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("primes:", primes)

	// slices
	var s []int = primes[1:4]
	fmt.Println("primes slice:", s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println("\nnames array:", names)

	b := names[0:2]
	c := names[1:3]
	fmt.Println("b slice:", b)
	fmt.Println("c slice:", c)

	c[0] = "XXX"
	fmt.Println("\nafter modification...")
	fmt.Println("names array:", names)
	fmt.Println("b slice:", b)
	fmt.Println("c slice:", c)

	// slice literals
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("\nslice literal q:", q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println("slice literal r:", r)

	// slices continued
	var d [10]int
	d1 := d[0:10]
	d2 := d[:10]
	d3 := d[0:]
	d4 := d[:]

	// d === d1 === d2 === d3 === d4
	fmt.Println("\nd:", d)
	fmt.Println("d1:", d1)
	fmt.Println("d2:", d2)
	fmt.Println("d3:", d3)
	fmt.Println("d4:", d4)

	t := []int{2, 3, 5, 7, 11, 13}
	printSlice(t)

	t = t[:0]
	printSlice(t)

	t = t[:4]
	printSlice(t)

	t = t[2:]
	printSlice(t)

	var nilSlice []int
	fmt.Println("\n\nnilSlice:", nilSlice, len(nilSlice), cap(nilSlice))

	// slices with make
	e := make([]int, 5)    // len(e) = 5
	f := make([]int, 0, 5) // len(f) = 0, cap(f) = 5

	fmt.Println("\ne:", e, len(e), cap(e))
	fmt.Println("f:", f, len(f), cap(f))

	f = f[:cap(f)] // len(f) = 5, cap(f) = 5
	fmt.Println("f:", f, len(f), cap(f))

	f = f[1:]
	fmt.Println("f:", f, len(f), cap(f), "\n")

	// slices of slices
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// appending to a slice
	var s2 []int
	printSlice(s2)

	s2 = append(s2, 0)
	printSlice(s2)

	s2 = append(s2, 1, 2, 3, 4)
	printSlice(s2)

	fmt.Println()

	// range looping

	var pow []int = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for idx, el := range pow {
		fmt.Printf("\n2**%d = %d", idx, el)
	}

	fmt.Println()

	var pow2 []int = make([]int, 10)
	for idx := range pow2 {
		pow2[idx] = 1 << uint(idx) // == 2**i
	}

	for _, el := range pow2 {
		fmt.Printf("\n%d", el)
	}

	fmt.Println()

	// exercise
	picture := Pic(8, 8)

	for _, line := range picture {
		fmt.Println(line)
	}

	// maps
	var Map map[string]Coord = make(map[string]Coord)
	Map["Bell Labs"] = Coord{
		40.68433, -74.39967,
	}

	fmt.Println("\nmap:", Map)

	// map literals
	Map2 := map[string]Coord{
		"Google": Coord{
			37.42202, -122.08408,
		},
		"Apple": Coord{
			37.33182, -122.03118,
		},
	}

	Map3 := map[string]Coord{
		"New York City":   {40.71278, -74.00594},
		"Los Angeles":     {34.05223, -118.24368},
		"Chicago":         {41.87811, -87.62980},
		"Houston":         {29.76043, -95.36980},
		"Philadelphia":    {39.95233, -75.16379},
		"Pittsburgh":      {40.44062, -79.99589},
		"San Francisco":   {37.77493, -122.41942},
		"Washington D.C.": {38.90719, -77.03687},
	}

	fmt.Println("map literal:", Map2)
	fmt.Println("map literal 2:", Map3)

	// mutating maps
	Map4 := make(map[string]int)

	Map4["Answer"] = 42
	fmt.Println("\nThe value:", Map4["Answer"])

	Map4["Answer"] = 48
	fmt.Println("The value:", Map4["Answer"])

	delete(Map4, "Answer")
	fmt.Println("The value:", Map4["Answer"])

	val, ok := Map4["Answer"]
	fmt.Println("The value:", val, "Present?", ok, "\n")

	// functions
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))

	fmt.Println(compute(math.Pow), "\n")

	// closures
	var pos, neg func(int) int = adder(), adder()

	for i := 0; i < 10; i++ {
		fmt.Println(
			i,
			"->",
			"pos:", pos(i),
			"neg:", neg(-i),
		)
	}

	fmt.Println()

	// Fibonacci closure
	var fib func() int = fibonnaci()
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}

}
