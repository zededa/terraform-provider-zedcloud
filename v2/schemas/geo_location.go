package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

// Function to perform the following actions:
// (1) Translate GeoLocation resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func GeoLocationModel(d *schema.ResourceData) *models.GeoLocation {
	city, _ := d.Get("city").(string)
	country, _ := d.Get("country").(string)
	freeloc, _ := d.Get("freeloc").(string)
	hostname, _ := d.Get("hostname").(string)
	latlong, _ := d.Get("latlong").(string)
	loc, _ := d.Get("loc").(string)
	org, _ := d.Get("org").(string)
	postal, _ := d.Get("postal").(string)
	region, _ := d.Get("region").(string)
	underlayIP, _ := d.Get("underlay_ip").(string)
	return &models.GeoLocation{
		City:       city,
		Country:    country,
		Freeloc:    freeloc,
		Hostname:   hostname,
		Latlong:    latlong,
		Loc:        loc,
		Org:        org,
		Postal:     postal,
		Region:     region,
		UnderlayIP: underlayIP,
	}
}

func GeoLocationModelFromMap(m map[string]interface{}) *models.GeoLocation {
	city := m["city"].(string)
	country := m["country"].(string)
	freeloc := m["freeloc"].(string)
	hostname := m["hostname"].(string)
	latlong := m["latlong"].(string)
	loc := m["loc"].(string)
	org := m["org"].(string)
	postal := m["postal"].(string)
	region := m["region"].(string)
	underlayIP := m["underlay_ip"].(string)
	return &models.GeoLocation{
		City:       city,
		Country:    country,
		Freeloc:    freeloc,
		Hostname:   hostname,
		Latlong:    latlong,
		Loc:        loc,
		Org:        org,
		Postal:     postal,
		Region:     region,
		UnderlayIP: underlayIP,
	}
}

// Update the underlying GeoLocation resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetGeoLocationResourceData(d *schema.ResourceData, m *models.GeoLocation) {
	d.Set("city", m.City)
	d.Set("country", m.Country)
	d.Set("freeloc", m.Freeloc)
	d.Set("hostname", m.Hostname)
	d.Set("latlong", m.Latlong)
	d.Set("loc", m.Loc)
	d.Set("org", m.Org)
	d.Set("postal", m.Postal)
	d.Set("region", m.Region)
	d.Set("underlay_ip", m.UnderlayIP)
}

// Iterate through and update the GeoLocation resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetGeoLocationSubResourceData(m []*models.GeoLocation) (d []*map[string]interface{}) {
	for _, GeoLocationModel := range m {
		if GeoLocationModel != nil {
			properties := make(map[string]interface{})
			properties["city"] = GeoLocationModel.City
			properties["country"] = GeoLocationModel.Country
			properties["freeloc"] = GeoLocationModel.Freeloc
			properties["hostname"] = GeoLocationModel.Hostname
			properties["latlong"] = GeoLocationModel.Latlong
			properties["loc"] = GeoLocationModel.Loc
			properties["org"] = GeoLocationModel.Org
			properties["postal"] = GeoLocationModel.Postal
			properties["region"] = GeoLocationModel.Region
			properties["underlay_ip"] = GeoLocationModel.UnderlayIP
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the GeoLocation resource defined in the Terraform configuration
func GeoLocationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"city": {
			Description: `City`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"country": {
			Description: `Country code consisting of 2 capital letters as per ISO 3166-1 alpha2 standard`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"freeloc": {
			Description: `Free formatted location string`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"hostname": {
			Description: `Host name`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"latlong": {
			Description: `Deprecated field: comma, separated lat, long`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"loc": {
			Description: `Ordered pair of (latitude, longitude) separated by comma (,). Latitude is the horizontal component used for geographic positioning; it is the angle between 0° (the equator) and ±90° (north or south) at the poles measured in decimal degrees. It is the first value in an ordered pair. A negative number denotes a location south of the equator; a positive number is north. Longitude is the vertical component used for geographic positioning; it is the angle between 0° (the Prime Meridian) and ±180° (westward or eastward) measured in decimal degrees. It is the second number in an ordered pair. A negative number indicates a location west of Greenwich, England; a positive number east.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"org": {
			Description: `The name of the recipient, firm, or company at this geographical location.`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"postal": {
			Description: `Postal code (ZIP code for USA) of the geographical location`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"region": {
			Description: `Region`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"underlay_ip": {
			Description: `Single IP address, either in IPv4 or in IPv6 format`,
			Type:        schema.TypeString,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the GeoLocation resource
func GetGeoLocationPropertyFields() (t []string) {
	return []string{
		"city",
		"country",
		"freeloc",
		"hostname",
		"latlong",
		"loc",
		"org",
		"postal",
		"region",
		"underlay_ip",
	}
}
