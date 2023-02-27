package cmd

const (
	greetingBanner = `
  _____ _    _ _____  _____  _____ 
 / ____| |  | |  __ \|_   _|/ ____|
| |    | |__| | |__) | | | | (___  
| |    |  __  |  _  /  | |  \___ \ 
| |____| |  | | | \ \ _| |_ ____) |
 \_____|_|  |_|_|  \_\_____|_____/
`
)

func init() {

}

func Execute() error {
	println(greetingBanner)
	return nil
}
