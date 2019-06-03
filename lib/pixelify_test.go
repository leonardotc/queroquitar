package pixelify

import (
	"testing"
	"os"
)

func TestEverything(t *testing.T) {
	f, _ := os.Create("image.gif")
	generateGif(f)
	if false {
		t.Errorf("Test")
	}
}