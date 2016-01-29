package gsh

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os/user"
	"path/filepath"
	"strings"
)

// CurrentConfig holds the host configuration for the current run.
var CurrentConfig *Config

// Operation type defines various commands that can be used
type Operation struct {
	Name        string `json:"name"`
	ShellCmd    string `json:"shell-command"`
	IsStreaming bool   `json:"is-streaming"`
}

// HostConfig is used to store information for specific hosts
type HostConfig struct {
	Name                string   `json:"name"`
	Host                string   `json:"host"`
	Port                int      `json:"port"`
	SudoForConfig       bool     `json:"sudo-for-config"`
	SupportedOperations []string `json:"supported-operations"`
}

// Config defines the configuration for all hosts.
type Config struct {
	Name          string       `json:"name"`
	DefaultSuffix string       `json:"default-suffix"`
	Hosts         []HostConfig `json:"hosts"`
	Operations    []Operation  `json:"operations"`
}

// LoadConfigs iteratates through the user's home directory trying to find "gsh-*.json" files.
// Any it finds, it attempts to load as a Config.  Returns an array of Configs.
func LoadConfigs() []Config {
	var aConfs []Config

	// Start by getting the home dir
	homeDir := getHomeDir()

	// Now, search the pattern
	aFilePaths, err := filepath.Glob(homeDir + "/gsh-*.json")
	if err != nil {
		log.Fatal(err)
	}

	for _, strPath := range aFilePaths {
		fConf, err := ioutil.ReadFile(strPath)
		if err != nil {
			log.Fatalf("Unable to open file: %s; err=%v", strPath, err)
		}
		var conf Config
		json.Unmarshal(fConf, &conf)
		aConfs = append(aConfs, conf)
	}

	return aConfs
}

// GetHost returns a pointer to the host identified by strHostName.  If no host can be found,
// it returns nil
func (c *Config) GetHost(strHostName string) *HostConfig {
	for _, host := range c.Hosts {
		if strings.EqualFold(host.Name, strHostName) {
			return &host
		}
	}
	return nil
}

// GetOperation returns a pointer to the operation identified by strOpName.  If no operation can be found,
// it returns nil
func (c *Config) GetOperation(strOpName string) *Operation {
	for _, op := range c.Operations {
		if strings.EqualFold(op.Name, strOpName) {
			return &op
		}
	}
	return nil
}

// SupportsOp checks to see if the specified operation is supported by this host.
func (op *HostConfig) SupportsOp(strOp string) bool {
	for _, strSupOp := range op.SupportedOperations {
		if strings.EqualFold(strOp, strSupOp) {
			return true
		}
	}
	return false
}

func getHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

// ByName sorts host config by name
type ByName []HostConfig

func (x ByName) Len() int           { return len(x) }
func (x ByName) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func (x ByName) Less(i, j int) bool { return x[i].Name < x[j].Name }

type OperationsByName []Operation

func (x OperationsByName) Len() int           { return len(x) }
func (x OperationsByName) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func (x OperationsByName) Less(i, j int) bool { return x[i].Name < x[j].Name }
