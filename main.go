package main

import (
	"fmt"
	"github.com/golang-template/pkg/api"
	"github.com/golang-template/pkg/loggers/logrus"
)

func main() {
	// Initialising config
	conf, err := configInit("config")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	// Initialising Log system
	myLog := logrus.NewLogrusLogger()
	myLog.SetLogLevel(conf.Log.Level)
	err = myLog.SetLogOutput(conf.Log.Out, conf.Log.Name)
	if err != nil {
		fmt.Printf("Unable to set logoutput. Eerror is %v", err)
		return
	}

	myLog.Info("Log Sysyem Initialised.")

	router := api.NewApi(myLog)
	router.StartRouter("127.0.0.1:4141")

}
