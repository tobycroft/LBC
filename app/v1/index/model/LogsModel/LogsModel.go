package LogsModel

import (
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "logs"

func Api_insert(log string) bool {
	db := tuuz.Db().Table(table)
	data := make(map[string]interface{})
	data["log"] = log
	db.Data(data)
	_, err := db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}
