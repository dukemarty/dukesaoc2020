package main

type GroupAnswers struct {
	Raw []string
}

func (ga GroupAnswers) CountAny() int {
	m := make(map[int32]int)
	for _, a := range ga.Raw {
		for _, c := range a {
			m[c] = 1
		}
	}

	return len(m)
}

func (ga GroupAnswers) CountAll() int {
	m := make(map[int32]int)
	for _, a := range ga.Raw {
		for _, c := range a {
			if _, ok := m[c]; ok {
				m[c]++
			} else {
				m[c] = 1
			}
		}
	}

	res := 0
	for _, v := range m {
		if v == len(ga.Raw) {
			res++
		}
	}

	return res
}
