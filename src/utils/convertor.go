package utils

func fromBoolPtr(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}
