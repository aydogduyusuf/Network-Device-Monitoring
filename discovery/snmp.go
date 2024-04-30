package discovery

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/gosnmp/gosnmp"
)

// DiscoverSNMPDevices discovers SNMP-enabled devices on the local network
func DiscoverSNMPDevices() ([]string, error) {
	community := "public" // SNMP community string
	var discoveredDevices []string

	target := "192.168.1.1" // My modem IP address

	snmp := &gosnmp.GoSNMP{
		Target:    target,
		Community: community,
		Version:   gosnmp.Version2c,
		Timeout:   time.Duration(10) * time.Second,
	}

	err := snmp.Connect()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to SNMP agent: %v", err)
	}
	defer snmp.Conn.Close()

	// Perform SNMP walk to discover devices
	result, err := snmp.Get([]string{"1.3.6.1.2.1.4.22.1.2", "1.3.6.1.2.1.4.22.1.3"})
	if err != nil {
		return nil, fmt.Errorf("SNMP walk failed: %v", err)
	}

	for _, pdu := range result.Variables {
		ipAddr := pdu.Value.(string)
		if !strings.HasPrefix(ipAddr, target) {
			continue
		}
		ipAddr = strings.TrimPrefix(ipAddr, target+".")
		if ip := net.ParseIP(ipAddr); ip != nil {
			discoveredDevices = append(discoveredDevices, ip.String())
		}
	}

	return discoveredDevices, nil
}
