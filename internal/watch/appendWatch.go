package watch

import "fmt"

func AppendWatch(name string, target string, schema string, targetService []string) error {
	if schema == "" {
		schema = "http"
	}
	fmt.Printf(
		"SERVICE ADDED:\nName: %s\nTarget: %s\nSchema: %s\nTargets: %v",
		name,
		target,
		schema,
		targetService,
	)
	return nil
}
