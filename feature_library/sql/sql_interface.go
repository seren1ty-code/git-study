package feature_library

type DbInterface interface {
	CreateTable()
	DeleteRow()
	UpdateRow()
	InsertRow()
	GetRows()
	RedactBookByStruct()
}
