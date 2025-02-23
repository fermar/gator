package main

import (
	// "errors"
	"fmt"
	"log"
	"os/user"

	"github.com/fermar/gator/internal/config"
	"github.com/fermar/gator/internal/logging"
)

// import "fmt"

func main() {
	logging.Lg.EnLog()
	conf, err := config.Read()
	if err != nil {
		log.Fatalln(err)
	}
	username, err := user.Current()
	if err != nil {

		log.Fatalln(err)
	}
	err = conf.SetUser(username.Username)
	if err != nil {

		log.Fatalln(err)
	}
	fmt.Println(conf)
	fmt.Println("bye...")
	fmt.Println("bye...")
	fmt.Println("bye...")
	fmt.Println("bye...")
}
