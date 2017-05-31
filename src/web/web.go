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
//     Contact: Julien SENON <julien.senon@gmail.com

package web

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"html/template"
	"net/http"
)

// Present Information on Dedicated WebPortal

// Func to display all server on table view
func Index(res http.ResponseWriter, req *http.Request) {
	// var rs Server
	rs, err := db.GetAll()
	t, _ := template.ParseFiles("templates/index.html")

	t.Execute(res, rs)
	if err != nil {
		return
	}
}
