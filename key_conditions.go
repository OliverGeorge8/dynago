package dynago

import "github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"

type KeyConditionsOpts struct {
	expr expression.KeyConditionBuilder
}

func KeyConditions() *KeyConditionsOpts {
	return &KeyConditionsOpts{expr: expression.KeyConditionBuilder{}}
}

func (k *KeyConditionsOpts) EqualTo(fieldName string, value interface{}) *KeyConditionsOpts {
	cond := expression.KeyEqual(expression.Key(fieldName), expression.Value(value))
	if !k.expr.IsSet() {
		k.expr = cond
	} else {
		k.expr.And(cond)
	}
	return k

}

func (k *KeyConditionsOpts) AndKeyCondition(kcb *KeyConditionsOpts) *KeyConditionsOpts {
	k.expr.And(kcb.expr)
	return k

}

func (k *KeyConditionsOpts) BeginsWith(fieldName string, value string) *KeyConditionsOpts {

	cond := expression.KeyBeginsWith(expression.Key(fieldName), value)
	if !k.expr.IsSet() {
		k.expr = cond
	} else {
		k.expr.And(cond)
	}

	return k

}

func (k *KeyConditionsOpts) Between(fieldName string, lowerValue interface{}, higher interface{}) *KeyConditionsOpts {

	cond := expression.KeyBetween(expression.Key(fieldName), expression.Value(lowerValue), expression.Value(higher))
	if !k.expr.IsSet() {
		k.expr = cond
	} else {
		k.expr.And(cond)
	}

	return k

}

func (k *KeyConditionsOpts) LessThan(fieldName string, value interface{}) *KeyConditionsOpts {

	cond := expression.KeyLessThan(expression.Key(fieldName), expression.Value(value))
	if !k.expr.IsSet() {
		k.expr = cond
	} else {
		k.expr.And(cond)
	}

	return k

}

func (k *KeyConditionsOpts) LessThanOrEqual(fieldName string, value interface{}) *KeyConditionsOpts {

	cond := expression.KeyLessThanEqual(expression.Key(fieldName), expression.Value(value))
	if !k.expr.IsSet() {
		k.expr = cond
	} else {
		k.expr.And(cond)
	}

	return k

}

func (k *KeyConditionsOpts) GreaterThanOrEqual(fieldName string, value interface{}) *KeyConditionsOpts {

	cond := expression.KeyGreaterThanEqual(expression.Key(fieldName), expression.Value(value))
	if !k.expr.IsSet() {
		k.expr = cond
	} else {
		k.expr.And(cond)
	}

	return k

}

func (k *KeyConditionsOpts) GreaterThan(fieldName string, value interface{}) *KeyConditionsOpts {

	cond := expression.KeyGreaterThan(expression.Key(fieldName), expression.Value(value))
	if !k.expr.IsSet() {
		k.expr = cond
	} else {
		k.expr.And(cond)
	}

	return k

}
