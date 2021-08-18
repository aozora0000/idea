package idea

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type Psyche struct {
	Key []byte
}

func CreatePsyche(name string) (*Psyche, error) {
	Psyche := &Psyche{Key: CreateNewKey(time.Now())}
	err := Psyche.Save(name)
	return Psyche, err
}

func LoadPsyche(name string) (*Psyche, error) {
	b, err := ioutil.ReadFile(filepath.Join(getShelfDir(), name))
	if err != nil {
		return nil, err
	}
	return &Psyche{Key: b}, nil
}

func ExistPsyche(name string) bool {
	return Exists(filepath.Join(getShelfDir(), name))
}

func (p *Psyche) Save(name string) error {
	f, err := os.Create(filepath.Join(getShelfDir(), name))
	if err != nil {
		return err
	}
	defer f.Close()
	buf := bytes.NewBuffer(p.Key)
	io.Copy(f, buf)
	return nil
}
