package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	name    string
	surname string
	contact contactInfo
}

func main() {
	p1 := person{"Jack", "Anderson", contactInfo{"smt@smt.com", 12345}}
	fmt.Println(p1)
	p1.updateName("Jim")
	fmt.Println(p1)
}
func (p *person) updateName(n string) {
	(*p).name = n
}
