package replace

func Exec(vars *Variable) (string, error) {
	var res string
	switch vars.Subcommand {
	case "codeToEnter":
		res = vars.codeToEnter()
	case "enterToCode":
		res = vars.enterToCode()
	}
	return res, nil
}
