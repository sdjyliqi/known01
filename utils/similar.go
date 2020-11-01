package utils

import (
	"math"
)

func SimilarText(first, second string, percent *float64) int {
	var similarText func(string, string, int, int) int
	similarText = func(str1, str2 string, len1, len2 int) int {
		var sum, max int
		pos1, pos2 := 0, 0

		// Find the longest segment of the same section in two strings
		for i := 0; i < len1; i++ {
			for j := 0; j < len2; j++ {
				for l := 0; (i+l < len1) && (j+l < len2) && (str1[i+l] == str2[j+l]); l++ {
					if l+1 > max {
						max = l + 1
						pos1 = i
						pos2 = j
					}
				}
			}
		}

		if sum = max; sum > 0 {
			if pos1 > 0 && pos2 > 0 {
				sum += similarText(str1, str2, pos1, pos2)
			}
			if (pos1+max < len1) && (pos2+max < len2) {
				s1 := []byte(str1)
				s2 := []byte(str2)
				sum += similarText(string(s1[pos1+max:]), string(s2[pos2+max:]), len1-pos1-max, len2-pos2-max)
			}
		}

		return sum
	}

	l1, l2 := len(first), len(second)
	if l1+l2 == 0 {
		return 0
	}
	sim := similarText(first, second, l1, l2)
	if percent != nil {
		*percent = float64(sim*200) / float64(l1+l2)
	}
	return sim
}

//Levenshtein 两个字符串的编辑距离
func Levenshtein(a, b string) int {
	if a == b {
		return 0
	}

	r1 := []rune(a)
	r2 := []rune(b)

	if len(r1) == 0 {
		return len(r2)
	}

	if len(r2) == 0 {
		return len(r1)
	}

	line := make([]int, 0, len(r2)+1)
	nook := 0
	for i := 0; i < len(r1)+1; i++ {
		for j := 0; j < len(r2)+1; j++ {
			if i == 0 {
				line = append(line, j)
				continue
			}
			if j == 0 {
				nook = i - 1
				line[0] = i
				continue
			}
			oval := line[j]

			var v int
			if r1[i-1] != r2[j-1] {
				v = 1
			}
			line[j] = min(nook+v, line[j-1]+1, oval+1)
			nook = oval
		}
	}

	return line[len(r2)]
}

//SimilarDegree 字符串相似度
func SimilarDegree(a, b string) float64 {
	return 1 - float64(Levenshtein(a, b))/math.Max(float64(len([]rune(a))), float64(len([]rune(b))))
}

//min 获取小值
func min(a ...int) int {
	r := a[0]
	for _, v := range a {
		if r > v {
			r = v
		}
	}
	return r
}
