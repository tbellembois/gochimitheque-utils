package test

type Urn struct {
	Label string `json:"label"`
	Name  string `json:"name"`
}

type Value struct {
	Sval string `json:"sval"`
}

type Prop struct {
	Urn   Urn   `json:"urn"`
	Value Value `json:"value"`
}

type PCCompound struct {
	Props []Prop `json:"props"`
}

type PCCompounds struct {
	PC_Compounds []PCCompound `json:"PC_Compounds"`
}
