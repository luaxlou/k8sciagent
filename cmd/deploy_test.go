package cmd

import (
	"os"
	"testing"
)

func TestDeploy(t *testing.T) {

	df := `From busybox
CMD ["echo","hello world"]
`

	os.WriteFile("Dockerfile", []byte(df), 0666)
	Deploy("a", "b", "c", "d", "e")
}
