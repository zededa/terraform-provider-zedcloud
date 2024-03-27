package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/models"
)

// Function to perform the following actions:
// (1) Translate GPSLocation resource data into a schema model struct that will sent to the LM API for resource creation/updating
// (2) Translate LM API response object from (1) or from a READ operation into a model that can be used to mofify the underlying resource data in the Terrraform configuration
func GPSLocationModel(d *schema.ResourceData) *models.GPSLocation {
	altitude, _ := d.Get("altitude").(float64)
	horizontalReliabilityInt, _ := d.Get("horizontal_reliability").(int)
	horizontalReliability := int64(horizontalReliabilityInt)
	horizontalUncertainty, _ := d.Get("horizontal_uncertainty").(float64)
	latitude, _ := d.Get("latitude").(float64)
	logicalLabel, _ := d.Get("logical_label").(string)
	longitude, _ := d.Get("longitude").(float64)
	timestamp, _ := d.Get("timestamp").(interface{}) // interface{}
	verticalReliabilityInt, _ := d.Get("vertical_reliability").(int)
	verticalReliability := int64(verticalReliabilityInt)
	verticalUncertainty, _ := d.Get("vertical_uncertainty").(float64)
	return &models.GPSLocation{
		Altitude:              altitude,
		HorizontalReliability: horizontalReliability,
		HorizontalUncertainty: horizontalUncertainty,
		Latitude:              latitude,
		LogicalLabel:          logicalLabel,
		Longitude:             longitude,
		Timestamp:             timestamp,
		VerticalReliability:   verticalReliability,
		VerticalUncertainty:   verticalUncertainty,
	}
}

func GPSLocationModelFromMap(m map[string]interface{}) *models.GPSLocation {
	altitude := m["altitude"].(float64)
	horizontalReliability := int64(m["horizontal_reliability"].(int)) // int64 false false false
	horizontalUncertainty := m["horizontal_uncertainty"].(float64)
	latitude := m["latitude"].(float64)
	logicalLabel := m["logical_label"].(string)
	longitude := m["longitude"].(float64)
	timestamp := m["timestamp"].(interface{})
	verticalReliability := int64(m["vertical_reliability"].(int)) // int64 false false false
	verticalUncertainty := m["vertical_uncertainty"].(float64)
	return &models.GPSLocation{
		Altitude:              altitude,
		HorizontalReliability: horizontalReliability,
		HorizontalUncertainty: horizontalUncertainty,
		Latitude:              latitude,
		LogicalLabel:          logicalLabel,
		Longitude:             longitude,
		Timestamp:             timestamp,
		VerticalReliability:   verticalReliability,
		VerticalUncertainty:   verticalUncertainty,
	}
}

// Update the underlying GPSLocation resource data in the Terraform configuration using the resource model built from the CREATE/UPDATE/READ LM API request response
func SetGPSLocationResourceData(d *schema.ResourceData, m *models.GPSLocation) {
	d.Set("altitude", m.Altitude)
	d.Set("horizontal_reliability", m.HorizontalReliability)
	d.Set("horizontal_uncertainty", m.HorizontalUncertainty)
	d.Set("latitude", m.Latitude)
	d.Set("logical_label", m.LogicalLabel)
	d.Set("longitude", m.Longitude)
	d.Set("timestamp", m.Timestamp)
	d.Set("vertical_reliability", m.VerticalReliability)
	d.Set("vertical_uncertainty", m.VerticalUncertainty)
}

// Iterate through and update the GPSLocation resource data within a pagination response (typically defined in the items array field) retrieved from a READ operation for multiple LM resources
func SetGPSLocationSubResourceData(m []*models.GPSLocation) (d []*map[string]interface{}) {
	for _, GPSLocationModel := range m {
		if GPSLocationModel != nil {
			properties := make(map[string]interface{})
			properties["altitude"] = GPSLocationModel.Altitude
			properties["horizontal_reliability"] = GPSLocationModel.HorizontalReliability
			properties["horizontal_uncertainty"] = GPSLocationModel.HorizontalUncertainty
			properties["latitude"] = GPSLocationModel.Latitude
			properties["logical_label"] = GPSLocationModel.LogicalLabel
			properties["longitude"] = GPSLocationModel.Longitude
			properties["timestamp"] = GPSLocationModel.Timestamp
			properties["vertical_reliability"] = GPSLocationModel.VerticalReliability
			properties["vertical_uncertainty"] = GPSLocationModel.VerticalUncertainty
			d = append(d, &properties)
		}
	}
	return
}

// Schema mapping representing the GPSLocation resource defined in the Terraform configuration
func GPSLocationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"altitude": {
			Description: `Altitude from mean sea level`,
			Type:        schema.TypeFloat,
			Optional:    true,
		},

		"horizontal_reliability": {
			Description: `Horizontal reliability`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"horizontal_uncertainty": {
			Description: `Horizontal uncertainty`,
			Type:        schema.TypeFloat,
			Optional:    true,
		},

		"latitude": {
			Description: `Latitude`,
			Type:        schema.TypeFloat,
			Optional:    true,
		},

		"logical_label": {
			Description: `Logical label of the adapter used to get GPS coordinates`,
			Type:        schema.TypeString,
			Optional:    true,
		},

		"longitude": {
			Description: `Longitude`,
			Type:        schema.TypeFloat,
			Optional:    true,
		},

		"timestamp": {
			Description: `UTC timestamp`,
			Type:        schema.TypeMap, //GoType: interface{}
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},

		"vertical_reliability": {
			Description: `Vertical reliability`,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"vertical_uncertainty": {
			Description: `Vertical uncertainty`,
			Type:        schema.TypeFloat,
			Optional:    true,
		},
	}
}

// Retrieve property field names for updating the GPSLocation resource
func GetGPSLocationPropertyFields() (t []string) {
	return []string{
		"altitude",
		"horizontal_reliability",
		"horizontal_uncertainty",
		"latitude",
		"logical_label",
		"longitude",
		"timestamp",
		"vertical_reliability",
		"vertical_uncertainty",
	}
}
