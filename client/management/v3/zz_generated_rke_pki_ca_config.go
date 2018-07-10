package client

const (
	RKEPkiCaConfigType           = "rkePkiCaConfig"
	RKEPkiCaConfigFieldValidDays = "validDays"
)

type RKEPkiCaConfig struct {
	ValidDays int64 `json:"validDays,omitempty" yaml:"validDays,omitempty"`
}
