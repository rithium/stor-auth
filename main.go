package main

import (
	"flag"
	"log"
	"os"
	"github.com/rithium/stor-auth/model"
	"github.com/rithium/version"
	"github.com/rithium/stor-auth/config"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"github.com/urfave/negroni"
	"fmt"
	"time"
)

type Env struct {
	db model.Datastore
}

func init() {
	versionFlag := flag.Bool("v", false, "prints version")
	configFlag := flag.Bool("c", false, "dumps configuration")

	flag.Parse()

	if *versionFlag {
		log.Println("Stor Data", version.GetVersion())
		os.Exit(0)
	}

	config.LoadConfig()

	if *configFlag {
		log.Printf("HTTP:\t%+v\n", config.HttpServer)
		log.Printf("MySQL:\t%+v\n", config.MySQL)

		os.Exit(0)
	}
}

func main() {
	db, err := model.NewDb(config.MySQL)

	if err != nil {
		log.Panic(err)
	}

	if db == nil {
		log.Panic("mysql", err)
	}

	env := &Env{db}

	log.Println("Stor Auth", version.GetVersion())

	router := mux.NewRouter()

	router.HandleFunc("/health", env.handleHealth)

	router.HandleFunc("/key/{key}", env.handleGetKey).Methods("GET")
	router.HandleFunc("/key", env.handlePostKey).Methods("POST")

	n := negroni.New()

	// Convert panics to 500 responses
	n.Use(negroni.NewRecovery())

	// Pretty print REST requests
	//n.Use(negroni.NewLogger())

	n.UseHandler(router)

	addr := fmt.Sprintf("%s:%d", config.HttpServer.Uri, config.HttpServer.Port)

	serv := &http.Server{
		Addr:           addr,
		Handler:        n,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Binding HTTP on", addr)

	log.Fatal("http serv:", serv.ListenAndServe())
}

func (env *Env) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (env *Env) handleGetKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	key := vars["key"]

	result, err := env.db.FindActiveKey(key)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%+v", result)

	if result == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func (env *Env) handlePostKey(w http.ResponseWriter, r *http.Request) {
	apiKey, err := env.db.CreateApiKey()

	if err != nil {
		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
	}

	log.Printf("%+v", apiKey)

	json.NewEncoder(w).Encode(apiKey)
}



