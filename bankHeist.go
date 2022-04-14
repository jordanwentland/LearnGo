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
	if eludedGuards >= 50 {
		fmt.Println("You succesfully made it past the guards! Move to the vault!")
	} else {
		isHeistOn = false
		fmt.Println("Heist failed, the guards caught you.")
	}

	openedVault := rand.Intn(100)
	if isHeistOn == true && openedVault >= 70 {
		fmt.Println("The vault is open, grab what you can and run!")
	} else if isHeistOn == true {
		isHeistOn = false
		fmt.Println("Turns out i forgot how to crack a safe.")
	}

	leftSafely := rand.Intn(5)
	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("Alarm has been triggered!!! Run!")
		case 1:
			isHeistOn = false
			fmt.Println("The door locked behind you. GG")
		case 2:
			isHeistOn = false
			fmt.Println("Literally got clothes lined, youre dog water!")
		case 3:
			isHeistOn = false
			fmt.Println("you got shot. rip.")
		default:
			fmt.Println("You made it out! Quick, start the car!")
		}
	}

	if isHeistOn {
		amtStolen := 10000 + rand.Intn(1000000)
		fmt.Printf("You made it out with: $%v", amtStolen)
	}
}
