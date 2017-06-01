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
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
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
	// Cluster ID
	ClusterID string `json:"ClusterID"`
	// Cluster Size
	ClusterSize string `json:"ClusterSize"`
}

type DBConfig struct {
	Host     string
	User     string
	Port     string
	Password string
}

var Servers []DBConfig
var Node []GaleraServer

func GetInfo() ([]GaleraServer, error) {

	sheetData, err := ioutil.ReadFile("../config-node.json")
	s := []DBConfig{}
	err = json.Unmarshal(sheetData, &s)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// Reinit Servers slice
	var Servers []DBConfig

	// Loop on all node in config file json
	for i := range s {
		Servers = append(Servers, DBConfig{

			Host:     s[i].Host,
			Port:     s[i].Port,
			User:     s[i].User,
			Password: s[i].Password,
		})
	}

	// Reinit Node Value
	Node = []GaleraServer{}

	//Retrieve info dynamically
	for i := range Servers {
		dsn := builddsn(Servers[i].Host, Servers[i].Port, Servers[i].User, Servers[i].Password, "")

		db, err := sql.Open("mysql", dsn)
		if err != nil {
			fmt.Println("Failed to connect to database: %v", err)
		}
		err = db.Ping()
		if err != nil {
			fmt.Println("Failed to ping to database: %v", err)
			Node = append(Node, GaleraServer{
				NodeName:       Servers[i].Host,
				WsrepStatus:    "Error",
				Status:         "Error",
				WsrepConnected: "Error",
				Dist:           "Error",
				AvgQueue:       "Error",
				ControlPaused:  "Error",
				SendAvgQueue:   "Error",
				ClusterID:      "Error",
				ClusterSize:    "Error",
			})
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

		var clusterid string
		db.QueryRow("SHOW GLOBAL STATUS LIKE 'wsrep_cluster_conf_id'").Scan(&unused, &clusterid)

		var clustersize string
		db.QueryRow("SHOW GLOBAL STATUS LIKE 'wsrep_cluster_size'").Scan(&unused, &clustersize)

		Node = append(Node, GaleraServer{
			NodeName:       Servers[i].Host,
			WsrepStatus:    wsrep,
			Status:         status,
			WsrepConnected: wsrepconnected,
			Dist:           dist,
			AvgQueue:       avgqueue,
			ControlPaused:  controlpaused,
			SendAvgQueue:   sendavgqueue,
			ClusterID:      clusterid,
			ClusterSize:    clustersize,
		})

	}
	return Node, nil

}

func builddsn(host, port, user, pass, name string) string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", user, pass, host, port, name)
}
