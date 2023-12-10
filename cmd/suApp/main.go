package main

import (
	"suApp/internal/controller"
	"suApp/internal/view"
)

func main() {

	gui := &view.View{}
	cntl := &controller.Cntl{}
	gui.CreateWindow(cntl.BuildApp)
}
