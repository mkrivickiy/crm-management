package method

type TreasuryCounts struct {
	Clicks  int `json:"clicks"`  //The number of times the Treasury has been clicked on
	Views   int `json:"views"`   //The number of times the Treasury has been viewed
	Shares  int `json:"shares"`  //The number of times the Treasury has been shared
	Reports int `json:"reports"` //The number of times the Treasury has been reported
}
