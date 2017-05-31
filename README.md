# GalAirCheck

Simple Web Server in GO for Galera Cluster HealthCheck

### Prerequisite

You need to have:

* Go 1.8
* Go Environment properly set
* Galera Cluster

### Compilation

```sh
git clone https://github.com/jsenon/GalAirCheck.git
go build -o galaircheck
```

### Galera Node File Configuration

Edit and rename config-node.json.example to config-node.json
Replace value by your own

```sh
"Host": "MyNode",
"Port": "MyPort",
"User": "MyUser",
"Password": "MyPassword"
```

### Access

Access through favorite web browser on http://YOURIP:YOURPORT ie http://127.0.0.1:9030


### API

Auto Health Check API status available at /healthy/am-i-up

### ToDo

- [ ] API
- [x] Web part to view details
- [ ] API Doc
- [ ] Build Docker App
- [x] Add API monitoring
- [ ] Add Cluster ID
- [ ] Add Cluster Size





