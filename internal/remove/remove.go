package remove

import (
	"fmt"
	"os"
)

var home = os.Getenv("HOME")

func Remove(name string) error {
	return os.RemoveAll(fmt.Sprintf("%s/.goblaq/%s", home, name))
}
