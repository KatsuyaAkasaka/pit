package replace

import (
	"regexp"
)

func (v *Variable) codeToEnter() string {
	regxNewline := regexp.MustCompile(`\\r\\n|\\r|\\n`)
	return regxNewline.ReplaceAllString(v.Option.InputText, `
`)
}
