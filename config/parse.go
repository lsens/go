package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
)

func init() {
	var err error
	rootDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	if !FileExists(configFileName) {
		configFileName = filepath.Join(rootDir, configFileName)
	}
	Cfg, err = ParseConfig(configFileName, rootDir)
	if err != nil {
		log.Fatal(err)
	}
}

func ParseConfig(cfgPath string, rootDir string) (*Configs, error) {
	fd, err := os.Open(cfgPath)
	if err != nil {
		return nil, err
	}
	config := &Configs{}
	config.RuntimeParam.RootDir = rootDir
	err = yaml.NewDecoder(fd).Decode(config)
	if err != nil {
		return nil, err
	}
	if err := config.verification(); err != nil {
		return nil, err
	}
	return config, nil
}

func (cf *Configs) verification() error {
	// todo: 做字段的检测
	if cf.Logger.Path == "" {
		cf.Logger.Path = "logs"
	}
	cf.Logger.Path = filepath.Join(cf.RuntimeParam.RootDir, cf.Logger.Path)
	_, err := os.Stat(cf.Logger.Path)
	if err != nil {
		if mkErr := os.Mkdir(cf.Logger.Path, 0644); mkErr != nil {
			return mkErr
		}
	}
	return nil
}

// FileExists FileExists
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
