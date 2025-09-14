package model

type Funfact struct {
	Id      int
	FunFact string
}

type FunFactPost struct {
	Funfact string
}

type FunFactList []Funfact

type BasicResponse struct {
	Action string
	Message string
}