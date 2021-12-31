package replace

func Exec(vars *Variable) (string, error) {
	var res string
	var err error
	switch vars.Subcommand {
	case "codeToEnter":
		res = vars.codeToEnter()
	case "enterToCode":
		res = vars.enterToCode()
	case "protoNum":
		res, err = vars.protoNum()
	}
	return res, err
}
