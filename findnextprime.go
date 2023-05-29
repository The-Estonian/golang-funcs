package piscine

func IsPrime2(nb int) bool {
	if nb < 2 {
		return false
	}
	num := 2
	answer := true
	for num <= nb/2 {
		remainder := nb % num
		if remainder != 0 {
			num = num + 1
		} else {
			answer = false
			break
		}
	}
	return answer
}

func FindNextPrime(nb int) int {
	if IsPrime2(nb) {
		return nb
	}
	return FindNextPrime(nb + 1)
}
