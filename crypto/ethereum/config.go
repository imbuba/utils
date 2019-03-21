package ethereum

// Config struct
type Config struct {
	Protocol           string                   `json:"protocol"`
	Host               string                   `json:"host"`
	Port               uint16                   `json:"port"`
	Addresses          map[string]string        `json:"addresses"`
	Definitions        map[string]AbiDefinition `json:"definitions"`
	ProvidedGas        int                      `json:"providedGas"`
	TokenPrecision     int                      `json:"tokenPrecision"`
	Logging            bool                     `json:"logging"`
	CheckWorldBlockURL string                   `json:"checkWorldBlockUrl"`
	MinBlockDifference int64                    `json:"minBlockDifference"`
}

// AbiDefinition struct
type AbiDefinition struct {
	Abi     string   `json:"abi"`
	Address string   `json:"address"`
	Events  []string `json:"events"`
}
