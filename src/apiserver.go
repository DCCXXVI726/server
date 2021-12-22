package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/smtp"
	"os"
)

type ApiServer struct {
	logger *logrus.Logger
	router *mux.Router
	store  *Store
}

func NewServer() *ApiServer {
	return &ApiServer{
		logger: logrus.New(),
		router: mux.NewRouter(),
		store:  NewStore(),
	}
}

func (s *ApiServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return fmt.Errorf("can't configure logger: %s", err)
	}

	bindAddr := os.Getenv("PORT")
	if bindAddr == "" {
		bindAddr = "8080"
	}

	s.configureRouting()

	if err := s.configureStore(); err != nil {
		return fmt.Errorf("can't configure store: %s", err)
	}
	defer s.store.Close()

	s.logger.Info("starting server")
	return http.ListenAndServe(":"+bindAddr, s.router)
}

func (s *ApiServer) configureLogger() error {
	Loglevel := os.Getenv("LOGLEVEL")
	if Loglevel == "" {
		Loglevel = "debug"
	}

	level, err := logrus.ParseLevel(Loglevel)
	if err != nil {
		return fmt.Errorf("can't parse level logger for string %s : %s", Loglevel, err)
	}

	s.logger.SetLevel(level)
	return nil
}

func (s *ApiServer) configureRouting() {
	s.router.HandleFunc("/", s.handleHome())
	s.router.HandleFunc("/registration", s.handleRegistration()).Methods("POST")
	s.router.HandleFunc("/confirm", s.handleConfirmEmail()).Queries("email", "{email}")
}

func (s *ApiServer) configureStore() error {
	if err := s.store.Open(); err != nil {
		fmt.Errorf("can't open store: %s", err)
	}
	return nil
}
func (s *ApiServer) handleConfirmEmail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
func (s *ApiServer) handleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to homePage!")
	}
}

func (s *ApiServer) handleRegistration() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var rec request
		if err := json.NewDecoder(r.Body).Decode(&rec); err != nil {
			err = fmt.Errorf("can't encode body in registration: %s", err)
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		if _, err := s.store.FindUserByEmail(rec.Email); err == nil {
			err = fmt.Errorf("user with email %s already registered", rec.Email)
			s.error(w, r, http.StatusConflict, err)
			return
		}
		from := "DCCXXVI726726@gmail.com"
		pass := os.Getenv("EMAIL_PASS")
		to := []string{rec.Email}
		smtpHost := "smtp.gmail.com"
		smtpPort := "587"

		message := []byte("My super secret message.")

		auth := smtp.PlainAuth("", from, pass, smtpHost)

		err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
		if err != nil {
			s.error(w, r, http.StatusConflict, err)
			return
		}
		s.respond(w, r, 404, map[string]string{"OK": "OK"})
	}
}

func (s *ApiServer) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *ApiServer) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
