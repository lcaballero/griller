package task

import (
	"os"
	"path/filepath"
	"encoding/json"
	"errors"
)

/*
  A .griller JSON file can be located in the user's home directory
  which can provide the flag values for --dest and --remote.

  For example:

  {
    "Remote": "github.com/saber",
    "Dest": "$GOPATH/src/github.com/saber"
  }

 */
type Dotfile struct {
	Remote string // example: github.com/lcaballero
	Dest   string // example: github.com/lcaballero
}

const DefaultGrillerConf = ".griller"

var GrillerDoesNotExistError = errors.New("Couldn't locate .griller")

type DotLoader struct {
	Env     func(key string) (string, bool)
	DotName string
}

func NewDotLoader() *DotLoader {
	return &DotLoader{
		Env: os.LookupEnv,
		DotName: DefaultGrillerConf,
	}
}

func (d *DotLoader) Load() error {
	file, err := d.Read()
	if err != nil {
		return err
	}

	os.Setenv("GRILLER_DEST", os.ExpandEnv(file.Dest))
	os.Setenv("GRILLER_REMOTE", file.Remote)

	return nil
}

func (d *DotLoader) Read() (*Dotfile, error) {
	dot := NewDotLoader()
	f, err := dot.Open()
	if err != nil {
		return &Dotfile{}, err
	}
	defer f.Close()

	grill := &Dotfile{}
	err = json.NewDecoder(f).Decode(grill)
	if err != nil {
		return &Dotfile{}, err
	}

	return grill, nil
}

func (d *DotLoader) Home() (string, bool) {
	return d.Env("HOME")
}

func (d *DotLoader) Filename() (string, bool) {
	val, ok := d.Home()
	if !ok {
		return "", false
	}
	return filepath.Join(val, d.DotName), true
}

func (d DotLoader) Exists() (string, bool) {
	name, ok := d.Filename()
	if !ok {
		return "", false
	}
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return "", false
	}
	return name, true
}

func (d *DotLoader) Open() (*os.File, error){
	name, ok := d.Exists()
	if !ok {
		return nil, GrillerDoesNotExistError
	}
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	return f, nil
}
