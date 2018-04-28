// Package dynamo gives common bindings to model related activities from AWS DynamoDB
package dynamo

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// DefaultSession uses environment variables and standard AWS paths to generate a session
var DefaultSession = session.Must(session.NewSession())

// DefaultClient uses the default session to generate a client
var DefaultClient = dynamodb.New(DefaultSession)

// Model implements the reelvant actions for a dynamoDB model
type Model struct {
	client    *dynamodb.DynamoDB
	tableName *string
}

// Create saves a model to DynamoDB. Note that this is only for new instances
func (m *Model) Create() (*dynamodb.PutItemOutput, error) {
	var (
		parsed map[string]*dynamodb.AttributeValue
		err    error
	)
	if parsed, err = m.marshal(); err != nil {
		return nil, err
	}
	input := &dynamodb.PutItemInput{
		TableName: m.tableName,
		Item:      parsed,
	}
	return m.client.PutItem(input)
}

func (m *Model) marshal() (map[string]*dynamodb.AttributeValue, error) {
	return dynamodbattribute.MarshalMap(m)
}
