package main

// https://0a.io/chapter1/calculus-explained.html
// https://0a.io/chapter1/set-theory-and-axiomatic-systems-explained.html
func main() {

	// function is a black box, that takes in a input and gives back you a output
	// we use to map numbers to others numbers
	f(10) // gives 20

	// limit is the number you expect to get from a function or algebraic expression
	// is referring to the expectation of the ouput when X approaches a certain value

	// lim x→3 f(x) = 0
	// lim f(x) = 0 when x approaches 3

	// set is a collection of things, the things are the elements of the set
	// { 0, 1, 2, 3 } set of integers
}

// the set of numbers that you can put into a function
// is called domain of the function
func f(x int) int {
	return x + 10
}