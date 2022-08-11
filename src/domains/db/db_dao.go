package db

import (
	"context"
	"encoding/json"
	"fasthttp/datasources/pg"
	"fasthttp/datasources/redis"
	"fasthttp/utils/errors"
	"fasthttp/utils/timer"
	"fmt"
	"log"
)

const (
	queryUserSelectAll = "select array_to_json(array_agg(row_to_json(r.*))) as rowjson from (select * from oky.oky_mn.oky_user order by country limit 200)r"
)

func LoadAllUserNoCache() []byte {
	defer timer.Timer()()
	fmt.Print("From DB\n")
	rows, err := pg.PgPool.Query(context.Background(), queryUserSelectAll)
	if err != nil {
		fmt.Printf("%v", err)
		log.Fatal("error while executing query")
		errors.NewInternalServerError("database error", nil)
	}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
			errors.NewInternalServerError("database error", nil)
		}
		a, _ := json.Marshal(values[0])

		return a
	}

	return []byte("")
}

func LoadAllUser() []byte {
	defer timer.Timer()()
	usrString, rerr := redis.RedisService.GetRedis("LoadAllUser")
	if rerr != nil {
		errors.NewInternalServerError("unauthorized", rerr)
	}
	if usrString != "" {
		fmt.Print("From Redis\n")
		a, _ := json.Marshal(usrString)
		return a
	} else {

		rows, err := pg.PgPool.Query(context.Background(), queryUserSelectAll)
		if err != nil {
			fmt.Printf("%v", err)
			log.Fatal("error while executing query")
			errors.NewInternalServerError("database error", nil)
		}
		for rows.Next() {
			fmt.Print("From DB\n")
			values, err := rows.Values()
			if err != nil {
				log.Fatal("error while iterating dataset")
				errors.NewInternalServerError("database error", nil)
			}
			a, _ := json.Marshal(values[0])
			if rerr := redis.RedisService.SetRedis(redis.Params{Id: "LoadAllUser", Data: string(a)}); rerr != nil {
				errors.NewInternalServerError("try to save redis", rerr)
			}
			//convert DB types to Go types

			return a
		}
	}
	return []byte("")
}
