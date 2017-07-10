package query

type Operatior struct{}

type Condition interface {
	Operator() Operator
}
