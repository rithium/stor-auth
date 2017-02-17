package model

import (
	"fmt"
	"crypto/rand"
	"time"
	"log"
)

type ApiKey struct {
	Id	int64	`json:"id"`
	Key	string	`json:"key"`
	Active	bool	`json:"-"`
	Created	int64	`json:"-"`
}

func newApiKey()(*ApiKey) {
	entry := &ApiKey{}

	entry.Key = generateApiKey()
	entry.Active = true
	entry.Created = time.Now().Unix()

	return entry
}

func (db *DB)CreateApiKey()(*ApiKey, error) {
	apiKey := newApiKey()

	result, err := db.Exec(`insert into apiKey values (null,?,?,?)`,
		apiKey.Key,
		apiKey.Active,
		apiKey.Created)

	if err != nil {
		return nil, err
	}

	apiKey.Id, err = result.LastInsertId()

	if err != nil {
		return nil, err
	}

	return apiKey, nil
}

func (db*DB)FindActiveKey(key string)(*ApiKey, error) {
	result, err := db.Query(`select id, `+"`key`"+`, active, created from apikey where `+"`key`"+` = ? and active = 1`,
		key)

	if err != nil {
		return nil, err
	}

	var apiKey ApiKey

	defer result.Close()

	if result.Next() {
		log.Println("result next")

		result.Scan(&apiKey.Id, &apiKey.Key, &apiKey.Active, &apiKey.Created)

		return &apiKey, nil
	}

	return nil, nil
}

func (db *DB)KeyExists(key string)(bool, error) {
	result, err := db.Query(`select * from apikey where `+"`key`"+` = ?`,
		key)

	if err != nil {
		return false, err
	}

	defer result.Close()

	if result.Next() {
		return true, nil
	}

	return false, nil
}

func generateApiKey()(string) {
	b := make([]byte, 32)
	rand.Read(b)

	return fmt.Sprintf("%x", b)
}