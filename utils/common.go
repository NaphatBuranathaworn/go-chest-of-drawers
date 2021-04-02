package utils

type ArrayString []string

func (s ArrayString) Contains(str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
