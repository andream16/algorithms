package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(solution("210022", 3))
	fmt.Println(solution("1211", 10))
}

func solution(s string, b int) int {
	return solutionRec(s, b, 0, map[string]int{})
}

func solutionRec(s string, b int, counter int, seen map[string]int) int {
	counter += 1
	l := len(s)
	
	xy := getXY(s)
	x, _ := strconv.ParseInt(xy[0], b, 32)
	y, _ := strconv.ParseInt(xy[1], b, 32)	
	z := strconv.FormatInt(x-y, b)
	
	if l > len(z) {
		zs, _ := strconv.Atoi(z)
		z = fmt.Sprintf("%0"+strconv.Itoa(l)+"d", zs)
	}
	
	idx, ok := seen[z]
	if ok {
		return counter - idx
	}
	seen[z] = counter
	return solutionRec(z, b, counter, seen)
}

func getXY(s string) []string {
	yArr := strings.Split(s, "")
	sort.Strings(yArr)
	xArr := append(yArr[:0:0], yArr...)
	sort.Sort(sort.Reverse(sort.StringSlice(xArr)))
	return []string{strings.Join(xArr, ""), strings.Join(yArr, "")}
}
