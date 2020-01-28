package main

import (
	"flag"
	"os"
	"prisma-golang-vue/src/system/app"

	"github.com/joho/godotenv"
)

var port string

func init() {
	flag.StringVar(&port, "port", "4000", "Assigning the port that the server should listen on.")

	flag.Parse()

	if err := godotenv.Load("config.ini"); err != nil {
		panic(err)
	}

	envPort := os.Getenv("PORT")
	if len(envPort) > 0 {
		port = envPort
	}

}

func main() {
	// client := prisma.New(nil)
	// ctx := context.TODO()
	s := app.NewServer()

	s.Init(port)
	s.Start()
}
