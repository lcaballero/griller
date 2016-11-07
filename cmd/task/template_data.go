package task

import "encoding/json"

// Data here represents a few key data points used by the generator
// and is exported for ease of debugging.
type Data struct {
	PackageName string
	Remote      string
}

// String is a convenience for writing the Data instance in Json
func (d Data) String() string {
	bin, err := json.Marshal(&d)
	if err != nil {
		return err.Error()
	}
	return string(bin)
}
