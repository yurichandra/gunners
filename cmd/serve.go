package cmd

// Serve :nodoc:
func Serve() {
	loadEnv()
	bootstrap()
	initHTTP()
}
