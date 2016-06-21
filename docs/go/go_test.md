
##Go test

### 1 Test Functions
Each test file must import the testing package. Test functions have the following sig nature:

    func TestName(t *testing.T) {
    // ...
    }

e.g.

```
// Package word provides utilities for word games.
package word
// IsPalindrome reports whether s reads the same forward and backward.
// (Our first attempt.)
func IsPalindrome(s string) bool {
    for i := range s {
        if s[i] != s[len(s)1i]{
            return false
        }
    }
    return true
}
```
In the same directory, the file word_test.go contains two test functions named TestPalindrome
and TestNonPalindrome. Each checks that IsPalindrome gives the rig ht answer for a
single input and reports failures using t.Error:
```
package word
import "testing"
func TestPalindrome(t *testing.T) {
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
    if IsPalindrome("été") {
        t.Error(`IsPalindrome("été") = false`)
    }
}
func TestCanalPalindrome(t *testing.T) {
    input := "A man, a plan, a canal: Panama"
    if IsPalindrome(input) {
        t.Errorf(`IsPalindrome(%q) = false`, input)
    }
}
```

Run the following command:

    go test

output the result :

    PASS
    coverage: 3.3% of statements
    ok  	basic	0.130s

Satisfied, we ship the program, but no sooner have the launch party guests departed than the
bug reports start to arrive.

As a bonus, running go test is usually quicker than manually going through the steps
described in the bug rep ort, allowing us to iterate more rapid ly. If the test suite contains many
slow tests, we may make even faster progress if we’re selective about which ones we run.
The **-v** flag prints the name and execution time of each test in the package:

    go test v
=== RUN   TestIsPalindrome
--- PASS: TestIsPalindrome (0.00s)
=== RUN   TestNonPalindrome
--- PASS: TestNonPalindrome (0.00s)
=== RUN   TestFrenchPalindrome
--- PASS: TestFrenchPalindrome (0.00s)
=== RUN   TestCanalPalindrome
--- PASS: TestCanalPalindrome (0.00s)
PASS
ok      basic   0.093s




The **-run** flag , whos e argument is a regular expression, causes go test to run only those
tests whose function name matches the pattern:

    go test -v -run="IsPalindrome|Canal"
    === RUN   TestIsPalindrome
    --- PASS: TestIsPalindrome (0.00s)
    === RUN   TestCanalPalindrome
    --- PASS: TestCanalPalindrome (0.00s)
    PASS
    ok      basic   0.128s

Of course, once we’ve gotten the selected tests to pass, we should invoke go test with no flags
to run the entire test suite one last time before we commit the change.

#### 1.1 Randomized Testing
Table-driven tests are convenient for checking that a function works on inputs carefully
selected to exercise interesting cases in the log ic. Another approach, randomized testing,
explores a broader range of inputs by constructing inputs at random.

The example below uses the one approach: the randomPalindrome function generates
words that are known to be palindromes by construction.
e.g.


    func TestRandomPalindromes(t *testing.T) {
        // Initialize a pseudorandom number generator.
        seed := time.Now().UTC().UnixNano()
        t.Logf("Random seed: %d", seed)
        rng := rand.New(rand.NewSource(seed))
        for i := 0; i < 1000; i++ {
            p := randomPalindrome(rng)
            if !IsPalindrome(p) {
                t.Errorf("IsPalindrome(%q) = false", p)
            }
        }
    }

test case

    func TestRandomPalindromes(t *testing.T) {
        // Initialize a pseudorandom number generator.
        seed := time.Now().UTC().UnixNano()
        t.Logf("Random seed: %d", seed)
        rng := rand.New(rand.NewSource(seed))
        for i := 0; i < 1000; i++ {
            p := randomPalindrome(rng)
            if !IsPalindrome(p) {
                t.Errorf("IsPalindrome(%q) = false", p)
            }
        }
    }


#### 1.2 Testing a Command
The go test tool is useful for testing library packages, but with a little effort we can use it to
test commands as well. A package named main ordinarily produces an exec utable program,
but it can be imported as a library too.


### 2 profile
Go supports many kinds of profiling, each concerned with a different aspect of performance,
but all of them involve recording a sequence of events of interest, each of which has an accompanying
stack trace—the stack of function cal ls active at the moment of the event. The
go test tool has bui lt-in support for several kinds of profiling.

