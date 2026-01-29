package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

/*
 * Go only has one looping construct: the for loop.
 * The basic for loop has three components, separated by semicolons:
 * - the init statement: executed before the first iteration,
 * - the condition expression: evaluated before each iteration,
 * - the post statement: executed at the end of each iteration.
 *
 * The init statement will often be short variable declaration;
 * variables declared by the init statement are visible only in the scope of the for loop.
 *
 * The loop will stop iterating when the condition expression evaluates to false.
 *
 * Unlike in C, Java, or JavaScript, there are no parentheses surrounding the three components of the for loop,
 * and the braces { } are always required.
 *
 * The init and post statements are optional (drop the semi-colons to make it function as a while loop).
 *
 * You can omit the loop condition to make it an infinite loop.
 */

/*
 * Golang's if statements are similar to its for loops; the expression doesn't need paranetheses, but
 * the braces { } are required.
 *
 * Like for, the if statement can start with a short statement to execute before the condition.
 * Variables declared by the statement are only in scope until the end of the if block.
 *
 * Variables declared inside an if block are also available inside any else block that follows it.
 */

/*
 * A switch statement is a shorter way to write if-else statements. It runs the first case whose value is equal to the switch expression.
 * Unlike in C, Java, or JavaScript, Go's switch only runs the selected case, not all the cases that follow it. In effect, the break
 * statement that's needed at the end of each case in those languages is implied in Go.
 * Another important difference is that Go's switch cases need not be constants, and the values involved don't need to be integers.
 *
 * Switch cases are evaluated from top to bottom, stopping when a case succeeds.
 *
 * A switch without a condition is the same as switch true; this construct can be a clean way to write long if-else chains.
 */

/*
 * A defer statement defers execution of a function until the surrounding function returns.
 * The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
 *
 * Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
 */

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here though

	return lim
}

/*
 * Exercise: Given a number x, find the number z for which z*z is most nearly x.
 */

func nmsqrt(x float64) float64 {
	z := 1.0 // guess

	// Newton's method
	for math.Abs(z*z-x) >= 1e-14 {
		z -= (z*z - x) / (2 * z)
	}

	// return estimate
	return z
}

func main() {
	// basic for loop
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println("Sum from 0 to 9 is:", sum)

	// for loop without init and post statements (while loop)
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}

	fmt.Println(sum2)

	// infinite loop
	// for {
	// }

	// if statement function call
	fmt.Println(sqrt(2))
	fmt.Println(sqrt(-4))

	// if statement with a short statement function call
	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(3, 3, 20))

	// exercise
	fmt.Println("nmsqrt(2):", nmsqrt(2))
	fmt.Println("nmsqrt(9):", nmsqrt(9))
	fmt.Println("nmsqrt(16):", nmsqrt(16))

	// switch statements
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	fmt.Print("Saturday is ")
	switch today := time.Now().Weekday(); today {
	case time.Saturday:
		fmt.Println("today!!!")
	case time.Saturday - 1:
		fmt.Println("tomorrow!")
	case time.Saturday - 2:
		fmt.Println("in two days.")
	default:
		fmt.Println("too far away :(")
	}

	// switcg without a condition
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	// defer statement
	defer fmt.Println("World!")

	fmt.Print("Hello, ")

	// stacked defers
	defer fmt.Println("\nDone!")
	for i := 10; i > 0; i-- {
		defer fmt.Println(i)
	}

	defer fmt.Println("Counting...")

}
