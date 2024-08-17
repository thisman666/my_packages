package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

type Action string
type Page string

const (
	ActionRefresh Action = "refresh" // 刷新
	ActionRefund  Action = "refund"  // 请求退款
	ActionDelete  Action = "delete"  // 删除

	PageRequirements      Page = "requirements"
	PageRequirementDetail Page = "requirement_detail"
	PageWaybills          Page = "waybills"
	PageWaybillDetail     Page = "waybill_detail"
)

type Message struct {
	Action        Action
	MerchantId    uint64
	DriverId      uint64
	RequirementId uint64
	WaybillId     uint64
}

func PublishMqttTopics(mqttClient mqtt.Client, param Message, pages ...Page) {

	if len(pages) > 0 {

		for _, page := range pages {
			if param.MerchantId != 0 {
				topic := fmt.Sprintf("yls/%d/%s/%d/%d/%s", param.MerchantId, page, param.RequirementId, param.WaybillId, param.Action)

				mqttClient.Publish(topic, 0, false, "")

				log.Printf("pub [%s] %s\n", topic, "")
			}

		}

		for _, page := range pages {

			if param.DriverId != 0 {
				topic := fmt.Sprintf("yls/%d/%s/%d/%d/%s", param.DriverId, page, param.RequirementId, param.WaybillId, param.Action)

				mqttClient.Publish(topic, 0, false, "")

				log.Printf("pub [%s] %s\n", topic, "")
			}

		}

	}

	return
}
