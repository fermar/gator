package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/fermar/gator/internal/config"
	"github.com/fermar/gator/internal/database"
	"github.com/fermar/gator/internal/logging"
)

// import "fmt"

type state struct {
	db   *database.Queries
	conf *config.Config
}

func main() {
	logging.Lg.EnLog()
	stat := state{}
	var err error
	stat.conf, err = config.Read()
	if err != nil {
		log.Fatalln(err)
	}

	db, err := sql.Open("postgres", stat.conf.DbURL)
	if err != nil {
		log.Fatalln(err)
	}
	stat.db = database.New(db)
	var coms commands
	coms = commands{comandos: make(map[string]func(*state, command) error)}
	coms.register("login", handlerLogin)
	coms.register("register", handlerRegister)
	coms.register("reset", handlerReset)

	args := os.Args
	if len(args) < 2 {
		log.Fatalln("Faltan comandos de ejecución")
	}
	com := command{name: args[1], args: args[2:]}

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
