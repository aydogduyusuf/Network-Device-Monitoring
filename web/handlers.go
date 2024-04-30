package web

import (
	"github.com/aydogduyusuf/Network-Device-Monitoring/discovery"
	"github.com/aydogduyusuf/Network-Device-Monitoring/snmp"
	"github.com/gofiber/fiber/v2"
)

func HandleHome(c *fiber.Ctx) error {
	return c.SendString("Welcome to Network Device Monitor!")
}

// HandleSNMPDiscovery handles SNMP discovery requests
func HandleSNMPDiscovery(c *fiber.Ctx) error {
	devices, err := discovery.DiscoverSNMPDevices()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(devices)
}

// HandleSNMPDataRetrieval handles SNMP data retrieval requests
func HandleSNMPDataRetrieval(c *fiber.Ctx) error {
	ipAddr := c.Params("ipAddr")
	data, err := snmp.QuerySNMPDevice(ipAddr)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(data)
}
