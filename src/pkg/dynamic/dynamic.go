package dynamic

import (
	"fmt"
	"strconv"
)

var SIZE = []int{2, 3, 5, 6, 9}
var VALUE = []int{2, 4, 7, 10, 14}

const NAP_SIZE = 17

// use size <= 1
// 0,2,2,4,4,6,6,8,8,10,10,12,12,14,14,16,
// use size <= 2
// 0,2,4,4,6,8,8,10,12,12,14,16,16,18,20,20,
// use size <= 3
// 0,2,4,4,7,8,9,11,12,14,15,16,18,19,21,22,
// use size <= 4
// 0,2,4,4,7,10,10,12,14,14,17,20,20,22,24,24,
// use size <= 5
// 0,2,4,4,7,10,10,12,14,14,17,20,20,22,24,24,
func NapZaq(){
	napValues := [NAP_SIZE]int{}
	
	for i := 0; i < 5; i++ {
		for j := SIZE[i];j < NAP_SIZE; j++ {
			n := napValues[j - SIZE[i]] + VALUE[i]
			if n > napValues[j] {
				napValues[j] = n
			}
		}

		fmt.Println("use size <= " + strconv.Itoa(i + 1))
		for k := 1;k < NAP_SIZE; k++ {
			fmt.Print(strconv.Itoa(napValues[k]) + ",")
		}
		fmt.Println("")
	}
}

var COIN = []int{1, 5, 12}

const MAX_PAY = 20

// use coin <= 1
// 1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,
// use coin <= 5
// 1,2,3,4,1,2,3,4,5,2,3,4,5,6,3,4,
// use coin <= 12
// 1,2,3,4,1,2,3,4,5,2,3,1,2,3,3,4,
func LessCoinCountPay() {
	coinCounts := [MAX_PAY+1]int{}

	for i := 0; i < len(COIN); i++ {
		for j := 1; j <= MAX_PAY; j++ {
			if j < COIN[i] {
				continue;
			}
			
			newCount := coinCounts[j - COIN[i]] + 1
			if coinCounts[j] == 0 {
				coinCounts[j] = newCount
				continue;
			}

			if coinCounts[j] > newCount {
				coinCounts[j] = newCount
			}
		}

		fmt.Println("use coin <= " + strconv.Itoa(COIN[i]))
		for k := 1;k < NAP_SIZE; k++ {
			fmt.Print(strconv.Itoa(coinCounts[k]) + ",")
		}
		fmt.Println("")
	}
}