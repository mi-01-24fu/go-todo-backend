package addition

import "database/sql"

type AdditionTask interface {
	ADD()
}

type AdditionTaskImpl struct {
	DB *sql.DB
}

func NewAdditionTaskImpl(db *sql.DB) *AdditionTaskImpl {
	return &AdditionTaskImpl{
		DB: db,
	}
}

func (a AdditionTaskImpl) ADD() {}
