package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"processChecker/pkg/checker"
	"processChecker/pkg/config"
	"time"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "", "the config file")
	flag.Parse()

	b, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalln(err)
	}

	instances := new(config.Instances)
	err = json.Unmarshal(b, instances)
	if err != nil {
		log.Fatalln(err)
	}

	for i, _ := range *instances {
		checker.Job.Mu.Lock()
		checker.Job.Count++
		checker.Job.Mu.Unlock()
		go checker.Do(&(*instances)[i])
	}

	finished := make(chan bool)
	go func() {
		t := time.NewTicker(time.Millisecond * 500)
		for {
			select {
			case <-t.C:
				if checker.Job.Count < 1 {
					finished <- true
				}
			}
		}
	}()
	<-finished
}
