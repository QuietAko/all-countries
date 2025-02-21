package entity

type Country struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (c *Country) GetID() int {
	return c.ID
}