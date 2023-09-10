package status

import (
	"fmt"
	"net/http"
	"os"
	"text/tabwriter"
	"time"

	"github.com/traksel/goblaq/internal/watch"
)

var home = os.Getenv("HOME")

type Status struct {
	Status    int    `yaml:"status"`
	Timestamp string `yaml:"timestamp"`
}

func getNames() []string {
	var names []string
	dirs, _ := os.ReadDir(fmt.Sprintf("%s/.goblaq/", home))
	for _, dir := range dirs {
		names = append(names, dir.Name())
	}
	return names
}

func (s *Status) Get(name string) error {
	var d watch.Data
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)
	defer w.Flush()
	names := []string{name}
	if name == "all" {
		names = getNames()
	}
	fmt.Fprintf(
		w,
		"%s\t\t%s\t\t%s\n",
		"NAME",
		"STATUS",
		"TIMESTAMP",
	)
	for _, n := range names {
		file, err := d.Get(n)
		if err == nil {
			resp, _ := http.Get(fmt.Sprintf("%s://%s%s", file.Schema, file.Target, file.Path))
			currentTime := time.Now().Format("Mon Jan _2 15:04:05 MST 2006")
			fmt.Fprintf(w,
				"%s\t\t%d\t\t%s\n",
				file.Name,
				resp.StatusCode,
				currentTime,
			)
		}
	}
	return nil
}
