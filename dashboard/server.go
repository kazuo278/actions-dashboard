package main

import (
	"github/dashboard/infrastructure"
	"github/dashboard/controller"
)

func main() {
	infrastructure.Init()
	controller.Init()
}