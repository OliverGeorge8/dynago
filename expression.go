package dynago

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const (
	SORT_ASC  = 1
	SORT_DESC = -1
)

type Expression struct {
	keyConditions *KeyConditionsOpts
	filter        *FilterOpts
	projection    *ProjectionOpts
	update        *UpdateOpts
	indexName     *string
	lastIndex     map[string]types.AttributeValue
	limit         *int32
	sort          *string
}

func NewExpr() *Expression {
	return &Expression{}
}

func (q *Expression) build() (expression.Expression, error) {
	return expression.NewBuilder().WithKeyCondition(q.keyConditions.expr).WithFilter(q.filter.expr).WithProjection(q.projection.projection).WithUpdate(q.update.expr).Build()
}

func (q *Expression) WithUpdate(update *UpdateOpts) *Expression {
	q.update = update
	return q
}

func (q *Expression) WithKeyConditions(keyConds *KeyConditionsOpts) *Expression {
	q.keyConditions = keyConds
	return q
}

func (q *Expression) WithFilter(filter *FilterOpts) *Expression {
	q.filter = filter
	return q
}

func (f *Expression) WithProjection(opts *ProjectionOpts) *Expression {
	f.projection = opts
	return f
}

func (q *Expression) Index(name string) *Expression {
	q.indexName = aws.String(name)
	return q

}

func (f *Expression) Sort(sort string) {
	f.sort = aws.String(sort)
}

func (f *Expression) Limit(limit int32) {
	f.limit = aws.Int32(limit)
}

func (f *Expression) LastId(lastId map[string]types.AttributeValue) {
	f.lastIndex = lastId
}
