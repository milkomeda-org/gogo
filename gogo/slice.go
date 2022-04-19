// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2021/11/30.

package utils

func Intersect(slice1, slice2 []int) []int {
	m := make(map[int]int)
	nn := make([]int, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

// DiffSet set between 2 slice
// Returns the element of A that is not in B
func DiffSet(a []int, b []int) []int {
	m := make(map[int]int)
	nn := make([]int, 0)
	inter := Intersect(a, b)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range a {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}

// Sort a slice
// field returns a compare value
// desc whether to reverse
/*
   sorted, err := utils.Sort([]interface{}{hours}, func(v interface{}) int {
      return v.(*model.StudentClassHourResp).Time
   }, true)
*/
func Sort(slice []interface{}, field func(interface{}) int, desc bool) ([]interface{}, error) {
	if sorted, err := TimSort(slice, field); desc && err == nil {
		return reverse(sorted), err
	} else {
		return sorted, err
	}
}

func reverse(s []interface{}) []interface{} {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
