package dynago

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type BatchWriteOpts struct {
	dynago *Dynago
	input  []types.WriteRequest
}

func (b *BatchWriteOpts) Delete(key KeyImpl) *BatchWriteOpts {
	b.input = append(b.input, types.WriteRequest{DeleteRequest: &types.DeleteRequest{Key: key.Build()}})
	return b
}

func (b *BatchWriteOpts) BatchDelete(keys ...KeyImpl) *BatchWriteOpts {
	for _, key := range keys {
		b.Delete(key)
	}

	return b
}

func (b *BatchWriteOpts) Put(item interface{}) *BatchWriteOpts {

	marshalMap, err := attributevalue.MarshalMap(item)
	if err != nil {
		return b
	}
	b.input = append(b.input, types.WriteRequest{PutRequest: &types.PutRequest{Item: marshalMap}})
	return b
}

func (b *BatchWriteOpts) BatchPut(items ...interface{}) *BatchWriteOpts {
	for _, item := range items {
		b.Put(item)
	}
	return b
}

func (b *BatchWriteOpts) CommitWithCtx(ctx context.Context) (*dynamodb.BatchWriteItemOutput, error) {

	result, err := b.dynago.client.BatchWriteItem(ctx, &dynamodb.BatchWriteItemInput{RequestItems: map[string][]types.WriteRequest{
		b.dynago.tableName: b.input,
	}})
	if err != nil {
		return nil, err
	}

	return result, nil

}

func (b *BatchWriteOpts) Commit() (*dynamodb.BatchWriteItemOutput, error) {
	return b.CommitWithCtx(context.Background())
}
