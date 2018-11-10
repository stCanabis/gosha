package cmd

const msInsertDataToDb = `package bootstrap

import (
	"{ms-name}/dbmodels"
	"fmt"
	"os"
	"{ms-name}/core"
)

func FillDBTestData()  {

	isDropTables := false

	if (len(os.Args) > 1 && os.Args[1] == "drop") ||
		(len(os.Args) > 2 && os.Args[2] == "drop") {
		isDropTables = true
	}

	if isDropTables == true {

		core.Db.DropTableIfExists(
          dbmodels.Entity{},
		)

		fmt.Println("All tables removed")
		os.Exit(1)
	}

	core.Db.AutoMigrate(
		&dbmodels.Entity{},
	)

	// add fixtures

}`

var msTemplateInsertDataToDb = template{
    Path:    "./bootstrap/insert_data_to_db.go",
    Content: assignMsName(msInsertDataToDb),
}
