package services

import (
	"CurdOperation/Model"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type IReadJson interface {
	ReadJson() (Model.Jsondata, error)
}
type ReadjsonController struct{}

func ReadJsonCtor() *ReadjsonController {
	return &ReadjsonController{}
}

func (rj ReadjsonController) ReadJson() (Model.Jsondata, error) {
	var JsonData Model.Jsondata
	Data, err := os.ReadFile("appsettings.json")
	if err != nil {
		log.Println("Error while reading appsetting.json file  ", err)
		return JsonData, err
	}
	err = json.Unmarshal(Data, &JsonData)
	if err != nil {
		fmt.Println("Error while unmarshalling JSON:", err)
		panic(err)
	}
	return JsonData, err

}
