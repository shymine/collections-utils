package collectionutils

func Map[A any, B any](list []A, f func(A) B) []B {
	res := []B{}
	for _, el := range list {
		res = append(res, f(el))
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