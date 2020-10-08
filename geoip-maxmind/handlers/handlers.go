package handlers

import (
	"fmt"
	"net"

	"github.com/gofiber/fiber/v2"
	"github.com/oschwald/maxminddb-golang"
)

// See https://pkg.go.dev/github.com/oschwald/geoip2-golang#City for a full list of options you can use here to modify
// what data is returned for a specific IP.
type ipLookup struct {
	City struct {
		GeoNameID uint              `maxminddb:"geoname_id"`
		Names     map[string]string `maxminddb:"names"`
	} `maxminddb:"city"`
	Country struct {
		IsoCode string `maxminddb:"iso_code"`
	} `maxminddb:"country"`
	Location struct {
		AccuracyRadius uint16 `maxminddb:"accuracy_radius"`
	} `maxminddb:"location"`
}

var geoIpDb *maxminddb.Reader

func init() {
	// Load MaxMind DB file
	var err error
	geoIpDb, err = maxminddb.Open("GeoLite2-City.mmdb")
	if err != nil {
		fmt.Println("Unable to load 'GeoLite2-City.mmdb'.")
		panic(err)
	}
}

func GeoIp(c *fiber.Ctx) error {

	ipAddr := c.Params("ip", c.IP())

	ip := net.ParseIP(ipAddr)

	if ip == nil {
		return c.Status(400).JSON(map[string]string{"status": "error", "message": "Invalid IP address"})
	}

	record := new(ipLookup)
	err := geoIpDb.Lookup(ip, &record)
	if err != nil {
		return err
	}

	return c.JSON(record)
}
