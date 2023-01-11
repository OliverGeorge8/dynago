package dynago

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type BatchGetOpts struct {
	dynago     *Dynago
	input      []map[string]types.AttributeValue
	projection ProjectionOpts
}

func (b *BatchGetOpts) Get(key KeyImpl) {
	b.input = append(b.input, key.Build())
}

func (b *BatchGetOpts) Projection(projection ProjectionOpts) {
	b.projection = projection
}

func (b *BatchGetOpts) BatchGet(keys ...KeyImpl) {
	for _, key := range keys {
		b.Get(key)
	}
}

func (b *BatchGetOpts) CommitWithCtx(ctx context.Context, model interface{}) (*dynamodb.BatchGetItemOutput, error) {
	return b.dynago.client.BatchGetItem(ctx, &dynamodb.BatchGetItemInput{RequestItems: map[string]types.KeysAndAttributes{
		b.dynago.tableName: {Keys: b.input},
	}})

}

func (b *BatchGetOpts) Commit(model interface{}) (*dynamodb.BatchGetItemOutput, error) {
	return b.CommitWithCtx(context.Background(), model)
}