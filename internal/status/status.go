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

func getNames(name string) []string {
	var names = []string{name}
	if name == "all" {
		names = []string{}
		dirs, _ := os.ReadDir(fmt.Sprintf("%s/.goblaq/", home))
		for _, dir := range dirs {
			names = append(names, dir.Name())
		}
	}
	return names
}

func (s *Status) WriteStatus() {
	var d watch.Data
	names := getNames("all")
	for _, n := range names {
		data, err := d.Get(n)
		if err == nil {
			resp, _ := http.Get(fmt.Sprintf("%s://%s%s", data.Schema, data.Target, data.Path))
			currentTime := time.Now().Format("Mon Jan _2 15:04:05 MST 2006")
			s.Status = resp.StatusCode
			s.Timestamp = currentTime
			writeData, _ := yaml.Marshal(data)
			writeStatus, _ := yaml.Marshal(s)
			writeData = append(writeData, writeStatus...)
			os.WriteFile(fmt.Sprintf("%s/.goblaq/%s/data.yaml", home, n), writeData, os.ModePerm)
		}
	}
}

func (s *Status) Get(name string) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)
	defer w.Flush()
	names := getNames(name)
	fmt.Fprintf(
		w,
		"%s\t\t%s\t\t%s\n",
		"NAME",
		"STATUS",
		"TIMESTAMP",
	)
	for _, n := range names {
		file, _ := os.ReadFile(fmt.Sprintf("%s/.goblaq/%s/data.yaml", home, n))
		yaml.Unmarshal(file, s)
		fmt.Fprintf(w,
			"%s\t\t%d\t\t%s\n",
			n,
			s.Status,
			s.Timestamp,
		)
	}
}