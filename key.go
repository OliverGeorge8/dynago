package dynago

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type KeyImpl struct {
	pk interface{}
	sk interface{}
}

func Key(pk, sk interface{}) *KeyImpl {
	key := KeyImpl{}
	key.PartitionKey(pk).SortKey(sk)
	return &key
}

func (k *KeyImpl) Build() map[string]types.AttributeValue {
	marshalMap, err := attributevalue.MarshalMap(k)
	if err != nil {
		return nil
	}
	return marshalMap
}

func (k *KeyImpl) PartitionKey(pk interface{}) *KeyImpl {
	switch pk.(type) {
	case string:
		k.pk = aws.String(pk.(string))
	case int32:
		k.pk = aws.Int32(pk.(int32))
	}

	return k
}

func (k *KeyImpl) SortKey(sk interface{}) *KeyImpl {
	switch sk.(type) {
	case string:
		k.pk = aws.String(sk.(string))
	case int32:
		k.pk = aws.Int32(sk.(int32))
	}

	return k
}
