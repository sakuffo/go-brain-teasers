// https://learning.oreilly.com/library/view/go-brain-teasers/9781680509144/f_0017.xhtml#rcp.Raw
// Go Brain Teasers
// Puzzle 5

package main

import "fmt"

func main() {
	s := `a\tb` // backticks are used to create a raw string literal

	str := "a\tb" // double quotes are used to create a interpreted string literal

	// The code editor gives hints as to what will happen below
	// backticks will produce the literal string a\tb because that is in the backtick quotes
	// double quotes on the other hand will produce a tab character because the \t is interpreted as a tab
	// \t \n \r are all interpreted as special characters
	fmt.Println(s)
	fmt.Println(str)
}
