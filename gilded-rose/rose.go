// This package exports Item type and related UpdateQuality function. It's purpose is to
// update Items according to the assignment's Requirements Specification.
package rose

// Updater is the interface that wraps the Update method.
type Updater interface {
	// Update updates the Quality and Days remaining.
	Update()
}

// Item is an item sold by Gilded Rose.
type Item struct {
	Name    string
	SellIn  int
	Quality int
}

// Update updates the Quality and Days remaining.
func (i Item) Update() {
	// Do nothing
}

// UpdateQuality ages the item by a day, and updates the Quality of the item according to the Requirements Specification.
func UpdateQuality(items ...*Item) {
	for _, item := range items {
		item.Update()
	}
}
