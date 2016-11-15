package task

import (
	"bytes"
	"fmt"
	"github.com/lcaballero/griller/config"
	"github.com/lcaballero/griller/embedded"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// ErrNoTemplateFiles occurs when the template name doesn't map
// to any files.
var ErrNoTemplateFiles = fmt.Errorf("no template files found error")

// ErrTemplatePlateMissingParams occurs when template command is given
// without the rest of the required parameters.
var ErrTemplatePlateMissingParams = fmt.Errorf("template plate missing params error")

// Generate constructs a Gen instance with the given conf and
// immediately call Run returning the error.
func Generate(conf *config.Conf) error {
	return NewGen(conf).Run()
}

// Gen is project generator based on a configuration.
type Gen struct {
	conf *config.Conf
	log  *Log
}

// NewGen allocates a Gen instance capable of producing
// Go project generator based on the given config.
func NewGen(config *config.Conf) *Gen {
	return &Gen{
		conf: config,
		log:  NewLog(config),
	}
}

// TemplateData exposes some key values from the generator for debugging
// purposes.
func (g *Gen) TemplateData() Data {
	return Data{
		PackageName: g.conf.Template.Project,
		Remote:      g.conf.Remote,
	}
}

// Run carries out project generation, making directories and files,
// returning an error if one is produced during generation.
func (g *Gen) Run() error {
	tp := g.conf.Template
	if tp.Name == "" || tp.Project == "" {
		return ErrTemplatePlateMissingParams
	}

	g.log.Println("generating new project named:", g.conf.Template.Project)
	dest := filepath.Join(g.conf.Dest, g.conf.Template.Project)
	prefix, assets, dirs := g.TemplateAssets()

	g.log.Println("prefix:", prefix)
	g.log.Println("assets:", assets)
	g.log.Println("dirs:", dirs)

	if len(assets) <= 0 {
		g.log.Println("source template name:", g.conf.Template.Name)
		return ErrNoTemplateFiles
	}

	g.log.Println("creating root", dest)

	err := os.MkdirAll(dest, 0777)
	if err != nil {
		return err
	}

	err = g.mkdirs()
	if err != nil {
		return err
	}

	err = g.cp()
	if err != nil {
		return err
	}

	return nil
}

// TemplateAssets provides those assets for the configured template name.
func (g *Gen) TemplateAssets() (string, []string, []string) {
	names := make([]string, 0)
	dirSet := make(map[string]struct{}, 0)
	prefix := fmt.Sprintf("%s/", g.conf.Template.Name)
	n := len(prefix)
	assets := embedded.AssetNames()

	for _, name := range assets {
		value := name[n:]
		isForTemplate := strings.HasPrefix(name, prefix)

		dir := filepath.Dir(name)

		_, ok := dirSet[dir]
		isDir := !ok && dir != "" && dir != "." && dir != g.conf.Template.Name && isForTemplate

		if isDir {
			d := dir[n:]
			dirSet[d] = struct{}{}
		}

		if isForTemplate {
			names = append(names, value)
		}
	}

	dirs := make([]string, 0)
	for k := range dirSet {
		if k != "" {
			dirs = append(dirs, k)
		}
	}

	return prefix, names, dirs
}

// cp copies assets directly from the source to the destination producing
// an error if one occurs during generation.
func (g *Gen) cp() error {
	target := filepath.Join(g.conf.Dest, g.conf.Template.Project)
	prefix, assets, _ := g.TemplateAssets()

	for _, name := range assets {
		dest := filepath.Join(target, name)

		bin, err := embedded.Asset(prefix + name)
		if err != nil {
			return err
		}

		tpls := ".tpl.go"
		if strings.HasSuffix(dest, tpls) {
			dest = fmt.Sprintf("%s.go", dest[:len(dest) - len(tpls)])
			g.process(dest, bin)
			continue
		}

		err = g.write(dest, bin)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Gen) write(dest string, bin []byte) error {
	g.log.Println("creating file:", dest)
	return ioutil.WriteFile(dest, bin, 0777)
}

// mkdirs produces directories where all new files will exist.
func (g *Gen) mkdirs() error {
	_, _, dirs := g.TemplateAssets()
	target := filepath.Join(g.conf.Dest, g.conf.Template.Project)

	for _, name := range dirs {
		dest := filepath.Join(target, name)

		info, err := os.Stat(dest)
		if err == nil && !info.IsDir() {
			continue
		}
		if err != nil && !os.IsNotExist(err) {

		}

		g.log.Println("creating directory:", name, dest)

		err = os.Mkdir(dest, 0777)
		if err != nil {
			return err
		}

		g.log.Println("created directory:", dest)
	}
	return nil
}

// process template found in src and writes the value to the destination.
func (g *Gen) process(dest string, src []byte) error {
	g.log.Printf("processing template: %s", dest)

	p := template.New(g.conf.Template.Project)
	tp, err := p.Parse(string(src))
	if err != nil {
		g.log.Println(err)
		return err
	}

	data := g.TemplateData()
	g.log.Printf("data: %s", data)
	buf := bytes.NewBuffer([]byte{})
	err = tp.Execute(buf, data)
	if err != nil {
		g.log.Println(err)
		return err
	}

	return ioutil.WriteFile(dest, buf.Bytes(), 0777)
}
