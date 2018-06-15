package request

import "crm-management/systems/amocrm"

type Leads struct {
	Add    []amocrm.Lead `json:"add,omitempty"`
	Update []amocrm.Lead `json:"update,omitempty"`
}
