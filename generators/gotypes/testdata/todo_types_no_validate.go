// Item A to-do item.
type Item struct {
  // CreatedAt is the time the to-do item was created.
  CreatedAt time.Time `json:"created_at"`

  // ID is the id of the item. This field is read-only.
  ID int `json:"id"`

  // Text is the to-do item text. This field is required.
  Text string `json:"text"`
}

// AddItemInput params.
type AddItemInput struct {
  // Item is the item to add. This field is required.
  Item string `json:"item"`
}

// GetItemsOutput params.
type GetItemsOutput struct {
  // Items is the list of to-do items.
  Items []Item `json:"items"`
}

// RemoveItemInput params.
type RemoveItemInput struct {
  // ID is the id of the item to remove.
  ID int `json:"id"`
}


// oneOf returns true if s is in the values.
func oneOf(s string, values []string) bool {
  for _, v := range values {
		if s == v {
			return true
		}
	}
	return false
}
