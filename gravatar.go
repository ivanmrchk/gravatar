package gravatar

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"strings"
)

// This function takes in an email address as a parameter and returns a complete url to the image.
// for example: GetImg("example@example.com") will return this: http://www.gravatar.com/avatar/563dd2c62e66d42a3b4454c410a0b3ab?d=identicon
func GetImg(email string) string {
	email = strings.TrimSpace(email)
	email = strings.ToLower(email)

	h := md5.New()
	io.WriteString(h, email)

	finalBytes := h.Sum(nil)

	finalString := "http://www.gravatar.com/avatar/" + hex.EncodeToString(finalBytes) + "?d=identicon"

	return finalString
}

// This function takes in an email address as a parameter and returns just the hash, which you can use to your liking.
func GetHash(email string) string {
	email = strings.TrimSpace(email)
	email = strings.ToLower(email)

	h := md5.New()
	io.WriteString(h, email)

	finalBytes := h.Sum(nil)

	finalString := hex.EncodeToString(finalBytes)

	return finalString
}
