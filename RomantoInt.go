package main

import "fmt"

func main() {
	fmt.Println("jus tryna undestand one thing")
	romans := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	result := 0
	count := 0
	ccount := 0
	cccount := 0
	err := 0

	for _, r := range s {
		if r == 73 { //I
			count++
		}
	}
	if count > 3 {
		return 0
	}
	for _, r := range s {
		if r == 86 { //V
			ccount++
		}
	}
	if ccount > 1 {
		return 0
	}

	for _, r := range s {
		if r == 88 { //X
			cccount++
		}
	}
	if cccount > 4 {
		return 0
		result = 0
		err = 1
	}

	if romans == nil {
		return 0
	}

	if err == 0 {
		for i := 0; i < len(s); i++ {
			if romans[string(s[i])] == 0 {
				fmt.Println("Err, wrong symbol in the string:", string(s[i]))
				result = 0
				break
			}
			if i+2 < len(s) && romans[string(s[i])] == romans[string(s[i+1])] && romans[string(s[i+1])] < romans[string(s[i+2])] {
				fmt.Println("Err,wrong syntaxis")
				result = 0
				break
			}
			if i+3 < len(s) && romans[string(s[i])] == romans[string(s[i+1])] && romans[string(s[i])] == romans[string(s[i+2])] && romans[string(s[i])] == romans[string(s[i+3])] {
				fmt.Println("Err,wrong syntaxis")
				result = 0
				break
			}
			if i+1 < len(s) && s[i] == 86 && s[i+1] == 88 {
				fmt.Println("Wrong syntaxis")
				result = 0
				break
			}
			if i+1 < len(s) && romans[string(s[i])] < romans[string(s[i+1])] {
				result += romans[string(s[i+1])] - romans[string(s[i+1])] - romans[string(s[i])]
			} else {
				result += romans[string(s[i])]
			}

		}

	}
	return result
}
