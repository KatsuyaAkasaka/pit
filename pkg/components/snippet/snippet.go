package snippet

import (
	"fmt"

	"github.com/KatsuyaAkasaka/pit/pkg/components/snippet/golang"
)

func Exec(vars *Variable) (error, string) {
	var res string
	switch vars.Option.Language {
	case "go":
		res = golang.Exec(vars.Subcommand)
	}
	fmt.Println(res)
	return nil, res
}
