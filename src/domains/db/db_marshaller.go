package db

import (
	"encoding/json"
)

type PublicDb struct {
	Id       string `json:"id"`
	Location string `json:"location"`
}

func (db *Db) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicDb{
			Id:       db.Id,
			Location: db.Location,
		}
	}

	userJson, _ := json.Marshal(db)
	var pUser PublicDb
	json.Unmarshal(userJson, &pUser)

	return pUser
}

func (dbs Dbs) Marshall(isPublic bool) interface{} {
	result := make([]interface{}, len(dbs))
	for index, db := range dbs {
		result[index] = db.Marshall(isPublic)
	}
	return result
}
