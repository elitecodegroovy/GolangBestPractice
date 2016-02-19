package strcon

import (
	"bytes"
	"fmt"
	"time"
	"unicode"
)

type Person struct {
	FirstName, LastName string
	Dob                 time.Time
	Email, Location     string
}

//A person method
func (p Person) PrintName() {
	fmt.Printf("\n%s %s\n", p.FirstName, p.LastName)
}

//A person method
func (p Person) PrintDetails() {
	fmt.Printf("[Date of Birth: %s, Email: %s, Location: %s ]\n",
		p.Dob.String(), p.Email, p.Location)
}

type Admin struct {
	Person //type embedding for composition
	Roles  []string
}

type Member struct {
	Person //type embedding for composition
	Skills []string
}

// Swap characters case from upper to lower or lower to upper.
func SwapCase(str string) string {
	buf := &bytes.Buffer{}
	for _, r := range str {
		if unicode.IsUpper(r) {
			buf.WriteRune(unicode.ToLower(r))
		} else {
			buf.WriteRune(unicode.ToUpper(r))
		}
	}
	return buf.String()
}
