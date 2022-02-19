package choice

import (
	"berkantbegdilili.com/first-project/database"
)

func (c *choice) add() error {
	return nil
}

func (c *choice) getOneById(id uint) error {
	db := database.Connection()
	result := db.Where(choice{ID: id}).First(c)
	return result.Error
}

func (c *choice) getAll() error {

	return nil
}

func (c *choice) update() error {

	return nil
}

func (c *choice) delete() error {

	return nil
}
