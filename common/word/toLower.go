package word

func ToLowerCase(chr byte) byte {
	if chr >= 65 && chr <= 90 {
		chr += 32
	}

	return chr
}
