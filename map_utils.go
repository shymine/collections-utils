package collectionsutils

func Keys[A any](m map[A]any) []A {
	res := make([]A, len(m))
	for el, _ := range m {
		res = append(res, el)
	}
	return res
}

func Values[A any](m map[any]A) []A {
	res := make([]A, len(m))
	for _, el := range m {
		res = append(res, el)
	}
	return res
}

type keyValue[A any, B any] struct {
	Key   A
	Value B
}

func Pairs[A any, B any](m map[A]B) []keyValue[A, B] {
	res := make([]keyValue[A, B], len(m))
	for k, v := range m {
		res = append(res, keyValue[A, B]{Key: k, Value: v})
	}
	return res
}
