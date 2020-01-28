package main

import "prisma-golang-vue/src/system/app"

func main() {
	// client := prisma.New(nil)
	// ctx := context.TODO()
	s := app.NewServer()

	s.Init()
	s.Start()
}
