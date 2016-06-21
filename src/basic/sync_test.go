package basic_test

import (
	. "basic"
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestBank(t *testing.T) {
	// Deposit [1..1000] concurrently.
	var n sync.WaitGroup
	for i := 1; i <= 1000; i++ {
		n.Add(1)
		go func(amount int) {
			Deposit(amount)
			n.Done()
		}(i)
	}
	n.Wait()

	if got, want := Balance(), (1000+1)*1000/2; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
	for i := -1000; i <= -1; i++ {
		n.Add(1)
		go func(amount int) {
			Deposit(amount)
			n.Done()
		}(i)
	}
	n.Wait()

	if got, want := Balance(), (-1000-1)*1000/2+(1000+1)*1000/2; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}

func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome("detartrated") {
		t.Error(`IsPalindrome("detartrated") = false`)
	}
	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak") = false`)
	}
}

func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("palindrome") {
		t.Error(`IsPalindrome("palindrome") = true`)
	}
}

func TestFrenchPalindrome(t *testing.T) {
	if !IsPalindrome("été") {
		t.Error(`IsPalindrome("été") = false`)
	}
}
func TestCanalPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input) {
		t.Errorf(`IsPalindrome(%q) = false`, input)
	}
}

//good test cases
func Test4IsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false}, // nonpalindrome
		{"desserts", false},   // semipalindrome
	}
	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}

func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudorandom number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := RandomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}

}

func TestReflection(t *testing.T) {
	v := reflect.ValueOf(3) // a reflect.Value
	fmt.Println(v)          // "3"
	fmt.Printf("%v\n", v)   // "3"
	fmt.Println(v.String()) // NOTE: "<int Value>"

	tx := reflect.TypeOf(3)  // a reflect.Type
	fmt.Println(tx.String()) // "int"
	fmt.Println(tx)          // "int"

	//The inverse operat ion to reflect.ValueOf is the reflect.Value.Interface method. It
	//returns an interface{} holding the same concrete value as the reflect.Value:
	v = reflect.ValueOf(3) // a reflect.Value
	x := v.Interface()     // an interface{}
	i := x.(int)           // an int
	fmt.Printf("%d\n", i)  // "3"
}

func TestAny(t *testing.T) {
	var x int64 = 1
	var d time.Duration = 1 * time.Nanosecond
	fmt.Println(Any(x))
	fmt.Println(Any(d))
	fmt.Println(Any([]int64{x}))
	fmt.Println(Any([]time.Duration{d}))
}
