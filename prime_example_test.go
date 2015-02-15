package prime

import "fmt"

func ExampleGenerator() {

	// create a new prime generator with 2 and 3 already contained within it
	pg := NewGenerator()

	// get the current max prime and number of primes
	fmt.Println(pg.Head()) // 3 2

	// find and add the next prime
	fmt.Println(pg.Add()) // 5 3

	// check primality of numbers
	fmt.Println(pg.IsPrime(100)) // false
	fmt.Println(pg.IsPrime(101)) // true

    // note: at this point, this prime generator should contain all primes to
    // 11, but could contain more.

	// get the 1000th prime
	fmt.Println(pg.Get(1000)) // 7919

    // now we are guaranteed this prime generator has the first 1000 primes
    fmt.Println(pg.Head()) // 7919 1000

	// Output:
	// 3 2
	// 5 3
	// false
	// true
	// 7919
    // 7919 1000
}
