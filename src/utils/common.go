package utils

// 交集
func Intersection[T comparable](a, b []T) []T {
	set := make(map[T]struct{})
	for _, v := range a {
		set[v] = struct{}{}
	}
	result := make([]T, 0)
	for _, v := range b {
		if _, exists := set[v]; exists {
			result = append(result, v)
		}
	}
	return result
}

// 差集
func Difference[T comparable](a, b []T) []T {
	set := make(map[T]struct{})
	for _, v := range b {
		set[v] = struct{}{}
	}
	result := make([]T, 0)
	for _, v := range a {
		if _, exists := set[v]; !exists {
			result = append(result, v)
		}
	}
	return result
}

// 去重复
func Distinct[T comparable](slice []T) []T {
	seen := make(map[T]struct{})
	result := make([]T, 0, len(slice))
	for _, item := range slice {
		if _, exists := seen[item]; !exists {
			seen[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// 并集

func Union[T comparable](a, b []T) []T {
	seen := make(map[T]struct{})
	result := make([]T, 0, len(a)+len(b))

	// 添加第一个切片的所有元素
	for _, item := range a {
		if _, exists := seen[item]; !exists {
			seen[item] = struct{}{}
			result = append(result, item)
		}
	}

	// 添加第二个切片中未出现的元素
	for _, item := range b {
		if _, exists := seen[item]; !exists {
			seen[item] = struct{}{}
			result = append(result, item)
		}
	}

	return result
}

// 获取Key列表
func MapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

// 指定值是否在key中
func ContainsKey[K comparable, V any](mp map[K]V, key K) bool {
	_, exists := mp[key]
	return exists
}

// 切片包含某个值

func SliceContains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// 生成
// 泛型 Slice 函数
func SliceGenerate[T any, R any](src []T, transform func(T) R) []R {
	res := make([]R, len(src))
	for i, v := range src {
		res[i] = transform(v)
	}
	return res
}

// 泛型 Filter 函数
func FilterGenerate[T any](src []T, predicate func(T) bool) []T {
	var res []T
	for _, v := range src {
		if predicate(v) {
			res = append(res, v)
		}
	}
	return res
}
