package status

import (
	"crypto/tls"
	"crypto/x509"
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
	Message   string `yaml:"message"`
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
	d := watch.NewData()
	names := getNames("all")
	client := &http.Client{}
	for _, n := range names {
		data, err := d.Get(n)
		if err == nil {
			if data.Chain != "" {
				cert, _ := os.ReadFile(data.Chain)
				certsPool := x509.NewCertPool()
				certsPool.AppendCertsFromPEM(cert)
				client = &http.Client{
					Transport: &http.Transport{
						TLSClientConfig: &tls.Config{
							InsecureSkipVerify: true,
							RootCAs:            certsPool,
						},
					},
				}
			}
			resp, _ := client.Get(fmt.Sprintf("%s://%s%s", data.Schema, data.Target, data.Path))
			currentTime := time.Now().Format("Mon Jan _2 15:04:05 MST 2006")
			s.Status = resp.StatusCode
			s.Message = resp.Status
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
	d := watch.NewData()
	w.Init(os.Stdout, 8, 8, 8, '\t', 0)
	defer w.Flush()
	names := getNames(name)
	fmt.Fprintf(
		w,
		"%s\t%s\t%s\t%s\t%s\n",
		"NAME",
		"URL",
		"STATUS",
		"TIMESTAMP",
		"MESSAGE",
	)
	for _, n := range names {
		d.Get(n)
		file, _ := os.ReadFile(fmt.Sprintf("%s/.goblaq/%s/data.yaml", home, n))
		yaml.Unmarshal(file, s)
		fmt.Fprintf(w,
			"%s\t%s\t%d\t%s\t%s\n",
			n,
			fmt.Sprintf("%s://%s", d.Schema, d.Target),
			s.Status,
			s.Timestamp,
			s.Message,
		)
	}
}
