package lib

import (
	"fmt"
	"log"
	"path"
	"strings"

	vault "github.com/hashicorp/vault/api"
)

type Migrator struct {
	*vault.Client
	Root    string
	Version string
}

func (m Migrator) MetaPath(prefix string) string {
	if m.GetKVVersion() == "1" {
		return path.Join(m.Root, prefix)
	}
	return path.Join(m.Root, "metadata", prefix)
}

func (m Migrator) DataPath(prefix string) string {
	if m.GetKVVersion() == "1" {
		return path.Join(m.Root, prefix)
	}
	return path.Join(m.Root, "data", prefix)
}

func (m *Migrator) WriteData(data map[string]map[string]interface{}) error {
	for k, v := range data {
		var body map[string]interface{}
		if m.GetKVVersion() == "1" {
			body = v
		} else {
			body = map[string]interface{}{"data": v}
		}
		if _, err := m.Logical().Write(m.DataPath(k), body); err != nil {
			return err
		}
	}
	return nil
}

func (m *Migrator) GetKVVersion() string {
	if m.Version != "" {
		return m.Version
	}

	r, err := m.Sys().ListMounts()
	if err != nil {
		log.Fatal("Unable to list mounts", err)
	}
	mc, ok := r[fmt.Sprintf("%s/", m.Root)]
	if !ok {
		log.Fatal("Unable to find mount")
	}
	if v, ok := mc.Options["version"]; ok {
		return v
	}

	return "1"
}

func (m *Migrator) ReadData(prefix string) (map[string]interface{}, error) {
	data := map[string]interface{}{}

	r, err := m.Logical().List(m.MetaPath(prefix))
	if err != nil {
		return nil, err
	}
	for _, keyIf := range r.Data["keys"].([]interface{}) {
		key := path.Join(prefix, keyIf.(string))

		if strings.HasSuffix(keyIf.(string), "/") {
			r, err := m.ReadData(key)
			if err != nil {
				return nil, err
			}

			for k, v := range r {
				data[k] = v
			}
		} else {
			s, err := m.Logical().Read(m.DataPath(key))
			if err != nil {
				return nil, err
			}

			if m.GetKVVersion() == "1" {
				data[key] = s.Data
			} else {
				data[key] = s.Data["data"]
			}
		}
	}

	return data, nil
}
