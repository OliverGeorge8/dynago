package dynago

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"os"
)

type Dynago struct {
	client    *dynamodb.Client
	tableName string
}

var instance *Dynago

func Initialize() (*Dynago, error) {
	tableName := os.Getenv("DYNAMO_TABLE_NAME")
	cfg, err := config.LoadDefaultConfig(context.Background())

	if err != nil {
		return nil, err
	}
	client := dynamodb.NewFromConfig(cfg)

	return &Dynago{client: client, tableName: tableName}, nil
}

func Get() *Dynago {
	if instance == nil {
		initInstance, err := Initialize()

		if err != nil {
			panic(err)
		}
		instance = initInstance
	}
	return instance
}

func (d *Dynago) BatchWrite() *BatchWriteOpts {
	return &BatchWriteOpts{dynago: d}
}

func (d *Dynago) BatchGet() *BatchWriteOpts {
	return &BatchWriteOpts{dynago: d}
}

func (d *Dynago) TransactionWrite() *TransactionWriteOpts {
	return &TransactionWriteOpts{dynago: d}
}

func (d *Dynago) TransactionGet() *TransactionWriteOpts {
	return &TransactionWriteOpts{dynago: d}
}

func (d *Dynago) GetItemWithCtx(ctx context.Context, key *KeyImpl, model interface{}) error {
	result, err := d.client.GetItem(ctx, &dynamodb.GetItemInput{TableName: aws.String(d.tableName), Key: key.Build()})
	if err != nil {
		return err
	}

	err = attributevalue.UnmarshalMap(result.Item, model)
	return err
}

func (d *Dynago) GetItem(key *KeyImpl, model IModel) error {
	return d.GetItemWithCtx(context.Background(), key, model)

}

func (d *Dynago) PutItemWithCtx(ctx context.Context, documentModel IModel) error {
	beforeInsertHook, ok := (documentModel).(BeforeInsert)

	if ok {
		beforeInsertHook.BeforeInsert()
	}
	marshaledMap, err := attributevalue.MarshalMap(documentModel)
	if err != nil {
		return err
	}

	_, err = d.client.PutItem(ctx, &dynamodb.PutItemInput{TableName: aws.String(d.tableName), Item: marshaledMap})
	if err != nil {
		return err
	}

	return nil
}
func (d *Dynago) PutItem(documentModel IModel) error {
	return d.PutItemWithCtx(context.Background(), documentModel)
}

func (d *Dynago) UpdateItemWithCtx(ctx context.Context, key *KeyImpl, update *UpdateOpts, model ...interface{}) (bool, error) {

	expr, err := NewExpr().WithUpdate(update).build()

	result, err := d.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName:                 aws.String(d.tableName),
		Key:                       key.Build(),
		UpdateExpression:          expr.Update(),
		ConditionExpression:       expr.Condition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ReturnValues:              types.ReturnValueAllNew})
	if err != nil {
		return false, err
	}

	if len(model) > 0 {
		err = attributevalue.UnmarshalMap(result.Attributes, model[0])
		if err != nil {
			return false, err
		}
	}

	return true, nil

}

func (d *Dynago) UpdateItem(key *KeyImpl, update *UpdateOpts, model ...interface{}) (bool, error) {
	return d.UpdateItemWithCtx(context.Background(), key, update, model)
}

func (d *Dynago) DeleteItemWithCtx(ctx context.Context, key KeyImpl) (*dynamodb.DeleteItemOutput, error) {
	return d.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{TableName: aws.String(d.tableName), Key: key.Build()})

}

func (d *Dynago) DeleteItem(key KeyImpl) (*dynamodb.DeleteItemOutput, error) {
	return d.DeleteItemWithCtx(context.Background(), key)
}

func (d *Dynago) QueryWithCtx(ctx context.Context, expr *Expression, out interface{}) error {

	exprExpr, err := expr.build()

	if err != nil {
		return err
	}
	queryInput := &dynamodb.QueryInput{
		TableName:                 aws.String(d.tableName),
		KeyConditionExpression:    exprExpr.KeyCondition(),
		FilterExpression:          exprExpr.Filter(),
		ExpressionAttributeNames:  exprExpr.Names(),
		ExpressionAttributeValues: exprExpr.Values(),
		ProjectionExpression:      exprExpr.Projection()}
	if expr.indexName != nil {
		queryInput.IndexName = expr.indexName
	}

	if expr.lastIndex != nil {
		queryInput.ExclusiveStartKey = expr.lastIndex
	}

	if expr.limit != nil {
		queryInput.Limit = expr.limit
	}
	result, err := d.client.Query(ctx, queryInput)
	if err != nil {
		return err
	}

	err = attributevalue.UnmarshalListOfMaps(result.Items, &out)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	return nil

}

func (d *Dynago) Query(filter *Expression, out interface{}) error {
	return d.QueryWithCtx(context.Background(), filter, out)
}
