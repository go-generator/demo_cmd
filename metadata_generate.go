package main

import (
	"context"
	"encoding/json"
	"flag"
	"github.com/go-generator/database"
	_ "github.com/go-generator/database/db_config"
	newConfig "github.com/go-generator/database/db_config"
	"github.com/go-generator/project"
	"io/ioutil"
	"log"
)

func SaveMetadataJson(projectStruct interface{}, filePath string) error { //s *SqlTablesData, conn *gorm.DB, tables []string, packageName, output string) {
	data, err := json.MarshalIndent(&projectStruct, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filePath, data, 0644) // Create and write files
	if err != nil {
		return err
	}
	return err
}
func JsonRead(filename string) (newConfig.DatabaseConfig, error) {
	var input newConfig.DatabaseConfig
	byteValue, err := ioutil.ReadFile(filename)
	if err != nil {
		return input, err
	}
	err = json.Unmarshal(byteValue, &input)
	if err != nil {
		return input, err
	}
	return input, nil
}
func WriteMetadata(dc newConfig.DatabaseConfig, optimize bool, seperateFile bool, inFile, outFile, modelFile string) error {
	err := dc.ValidateDatabaseConfig()
	if err != nil {
		return err
	}
	conn, err := dc.ConnectToSqlServer()
	if err != nil {
		return err
	}
	var t project.DatabaseAdapter
	t = &database.DefaultMetadataService{Config: dc}
	var u project.ProjectService
	u = &project.GoMongoProjectService{}
	projectStruct, err := u.CreateProjectByAdapter(context.Background(), inFile, t, conn)
	if err != nil {
		return err
	}
	if seperateFile {
		projectStruct.ModelsFile = modelFile
		err = SaveMetadataJson(projectStruct.Models, modelFile)
		if err != nil {
			return err
		}
		projectStruct.Models = nil
		err = SaveMetadataJson(projectStruct, outFile)
		if err != nil {
			return err
		}
	} else {
		err = SaveMetadataJson(projectStruct, outFile)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	var i = flag.String("i", "goWithoutModel.json", "input json")
	var o = flag.String("o", "metadata.json", "output json")
	var c = flag.String("c", "dbConfig.json", "config json")
	var m = flag.String("m", "", "separated metadata output")

	flag.Parse()
	dbConfig, err := JsonRead(*c)
	if err != nil {
		log.Panicln(err)
	}
	err = WriteMetadata(dbConfig, false, *m != "", *i, *o, *m)
	if err != nil {
		log.Panicln(err)
	}
	log.Print("Successful")

}
