package goutil

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
)

// ReadJSON Unmarshal json string into struct
func ReadJSON(s string, result interface{}) error {
	return json.Unmarshal([]byte(s), result)
}

// ReadJSONIO Unmarshal json stream into struct
func ReadJSONIO(r io.ReadCloser, result interface{}) error {
	defer r.Close()
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, result)
}

// ReadJSONFromFile Unmarshal json file into struct
func ReadJSONFromFile(filename string, result interface{}) error {
	rawFileData, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(rawFileData, result)
}

// WriteJSONToFile Write json marshalable struct to file
func WriteJSONToFile(filename string, data interface{}, indent int) error {
	rawData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// make json "pretty"
	var buf bytes.Buffer
	pad := strings.Repeat(" ", indent)
	err = json.Indent(&buf, rawData, "", pad)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, buf.Bytes(), 0644)
}
