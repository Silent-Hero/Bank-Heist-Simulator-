package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {

  rand.Seed(time.Now().UnixNano())
  isHeistOn := true
  eludedGuards := rand.Intn(100)

  // fmt.PrintLn("MONEY HEIST")
  // First Conditional (Sneak past guards)
  if (eludedGuards >= 50) {
    fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
  } else {
    isHeistOn = false
    fmt.Println("Plan a better diguise next time?")
  }
  //  Opening the Vault 
  //  opening the vault is harder than sneaking past the guards and we only have a 30% chance of doing it.
  openedVault := rand.Intn(100)

  if (isHeistOn == true && openedVault >= 70) {
    fmt.Println("Grab and Go!")
  } else if (isHeistOn == false && openedVault < 70 ) {
    fmt.Println("We won't will be able to open the vault. :( ")
  }

  // Leaving with the Money 
  leftSafetly := rand.Intn(5)
  if (isHeistOn) {
      switch (leftSafetly) {
        case 0:
        isHeistOn = false 
        fmt.Println("Turns out the guard spotted you...") 
        case 1:
        isHeistOn = false 
        fmt.Println("Turns out vault doors don't open from the inside...") 
        case 2:
        isHeistOn = false 
        fmt.Println("Turns out vault doors when open the room will get flooded by water...") 
        case 3:
        isHeistOn = true 
        fmt.Println("Turns out vault doors will open with ease...") 
        case 4:
        isHeistOn = true 
        fmt.Println("Turns out vault doors will open if we have great hacker and someone to stablize the pipes to keep the balance of the room to prevent the flood...") 
        default:
        fmt.Println("Get in the Chopper!")
      }
      // Amount stolen from the vault 
      amountStolen := 10000 + rand.Intn(1000000);
      fmt.Println("Amount stolen: ", amountStolen)
  }

  fmt.Println("Is the heist ready: ", isHeistOn)

}
