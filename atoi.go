package piscine

func atoiConverter3(n string) int {
	if n == "0" {
		return 0
	} else if n == "1" {
		return 1
	} else if n == "2" {
		return 2
	} else if n == "3" {
		return 3
	} else if n == "4" {
		return 4
	} else if n == "5" {
		return 5
	} else if n == "6" {
		return 6
	} else if n == "7" {
		return 7
	} else if n == "8" {
		return 8
	} else if n == "9" {
		return 9
	} else {
		return -1
	}
}

func Atoi(s string) int {
	sum := 0
	stringTrigger := false
	negativePositive := false

	// Check for +/- marks and return zero if matches
	if string(s[0]) == "+" && string(s[1]) == "+" ||
		string(s[0]) == "-" && string(s[1]) == "-" ||
		string(s[0]) == "+" && string(s[1]) == "-" ||
		string(s[0]) == "-" && string(s[1]) == "+" {
		return 0
	}

	// check for +/- mark and cut string if matches
	if string(s[0]) == "+" || string(s[0]) == "-" {
		if string(s[0]) == "-" {
			negativePositive = true
		}
		s = s[1:]
	}

	// loop thru all letters
	for i := 0; i < len(s); i++ {
		// check if all letters are numbers
		if atoiConverter3(string(s[i])) == -1 {
			stringTrigger = true
		}

		power := 1
		if i < len(s) {
			for j := 1; j < len(s)-i; j++ {
				power *= 10
			}
		}
		sum += atoiConverter3(string(s[i])) * power
	}

	if stringTrigger {
		return 0
	} else {
		if negativePositive {
			return sum * -1
		} else {
			return sum
		}
	}
}
