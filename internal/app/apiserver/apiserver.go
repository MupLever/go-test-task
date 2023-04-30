package apiserver

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	config *Config
	router *mux.Router
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {

	s.configureRouter()
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/", s.handleMain())

}

func (s *APIServer) handleMain() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello")
		tmpl, _ := template.ParseFiles("templates/main_page.html")
		tmpl.Execute(w, "")
	}
}
