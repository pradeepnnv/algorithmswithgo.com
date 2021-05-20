package module01

import (
	"math"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) int {
	var dec int
	for i := 0; i < len(value); i++ {
		v := convHexToInt(string(value[len(value)-i-1]))
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Println(v)
		// fmt.Println(int(math.Pow(float64(base), float64(i))))
		dec += v * (int(math.Pow(float64(base), float64(i))))
	}
	return dec
}

func convHexToInt(h string) int {
	switch h {
	case "0":
		return 0
	case "1":
		return 1
	case "2":
		return 2
	case "3":
		return 3
	case "4":
		return 4
	case "5":
		return 5
	case "6":
		return 6
	case "7":
		return 7
	case "8":
		return 8
	case "9":
		return 9
	case "A":
		return 10
	case "B":
		return 11
	case "C":
		return 12
	case "D":
		return 13
	case "E":
		return 14
	case "F":
		return 15
	default:
		return -1
	}
}
