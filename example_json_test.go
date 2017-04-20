package goutil

func ExampleReadJSON() {
	type Product struct {
		Name  string
		Price float64 `json:",string"`
	}

	jsonString := `{"name":"Galaxy Nexus", "price":"3460.00"}`
	var product Product
	_ = ReadJSON(jsonString, &product)
}

func ExampleReadJSONFromFile() {
	type Product struct {
		Name  string
		Price float64 `json:",string"`
	}

	var product Product
	_ = ReadJSONFromFile("./path/to/file.json", &product)
}

func ExampleWriteJSONToFile() {
	type Product struct {
		Name  string
		Price float64 `json:",string"`
	}

	product := Product{
		Name:  "Galaxy Nexus",
		Price: 3460.00,
	}
	_ = WriteJSONToFile("./path/to/file.json", product, 4)
}
