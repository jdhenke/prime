package prime

import "testing"

func TestIsPrimeForFirst10Primes(t *testing.T) {
	pg := NewGenerator()
	for _, v := range []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29} {
		if !pg.IsPrime(v) {
			t.Errorf("%v is prime but IsPrime(%v) returned false.", v, v)
		}
	}
}

func TestIsPrimeForFirst10NonPrimes(t *testing.T) {
	pg := NewGenerator()
	for _, v := range []int{4, 6, 8, 9, 10, 12, 14, 15, 16, 18} {
		if pg.IsPrime(v) {
			t.Errorf("%v is not prime, but IsPrime(%v) returned true.", v, v)
		}
	}
}

func TestNegativeInputs(t *testing.T) {
	pg := NewGenerator()
	if pg.IsPrime(-1) {
		t.Error("-1 is not prime, but IsPrime(-1) returned true.")
	}
}

func TestZeroInput(t *testing.T) {
	pg := NewGenerator()
	if pg.IsPrime(0) {
		t.Error("0 is not prime, but IsPrime(0) returned true.")
	}
}

func TestRandomPrime(t *testing.T) {
	pg := NewGenerator()
	if !pg.IsPrime(101) {
		t.Error("101 is prime, but IsPrime(101) returned false.")
	}
}

func TestRandomNonPrime(t *testing.T) {
	pg := NewGenerator()
	if pg.IsPrime(125) {
		t.Error("125 is not prime, but IsPrime(125) returned true.")
	}
}

func TestInitialHead(t *testing.T) {
	pg := NewGenerator()
	p, n := pg.Head()
	if p != 3 {
		t.Errorf("Head() on new generator should have p=3 but is %v", p)
	}
	if n != 2 {
		t.Errorf("Head() on new generator should have n=2 but is %v", n)
	}
}

func TestMultipleCallsToHead(t *testing.T) {
	pg := NewGenerator()
	for i := 0; i < 10; i += 1 {
		p, n := pg.Head()
		if p != 3 || n != 2 {
			t.Errorf("Call #%v to Head() should be (3, 2) but was (%v, %v)", i, p, n)
		}
	}
}

func TestAdd(t *testing.T) {
	pg := NewGenerator()
	for i, v := range []int{5, 7, 11, 13, 17, 19, 23, 29} {
		if p, n := pg.Add(); p != v || i+3 != n {
			t.Errorf("Call #%v to Add() should be (%v, %v) but was (%v, %v)", i, v, i+3, p, n)
		}
	}
}

func TestGet(t *testing.T) {
	pg := NewGenerator()
	if p := pg.Get(1); p != 2 {
		t.Errorf("Get(0) should be 2 but was %v", p)
	}
	if p := pg.Get(10); p != 29 {
		t.Errorf("Get(9) should be 29 but was %v", p)
	}
	if p := pg.Get(2); p != 3 {
		t.Errorf("Get(1) should be 3 but was %v", p)
	}
	if p := pg.Get(3); p != 5 {
		t.Errorf("Get(1) should be 3 but was %v", p)
	}
}

func TestSqrt(t *testing.T) {
	for _, test := range []struct {
		i int
		o int
	}{
		{1, 1},
		{2, 2},
		{3, 2},
		{4, 2},
		{5, 3},
		{6, 3},
		{7, 3},
		{8, 3},
		{9, 3},
		{10, 4},
	} {
		if s := sqrt(test.i); s != test.o {
			t.Errorf("Sqrt(%v) was %v not %v", test.i, s, test.o)
		}
	}
}
