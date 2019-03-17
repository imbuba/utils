package database

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func sprintf(format string, params map[string]string) string {
	for key, val := range params {
		format = strings.Replace(format, "%{"+key+"}%", val, -1)
	}
	return format
}

func sprintSlice(slice []interface{}) string {
	var b strings.Builder
	b.WriteRune('[')
	for i, v := range slice {
		if i != 0 {
			b.WriteString(", ")
		}
		if str, ok := v.(fmt.Stringer); ok {
			b.WriteString(str.String())
		} else {
			b.WriteString(fmt.Sprintf("%#+v", v))
		}
	}
	b.WriteRune(']')
	return b.String()
}

// ParseConfig parses config with a given path
func ParseConfig(path string) (*Database, error) {
	result := &Database{}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Database describes db connection
type Database struct {
	Hosts      []string `json:"hosts"`
	Ports      []string `json:"ports"`
	Protocol   string   `json:"protocol"`
	DBName     string   `json:"dbName"`
	AppName    string   `json:"appName"`
	User       string   `json:"user"`
	Pass       string   `json:"pass"`
	Connection string   `json:"connection"`
}

func (d *Database) toConnectionMap() map[string]string {
	return map[string]string{
		"User":     d.User,
		"Password": d.Pass,
		"Host":     d.defaultHost(),
		"Port":     d.defaultPort(),
		"Protocol": d.Protocol,
		"DBName":   d.DBName,
		"AppName":  d.AppName,
	}
}

func (d *Database) defaultHost() string {
	if len(d.Hosts) > 0 {
		return d.Hosts[0]
	}
	return ""
}

func (d *Database) defaultPort() string {
	if len(d.Ports) > 0 {
		return d.Ports[0]
	}
	return ""
}

// ConnectionString do smth
func (d *Database) ConnectionString() string {
	return sprintf(d.Connection, d.toConnectionMap())
}
