// Package db GalAirCheck.
//
// the purpose of this package is to provide DB Interface to Galera Cluster
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http
//     Host: localhost
//     BasePath:
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Julien SENON <julien.senon@gmail.com>
//
package db

import (
	// "fmt"
	"github.com/go-sql-driver/mysql"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Server represents the asset for this application
//
// A Server have multiple information to be stored.
//
// swagger:model server
type Server struct {
	// ID Server Generated
	ID bson.ObjectId `json:"id" bson:"_id,omitempty"`
	// Server Name
	NodeName string `json:"NodeName"`
	// Server WSREP
	WsrepStatus string `json:"WsrepStatus"`
	// Server Dist
	Dist string `json:"Dist"`
	// Server Status
	Status string `json:"Status"`
	// Server Avg Message Queue
	AvgQueue string `json:"AvgQueue"`
	// Server Latency
	Latency string `json:"Latency"`
}

// Init connection to MongoDB
func init() {

	session, err := mgo.Dial("localhost")

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db = session.DB("cloudtab")
}

func GetAll(server string) ([]Server, error) {

}
