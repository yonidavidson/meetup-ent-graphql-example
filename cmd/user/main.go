package main

import (
	"context"
	"log"
	"net/http"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/alecthomas/kong"

	"user"

	_ "github.com/mattn/go-sqlite3"

	"user/ent"
	"user/ent/migrate"
)

func main() {
	var cli struct {
		Addr  string `name:"address" default:":8081" help:"Address to listen on."`
		Debug bool   `name:"debug" help:"Enable debugging mode."`
	}
	kong.Parse(&cli)
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatal("opening ent client", err)
	}
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("running schema migration", err)
	}
	if cli.Debug {
		client = client.Debug()
	}
	srv := handler.NewDefaultServer(user.NewSchema(client))
	srv.Use(entgql.Transactioner{TxOpener: client})
	http.Handle("/",
		playground.Handler("User", "/query"),
	)
	http.Handle("/query", srv)
	log.Println("listening on", cli.Addr)
	if err := http.ListenAndServe(cli.Addr, nil); err != nil {
		log.Fatal("http server terminated", err)
	}
}
