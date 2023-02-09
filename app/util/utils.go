package util

import (
	"log"
	"strconv"

	"golang.org/x/exp/slices"
)

var (
	masterConst = []string{"51", "52", "53", "54", "55"}
	visaConst   = []string{"4"}
	amexConst   = []string{"34", "37"}
)

func validadorLuhn(cardNumber int64, master bool, visa bool, amex bool) bool {

	strCardNumber := strconv.Itoa(int(cardNumber))

	if master && slices.Contains(masterConst, strCardNumber[:2]) {
		return false
	}

	if amex && slices.Contains(amexConst, strCardNumber[:2]) {
		return false
	}

	if visa && slices.Contains(visaConst, strCardNumber[:1]) {
		return false
	}

	number := strCardNumber[:len(strconv.Itoa(int(cardNumber)))-1]
	lastDigit, _ := strconv.Atoi(strCardNumber[len(strCardNumber)-1:])

	sum := 0

	for i := 0; i < len(number); i++ {
		multiplier := 2

		if i%2 == 1 {
			multiplier = 1
		}

		currNum, erro := strconv.Atoi(number[i : i+1])

		if erro != nil {
			log.Fatal(erro)
		}

		currNum = currNum * multiplier

		if currNum < 10 {
			sum += currNum
		} else {
			//	strNum := strconv.Itoa(currNum)

			sumChar := currNum%10 + currNum/10
			sum += sumChar
		}
	}

	return ((10-(sum%10))%10 == lastDigit)

}
