package gonkey

import (
	"testing"
)

func TestCreateWarehouse(t *testing.T) {
	TestRunner{
		caseDirs: []string{"cases/warehouses"},
	}.Run(t)
}
