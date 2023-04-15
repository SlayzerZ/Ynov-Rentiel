package hangman

func Atoi(s string) int {
	var numberString string
	var numberReverse string
	var number int = 0
	var coef int = 1
	var negativeNumber bool = false
	var first bool = true

	allNumber := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	allNumberString := [...]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	if len(s) == 0 {
		return 0
	}

	// ==> Loop to convert string number into int number
	for i := 0; i <= len(s)-1; i++ {
		if (s[i] >= 48 && s[i] <= 57) || s[i] == 45 {
			if s[i] == 45 && first == true {
				if i <= len(s)-2 && (s[i+1] == 45 || s[i+1] == 43) {
					return 0
				}
				negativeNumber = true
				first = false
			} else if first == false && (s[i] == 45 || s[i] == 43) {
				return 0
			} else {
				numberString += string(s[i])
				first = false
			}
		} else {
			if s[len(s)-1] == 43 || s[len(s)-1] == 45 {
				return 0
			} else if i <= len(s)-2 && s[0] == 43 && s[1] == 43 {
				return 0
			} else if s[i] != 43 {
				return 0
			}
		}
	}

	// ==> Reversr
	for j := len(numberString) - 1; j >= 0; j-- {
		numberReverse += string(numberString[j])
	}

	for k := 0; k <= len(numberReverse)-1; k++ {
		for l := 0; l <= len(allNumberString)-1; l++ {
			if allNumberString[l] == string(numberReverse[k]) {
				number += (allNumber[l] * coef)
				coef *= 10
			}
		}
	}

	if negativeNumber {
		number = number * (-1)
	}
	return number
}
