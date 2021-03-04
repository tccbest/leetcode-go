package _013_roman_to_integer

var romanToIntMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	if s == "" {
		return 0
	}
	num, last, total := 0, 0, 0
	for i := 0; i < len(s); i++ {
		char := s[len(s)-(i+1) : len(s)-i]
		num = romanToIntMap[char]
		if num < last {
			total = total - num
		} else {
			total = total + num
		}
		last = num
	}
	return total
}
