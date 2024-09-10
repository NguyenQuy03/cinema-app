package mailUtil

import "regexp"

func IsValidEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regex pattern
	re := regexp.MustCompile(emailRegex)

	// Use the regex pattern to match the email address
	return re.MatchString(email)
}
