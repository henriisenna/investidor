package main

func main() {

	// setup various configuration for app
	config.LoadAllConfigs(".env")

	server.Serve()
}
