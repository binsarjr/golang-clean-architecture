package main

import (
	"giapps/newapp/infrastructure"
)

func main() {
	infrastructure.NewHTTPServer().Start()
}
