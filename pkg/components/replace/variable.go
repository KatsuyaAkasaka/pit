package replace

type Options struct {
	InputText string
}

type Variable struct {
	Subcommand string
	Option     Options
}
