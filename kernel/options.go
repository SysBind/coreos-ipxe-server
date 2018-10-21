package kernel

import (
	"bytes"
	"fmt"
)

type Options struct {
	IgnitionConfig     string   `json:"ignition_config"`
	Console         []string `json:"console"`
	CoreOSAutologin string   `json:"coreos_autologin"`
	Root            string   `json:"root"`
	RootFstype      string   `json:"rootfstype"`
	SSHKey          string   `json:"sshkey"`
	Version         string   `json:"version"`
	ignitionConfigUrl  string
}

func New() *Options {
	return &Options{}
}

func (o *Options) SetIgnitionConfigUrl(url string) {
	o.ignitionConfigUrl = url
}

func (o *Options) String() string {
	var options bytes.Buffer
	if o.RootFstype != "" {
		options.WriteString(fmt.Sprintf(" rootfstype=%s", o.RootFstype))
	}
	for _, c := range o.Console {
		options.WriteString(fmt.Sprintf(" console=%s", c))
	}
	if o.ignitionConfigUrl != "" {
		options.WriteString(fmt.Sprintf(" coreos.first_boot=1 coreos.config.url=%s", o.ignitionConfigUrl))
	}
	if o.CoreOSAutologin != "" {
		options.WriteString(fmt.Sprintf(" coreos.autologin=%s", o.CoreOSAutologin))
	}
	if o.SSHKey != "" {
		options.WriteString(fmt.Sprintf(" sshkey=\"%s\"", o.SSHKey))
	}
	if o.Root != "" {
		options.WriteString(fmt.Sprintf(" root=%s", o.Root))
	}
	return options.String()
}
