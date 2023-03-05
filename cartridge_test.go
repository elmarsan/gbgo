package main

import (
	"strings"
	"testing"
)

func TestCartridge(t *testing.T) {

	t.Run("Load", func(t *testing.T) {
		c := &Cartridge{}

		err := c.load("rom/dmg-acid2.gb")
		if err != nil {
			t.Error(err)
		}

		if strings.Compare(c.header.title, "MG-ACID2") != 0 {
			t.Errorf("Wrong title (%s), expected (MG-ACID2)", c.header.title)
		}
	})
}
