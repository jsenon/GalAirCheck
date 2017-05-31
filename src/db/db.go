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
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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
	// Server WSREP Connected
	WsrepConnected string `json:"WsrepConnected"`
	// Server Dist
	Dist string `json:"Dist"`
	// Server Status
	Status string `json:"Status"`
	// Server Avg Message Queue
	AvgQueue string `json:"AvgQueue"`
	// Server Latency
	Latency string `json:"Latency"`
	// Control Paused
	ControlPaused string `json:"ControlPaused"`
	// Send Queue Average
	SendAvgQueue string `json:"SendAvgQueue"`
}

type DBConfig struct {
	Host     string
	User     string
	Port     string
	Password string
}

var Servers []DBConfig
var Node []GaleraServer

// Init connection to Galera
// func init() {
// 	// sql.Open don't open connection
// 	db, err := sql.Open("mysql", "maxscale:CIHblhmzv74eMYPjhUHO@tcp(fr0-ac-cmp-n01.cloud.airbus.corp:3306)/")
// 	if err != nil {
// 		fmt.Println("Failed to connect to database: %v", err)
// 	}
// 	err = db.Ping()
// 	if err != nil {
// 		panic(err.Error()) // proper error handling instead of panic in your app
// 	}
// 	var version string
// 	// QueryRow will open connection
// 	db.QueryRow("SELECT VERSION()").Scan(&version)
// 	fmt.Println("Connected to:", version)
// }

func GetInfo() ([]GaleraServer, error) {
	//Static Galera Node Server
	Servers = []DBConfig{}

	//Static Info for test Purpose
	Node = []GaleraServer{}

	//Retrieve info dynamically
	for i := range Servers {
		dsn := builddsn(Servers[i].Host, Servers[i].Port, Servers[i].User, Servers[i].Password, "")
		fmt.Println(dsn)

		db, err := sql.Open("mysql", dsn)
		if err != nil {
			fmt.Println("Failed to connect to database: %v", err)
		}
		err = db.Ping()
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		var version string
		db.QueryRow("SELECT VERSION()").Scan(&version)

		var wsrep string
		var unused string
		db.QueryRow("SHOW GLOBAL STATUS LIKE 'wsrep_ready'").Scan(&unused, &wsrep)

		var status string
		db.QueryRow("SHOW GLOBAL STATUS LIKE 'wsrep_local_state_comment'").Scan(&unused, &status)

		var wsrepconnected string
		db.QueryRow("SHOW GLOBAL STATUS LIKE 'wsrep_connected'").Scan(&unused, &wsrepconnected)

		var dist string
		db.QueryRow("SHOW STATUS LIKE 'wsrep_cert_deps_distance'").Scan(&unused, &dist)

		var avgqueue string
		db.QueryRow("SHOW STATUS LIKE 'wsrep_local_recv_queue_avg'").Scan(&unused, &avgqueue)

		var sendavgqueue string
		db.QueryRow("SHOW STATUS LIKE 'wsrep_local_send_queue_avg'").Scan(&unused, &sendavgqueue)

		var controlpaused string
		db.QueryRow("SHOW STATUS LIKE 'wsrep_flow_control_paused'").Scan(&unused, &controlpaused)

		Node = append(Node, GaleraServer{
			NodeName:       Servers[i].Host,
			WsrepStatus:    wsrep,
			Status:         status,
			WsrepConnected: wsrepconnected,
			Dist:           dist,
			AvgQueue:       avgqueue,
			ControlPaused:  controlpaused,
			SendAvgQueue:   sendavgqueue,
		})

	}
	return Node, nil

}

func builddsn(host, port, user, pass, name string) string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", user, pass, host, port, name)
}
