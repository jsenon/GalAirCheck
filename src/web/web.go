// Package web CloudTab.
//
// the purpose of this package is to provide Web HTML Interface
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http
//     Host: localhost
//     BasePath: /
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Julien SENON <julien.senon@gmail.com>

package web

import (
	"db"
	"fmt"
	"html/template"
	"net/http"
)

// TO DO
// Have list of servers, username and password used to connect
var Server db.DBConfig

// Present Information on Dedicated WebPortal

// Func to display all server
func Index(res http.ResponseWriter, req *http.Request) {

	fmt.Println("in web.go:", Server)

	rs, err := db.GetInfo()
	fmt.Println("rs in web.go:", rs)

	t, _ := template.ParseFiles("templates/index.html")
	fmt.Println(t)
	t.Execute(res, rs)
	if err != nil {
		return
	}
}

func Toto(res http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")

	t.Execute(res, req)

}

func Login(res http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("templates/login.html")

	t.Execute(res, req)

}
