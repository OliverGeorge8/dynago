package dynago

import "github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"

type FilterOpts struct {
	expr    expression.ConditionBuilder
	lastOpt string
}

func Filter() *FilterOpts {
	return &FilterOpts{lastOpt: "-", expr: expression.ConditionBuilder{}}
}

func (f *FilterOpts) EqualTo(fieldName string, value interface{}) *FilterOpts {
	if !f.expr.IsSet() {
		f.expr = expression.Equal(expression.Name(fieldName), expression.Value(value))
	} else if f.lastOpt == "and" {
		f.expr.And(expression.Equal(expression.Name(fieldName), expression.Value(value)))
	} else if f.lastOpt == "or" {
		f.expr.Or(expression.Equal(expression.Name(fieldName), expression.Value(value)))
	}

	return f
}

func (f *FilterOpts) And() *FilterOpts {
	f.lastOpt = "and"
	return f
}

func (f *FilterOpts) AndFilter(filter *FilterOpts) *FilterOpts {
	f.expr.And(filter.expr)
	return f
}

func (f *FilterOpts) OrFilter(filter *FilterOpts) *FilterOpts {
	f.expr.Or(filter.expr)
	return f

}

func (f *FilterOpts) Or() *FilterOpts {
	f.lastOpt = "or"
	return f

}

func (f *FilterOpts) Not() *FilterOpts {
	f.expr = f.expr.Not()
	return f
}

func (f *FilterOpts) NotEqualTo(fieldName string, value interface{}) *FilterOpts {

	if f.expr.IsSet() {
		f.expr = expression.NotEqual(expression.Name(fieldName), expression.Value(value))

	} else if f.lastOpt == "and" {
		f.expr.And(f.expr, expression.NotEqual(expression.Name(fieldName), expression.Value(value)))

	} else if f.lastOpt == "or" {
		f.expr.Or(f.expr, expression.NotEqual(expression.Name(fieldName), expression.Value(value)))
	}

	return f
}

func (f *FilterOpts) Exists(fieldName string) *FilterOpts {

	if f.expr.IsSet() {
		f.expr = expression.AttributeExists(expression.Name(fieldName))

	} else if f.lastOpt == "and" {
		f.expr.And(expression.AttributeExists(expression.Name(fieldName)))

	} else if f.lastOpt == "or" {
		f.expr.Or(expression.AttributeExists(expression.Name(fieldName)))
	}

	return f

}

func (f *FilterOpts) NotExists(fieldName string) *FilterOpts {
	if f.expr.IsSet() {
		f.expr = expression.AttributeNotExists(expression.Name(fieldName))

	} else if f.lastOpt == "and" {
		f.expr.And(expression.AttributeNotExists(expression.Name(fieldName)))

	} else if f.lastOpt == "or" {
		f.expr.Or(expression.AttributeNotExists(expression.Name(fieldName)))
	}
	return f

}

func (f *FilterOpts) LessThan(fieldName string, value interface{}) *FilterOpts {
	f.expr.IsSet()
	f.expr.And(f.expr, expression.LessThan(expression.Name(fieldName), expression.Value(value)))

	return f

}
func (f *FilterOpts) LessThanOrEqual(fieldName string, value interface{}) *FilterOpts {

	f.expr.And(f.expr, expression.LessThanEqual(expression.Name(fieldName), expression.Value(value)))
	return f
}

func (f *FilterOpts) GreaterThanOrEqual(fieldName string, value interface{}) *FilterOpts {

	f.expr.And(f.expr, expression.GreaterThanEqual(expression.Name(fieldName), expression.Value(value)))
	return f

}
func (f *FilterOpts) GreaterThan(fieldName string, value interface{}) *FilterOpts {

	f.expr.And(f.expr, expression.GreaterThan(expression.Name(fieldName), expression.Value(value)))
	return f

}

func (f *FilterOpts) In(fieldName string, value []interface{}) *FilterOpts {

	f.expr.And(f.expr, expression.In(expression.Name(fieldName), expression.Value(value)))
	return f

}
func (f *FilterOpts) NotIn(fieldName string, value []interface{}) *FilterOpts {

	f.expr.And(f.expr, expression.Not(expression.In(expression.Name(fieldName), expression.Value(value))))
	return f

}

func (f *FilterOpts) Contains(fieldName string, value interface{}) *FilterOpts {
	return f

}

func (f *FilterOpts) BeginsWith(fieldName string, value interface{}) *FilterOpts {
	return f
}
