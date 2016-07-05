package task

import (
	"github.com/lcaballero/griller/config"
	"path/filepath"
	"os"
	"github.com/lcaballero/griller/embedded"
	"io/ioutil"
	"text/template"
	"strings"
	"bytes"
)

// Generate constructs a Gen instance with the given conf and
// immediately call Run returning the error.
func Generate(conf *config.Conf) error {
	return NewGen(conf).Run()
}

// Gen is project generator based on a configuration.
type Gen struct {
	conf *config.Conf
	log *Log
}

// NewGen allocates a Gen instance capable of producing
// Go project generator based on the given config.
func NewGen(config *config.Conf) *Gen {
	return &Gen{
		conf: config,
		log: NewLog(config),
	}
}

func (g *Gen) TemplateData() Data {
	return Data{
		PackageName: g.conf.Name,
		Remote: g.conf.Remote,
	}
}

// Run carries out project generation, making directories and files,
// returning an error if one is produced during generation.
func (g *Gen) Run() error {
	g.log.Println("Generating Project: ", g.conf.Name)
	dest := filepath.Join(g.conf.Dest, g.conf.Name)

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

// cp copies assets directly from the source to the destination producing
// an error if one occurs during generation.
func (g *Gen) cp() error {
	names := embedded.AssetNames()
	target := filepath.Join(g.conf.Dest, g.conf.Name)
	for _, name := range names {
		dest := filepath.Join(target, name)

		bin, err := embedded.Asset(name)
		if err != nil {
			return err
		}

		if strings.HasSuffix(dest, ".tpl.go") {
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
	names := embedded.AssetNames()
	target := filepath.Join(g.conf.Dest, g.conf.Name)
	for _, name := range names {
		dest := filepath.Join(target, name)
		dir := filepath.Dir(dest)

		if dir == target {
			continue
		}

		g.log.Println("creating directory:", dir)

		err := os.Mkdir(dir, 0777)
		if err != nil {
			return err
		}
	}
	return nil
}

// process template found in src and writes the value to the destination.
func (g *Gen) process(dest string, src []byte) error {
	g.log.Printf("processing template: %s", dest)

	p := template.New(g.conf.Name)
	tp, err := p.Parse(string(src))
	if err != nil {
		g.log.Println(err)
		return err
	}

	buf := bytes.NewBuffer([]byte{})
	err = tp.Execute(buf, g.TemplateData())
	if err != nil {
		g.log.Println(err)
		return err
	}

	return ioutil.WriteFile(dest, buf.Bytes(), 0777)
}
