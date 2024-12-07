package gonkey

import (
	"testing"
)

func TestItemsCreate(t *testing.T) {
	TestRunner{
		caseDirs: []string{"cases/items/create"},
	}.Run(t)
}

func TestGetItemsQuantity(t *testing.T) {
	TestRunner{
		caseDirs: []string{"cases/items/quantity"},
	}.Run(t)
}
