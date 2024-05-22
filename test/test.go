package main

type Total struct {
	Data    []Info `json:"data"`
	Message string `json:"message"`
}
type Info struct {
	Desc            string      `json:"desc"`
	DeviceGroupId   int         `json:"deviceGroupId"`
	EnvAttr         EnvAttr     `json:"envAttr"`
	EventId         string      `json:"eventId"`
	EventName       string      `json:"eventName"`
	EventTypeId     int         `json:"eventTypeId"`
	InteractionAttr interface{} `json:"interactionAttr"`
	PageAttr        interface{} `json:"pageAttr"`
	PerformanceAttr interface{} `json:"performanceAttr"`
	ProjectId       int         `json:"projectId"`
	ScreenDirect    int         `json:"screenDirect"`
	SessionId       int         `json:"sessionId"`
	UserId          int         `json:"userId"`
}
type EnvAttr struct {
	City            Detail `json:"city"`
	Country         Detail `json:"country"`
	DeviceType      Detail `json:"deviceType"`
	DisplayHeight   Detail `json:"displayHeight"`
	DisplayWidth    Detail `json:"displayWidth"`
	Ip              Detail `json:"ip"`
	OperatingSystem Detail `json:"operatingSystem"`
	Province        Detail `json:"province"`
}
type Detail struct {
	DisplayName string
	Order       string
	Value       interface{}
}

//func main() {
//	file, err := os.ReadFile("./test/data.json")
//	if err != nil {
//		fmt.Println(err)
//	}
//	var total Total
//	err = json.Unmarshal(file, &total)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(total.Data[0].EnvAttr.Country.Value)
//
//}
