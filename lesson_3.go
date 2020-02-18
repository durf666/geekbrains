package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Car struct {
	Model          string
	ProductionYear int
	TrunkVolume    int
	EngineLaunched bool
	WindowsOpened  bool
	TrunkFill      int
}
type Lorry struct {
	Model          string
	ProductionYear int
	BodyVolume     int
	EngineLaunched bool
	WindowsOpened  bool
	BodyFill       int
}
type Query struct {
	data []int
}

func (q *Query) push(i int) {
	q.data = append(q.data, i)
}
func (q *Query) pop() (i int) {
	if len(q.data) == 0 {
		return 0
	} else {
		i = q.data[0]
		q.data = append(q.data)[1:]
		return i
	}
}

type Phonebook struct {
	Name    string
	Phone   string
	Address string
	Age     int
}

func (record *Phonebook) update(name, phone, address string, age int) {
	record.Name = name
	record.Phone = phone
	record.Address = address
	record.Age = age
	a, _ := json.Marshal(record)
	fmt.Println("ssss")
	fmt.Println(string(a))
	ioutil.WriteFile(name+".json", a, 0744)

}

func main() {
	car1 := Car{"tesla model S", 2013, 300, false, false, 0}
	car2 := Car{"VAZ 2101", 1980, 600, false, true, 350}
	lorry1 := Lorry{"URAL", 1974, 5600, false, false, 0}
	lorry1.BodyFill = 4000
	car1.EngineLaunched = true
	fmt.Sprintf("%+v", car1)
	fmt.Printf("%+v", car2)
	fmt.Printf("%+v", lorry1)
	var zzz Phonebook
	var sss Query
	var aa int
	sss.push(1)
	sss.push(2)
	sss.push(3)
	sss.push(4)
	fmt.Println("")
	fmt.Println(sss)
	aa = sss.pop()
	fmt.Println(sss)
	fmt.Println(aa)
	zzz.update("Vasya", "9031234567", "Gogol_st.", 21)
	// zzz = Phonebook{"Vasya", "903123456", "Gogol_st."}
	// dd, _ := (json.Marshal(zzz))
	// fmt.Println(dd)
}
