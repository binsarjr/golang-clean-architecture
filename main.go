package main

import (
	"giapps/servisin/infrastructure"
)

func main() {
	infrastructure.NewHTTPServer().Start()
}
