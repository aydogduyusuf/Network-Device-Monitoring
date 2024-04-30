package web

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	// Home route
	app.Get("/", HandleHome)

	// SNMP discovery route
	app.Get("/snmp/discover", HandleSNMPDiscovery)

	// SNMP data retrieval route
	app.Get("/snmp/devices/:ipAddr", HandleSNMPDataRetrieval)
}
