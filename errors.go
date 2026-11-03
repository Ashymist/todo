package todo

import "fmt"

type ItemDoesntExist struct {
	ItemId int
}

func (r *ItemDoesntExist) Error() string {
	return fmt.Sprintf("Error: Item %d does not exist", r.ItemId)
}
