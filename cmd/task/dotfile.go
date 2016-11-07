package task

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

// Dotfile represents a configuration file (ie .griller JSON file) can
// be located from the user's home directory
// which can provide the flag values for --dest and --remote.
//
// For example:
//
// {
//   "Remote": "github.com/saber",
//   "Dest": "$GOPATH/src/github.com/saber"
// }
type Dotfile struct {
	Remote string // example: github.com/lcaballero
	Dest   string // example: github.com/lcaballero
}

// DefaultGrillerConf is the name of the file a user can add to their
// home directory and it's name defaults to ~/.griller
const DefaultGrillerConf = ".griller"

// ErrGrillerDoesNotExist occurs internally when the .griller file can
// not be located.
var ErrGrillerDoesNotExist = errors.New("error ~/.griller does not exist")

// DotLoader is a struct used to read various 'dot' files.
type DotLoader struct {
	Env     func(key string) (string, bool)
	DotName string
}

// NewDotLoader creates a DotLoader that will read values from the
// environment and look for the DefaultGrillerConf file.
func NewDotLoader() *DotLoader {
	return &DotLoader{
		Env:     os.LookupEnv,
		DotName: DefaultGrillerConf,
	}
}

// Load injects GRILLER_DEST and GRILLDER_REMOVE in the environment from
// those values found in the griller file.
func (d *DotLoader) Load() error {
	file, err := d.Read()
	if err != nil {
		return err
	}

	os.Setenv("GRILLER_DEST", os.ExpandEnv(file.Dest))
	os.Setenv("GRILLER_REMOTE", file.Remote)

	return nil
}

// Read parses the griller file found and deserializes into a Dotfile.
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

// Home returns the value HOME in the environment.
func (d *DotLoader) Home() (string, bool) {
	return d.Env("HOME")
}

// Filename joins the value of HOME and the griller file name.
func (d *DotLoader) Filename() (string, bool) {
	val, ok := d.Home()
	if !ok {
		return "", false
	}
	return filepath.Join(val, d.DotName), true
}

// Exists checks that the griller file does indeed exist.
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

// Open opens the griller file for reading.
func (d *DotLoader) Open() (*os.File, error) {
	name, ok := d.Exists()
	if !ok {
		return nil, ErrGrillerDoesNotExist
	}
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	return f, nil
}
