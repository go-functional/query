package query

type Builder interface {
	Select() Select
	SelectDistinct() Select
	SelectCount() Select
	// Update() Update
	// Delete() Delete
	// InsertInto() InsertInto
	// CreateDatabase() CreateDatabase
	// AlterDatabase() AlterDatase
	// CreateTable() CreateTable
	// AlterTable() AlterTable
	// DropTable() DropTable
	// CreateIndex() CreateIndex
	// DropIndex() DropIndex
}

type Select interface {
	As(string) Rows
	Columns(...string) Rows
	All() Rows
}

type Rows interface {
	FromTable(string) Table
	FromSelect(Select) Table
}

type Table interface {
	Where(...Condition)
}
