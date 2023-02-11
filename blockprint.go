package main

import (
	"fmt"
	"sort"
	"strings"

	figure "github.com/common-nighthawk/go-figure"
	"github.com/tucats/gopackages/app-cli/cli"
	"github.com/tucats/gopackages/app-cli/settings"
	"github.com/tucats/gopackages/app-cli/tables"
)

var fontSet []string

func FontsAction(c *cli.Context) error {
	initFonts()
	t, _ := tables.New([]string{"Font Name"})
	for _, font := range fontSet {
		t.AddRowItems(font)
	}

	t.SetPagination(0, 0)
	t.Print("text")

	return nil
}

func PrintAction(c *cli.Context) error {
	font := settings.Get("default.font")
	if font == "" {
		font = "standard"
		settings.Set("default.font", font)
		settings.Save()
	}

	if c.WasFound("font") {
		font, _ = c.String("font")
	}

	text := c.FindGlobal().Parameter(0)

	return blockPrint(font, text)
}

func blockPrint(font, text string) error {
	initFonts()

	if font == "" {
		font = settings.Get("default.font")
	}

	if font == "" {
		font = "standard"
	}

	if !isFont(font) {
		return fmt.Errorf("no such font: %s", font)
	}

	myFigure := figure.NewFigure(text, font, true)
	fmt.Println(myFigure.String())

	return nil
}

func initFonts() {
	if fontSet == nil {
		fontSet = figure.AssetNames()
		sort.Strings(fontSet)

		for idx := 0; idx < len(fontSet); idx++ {
			fontSet[idx] = strings.TrimPrefix(strings.TrimSuffix(fontSet[idx], ".flf"), "fonts/")
		}
	}
}

func isFont(name string) bool {
	for _, font := range fontSet {
		if strings.EqualFold(name, font) {
			return true
		}
	}

	return false
}
