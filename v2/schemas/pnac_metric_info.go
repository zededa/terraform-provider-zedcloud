package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zededa/terraform-provider-zedcloud/v2/models"
)

func SetPNACMetricInfoSubResourceData(m []*models.PNACMetricInfo) (d []*map[string]interface{}) {
	for _, v := range m {
		if v != nil {
			properties := make(map[string]interface{})
			properties["logical_label"] = v.LogicalLabel
			properties["eapol_frames_rx"] = v.EapolFramesRx
			properties["eapol_frames_tx"] = v.EapolFramesTx
			properties["eapol_start_frames_tx"] = v.EapolStartFramesTx
			properties["eapol_logoff_frames_tx"] = v.EapolLogoffFramesTx
			properties["eapol_resp_frames_tx"] = v.EapolRespFramesTx
			properties["eapol_req_id_frames_rx"] = v.EapolReqIDFramesRx
			properties["eapol_req_frames_rx"] = v.EapolReqFramesRx
			properties["invalid_eapol_frames_rx"] = v.InvalidEapolFramesRx
			properties["eap_length_error_frames_rx"] = v.EapLengthErrorFramesRx
			d = append(d, &properties)
		}
	}
	return
}

func PNACMetricInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"logical_label": {
			Description: `Logical label identifying the network port`,
			Type:        schema.TypeString,
			Computed:    true,
		},
		"eapol_frames_rx": {
			Description: `Total EAPOL frames received`,
			Type:        schema.TypeString,
			Computed:    true,
		},
		"eapol_frames_tx": {
			Description: `Total EAPOL frames transmitted`,
			Type:        schema.TypeString,
			Computed:    true,
		},
		"eapol_start_frames_tx": {
			Description: `EAPOL-Start frames transmitted`,
			Type:        schema.TypeString,
			Computed:    true,
		},
		"eapol_logoff_frames_tx": {
			Description: `EAPOL-Logoff frames transmitted`,
			Type:        schema.TypeString,
			Computed:    true,
		},
		"eapol_resp_frames_tx": {
			Description: `EAP-Response frames transmitted`,
			Type:        schema.TypeString,
			Computed:    true,
		},
		"eapol_req_id_frames_rx": {
			Description: `EAP-Request Identity frames received`,
			Type:        schema.TypeString,
			Computed:    true,
		},
		"eapol_req_frames_rx": {
			Description: `EAP-Request frames received`,
			Type:        schema.TypeString,
			Computed:    true,
		},
		"invalid_eapol_frames_rx": {
			Description: `Invalid or malformed EAPOL frames received`,
			Type:        schema.TypeString,
			Computed:    true,
		},
		"eap_length_error_frames_rx": {
			Description: `EAPOL frames with length errors received`,
			Type:        schema.TypeString,
			Computed:    true,
		},
	}
}
