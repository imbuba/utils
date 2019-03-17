package transport

import "sync"

// LogFlag alias
type LogFlag int

// LogPathes struct
type LogPathes struct {
	sync.RWMutex
	FullLogPathes map[string]LogFlag `json:"fullLogPathes"`
}

// LogFlags
const (
	ToLog LogFlag = 1 << iota
	FullLog
)

// Server describes server config
type Server struct {
	Host        string `json:"host"`
	Port        string `json:"port"`
	Secret      string `json:"secret"`
	AdminSecret string `json:"adminSecret"`
	Logging     bool   `json:"logging"`
	Local       bool   `json:"local"`
}

// PathesLogger interface
type PathesLogger interface {
	GetLogFlag(path string) LogFlag
	SetLogFlag(path string, flag LogFlag)
}

// GetLogFlag returns LogFlag
func (srv *LogPathes) GetLogFlag(path string) (result LogFlag) {
	var ok bool
	srv.RLock()
	if srv.FullLogPathes != nil {
		result, ok = srv.FullLogPathes[path]
	}
	srv.RUnlock()
	if !ok {
		result = ToLog
	}
	return result
}

// SetLogFlag changes log level for path
func (srv *LogPathes) SetLogFlag(path string, flag LogFlag) {
	srv.Lock()
	if srv.FullLogPathes == nil {
		srv.FullLogPathes = make(map[string]LogFlag)
	}
	srv.FullLogPathes[path] = flag
	srv.Unlock()
}
