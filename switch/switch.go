package main

import "fmt"

func main() {
  i := 4
  switch i {
  case 1:
    fmt.Println("Oh... It's you." )
  case 2:
    fmt.Println("It's been a long time. How have you been?")
  case 3:
    fmt.Println("I've been really busy being dead. You know, after you MURDERED ME.")
  case 4:
    fmt.Println("Okay. Look. We both said a lot of things that you're going to regret. But I think we can put our differences behind us.\n\nFor science.\n\nYou monster.")

  }


  i = 1
  switch {
  case i==1:
    fmt.Println("Same thing..." )
  case false:
    fmt.Println("but cooler")
  case true:
    fmt.Println("but more useless way")
  case i!=1:
    fmt.Println("fck JavaScript")

  }
}
