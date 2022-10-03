package utils

func FromBoolPtr(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

func FromStrPtr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func FromIntPtr(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}
