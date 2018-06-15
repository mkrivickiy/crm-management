package method

type Team struct {
	GroupID    int      `json:"group_id"`    //The team's numeric ID.
	Name       string   `json:"name"`        //The team's name.
	CreateDate int      `json:"create_date"` //The date and time the team was created in Epoch seconds.
	UpdateDate int      `json:"update_date"` //The date and time the team was last updated in Epoch seconds.
	Tags       []string `json:"tags"`        //A list of tags describing the team.
}
