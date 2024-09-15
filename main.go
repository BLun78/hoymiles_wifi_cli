package main

import (
	"encoding/json"
	"fmt"
	"github.com/BLun78/hoymiles_wifi"
	"github.com/BLun78/hoymiles_wifi/common"
	"github.com/BLun78/hoymiles_wifi/hoymiles/models"
	"time"
)

func main() {
	client := hoymiles_wifi.NewClient("10.10.34.137", common.DTU_PORT)

	GetRealDataNew(client)
	GetRealData(client)
	GetConfig(client)

}

func GetRealDataNew(client *hoymiles_wifi.ClientData) {
	fmt.Println("GetRealDataNew...")
	request := &models.RealDataNewReqDTO{}
	// int32 would not be Year 2038 safe
	// See https://en.wikipedia.org/wiki/Year_2038_problem
	// Not 100% sure if the models are self-defined or provided by hoymiles
	request.Time = int32(time.Now().Unix())
	request.TimeYmdHms = time.Now().Format("2006-01-02 15:04:05")

	result, err := client.GetRealDataNew(request)
	if err != nil {
		fmt.Println(err)
	}
	j, _ := json.MarshalIndent(result, "", "    ")

	fmt.Println(string(j))
}

func GetRealData(client *hoymiles_wifi.ClientData) {
	fmt.Println("GetRealData...")
	request := &models.RealDataReqDTO{}
	// int32 would not be Year 2038 safe
	// See https://en.wikipedia.org/wiki/Year_2038_problem
	// Not 100% sure if the models are self-defined or provided by hoymiles
	request.Time = int32(time.Now().Unix())
	request.TimeYmdHms = time.Now().Format("2006-01-02 15:04:05")

	result, err := client.GetRealData(request)
	if err != nil {
		fmt.Println(err)
	}
	j, _ := json.MarshalIndent(result, "", "    ")

	fmt.Println(string(j))
}

func GetConfig(client *hoymiles_wifi.ClientData) {
	fmt.Println("GetConfig...")
	request := &models.GetConfigReqDTO{}
	// int32 would not be Year 2038 safe
	// See https://en.wikipedia.org/wiki/Year_2038_problem
	// Not 100% sure if the models are self-defined or provided by hoymiles
	request.Time = uint32(time.Now().Unix()) - 60

	result, err := client.GetConfig(request)
	if err != nil {
		fmt.Println(err)
	}
	j, _ := json.MarshalIndent(result, "", "    ")

	fmt.Println(string(j))
}
