package model

type Funfact struct {
	Id      int
	FunFact string
}

type FunFactPost struct {
	Funfact string
}

type FunFactList []Funfact
