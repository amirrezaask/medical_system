package main

import (
	"encoding/json"
	"fmt"
	"medical_system/config"
	"medical_system/database/models"
	"medical_system/handlers"
	"medical_system/services"
	"medical_system/transport/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{}

// to run the app
// arg 1 is host:port
var serve = &cobra.Command{
	Use:   `serve`,
	Short: `serves a server from configuration`,
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Instance
		server := http.NewHTTPServer()
		db, err := models.Open("sqlite3", "file:ent?&_fk=1")
		if err != nil {
			panic(err)
		}

		http.RouteRegisterers = append(http.RouteRegisterers, handlers.NewUsersHandler(services.NewUserService(db.User, &services.AuthService{JWTSecret: []byte("this is a secret")})))
		if err := server(cfg.Servers.Get("http").Addr); err != nil {
			panic(err)
		}
	},
}

// application related commands
// handlers => list of application routes
// config => list of configuration.
var app = &cobra.Command{
	Use:   `app`,
	Short: ``,
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var app_handlers = &cobra.Command{
	Use:   `handlers`,
	Short: `list of handlers.`,
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		http.Routes(os.Stdout)
	},
}

var app_config = &cobra.Command{
	Use:   `config`,
	Short: `list of configurations.`,
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		bs, _ := json.MarshalIndent(config.Instance, "", " ")
		fmt.Println(string(bs))
	},
}

// to do db stuff
// models generate => ent generate
// models describe => database drivers + ent describe
// migration new => ent generate and new migration from ent
// migration apply => migration:new plus running on DB --get db connection name
var db = &cobra.Command{}

func main() {
	app.AddCommand(app_handlers)
	app.AddCommand(app_config)
	root.AddCommand(serve, app)
	if err := root.Execute(); err != nil {
		panic(err)
	}
}
