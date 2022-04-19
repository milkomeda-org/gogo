// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2021/12/8.
// TimSort for Go

package utils

import (
	"fmt"
	"math"
)

func binarySearch(data []interface{}, item interface{}, field func(interface{}) int, start int, end int) int {
	if start == end {
		if field(data[start]) > field(item) {
			return start
		}
		return start + 1
	}

	if start > end {
		return start
	}

	mid := int(math.Round((float64(start + end)) / 2))
	if field(data[mid]) < field(item) {
		return binarySearch(data, item, field, mid+1, end)
	}
	if field(data[mid]) > field(item) {
		return binarySearch(data, item, field, start, mid-1)
	}

	return mid
}

func insertionSort(data []interface{}, field func(interface{}) int) []interface{} {
	for i := 1; i < len(data); i++ {
		var result []interface{}
		value := data[i]
		pos := binarySearch(data, value, field, 0, i-1)

		result = append(result, data[0:pos]...)
		result = append(result, value)
		result = append(result, data[pos:i]...)
		result = append(result, data[i+1:]...)

		data = result
	}

	return data
}

func merge(left []interface{}, right []interface{}, field func(interface{}) int) []interface{} {
	var result []interface{}

	if len(left) == 0 {
		return right
	}
	if len(right) == 0 {
		return left
	}

	if field(left[0]) < field(right[0]) {
		result = merge(left[1:], right, field)
		result = append([]interface{}{left[0]}, result...)
	} else {
		result = merge(left, right[1:], field)
		result = append([]interface{}{right[0]}, result...)
	}

	return result
}

// TimSort sort elements using TimSort algorithm.
func TimSort(elements []interface{}, field func(interface{}) int) ([]interface{}, error) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	if len(elements) < 2 {
		return elements, nil
	}

	var runs [][]interface{}
	var sortedRuns [][]interface{}
	newRun := []interface{}{elements[0]}

	for i := 1; i < len(elements); i++ {
		if i == len(elements)-1 {
			newRun = append(newRun, elements[i])
			runs = append(runs, newRun)
			break
		}

		if field(elements[i]) < field(elements[i-1]) {
			runs = append(runs, newRun)
			newRun = []interface{}{elements[i]}
		} else {
			newRun = append(newRun, elements[i])
		}
	}

	for _, item := range runs {
		sortedRuns = append(sortedRuns, insertionSort(item, field))
	}
	var sortedArray []interface{}
	for _, item := range sortedRuns {
		sortedArray = merge(sortedArray, item, field)
	}

	return sortedArray, nil
}
