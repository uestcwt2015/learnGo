package main

import (
	"fmt"
)

func manacher(s string) string {
	tmp := make([]rune, 0)

	// handle string
	tmp = append(tmp, '#')

	for _, c := range s {
		tmp = append(tmp, c)
		tmp = append(tmp, '#')
	}

	p := make([]int, len(tmp))
	c, r := 0, 0
	// fmt.Println(tmp, p)
	for i := 1; i < len(p); i++ {
		// fmt.Println(i)
		mirror := 2*c - i

		if r > i {
			p[i] = min(p[mirror], r-i)
		} else {
			p[i] = 0
		}

		for i-p[i]-1 >= 0 && i+p[i]+1 < len(tmp) && tmp[i-p[i]-1] == tmp[i+p[i]+1] {
			p[i]++
		}

		if i+p[i] > r {
			c = i
			r = i + p[i]
		}
	}

	res := make([]rune, 0)
	maxLen, center := 0, 0

	for i, v := range p {
		if v > maxLen {
			center = i
			maxLen = v
		}
	}

	for _, v := range tmp[center-maxLen : center+maxLen] {
		if v != '#' {
			res = append(res, v)
		}
	}

	return string(res)
}

// func manacher(s string) string {
// 	tmp := make([]rune, 0)
// 	res := make([]rune, 0)
// 	tmp = append(tmp, '#')

// 	for _, c := range s {
// 		tmp = append(tmp, c)
// 		tmp = append(tmp, '#')
// 	}

// 	dp := make([]int, len(tmp))
// 	pos, maxRight := 0, 0
// 	center, maxLen := 0, 0

// 	for i := range dp {
// 		if i < maxLen {
// 			dp[i] = Min(dp[2*pos-i], maxRight-i)
// 		} else {
// 			dp[i] = 1
// 		}

// 		for i-dp[i] >= 0 && i+dp[i] < len(tmp) && tmp[i-dp[i]] == tmp[i+dp[i]] {
// 			dp[i]++
// 		}

// 		if dp[i]+i-1 > maxRight {
// 			maxRight = dp[i] + i - 1
// 			pos = i
// 		}

// 		if maxLen < dp[i]-1 {
// 			maxLen = dp[i] - 1
// 			center = i
// 		}
// 	}

// 	for _, c := range string(tmp[center-maxLen : center+maxLen]) {
// 		if c != '#' {
// 			res = append(res, c)
// 		}
// 	}

// 	return string(res)
// }

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	fmt.Println(manacher("babad"))
}
