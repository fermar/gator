package main

import (
	// "errors"
	"fmt"
	"log"
	"os"

	// "os/user"

	"github.com/fermar/gator/internal/config"
	"github.com/fermar/gator/internal/logging"
)

// import "fmt"

type state struct {
	conf config.Config
}

func main() {
	logging.Lg.EnLog()
	var stat = state{}
	var err error
	stat.conf, err = config.Read()
	if err != nil {
		log.Fatalln(err)
	}
	var coms commands
	coms = commands{comandos: make(map[string]func(*state, command) error)}
	coms.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		log.Fatalln("Faltan comandos de ejecución")
	}
	var com = command{name: args[1], args: args[2:]}

	err = coms.run(&stat, com)
	if err != nil {
		log.Fatalln(err)
	}
	// username, err := user.Current()
	// if err != nil {
	//
	// 	log.Fatalln(err)
	// }
	// // err = conf.SetUser(username.Username)
	// if err != nil {
	//
	// 	log.Fatalln(err)
	// }
	// fmt.Println(conf)
	fmt.Println("bye...")
}
