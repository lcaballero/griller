package task

import (
	"github.com/lcaballero/griller/config"
	"path/filepath"
	"os"
	"log"
	"github.com/lcaballero/griller/embedded"
	"io/ioutil"
)

func Generate(conf *config.Conf) error {
	g := &Gen{
		conf: conf,
	}
	return g.Run()
}

type Gen struct {
	conf *config.Conf
}

func (g *Gen) Run() error {
	log.Println("Generating Project: ", g.conf.Name)
	dest := filepath.Join(g.conf.Dest, g.conf.Name)
	err := os.MkdirAll(dest, 0777)
	if err != nil {
		return err
	}

	names := embedded.AssetNames()
	for _, name := range names {
		dest := filepath.Join(g.conf.Dest, name)

		info, err := embedded.AssetInfo(name)
		if err != nil {
			log.Println(err)
			continue
		}

		if info.IsDir() {
			dir := filepath.Join(g.conf.Dest, name)
			os.Mkdir(dir, 0666)
			continue
		}

		bin, err := embedded.Asset(name)
		if err != nil {
			log.Println(err)
			continue
		}
		err = ioutil.WriteFile(dest, bin, 0777)

		if err != nil {
			log.Println(err)
			continue
		}
	}

	return nil
}
