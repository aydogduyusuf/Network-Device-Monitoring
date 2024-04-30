package snmp

import (
	"fmt"
	"time"

	"github.com/gosnmp/gosnmp"
)

// QuerySNMPDevice queries an SNMP-enabled device for system information
func QuerySNMPDevice(ipAddr string) (map[string]interface{}, error) {
	community := "public" // SNMP community string

	snmp := &gosnmp.GoSNMP{
		Target:    ipAddr,
		Port:      161, // SNMP default port
		Community: community,
		Version:   gosnmp.Version2c,
		Timeout:   time.Duration(2) * time.Second,
	}

	err := snmp.Connect()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to SNMP agent at %s: %v", ipAddr, err)
	}
	defer snmp.Conn.Close()

	// Example: query system description (OID: 1.3.6.1.2.1.1.1.0)
	result, err := snmp.Get([]string{
		".1.3.6.1.2.1.1.1.0", // System description
		".1.3.6.1.2.1.1.3.0", // System uptime
		".1.3.6.1.2.1.1.4.0", // System contact
	})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve system description from SNMP agent at %s: %v", ipAddr, err)
	}

	data := make(map[string]interface{})
	for _, variable := range result.Variables {
		// Extract OID and value
		oid := variable.Name
		//value := gosnmp.ToBigInt(variable.Value)
		value := variable.Value

		// Add data to map
		data[oid] = value
	}

	return data, nil
}
