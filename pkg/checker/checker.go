package checker

import (
	"fmt"
	"log"
	"os"
	"processChecker/pkg/config"
	"processChecker/pkg/mail"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

type job struct {
	Mu    sync.Mutex
	Count int
}

var Job = new(job)

func readPid(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() { _ = file.Close() }()

	info, err := file.Stat()
	if err != nil {
		log.Fatalln(err)
	}
	if info.IsDir() {
		log.Fatalln(fmt.Errorf("read pid from %s: cannot be a directory", path))
	}

	if info.Size() > 10 {
		log.Fatalln(fmt.Errorf("read pid from %s: too large file", path))
	}

	var b = make([]byte, 10)
	l, err := file.Read(b)
	if err != nil {
		log.Fatalln(err)
	}

	r, err := strconv.Atoi(strings.TrimSpace(string(b[:l])))
	if err != nil {
		log.Fatalln(err)
	}
	return r
}

func processExists(pid int) bool {
	process, err := os.FindProcess(pid)
	if err != nil {
		log.Fatalln(err)
	}
	err = process.Signal(syscall.Signal(0))
	if err != nil {
		if err.Error() == "operation not permitted" {
			log.Fatalln(err)
		}
		// TODO handle other unexpected errors
		return false
	}
	return true
}

func Do(config *config.Instance) {
	pid := readPid(config.Process.PidFile)
	t := time.NewTicker(config.Duration * time.Second)
	for {
		select {
		case <-t.C:
			if !processExists(pid) {
				err := mail.Do(config)
				if err != nil {
					log.Println(err)
				}
				// TODO use database to check the mail sent or not and
				//      do not exit this goroutine
				Job.Mu.Lock()
				Job.Count--
				Job.Mu.Unlock()
				return
			} else {
				log.Printf("running checker for %s, PID: %d ...",
					config.Process.Name, pid)
			}
		}
	}
}
