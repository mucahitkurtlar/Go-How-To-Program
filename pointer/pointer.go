package main

import "fmt"

func changeString(str *string){
  *str = "Mumdan çıkmış yangın."
}

func main(){
  var str string = "Dedeye sahip çıkalım."
  fmt.Println(str)
  changeString(&str)
  fmt.Println(str)
}
