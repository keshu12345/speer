package main

import (
	"flag"
	"log"
	"os"

	"github.com/keshu12345/notes/config"
	"github.com/keshu12345/notes/db"
	"github.com/keshu12345/notes/notesService"

	// "github.com/keshu12345/notes/noteService"
	"github.com/keshu12345/notes/server/router"
	"go.uber.org/fx"
)

var configDirPath = flag.String("config", "", "path for config dir")

func main() {

	flag.Parse()
	log.New(os.Stdout, "", 0)
	app := fx.New(
		config.NewFxModule(*configDirPath, ""),
		// noteService.Module,
		router.Module,
		db.Module,
		notesService.Module,
	)
	app.Run()
}
