package piscine

func atoiConverter(n string) int {
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

func BasicAtoi2(s string) int {
	sum := 0
	stringTrigger := false
	for i := 0; i < len(s); i++ {
		if atoiConverter(string(s[i])) == -1 {
			stringTrigger = true
		}
		power := 1
		if i < len(s) {
			for j := 1; j < len(s)-i; j++ {
				power *= 10
			}
		}
		sum += atoiConverter(string(s[i])) * power
	}
	if stringTrigger {
		return 0
	} else {
		return sum
	}
}
