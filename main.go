package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/YAWAL/ETLUNL/L"

	"github.com/YAWAL/ETLUNL/E"
	"github.com/YAWAL/ETLUNL/T"
	"github.com/YAWAL/ETLUNL/config"
	. "github.com/YAWAL/ETLUNL/logging"
)

func main() {

	configFile := flag.String("config", "./config.json", "config file path")
	flag.Parse()
	// read, config from form
	conf, err := config.ReadConfig(configFile)
	if err != nil {
		Log.Errorf("Cannot load config: ", err.Error())
		return
	}

	rawData, err := E.Extract("E/ir.csv")
	if err != nil {
		log.Println(err)
	}

	for _, row := range rawData {
		fmt.Println(row)
	}

	preparedDate, err := T.Transform(rawData)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("PREPARED:")
	for _, row := range preparedDate {
		fmt.Println(row)
	}

	pg, err := L.Postgres(conf.Postgres)
	if err != nil {
		log.Println(err)
	}
	pg.Close()

}
