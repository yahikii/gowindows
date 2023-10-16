package connection

import (
	"github.com/masterzen/winrm"
)

type KerberosConfig struct {
	Realm         string
	KrbConfigFile string
}

const (
	// Default kerberos values
	kerberosProtocolDefault = "http"
)

// winRMKerberosParams returns the neccessary parameters
// to pass into the kerberos winrm connection
func winRMKerberosParams(config *WinRMConfig) *winrm.Parameters {

	// Init default parameters
	params := winrm.DefaultParameters

	// Set the protocol
	kerberosProtocol := kerberosProtocolDefault
	if config.WinRMUseTLS {
		kerberosProtocol = "https"
	}

	// Configure kerberos transporter
	params.TransportDecorator = func() winrm.Transporter {
		return &winrm.ClientKerberos{
			Username: config.WinRMUsername,
			Password: config.WinRMPassword,
			Hostname: config.WinRMHost,
			Realm:    config.WinRMKerberos.Realm,
			Port:     config.WinRMPort,
			Proto:    kerberosProtocol,
			KrbConf:  config.WinRMKerberos.KrbConfigFile,
		}
	}

	return params
}
