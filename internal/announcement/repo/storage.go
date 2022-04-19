package store

type (
	Storage interface {
		Create([]byte) (int64, error)
		Update(id int, path string) ([]byte, error)
		Get(id int64, opt []string) ([]byte, error)
		GetList(page, limit int, sortBy, sort string) ([]byte, error)
	}
)
