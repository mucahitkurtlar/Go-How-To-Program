package main

import "fmt"

type person struct {
  name string
  surname string
  age int
}

func (p *person)speak(){
  fmt.Println(p.name, p.surname, "çevir oğlum!")
}

func main(){
  newPerson := person{"Antonio", "Banderas", 40}
  fmt.Println(newPerson.name, newPerson.surname, newPerson.age)
  newPerson.speak()
}
