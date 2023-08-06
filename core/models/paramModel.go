package models

type Params struct {
	Id int `uri:"id"`
}

type FileParams struct {
	Id   int    `uri:"id"`
	Name string `uri:"name"`
}
