package models

type Person struct {
	Name   string  `json:"name"`
	Role   string  `json:"role"`
	Weight float32 `json:"weight"`
}

type ConnectPerson struct {
	PersonA string  `json:"personA"`
	PersonB string  `json:"personB"`
	Weight  float32 `json:"weight"`
}

type Vertex struct {
	V1   interface{} `json:"v1"`
	Edge interface{} `json:"edge"`
	V2   interface{} `json:"v2"`
}
