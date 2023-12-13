package services

import (
	"archive/actions/utils/interfaces"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type ResponseMetaData struct {
	Meta *interfaces.MetaData `json:"meta"`
	Data interface{}          `json:"data"`
}

var (
	pathName = "config"
	nameFile = "message.yml"
)

func ReadFile(tag string) (*interfaces.MetaData, error) {
	configFile := filepath.Join(pathName, nameFile)

	contentFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var yamlConfig map[string]interface{}
	err = yaml.Unmarshal(contentFile, &yamlConfig)
	if err != nil {
		return nil, err
	}

	jsonContent, err := json.Marshal(yamlConfig[tag])
	if err != nil {
		return nil, err
	}

	var metaData interfaces.MetaData
	err = yaml.Unmarshal(jsonContent, &metaData)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Println(metaData)

	return &metaData, nil
}

func BuildResponseHandler(tag string, target interface{}) (*ResponseMetaData, error) {
	metaData, err := ReadFile(tag)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var responseMetaData *ResponseMetaData

	responseMetaData = &ResponseMetaData{
		Meta: metaData,
		Data: target,
	}

	return responseMetaData, nil
}
