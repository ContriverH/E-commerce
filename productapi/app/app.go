package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/codernishchay/productapi/app/handler"
	"github.com/codernishchay/productapi/config"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Router *mux.Router
	DB     *mongo.Database
}

func ConfigAndRunApp(config *config.Config) {
	app := new(App)
	app.Initialize(config)
	app.Run("127.0.0.1:8081")
}

func (app *App) Run(host string) {
	// use signals for shutdown server gracefully.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt, os.Kill)
	go func() {
		log.Fatal(http.ListenAndServe(host, app.Router))
	}()
	log.Printf("Server is listning on http://%s\n", host)
	sig := <-sigs
	log.Println("Signal: ", sig)

	log.Println("Stoping MongoDB Connection...")
	app.DB.Client().Disconnect(context.Background())
}

func (app *App) Initialize(config *config.Config) {
	// app.DB = db.InitialConnection("golang", config.MongoURI())
	// app.createIndexes()
	app.DB = InitialConnection("golang", "mongodb://localhost:27017")
	app.Router = mux.NewRouter()
	app.setRouters()
}

func (app *App) setRouters() {
	app.Post("/addproduct", app.handleRequest(handler.Createproduct))
	// 	app.Patch("/person/{id}", app.handleRequest(handler.UpdatePerson))
	// 	app.Put("/person/{id}", app.handleRequest(handler.UpdatePerson))
	// 	app.Get("/person/{id}", app.handleRequest(handler.GetPerson))
	app.Get("/product", app.handleRequest(handler.GetProducts))
	// 	app.Get("/person", app.handleRequest(handler.GetPersons), "page", "{page}")
	// }
}

func (app *App) Post(route string, endpoint http.HandlerFunc, queries ...string) {
	app.Router.HandleFunc(route, endpoint).Methods("POST").Queries(queries...)
}
func (app *App) Get(route string, endpoint http.HandlerFunc, queries ...string) {
	app.Router.HandleFunc(route, endpoint).Methods("POST").Queries(queries...)
}
func (app *App) Put(route string, endpoint http.HandlerFunc, queries ...string) {
	app.Router.HandleFunc(route, endpoint).Methods("POST").Queries(queries...)
}
func (app *App) Patch(route string, endpoint http.HandlerFunc, queries ...string) {
	app.Router.HandleFunc(route, endpoint).Methods("POST").Queries(queries...)
}
func (app *App) Delete(route string, endpoint http.HandlerFunc, queries ...string) {
	app.Router.HandleFunc(route, endpoint).Methods("POST").Queries(queries...)
}

type RequestHandlerFunction func(db *mongo.Database, r http.ResponseWriter, w *http.Request)

func (app *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(app.DB, w, r)
	}
}
