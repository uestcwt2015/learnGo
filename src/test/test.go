package main

import "fmt"

func convert(s string, numRows int) string {
	var ret = make([][]rune, 0)
	var res = make([]rune, 0)
	for i := 0; i < numRows; i++ {
		child := make([]rune, len(s))
		ret = append(ret, child)
	}
	x, y := 0, 0
	for i, v := range s {
		n := 2*numRows - 2
		ret[y][x] = v
		// fmt.Println(x, y, v, i%n < numRows-1)
		if i%n < numRows-1 {
			y++
		} else {
			x++
			y--
		}
		// fmt.Println(x, y)
	}
	// fmt.Println(ret)
	for i := 0; i < numRows; i++ {
		length := len(ret[i])
		for j := 0; j < length && len(res) < len(s); j++ {
			if ret[i][j] != 0 {
				res = append(res, ret[i][j])
			}
		}
	}
	fmt.Println(string(res))
	return string(res)
}

func main() {
	convert("PAYPALISHIRING", 3)
}
