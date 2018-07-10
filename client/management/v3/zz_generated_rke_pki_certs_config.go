package client

const (
	RKEPkiCertsConfigType           = "rkePkiCertsConfig"
	RKEPkiCertsConfigFieldValidDays = "validDays"
)

type RKEPkiCertsConfig struct {
	ValidDays int64 `json:"validDays,omitempty" yaml:"validDays,omitempty"`
}
