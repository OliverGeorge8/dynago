package dynago

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type TransactionWriteOpts struct {
	dynago *Dynago
	items  *dynamodb.TransactWriteItemsInput
}

func (t *TransactionWriteOpts) Update(key KeyImpl, update UpdateOpts) {

	t.items.TransactItems = append(t.items.TransactItems, types.TransactWriteItem{Update: &types.Update{Key: key.Build(), TableName: t.dynago.tableName}})
}

func (t *TransactionWriteOpts) Delete(key KeyImpl, update UpdateOpts) {
	t.items.TransactItems = append(t.items.TransactItems, types.TransactWriteItem{Delete: &types.Delete{Key: key.Build(), TableName: t.dynago.tableName}})

}

func (t *TransactionWriteOpts) Put(key KeyImpl, item interface{}) {
	t.items.TransactItems = append(t.items.TransactItems, types.TransactWriteItem{Delete: &types.Delete{Key: key.Build(), TableName: t.dynago.tableName}})
}

func (t *TransactionWriteOpts) ConditionCheck(query Expression) {

}

func (t *TransactionWriteOpts) CommitWithCtx(ctx context.Context) (*dynamodb.TransactWriteItemsOutput, error) {
	items, err := t.dynago.client.TransactWriteItems(ctx, t.items)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (t *TransactionWriteOpts) Commit() (*dynamodb.TransactWriteItemsOutput, error) {
	return t.CommitWithCtx(context.Background())
}
