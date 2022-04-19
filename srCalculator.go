package main

import (
	"fmt"
)

func srToRank(sr int) string {
	var rank string
	if sr < 1500 {
		rank = "Bronze"
	} else if sr > 1500 && sr < 2000 {
		rank = "Silver"
	} else if sr > 2000 && sr < 2500 {
		rank = "Gold"
	} else if sr > 2500 && sr < 3000 {
		rank = "Platinum"
	} else if sr > 3000 && sr < 3500 {
		rank = "Diamond"
	} else if sr > 3500 && sr < 4000 {
		rank = "Master"
	} else if sr > 4000 && sr < 5000 {
		rank = "GrandMaster"
	}
	return rank
}

func main() {
	var sr = 2765
	var rank string
	rank = srToRank(sr)
	fmt.Println(rank)
}
