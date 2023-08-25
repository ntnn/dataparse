package dataparse

import (
	"fmt"
	"net"
)

func (v Value) IP() (net.IP, error) {
	switch typed := v.Data.(type) {
	case []byte:
		if len(typed) == 4 {
			return net.IPv4(typed[0], typed[1], typed[2], typed[3]), nil
		}
	default:
		s, err := v.String()
		if err != nil {
			return nil, fmt.Errorf("dataparse: error turning %q into string to parse: %w", v.Data, err)
		}
		if ip := net.ParseIP(s); ip != nil {
			return ip, nil
		}
	}
	return nil, fmt.Errorf("dataparse: error parsing %q as IP", v.Data)
}

func (v Value) MustIP() net.IP {
	val, _ := v.IP()
	return val
}

func (v Value) MAC() (net.HardwareAddr, error) {
	s, err := v.String()
	if err != nil {
		return nil, fmt.Errorf("dataparse: error turning %q into string to parse: %w", v.Data, err)
	}
	return net.ParseMAC(s)
}

func (v Value) MustMAC() net.HardwareAddr {
	val, _ := v.MAC()
	return val
}
