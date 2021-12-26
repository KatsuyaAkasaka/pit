package replace

import "strings"

func (v *Variable) enterToCode() string {
	return strings.Join(strings.Split(v.Option.InputText, "\n"), `\n`)
}
