package main

import (
	"flag"
)

func init() {
	flag.IntVar(&API_PORT, "port", 10777, "api port")
	flag.StringVar(&CONFIG_FILE, "conf", "./service.conf", "config file")
	flag.StringVar(&DATASTORE_DIR, "datastore", "./datastore/", "datastore dir")
	flag.Parse()
}
