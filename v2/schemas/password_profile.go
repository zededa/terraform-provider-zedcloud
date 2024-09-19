package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func PasswordProfileModel(d *schema.ResourceData) *models.PasswordProfile {
	maxLengthInt, _ := d.Get("max_length").(int)
	maxLength := int64(maxLengthInt)
	maxPasswordAgeInt, _ := d.Get("max_password_age").(int)
	maxPasswordAge := int64(maxPasswordAgeInt)
	minLengthInt, _ := d.Get("min_length").(int)
	minLength := int64(minLengthInt)
	minLowercaseCharsInt, _ := d.Get("min_lowercase_chars").(int)
	minLowercaseChars := int64(minLowercaseCharsInt)
	minNumericCharsInt, _ := d.Get("min_numeric_chars").(int)
	minNumericChars := int64(minNumericCharsInt)
	minPasswordAgeInt, _ := d.Get("min_password_age").(int)
	minPasswordAge := int64(minPasswordAgeInt)
	minSymbolCharsInt, _ := d.Get("min_symbol_chars").(int)
	minSymbolChars := int64(minSymbolCharsInt)
	minUppercaseCharsInt, _ := d.Get("min_uppercase_chars").(int)
	minUppercaseChars := int64(minUppercaseCharsInt)
	numPrevPasswordCheckInt, _ := d.Get("num_prev_password_check").(int)
	numPrevPasswordCheck := int64(numPrevPasswordCheckInt)
	passwordExpiryNotificationPeriodInSecondsInt, _ := d.Get("password_expiry_notification_period_in_seconds").(int)
	passwordExpiryNotificationPeriodInSeconds := int64(passwordExpiryNotificationPeriodInSecondsInt)
	return &models.PasswordProfile{
		MaxLength:            maxLength,
		MaxPasswordAge:       maxPasswordAge,
		MinLength:            minLength,
		MinLowercaseChars:    minLowercaseChars,
		MinNumericChars:      minNumericChars,
		MinPasswordAge:       minPasswordAge,
		MinSymbolChars:       minSymbolChars,
		MinUppercaseChars:    minUppercaseChars,
		NumPrevPasswordCheck: numPrevPasswordCheck,
		PasswordExpiryNotificationPeriodInSeconds: passwordExpiryNotificationPeriodInSeconds,
	}
}

func PasswordProfileModelFromMap(m map[string]interface{}) *models.PasswordProfile {
	maxLength := int64(m["max_length"].(int))                                                                     // int64
	maxPasswordAge := int64(m["max_password_age"].(int))                                                          // int64
	minLength := int64(m["min_length"].(int))                                                                     // int64
	minLowercaseChars := int64(m["min_lowercase_chars"].(int))                                                    // int64
	minNumericChars := int64(m["min_numeric_chars"].(int))                                                        // int64
	minPasswordAge := int64(m["min_password_age"].(int))                                                          // int64
	minSymbolChars := int64(m["min_symbol_chars"].(int))                                                          // int64
	minUppercaseChars := int64(m["min_uppercase_chars"].(int))                                                    // int64
	numPrevPasswordCheck := int64(m["num_prev_password_check"].(int))                                             // int64
	passwordExpiryNotificationPeriodInSeconds := int64(m["password_expiry_notification_period_in_seconds"].(int)) // int64
	return &models.PasswordProfile{
		MaxLength:            maxLength,
		MaxPasswordAge:       maxPasswordAge,
		MinLength:            minLength,
		MinLowercaseChars:    minLowercaseChars,
		MinNumericChars:      minNumericChars,
		MinPasswordAge:       minPasswordAge,
		MinSymbolChars:       minSymbolChars,
		MinUppercaseChars:    minUppercaseChars,
		NumPrevPasswordCheck: numPrevPasswordCheck,
		PasswordExpiryNotificationPeriodInSeconds: passwordExpiryNotificationPeriodInSeconds,
	}
}

func SetPasswordProfileResourceData(d *schema.ResourceData, m *models.PasswordProfile) {
	d.Set("max_length", m.MaxLength)
	d.Set("max_password_age", m.MaxPasswordAge)
	d.Set("min_length", m.MinLength)
	d.Set("min_lowercase_chars", m.MinLowercaseChars)
	d.Set("min_numeric_chars", m.MinNumericChars)
	d.Set("min_password_age", m.MinPasswordAge)
	d.Set("min_symbol_chars", m.MinSymbolChars)
	d.Set("min_uppercase_chars", m.MinUppercaseChars)
	d.Set("num_prev_password_check", m.NumPrevPasswordCheck)
	d.Set("password_expiry_notification_period_in_seconds", m.PasswordExpiryNotificationPeriodInSeconds)
}

func SetPasswordProfileSubResourceData(m []*models.PasswordProfile) (d []*map[string]interface{}) {
	for _, PasswordProfileModel := range m {
		if PasswordProfileModel != nil {
			properties := make(map[string]interface{})
			properties["max_length"] = PasswordProfileModel.MaxLength
			properties["max_password_age"] = PasswordProfileModel.MaxPasswordAge
			properties["min_length"] = PasswordProfileModel.MinLength
			properties["min_lowercase_chars"] = PasswordProfileModel.MinLowercaseChars
			properties["min_numeric_chars"] = PasswordProfileModel.MinNumericChars
			properties["min_password_age"] = PasswordProfileModel.MinPasswordAge
			properties["min_symbol_chars"] = PasswordProfileModel.MinSymbolChars
			properties["min_uppercase_chars"] = PasswordProfileModel.MinUppercaseChars
			properties["num_prev_password_check"] = PasswordProfileModel.NumPrevPasswordCheck
			properties["password_expiry_notification_period_in_seconds"] = PasswordProfileModel.PasswordExpiryNotificationPeriodInSeconds
			d = append(d, &properties)
		}
	}
	return
}

func PasswordProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"max_length": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"max_password_age": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"min_length": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"min_lowercase_chars": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"min_numeric_chars": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"min_password_age": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"min_symbol_chars": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"min_uppercase_chars": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"num_prev_password_check": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},

		"password_expiry_notification_period_in_seconds": {
			Description: ``,
			Type:        schema.TypeInt,
			Optional:    true,
		},
	}
}

func GetPasswordProfilePropertyFields() (t []string) {
	return []string{
		"max_length",
		"max_password_age",
		"min_length",
		"min_lowercase_chars",
		"min_numeric_chars",
		"min_password_age",
		"min_symbol_chars",
		"min_uppercase_chars",
		"num_prev_password_check",
		"password_expiry_notification_period_in_seconds",
	}
}
