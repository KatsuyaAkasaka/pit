package golang

func Exec(subcommand string) string {
	switch subcommand {
	case "iferr":
		return IfErr()
	default:
		return ""
	}
}
