package dynago

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
)

type UpdateOpts struct {
	expr      expression.UpdateBuilder
	returnNew bool
}

func Update() *UpdateOpts {
	return &UpdateOpts{}
}

func (u *UpdateOpts) Set(fieldName string, value interface{}) *UpdateOpts {

	u.expr.Set(expression.Name(fieldName), expression.Value(value))
	return u
}

func (u *UpdateOpts) Increment(fieldName string, value interface{}) *UpdateOpts {

	u.expr.Set(expression.Name(fieldName), expression.Name(fieldName).Plus(expression.Value(value)))
	return u
}

func (u *UpdateOpts) Decrement(fieldName string, value interface{}) *UpdateOpts {

	u.expr.Set(expression.Name(fieldName), expression.Name(fieldName).Minus(expression.Value(value)))
	return u
}

func (u *UpdateOpts) Delete(fieldName string, value interface{}) *UpdateOpts {
	u.expr.Delete(expression.Name(fieldName), expression.Value(value))
	return u
}

func (u *UpdateOpts) Remove(fieldName string, value interface{}) *UpdateOpts {
	u.expr.Remove(expression.Name(fieldName))
	return u
}

func (u *UpdateOpts) ListAppend(fieldName string, value interface{}) *UpdateOpts {
	u.expr.Set(expression.Name(fieldName), expression.ListAppend(expression.Name(fieldName), expression.Value(value)))
	return u
}

func (u *UpdateOpts) ReturnNew() *UpdateOpts {
	u.returnNew = true
	return u
}
