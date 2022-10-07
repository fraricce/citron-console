package main

import (
	"log"

	"strconv"
)

type Service struct {
	dbdriver   string
	dbpath     string
	repository Repo
}

func (service Service) Init() bool {
	service.repository = Repo{dbdriver: service.dbdriver, dbpath: service.dbpath}
	service.repository.InitDatabase()
	return true
}

func (service Service) AddEntity(entityName string, text string) (bool, string) {
	log.Println("AddEntity")
	var eCode int

	switch *addEntity {
	case "task":
		eCode = 0
		break
	case "note":
		eCode = 1
		break
	}

	item := entity{text: text, done: false, entity: eCode}
	if !service.repository.CreateEntity(item) {
		return false, entityName
	}
	return true, entityName
}

func (service Service) ListEntities(entityName string) [][]interface{} {
	var eCode int

	switch entityName {
	case "task":
		eCode = 0
		break
	case "note":
		eCode = 1
		break
	}

	return service.repository.FindAllEntities(eCode)
}

func (service Service) RemoveEntity(entityId string) int64 {
	eId, _ := strconv.Atoi(*deleteEntityId)
	return service.repository.DeleteEntity(eId)
}

func (service Service) SetEntityStatus(entityId string, done int) int64 {
	eId, _ := strconv.Atoi(entityId)
	return service.repository.UpdateEntityStatus(eId, done)
}
