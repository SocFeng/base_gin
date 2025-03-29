package utils

import (
	"fmt"
	"testing"
)

// 去重复
func TestDistinct(t *testing.T) {
	isd := []any{1, 2, 311, 2, 2, 2, "ab", "ab", "c"}
	res := Distinct(isd)
	fmt.Println(res)
}

// 交集
func TestIntersection(t *testing.T) {
	set1 := []any{1, 2, "ab", "c"}
	set2 := []any{3, 2, "AC", "c"}
	res := Intersection(set1, set2)
	fmt.Println(res)
}

// 差集
func TestDifference(t *testing.T) {
	set1 := []any{1, 2, "ab", "c"}
	set2 := []any{3, 2, "AC", "c"}
	res := Difference(set1, set2)
	fmt.Println(res)
}

// 交集
func TestUnion(t *testing.T) {
	set1 := []any{1, 2, "ab", "c"}
	set2 := []any{3, 2, "AC", "c"}
	res := Union(set1, set2)
	fmt.Println(res)
}

// 获取MapKeys
func TestMapKeys(t *testing.T) {
	mp := map[string]int{"a": 111, "b": 333, "fhahgs": 444}
	res := MapKeys(mp)

	fmt.Println(res)
}

// key存不存在
func TestContainsKey(t *testing.T) {
	mp := map[any]int{"a": 111, "b": 333, "fhahgs": 444, 1: 44}
	res := ContainsKey(mp, 1)
	fmt.Println(res)
}

// 切片包含某个值
func TestSliceContains(t *testing.T) {
	set2 := []any{3, 2, "AC", "c"}
	res := SliceContains(set2, 1)
	fmt.Println(res)
}

// 生成函数
func TestSliceGenerate(t *testing.T) {
	set2 := []int{1, 2, 3, 4, 5}
	res := SliceGenerate(set2, func(num int) int { return num * num })
	fmt.Println(res)
}

// 过滤函数
func TestFilterGenerate(t *testing.T) {
	set2 := []int{1, 2, 3, 4, 5}
	res := FilterGenerate(set2, func(num int) bool {
		if num%2 == 0 {
			return true
		}
		return false
	})
	fmt.Println(res)
}
