package service

import (
	"RuntimeAlertProducer/internal/models"
	"fmt"
	"github.com/ZscalerCWP/cwp-alert-common/pkg/kafka/producerUtil"
	v1 "github.com/ZscalerCWP/cwp-int-shared-contracts/pkg/alert/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
)

func ProcessRuntimeAlerts(alerts []models.RuntimeAlert, integrationId string) {

	producerConfigs := producerUtil.ProducerConfigurations{
		UserName:   "",
		Password:   "",
		Brokers:    "localhost:9092",
		BatchSize:  1024,
		LingerMs:   1000,
		Protocol:   "PLAINTEXT",
		Mechanisms: "PLAIN",
	}

	producerClient, producerClientErr := producerUtil.NewProducerClient(&producerConfigs)

	for _, alert := range alerts {
		m := make(map[string]interface{})
		m["ID"] = alert.ID
		m["Alert_ID"] = alert.Alert_ID
		m["ALERT_STATUS"] = alert.ALERT_STATUS
		m["ALERT_CREATED_DT"] = alert.ALERT_CREATED_DT
		m["ALERT_UPDATED_DT"] = alert.ALERT_UPDATED_DT
		m["CSP"] = alert.CSP
		m["ACCOUNT_ID"] = alert.ACCOUNT_ID
		m["RESOURCE_NAME"] = alert.RESOURCE_NAME
		m["RESOURCE_ID"] = alert.RESOURCE_ID
		m["POLICY_ID"] = alert.POLICY_ID
		m["POLICY_NAME"] = alert.POLICY_NAME
		m["CSP_EVENT_TYPE"] = alert.CSP_EVENT_TYPE

		msg, _ := structpb.NewStruct(m)
		composedMsg := v1.Alert{Attributes: msg, Type: v1.Alert_RunTime}
		alert2 := []*v1.Alert{&composedMsg}
		am := v1.AlertMessage{TenantId: "t1", IntegrationId: integrationId, Alerts: alert2}
		data, _ := proto.Marshal(&am)

		header := make(map[string]string)
		header["traceid"] = "TestId"

		fmt.Println(data)

		topicEvent := producerUtil.IntegrationTopicEvent{
			Message: data,
			Header:  header,
		}
		topic := "test2"
		if producerClientErr == nil {
			producerClient.SendEventToIntegration(&topic, topicEvent)
		}
	}

}
