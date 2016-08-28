package task

import "encoding/json"

type Data struct {
	PackageName string
	Remote      string
}

func (d Data) String() string {
	bin, err := json.Marshal(&d)
	if err != nil {
		return err.Error()
	}
	return string(bin)
}
