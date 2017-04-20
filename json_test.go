package goutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testProduct struct {
	Name  string
	Price float64 `json:",string"`
}

func TestReadJSON(t *testing.T) {
	jsonString := `{"name":"Galaxy Nexus", "price":"3460.00"}`

	var product testProduct
	expected := testProduct{
		Name:  "Galaxy Nexus",
		Price: 3460.00,
	}
	err := ReadJSON(jsonString, &product)
	assert.NoError(t, err)
	assert.Equal(t, expected, product)
}
