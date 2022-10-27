package main

import "fmt"

type Article struct {
	Title  string
	Author string
}

func (a Article) String() string {
	return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
}

type Stringer interface {
	String() string
}

func main() {
	a := Article{
		Title:  "Entendendo Interface em Go",
		Author: "Tenille Martins",
	}
	Print(a)
}

func Print(s Stringer) {
	fmt.Println(s.String())
}
