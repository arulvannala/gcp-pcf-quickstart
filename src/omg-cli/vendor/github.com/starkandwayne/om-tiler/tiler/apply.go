package tiler

import (
	"fmt"
	"path/filepath"

	"github.com/starkandwayne/om-tiler/pattern"
)

func (c *Tiler) Apply(p pattern.Pattern) error {
	if err := p.Validate(true); err != nil {
		return err
	}

	err := c.client.ConfigureAuthentication()
	if err != nil {
		return err
	}

	err = c.configureDirector(p.Director)
	if err != nil {
		return err
	}

	for _, tile := range p.Tiles {
		if err = c.ensureFilesUploaded(tile); err != nil {
			return err
		}

		err = c.client.StageProduct(tile)
		if err != nil {
			return err
		}

		err = c.configureProduct(tile)
		if err != nil {
			return err
		}
	}

	err = c.client.ApplyChanges()
	if err != nil {
		return err
	}

	return nil
}

func (c *Tiler) ensureFilesUploaded(t pattern.Tile) error {
	ok, err := c.client.FilesUploaded(t)
	if err != nil {
		return err
	}
	if ok {
		c.logger.Printf("files for %s/%s already uploaded skipping download",
			t.Name, t.Version)
		return nil
	}

	product, err := c.mover.Get(t.Product)
	if err != nil {
		return err
	}

	if err = c.client.UploadProduct(product); err != nil {
		return err
	}

	stemcell, err := c.mover.Get(t.Stemcell)
	if err != nil {
		return err
	}

	return c.client.UploadStemcell(stemcell)
}

func (c *Tiler) configureProduct(t pattern.Tile) error {
	tpl, err := t.ToTemplate().Evaluate(true)
	if err != nil {
		return err
	}

	return c.client.ConfigureProduct(tpl)
}

func (c *Tiler) configureDirector(d pattern.Director) error {
	tpl, err := d.ToTemplate().Evaluate(true)
	if err != nil {
		return err
	}

	return c.client.ConfigureDirector(tpl)
}

func findFileInDir(dir, glob string) (string, error) {
	files, err := filepath.Glob(filepath.Join(dir, glob))
	if err != nil {
		return "", err
	}
	if len(files) != 1 {
		return "", fmt.Errorf("no file found for %s in %s", glob, dir)
	}
	return files[0], nil
}
