package main

import "fmt"

func madonna()string{
  fmt.Print("Fck ")
  return "Turmp!"
}

type mensch struct {
  name string
}

type uberMensch struct {
  mensch
  ability string
}

func (u *uberMensch)fly(){
  fmt.Println("Uçuyorum şuan.")
}


func main(){
  var str string = "Where is your head at?"
  fmt.Println(str)
  //---------------------------------------
  for i := 0; i < 5; i++{
    fmt.Println("Do it again!")
  }
  //---------------------------------------
  defer fmt.Println("Hung up.")
  var arr0 [5]int
  for i := 0; i < 5; i++{
    arr0[i] = i
  }
  fmt.Println(arr0)
  arr1 := [5]int{1, 2, 3, 4, 5}
  fmt.Println(arr1)
  var slice []int = arr1[2:4]
  fmt.Println(slice)
  fmt.Println(madonna())
  person0 := mensch{"Recep"}
  fmt.Println(person0.name)
  person1 := uberMensch{}
  person1.name = "Nietzsche"
  person1.ability = "Overthinking"
  fmt.Println(person1)
  person1.fly()
  for i, a := range arr0 {
    fmt.Println(i, a)
  }
}
