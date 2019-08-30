package allchown

import (
	"testing"
)

func TestChange(t *testing.T) {
	if err := Change(".", 502, 20); err != nil {
		t.Error(err)
	}
}

func TestChangeAs(t *testing.T) {
	if err := ChangeAs(".", "../n.txt"); err != nil {
		t.Error(err)
	}
}
