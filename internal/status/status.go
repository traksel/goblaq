package status

import (
	"fmt"
	"net/http"
	"os"
	"text/tabwriter"
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
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)
	defer w.Flush()
	if name == "all" {
		names, _ := os.ReadDir(fmt.Sprintf("%s/.goblaq/", home))
		for id, n := range names {
			file, err := d.Get(n.Name())
			if err == nil {
				resp, _ := http.Get(fmt.Sprintf("%s://%s%s", file.Schema, file.Target, file.Path))
				currentTime := time.Now().Format("Mon Jan _2 15:04:05 MST 2006")
				if id == 0 {
					fmt.Fprintf(
						w,
						"%s\t%s\t%s\n",
						"NAME",
						"STATUS",
						"TIMESTAMP",
					)
				}
				fmt.Fprintf(w,
					"%s\t%d\t%s\n",
					file.Name,
					resp.StatusCode,
					currentTime,
				)
				writeData, _ := yaml.Marshal(file)
				writeData2, _ := yaml.Marshal(s.update(resp.StatusCode, currentTime))
				writeData = append(writeData, writeData2...)
				os.WriteFile(fmt.Sprintf("%s/.goblaq/%s/data.yaml", home, name), writeData, os.ModePerm)
			}
		}
	} else {
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
			writeData, _ := yaml.Marshal(file)
			writeData2, _ := yaml.Marshal(s.update(resp.StatusCode, currentTime))
			writeData = append(writeData, writeData2...)
			os.WriteFile(fmt.Sprintf("%s/.goblaq/%s/data.yaml", home, name), writeData, os.ModePerm)
		}
	}
	return nil
}
