package conversor

func DecToHex(decimal int) []string {
	num := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}
	var hex []string
	for decimal > 0 {
		hex = append([]string{num[decimal%16]}, hex...)
		decimal /= 16
	}
	return hex
}

func DecToBin(decimal int) []int {
	var bin []int
	for decimal > 0 {
		if decimal%2 == 0 {
			bin = append([]int{0}, bin...)
		} else {
			bin = append([]int{1}, bin...)
		}
		decimal /= 2
	}
	return bin
}

func BinToDec(binario string) int {
	bin := sortArray(binario)
	var dec int

	for i := len(bin) - 1; i >= 0; i-- {
		if bin[i] == "1" {
			dec += pow(2, i)
		}
	}

	return dec
}

func sortArray(bin string) []string {
	array := []rune(bin)
	var binArray []string
	for i, j := 0, len(bin)-1; i < j; i, j = i+1, j-1 {
		binArray[i] = string(array[j])
	}
	return binArray
}

func pow(num, pow int) int {
	if pow == 0 {
		return 1
	} else {
		result := num
		for i := 1; i < pow; i++ {
			result = result * num
		}
		return result
	}
}
