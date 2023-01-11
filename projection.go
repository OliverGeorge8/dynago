package dynago

import "github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"

type ProjectionOpts struct {
	projection expression.ProjectionBuilder
}

func Projection(attributes ...string) *ProjectionOpts {
	var names []expression.NameBuilder
	for _, v := range attributes {
		names = append(names, expression.Name(v))
	}
	projBuilder := expression.AddNames(expression.ProjectionBuilder{}, names...)

	return &ProjectionOpts{projection: projBuilder}
}
