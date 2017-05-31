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
	"fmt"
	// "github.com/go-sql-driver/mysql"
	// "log"
)

// Server represents the asset for this application
//
// A Server have multiple information to be stored.
//
// swagger:model GaleraServer
type GaleraServer struct {
	// // ID Server Generated
	// ID string `json:"ID"`

	// Server Name
	NodeName string `json:"NodeName"`
	// Server WSREP
	WsrepStatus string `json:"WsrepStatus"`
	// Server Dist
	Dist float64 `json:"Dist"`
	// Server Status
	Status string `json:"Status"`
	// Server Avg Message Queue
	AvgQueue float64 `json:"AvgQueue"`
	// Server Latency
	Latency float64 `json:"Latency"`
}

type DBConfig struct {
	Host     string
	User     string
	Port     int
	Password string
}

// Init connection to MongoDB
func init() {

}

func GetInfo() (GaleraServer, error) {
	Server := DBConfig{
		Host:     "fr0-ac-cmp-n01.cloud.airbus.corp",
		Port:     3306,
		User:     "maxscale",
		Password: "CIHblhmzv74eMYPjhUHO",
	}
	fmt.Println("in db.go:", Server)
	Node := GaleraServer{
		NodeName:    "fr0-ac-cmp-n01.cloud.airbus.corp",
		WsrepStatus: "ON",
		Dist:        55,
		Status:      "Sync",
		AvgQueue:    70,
		Latency:     12,
	}
	fmt.Println("Galera Output Example", Node)
	return Node, nil

}
