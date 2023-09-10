package status

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/traksel/goblaq/internal/watch"
	"gopkg.in/yaml.v3"
)

var home = os.Getenv("HOME")

type Status struct {
	Status    int    `yaml:"status"`
	Timestamp string `yaml:"timestamp"`
}

func (s *Status) update(status int, currentTime string) *Status {
	s.Status = status
	s.Timestamp = currentTime
	return s
}

func (s *Status) Get(name string) error {
	var d watch.Data
	file, err := d.Get(name)
	if err == nil {
		resp, _ := http.Get(fmt.Sprintf("%s://%s%s", file.Schema, file.Target, file.Path))
		currentTime := time.Now().Format("Mon Jan _2 15:04:05 MST 2006")
		fmt.Printf(
			"NAME\tSTATUS\tTIMESTAMP\n%s\t%d\t%v",
			file.Name,
			resp.StatusCode,
			currentTime,
		)
		writeData, _ := yaml.Marshal(s.update(resp.StatusCode, currentTime))
		writeData2, _ := yaml.Marshal(file)

		writeData = append(writeData, writeData2...)
		os.WriteFile(fmt.Sprintf("%s/.goblaq/%s/data.yaml", home, name), writeData, os.ModePerm)
	}
	return nil
}
