package config

import "chemex/utils"

var ServiceRestarting = false

//ServiceName service name
var ServiceName = "chemex"

// Port port of server to open and listen to connections
var Port = utils.GetEnv("PORT", "10000")

// SQLHost host of mongodb to connect
var SQLHost = utils.GetEnv("SQLHost", "localhost")

//SQLPort is port of mysql
var SQLPort = utils.GetEnv("SQLPort", "3306")

// SQLDBName name of current project database
var SQLDBName = utils.GetEnv("SQLDBName", "chemex")
//SQLUserName is sql username
//var SQLUserName = utils.GetEnv("SQLUserName", "root")
var SQLUserName = utils.GetEnv("SQLUserName", "root")

//SQLPassword is sql password
//var SQLPassword = utils.GetEnv("SQLPassword", "root")
var SQLPassword = utils.GetEnv("SQLPassword", "558822lyz")
