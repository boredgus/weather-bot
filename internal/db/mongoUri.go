package db

import (
	"bytes"
	"fmt"
	"subscription-bot/config"
	"text/template"
)

const MongoDBUriTemplate = "mongodb://{{.Username}}:{{.Password}}@{{.Server}}:27017"

type MongoDBUri struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Server   string `json:"server"`
}

func NewMongoDbUri() MongoDBUri {
	config := config.GetConfig()
	return MongoDBUri{
		Username: config.MongoDBUsername,
		Password: config.MongoDBPassword,
		Server:   config.DbServer,
	}
}

func (r MongoDBUri) newTemplate() (*template.Template, error) {
	tmpl, err := template.New("mongo_db_uri").Parse(MongoDBUriTemplate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse mongo db uri template: %w", err)
	}
	return tmpl, nil
}

func (r MongoDBUri) Generate() (string, error) {
	template, err := r.newTemplate()
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	err = template.Execute(&buffer, r)
	if err != nil {
		return "", fmt.Errorf("failed to execute weather mongo db uri template: %w", err)
	}
	return buffer.String(), nil
}
