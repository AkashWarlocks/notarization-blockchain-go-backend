package types

type SetData struct {
	Documemt string `json:"document"`
	User string `json:"user"`
	DocumentName string `json:"documentName"`
}

type ParametersType struct  {
	Datatype string
	Value string
}

type SetDataOutput struct {
	Hash string
	Timestamp string
}
