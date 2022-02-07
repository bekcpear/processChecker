package mail

import (
	"encoding/json"
	"log"
	"os"
	"processChecker/pkg/config"
	"testing"
)

func TestDo(t *testing.T) {
	c := new(config.Instances)

	b, err := os.ReadFile("/home/ryan/Tmp/processChecker/test.json")
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(b, c)
	if err != nil {
		log.Fatalln(err)
	}

	e := Do(&(*c)[0])
	if e != nil {
		log.Fatalln(e)
	}
}
