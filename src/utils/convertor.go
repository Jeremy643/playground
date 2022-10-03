package utils

func fromBoolPtr(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

func fromStrPtr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
