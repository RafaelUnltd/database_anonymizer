package anonymizer

import (
	"crypto/rand"
	"math/big"
	"strings"
)

// replaceField replaces the field value with a fixed mask
func replaceField(field *string, mask string) {
	*field = mask
}

// hideField hides the entire field value
func hideField(field *string) {
	length := len(*field)
	if length == 0 {
		return
	}
	*field = strings.Repeat("X", length)
}

// partiallyHideField hides all but the first 5 characters of a field
func partiallyHideField(field *string) {
	length := len(*field)
	if length <= 5 {
		hideField(field)
		return
	}
	*field = (*field)[0:5] + strings.Repeat("X", length-5)
}

// maskField masks a field based on the provided mask
// The mask can contain the following characters:
//
// - # for a random number
//
// - @ for a random letter
//
// - * for a random number or letter
//
// Any other character is considered a literal
func maskField(field *string, mask string) error {
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	charnumset := "0123456789" + charset

	if len(mask) == 0 {
		return ErrMaskEmpty
	}

	masked := ""
	for _, char := range mask {
		switch char {
		case '#':
			value, err := rand.Int(rand.Reader, big.NewInt(10))
			if err != nil {
				return err
			}
			masked += value.String()
		case '+':
			length := int64(len(charset))
			value, err := rand.Int(rand.Reader, big.NewInt(length))
			if err != nil {
				return err
			}
			masked += string(charset[value.Int64()])
		case '*':
			length := int64(len(charnumset))
			value, err := rand.Int(rand.Reader, big.NewInt(length))
			if err != nil {
				return err
			}
			masked += string(charnumset[value.Int64()])
		default:
			masked += string(char)
		}
	}

	*field = masked
	return nil
}
