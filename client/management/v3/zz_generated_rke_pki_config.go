package client

const (
	RKEPkiConfigType       = "rkePkiConfig"
	RKEPkiConfigFieldCA    = "ca"
	RKEPkiConfigFieldCerts = "certs"
)

type RKEPkiConfig struct {
	CA    *RKEPkiCaConfig    `json:"ca,omitempty" yaml:"ca,omitempty"`
	Certs *RKEPkiCertsConfig `json:"certs,omitempty" yaml:"certs,omitempty"`
}
