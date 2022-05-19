package main

import (
	"log"

	"strconv"
)

func AddEntity(entityName string, text string) (bool, string) {
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
	if !CreateEntity(item) {
		return false, entityName
	}
	return true, entityName
}

func ListEntities(entityName string) [][]interface{} {
	var eCode int

	switch entityName {
	case "task":
		eCode = 0
		break
	case "note":
		eCode = 1
		break
	}

	return FindAllEntities(eCode)
}

func RemoveEntity(entityId string) int64 {
	eId, _ := strconv.Atoi(*deleteEntityId)
	return DeleteEntity(eId)
}

func SetEntityStatus(entityId string, done int) int64 {
	eId, _ := strconv.Atoi(entityId)
	return UpdateEntityStatus(eId, done)
}
