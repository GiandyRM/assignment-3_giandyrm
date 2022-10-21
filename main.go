package main

import (
	"assignment-3_giandyrm/routers"
)

func main() {
	routers.StartServer().Run(":8000")
}
