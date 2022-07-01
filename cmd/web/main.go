package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/evolutionengine/bookings/cmd/pkg/config"
	"github.com/evolutionengine/bookings/cmd/pkg/handlers"
	"github.com/evolutionengine/bookings/cmd/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const portNo string = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	// change this to true later on
	app.InProduction = false

	// initialize session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalln("Error fetching template cache", err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	// Passing app config data to `templates`
	render.NewTemplates(&app)

	// Passing app config data to `handlers`
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	fmt.Printf("Starting server on http://127.0.0.1%s\n", portNo)

	srv := &http.Server{
		Addr:    portNo,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
