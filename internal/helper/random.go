package helper

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strconv"
	"time"
)

func GenerateRandomState() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func GenerateUniqueSlug(base string) string {
	suffix := strconv.FormatInt(time.Now().Unix(), 10)
	return fmt.Sprintf("%s-%s", base, suffix)
}
