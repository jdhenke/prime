// Package prime defines an interface and implementation of a construct that is
// useful for finding primes and determing primality.
package prime

// A Generator finds primes in ascending order and determines numbers' primality.
type Generator interface {

	// Returns p, the largest prime and n, the number of primes this generator
	// contains.
	Head() (p, n int)

	// Finds and adds the next prime to this Generator. Returns p, the prime that
	// was added and n, the number of primes the generator now contains.
	Add() (p, n int)

	// Returns the nth prime, one indexed, e.g. Get(3) returns 5. Guarantees the
	// first n primes will be added to this generator. n must be >= 1.
	Get(n int) (p int)

	// Returns true iff n is prime. Guarantees all primes <= sqrt(n) will be added
	// to this generator.
	IsPrime(n int) bool
}

// Returns a generator with primes 2 and 3 already added.
func NewGenerator() Generator {
	return &generator{
		primesMap:  map[int]bool{2: true, 3: true},
		primesList: []int{2, 3},
	}
}

type generator struct {
	primesMap  map[int]bool
	primesList []int
}

func (g *generator) Head() (p, n int) {
	n = len(g.primesList)
	p = g.primesList[n-1]
	return
}

func (g *generator) Add() (p, n int) {
	x, _ := g.Head()
	for x += 2; !g.IsPrime(x); x += 2 {
	}
	g.primesMap[x] = true
	g.primesList = append(g.primesList, x)
	return g.Head()
}

func (g *generator) Get(n int) (p int) {
	for _, nPrimes := g.Head(); nPrimes < n; _, nPrimes = g.Add() {
	}
	return g.primesList[n-1]
}

func (g *generator) IsPrime(x int) bool {
	if isPrime, ok := g.primesMap[x]; ok {
		return isPrime
	}
	if x <= 1 {
		return false
	}
	s := sqrt(x)
	for p, _ := g.Head(); p < s; p, _ = g.Add() {
	}
	for i := 0; i < len(g.primesList) && g.primesList[i] <= s; i += 1 {
		if p := g.primesList[i]; x%p == 0 {
			g.primesMap[p] = false
			return false
		}
	}
	g.primesMap[x] = true
	return true
}

func sqrt(x int) int {
	l, r := 0, x+3
	m := (l + r) / 2
	for !(m*m >= x && (m-1)*(m-1) < x) {
		if m*m < x {
			l = m
		} else {
			r = m
		}
		m = (l + r) / 2
	}
	return m
}
