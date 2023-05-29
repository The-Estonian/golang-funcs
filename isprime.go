package piscine

func IsPrime(nb int) bool {
	if nb == 1 || nb < 0 {
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
