package main

import (
	"canopy/internal/configs"
)

func main() {
	a := configs.AppConfig{}
	a.Setter()
	a.Getter()
}
