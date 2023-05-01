package apiserver

import (
	"html/template"
	"net/http"

	"github.com/MupLever/go-test-task/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.logger.Info("starting API Server...")
	s.configureRouter()
	if err := s.configureStore(); err != nil {
		return nil
	}
	return http.ListenAndServe(s.config.BindAddr, s.router)
}
func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)

	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/", s.handleMain())
}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}
	s.store = st
	return nil
}

func (s *APIServer) handleMain() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "Hello")
		tmpl, _ := template.ParseFiles("templates/main_page.html")
		tmpl.Execute(w, "")
	}
}
