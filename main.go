package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kazijawad/PhotoGallery/controllers"
	"github.com/kazijawad/PhotoGallery/models"
)

const (
	host     = "ec2-54-157-88-70.compute-1.amazonaws.com"
	port     = 5432
	user     = "czbyfpbnovxxsj"
	password = "3ac3643181cc898fa2f80f3b9f77d4cdeb514062a5e3780d30493eaa7f9cb193"
	dbname   = "d2jd17qkkc675u"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname)
	services, err := models.NewServices(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer services.Close()
	services.AutoMigrate()

	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers(services.User)

	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	r.Handle("/login", usersC.LoginView).Methods("GET")
	r.HandleFunc("/login", usersC.Login).Methods("POST")
	r.HandleFunc("/cookietest", usersC.CookieTest).Methods("GET")
	http.ListenAndServe(":3000", r)
}
