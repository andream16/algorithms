package main

import "fmt"

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {

	res := []int{}
	seen1, seen2 := map[int]struct{}{}, map[int]struct{}{}

	for _, a1 := range arr1 {
		seen1[a1] = struct{}{}
	}

	for _, a2 := range arr2 {
		_, ok := seen1[a2]
		if !ok {
			continue
		}
		seen2[a2] = struct{}{}
		for idx, a3 := range arr3 {
			_, ok1 := seen1[a3]
			_, ok2 := seen2[a3]
			if ok1 && ok2 {
				res = append(res, a3)
				arr3 = append(arr3[:idx], arr3[idx+1:]...)
				break
			}
		}
	}

	return res

}

func main() {

	a1, a2, a3 := []int{1, 2, 3, 4, 5}, []int{1, 2, 5, 7, 9}, []int{1, 3, 4, 5, 8}

	fmt.Println(arraysIntersection(a1, a2, a3))
}
