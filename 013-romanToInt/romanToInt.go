package romanToInt

import "strings"

const  (
	RomanNumeralI = "I"
	RomanNumeralV = "V"
	RomanNumeralX = "X"
	RomanNumeralL = "L"
	RomanNumeralC = "C"
	RomanNumeralD = "D"
	RomanNumeralM = "M"
)

var (
	romanNumeralToInt = map[string]int{
		RomanNumeralI: 1,
		RomanNumeralV: 5,
		RomanNumeralX: 10,
		RomanNumeralL: 50,
		RomanNumeralC: 100,
		RomanNumeralD: 500,
		RomanNumeralM: 1000,
	}
)

func romanToInt(s string) int {
	result := 0
	for i, romanNumStr := range s {
		if _ , ok := romanNumeralToInt[string(romanNumStr)]; !ok {
			return 0
		}
		if i != len(s) - 1 {
			if romanNumeralToInt[string(s[i])] < romanNumeralToInt[string(s[i+1])] {
				result -= romanNumeralToInt[string(s[i])]
			} else {
				result += romanNumeralToInt[string(s[i])]
			}
		}
	}
	result += romanNumeralToInt[string(s[len(s)-1])]
	return result
}

func romanToIntByte(s string) int {
	dict := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := 0
	for i := 0; i < len(s)-1; i++ {
		if _ , ok := dict[s[i]]; !ok {
			return 0
		}
		if i != len(s) - 1 {
			if dict[s[i]] < dict[s[i+1]] {
				result -= dict[s[i]]
			} else {
				result += dict[s[i]]
			}
		}
	}
	result += dict[s[len(s)-1]]
	return result
}

func romanToIntBrute(s string) int {
	result := 0
	romanNumerals := strings.Split(s, "")
	for i, romanNumStr := range romanNumerals {
		romanNumInt := 0
		romanNumInt, ok := romanNumeralToInt[romanNumStr]
		if !ok {
			return 0
		}
		if i + 1 < len(romanNumerals) {
			switch romanNumStr {
			case RomanNumeralI:
				if romanNumerals[i + 1] == RomanNumeralV ||  romanNumerals[i + 1] == RomanNumeralX {
					result = result - romanNumInt
					continue
				}
			case RomanNumeralX:
				if romanNumerals[i + 1] == RomanNumeralL ||  romanNumerals[i + 1] == RomanNumeralC {
					result = result - romanNumInt
					continue
				}
			case RomanNumeralC:
				if romanNumerals[i + 1] == RomanNumeralD ||  romanNumerals[i + 1] == RomanNumeralM {
					result = result - romanNumInt
					continue
				}
			}
		}
		result = result + romanNumInt
	}

	return result
}
