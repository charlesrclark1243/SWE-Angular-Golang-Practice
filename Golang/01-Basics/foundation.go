// programs start running in the main package.
package main

// import necessary packages using import paths.
//
// by convention, package name is the same as the last element of the import path.
// for example, package "math/rand" has the package name "rand".
// import (
// 	"fmt"
// 	"math/rand"
// )

// you can also write multiple import statements like this:
// import "fmt"
// import "math"

// a name is exported if it begins with a capital letter.
// when importing a package, you can refer only to its exported names.
// any "unexported" names aren't accessible outside the package (like "private" or "protected" in other languages).

// functions can take >= 0 arguments.
// the function `add` below takes two parameters of type "int".
// notice the type comes after the variable name.

import "fmt"

func add(x int, y int) int {
	return x + y
}

// when two or more consecutive parameters share the same type, you can omit the
// type from all but the last parameter. See the function `subtract` below.

func subtract(x, y int) int {
	return x - y
}

// a function can return any number of results. The function `swap` below retturns two strings.
func swap(x, y string) (string, string) {
	return y, x
}

// return values can be named. If so, they are treated as variables defined at the top of the function.
// these names should be used to document the meaning of the return values.
// a return statement without arguments returns the named return values and is called a
// "naked" return.
// naked return statements should only be used in short functions, as in the example below;
// they can harm readability in longer functions.

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	return
}

// the var statement declares a list of variables; the type is last.
// a var statement can be at the package or function level.

var packageLevelVar int // declared at package level

// a var declaration can include an initial value, which will be assigned to the variable.
// if an initial value is present, the type can be omitted; the variable will take the type of the initial value.

// inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
// outside a function, every statement begins with a keyword (var, func, etc.) so the := construct is not available.

// the basic types are:
//  - bool
//  - string
//  - int, int8, int16, int32, int64
//  - uint, uint8, uint16, uint32, uint64, uintptr
//  - byte (alias for uint8)
//  - rune (alias for int32, represents a Unicode code point)
//  - float32, float64
//  - complex64, complex128
// int, uint, and uintptr are architecture dependent (32 or 64 bits)
// it's recommended to use int unless you have specific reasons to use a sized or unsigned integer type.

/*
* Variables declared without being initialized are given a zero value:
* - 0 for numeric types
* - false for the boolean type
* - "" (the empty string) for strings
 */

/*
* The experession T(v) converts value v to type T.
* unlike in C, Go assignment between items of different types require explicit conversion.
*
 */

/*
 * When declaring a variable without specifying an explicit type (either using := or var =), the variable's type is inferred
 * from the value on the right-hand side.
 *
 * When the right-hand sidde is typed, the new variable is of the same type.
 * When the right-hand side contains an untyped numeric constant (i.e., a number without an explicit type), the new variable
 * may be an int, float64, or complex128, depending on the precision of the constant.
 */

/*
 * Constants are declared like variables, but using const keyword.
 * Constants can be character, string, boolean, or numeric values.
 * Constants cannot be declared using the := syntax.
 */

/*
 * Numeric constants are high-precision values.
 * An untyped constant takes the type needed by its context.
 */
const (
	big   = 1 << 100  // shift left 1 by 100 places (binary number with 1 followed by 100 zeros)
	small = big >> 99 // shift it right again 99 places, resulting in 2 (10, or 1 << 1).
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(add(42, 13))      // add function
	fmt.Println(subtract(42, 13)) // subtract function

	a, b := swap("hello", "world") // swap function
	fmt.Println(a, b)

	fmt.Println(split(17))       // split function
	fmt.Println(packageLevelVar) // package level variable

	var functionLevelVar = 42     // declared and initialized at function level
	fmt.Println(functionLevelVar) // function level variable

	short := "short variable declaration" // short variable declaration
	fmt.Println(short)

	// data types
	var (
		toBe   bool       = false
		maxInt uint       = 1<<64 - 1
		z      complex128 = complex(1, 2)
	)
	fmt.Printf("Type: %T Value: %v\n", toBe, toBe)
	fmt.Printf("Type: %T Value: %v\n", maxInt, maxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// uninitialized variables
	var i int
	var f float64
	var boolean bool
	var s string
	fmt.Printf("Zero values - int: %d, float64: %f, bool: %t, string: '%s'\n", i, f, boolean, s)

	// type conversion
	var j int = 42
	var floating float64 = float64(j)
	var u uint = uint(floating)

	// simplified type conversions
	k := 25
	floating2 := float64(k)
	u2 := uint(floating2)

	fmt.Printf("Type conversion - int: %d, float64: %f, uint: %d\n", j, floating, u)
	fmt.Printf("Simplified type conversion - int: %d, float64: %f, uint: %d\n", k, floating2, u2)

	// right-hand typed
	var m int = 27
	n := m //  n is an int

	// right-hand untyped
	q := 56           // q is an int
	p := 3.142        // p is a float64
	g := 0.867 + 0.5i // g is a complex128

	fmt.Printf("Right-hand typed - m: %d, n: %d\n", m, n)
	fmt.Printf("Right-hand untyped - q: %d, p: %f, g: %v\n", q, p, g)

	// constants
	const hello = "Hello, World!"
	fmt.Println(hello)

	const truth = true
	fmt.Println("Go is a language?", truth)

	// numeric constants
	fmt.Println(needInt(small))
	// fmt.Println(needInt(big)) // needInt(big) would cause an error: the constant overflows int
	fmt.Println(needFloat(small))
	fmt.Println(needFloat(big))
}
