package db

type Db struct {
	Location string `json:"location"`
	Id       string `json:"id"`
}
type Dbs []Db
