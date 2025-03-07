package main

import (
	"context"
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
	logging.Lg.Logger.Printf("conexion: %v\n", stat.conf.DbURL)
	db, err := sql.Open("postgres", stat.conf.DbURL)
	if err != nil {
		log.Fatalln(err)
	}
	stat.db = database.New(db)
	coms := commands{comandos: make(map[string]func(*state, command) error)}

	coms.register("login", handlerLogin)
	coms.register("register", handlerRegister)
	coms.register("reset", handlerReset)
	coms.register("users", handlerUsers)
	coms.register("agg", handlerAgg)
	coms.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	coms.register("feeds", handlerFeeds)
	coms.register("follow", middlewareLoggedIn(handlerFollow))
	coms.register("following", middlewareLoggedIn(handlerFollowing))
	coms.register("unfollow", middlewareLoggedIn(handlerUnfollow))

	args := os.Args
	if len(args) < 2 {
		log.Fatalln("Faltan comandos de ejecución")
	}
	com := command{name: args[1], args: args[2:]}

	err = coms.run(&stat, com)
	if err != nil {
		fmt.Printf("error en ejecucion del comando %v\n", com.name)
		log.Fatalln(err)
	}
	fmt.Println("bye...")
}

func middlewareLoggedIn(
	handler func(s *state, cmd command, user database.User) error,
) func(*state, command) error {
	return func(s *state, cmd command) error {
		usrInfo, err := s.db.GetUser(context.Background(), s.conf.CurrentUserName)
		if err != nil {
			return err
		}
		return handler(s, cmd, usrInfo)
	}
}
