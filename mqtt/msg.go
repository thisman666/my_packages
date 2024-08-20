package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

type Action string

const (
	ActionRefresh Action = "refresh" // 刷新
	ActionRefund  Action = "refund"  // 请求退款
	ActionDelete  Action = "delete"  // 删除
)

type Message struct {
	Action        Action
	MerchantId    uint64
	DriverId      uint64
	RequirementId uint64
	WaybillId     uint64
}

func PublishMqttTopics(mqttClient mqtt.Client, param Message) {
	fmt.Println()

	if param.MerchantId != 0 {
		topic := fmt.Sprintf("yls/%d/%d/%d/%s", param.MerchantId, param.RequirementId, param.WaybillId, param.Action)

		mqttClient.Publish(topic, 0, false, "")

		log.Printf("pub [%s] %s\n", topic, "")
	}

	if param.DriverId != 0 {
		topic := fmt.Sprintf("yls/%d/%d/%d/%s", param.DriverId, param.RequirementId, param.WaybillId, param.Action)

		mqttClient.Publish(topic, 0, false, "")

		log.Printf("pub [%s] %s\n", topic, "")
	}

	return
}
