package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kazijawad/Photoshoot/controllers"
	"github.com/kazijawad/Photoshoot/models"
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
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	// us.DestructiveReset()
	us.AutoMigrate()

	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers(us)

	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	http.ListenAndServe(":3000", r)
}
