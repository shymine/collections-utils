package collectionutils

func Map[A any, B any](list []A, f func(A) B) []B {
	res := make([]B, len(list))
	for i, el := range list {
		res[i] = f(el)
	}
	return res
}

func Reduce[A any, B any](list []A, init B, f func(B, A) B) B {
	res := init
	for _, el := range list {
		res = f(res, el)
	}
	return res
}

func Filter[A any](list []A, f func(A) bool) []A {
	res := []A{}
	for _, el := range list {
		if f(el) {
			res = append(res, el)
		}
	}
	return res
}

func Intersection[A comparable](a []A, b []A) []A {
	res := []A{}
	for _, el := range a {
		if Contains(b, el) {
			res = append(res, el)
		}
	}
	return res
}

func Contains[A comparable](list []A, elem A) bool {
	for _, el := range list {
		if el == elem {
			return true
		}
	}
	return false
}