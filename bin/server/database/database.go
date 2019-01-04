package database

//Database is structure of database manipulator
type Database struct {
}

//NewDatabase return new structure of database manipulator
func NewDatabase() (Database, error) {
	out := Database{}

	return out, nil
}
