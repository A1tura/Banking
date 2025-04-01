package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func PasswordLeaked(password string) (bool, int) {
	hasher := sha1.New()
	hasher.Write([]byte(password))
	hash := hex.EncodeToString(hasher.Sum(nil))

	prefix := strings.ToUpper(hash[:5])
	suffix := strings.ToUpper(hash[5:])

	req, err := http.Get("https://api.pwnedpasswords.com/range/" + prefix)
	if err != nil {
		return false, 0
	}
	res, err := io.ReadAll(req.Body)
	if err != nil {
		return false, 0
	}
	defer req.Body.Close()

	for _, line := range strings.Split(string(res), "\n") {
        parts := strings.Split(strings.TrimSpace(line), ":")
		if parts[0] == suffix {
			leaked, err := strconv.Atoi(parts[1])
			if err != nil {
				return true, 0
			}
			return true, leaked
		}
	}

	return false, 0
}

func VerifyPasswordStrength(password string) bool {
	if len(password) < 8 && len(password) > 128 {
		return false
	}

	return true
}
