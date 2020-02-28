package main

import (
	"fmt"
	"sort"
)

type Knight struct {
	x, y int
}

func (k *Knight) availableMoves() {

}

func (k *Knight) setPosition(x, y int) {
	k.x = x
	k.y = y
}

type Pawn struct {
	x, y int
}

func (p *Pawn) setPosition(x, y int) {
	p.x = x
	p.y = y
}

type Figure interface {
	setPosition(x, y int)
}

func setOnes(figures ...Figure) {
	for _, figure := range figures {
		figure.setPosition(1, 1)
	}
}

type Record struct {
	Name    string
	Phone   string
	Address string
	Age     int
}

type Phonebook []Record

func main() {
	var a Knight
	a.x = 3
	a.setPosition(4, 7)
	var b Pawn
	b.setPosition(2, 2)
	fmt.Println(a)
	fmt.Println(b)
	setOnes(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
	v := Record{"vasya", "1234567", "ul mira", 33}
	p := Record{"anya", "2345678", "red squre", 18}
	var book []Record
	book = append(book, v)
	book = append(book, p)
	fmt.Println(book)
	sort.Slice(book, func(i int, j int) bool {
		return book[i].Name < book[j].Name
	})
	fmt.Println(book)
}
