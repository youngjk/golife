package main

import "fmt"

type phone struct {
	name string
	modelNumber int
	phoneNumber int
	owner string
}

func newPhone(n string, mn int, pn int, o string) *phone {
	nPhone := phone{name: n, modelNumber: mn, phoneNumber: pn, owner: o}
	return &nPhone
}

func main() {
	fmt.Println(phone{"iPhone", 1, 1111111111, "Bob"})
	fmt.Println(phone{name: "Samsung Note 10", modelNumber: 2, phoneNumber: 2222222222, owner: "Billy"})
	fmt.Println(newPhone("Google Pixel 4", 3, 3333333333, "John"))

	phone1 := phone{"iPhone 11 Max Pro", 4, 4444444444, "Matt"}
	fmt.Println(phone1)
	phone2 := &phone1
	fmt.Println(phone2.phoneNumber)
	phone2.phoneNumber = 5555555555
	fmt.Println(phone2.phoneNumber)
}