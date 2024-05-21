// This package exports Item type and related UpdateQuality function. It's purpose is to
// update Items according to the assignment's Requirements Specification.
package rose

import "fmt"

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

// String outputs text representation of an Item.
func (i *Item) String() string {
	return fmt.Sprintf("%s\t%d\t%d\t\n", i.Name, i.SellIn, i.Quality)
}

// minQuality and maxQuality spcify upper and lowerbounds for Item.Qualty value on update.
// Item "Sulfuras" is the only exception and has a cosntant Quality of 80
const (
	minQuality = 0
	maxQuality = 50
)

// category is Enum used to collect possible Item categories
type category int

// possible Item categories
const (
	normalCategory = iota
	agedBrieCategory
	sulfurasCategory
	backstageCategory
	conjuredCategory
)

// itemCategory is responsible for determining Item category from it's Name field.
// This function exsits to isolate the logic of category determination
func itemCategory(name string) (category category) {
	// more logic could be placed here in case we need more complex rules
	// for name <> category
	switch name {
	case "Aged Brie":
		category = agedBrieCategory
	case "Sulfuras, Hand of Ragnaros":
		category = sulfurasCategory
	case "Backstage passes to a TAFKAL80ETC concert":
		category = backstageCategory
	case "Conjured Mana Cake":
		category = conjuredCategory
	default:
		category = normalCategory
	}
	return
}

// normal is a normal item.
type normal struct {
	*Item
}

// Update Quality and SellIn fields.
func (i normal) Update() {
	var qualityAdj int
	if i.SellIn > 0 {
		qualityAdj = -1
	} else {
		qualityAdj = -2
	}
	i.Quality = max((i.Quality + qualityAdj), minQuality)
	i.SellIn--
}

// agedBrie is an "Aged Brie".
type agedBrie struct {
	*Item
}

// Update Quality and SellIn fields.
func (i agedBrie) Update() {
	var qualityAdj int
	if i.SellIn > 0 {
		qualityAdj = 1
	} else {
		qualityAdj = 2
	}
	i.Quality = min((i.Quality + qualityAdj), maxQuality)
	i.SellIn--
}

// sulfuras is an "Sulfuras".
type sulfuras struct {
	*Item
}

// Update Quality and SellIn fields. Sulfuras is an execption - we don't udapte it's fields.
func (i sulfuras) Update() {
	// Do nothing
	return
}

// Backstage is "Backstage passes to a TAFKAL80ETC concert".
type backstage struct {
	*Item
}

// Update Quality and SellIn fields.
func (i backstage) Update() {
	var qualityAdj int
	if i.SellIn > 10 {
		qualityAdj = 1
	} else if i.SellIn > 5 {
		qualityAdj = 2
	} else if i.SellIn > 0 {
		qualityAdj = 3
	} else {
		i.Quality = 0
	}
	i.Quality = min((i.Quality + qualityAdj), maxQuality)
	i.SellIn--
}

// conjured is a "Conjured Mana Cake".
type conjured struct {
	*Item
}

// Update Quality and SellIn fields.
func (i conjured) Update() {
	var qualityAdj int
	if i.SellIn > 0 {
		qualityAdj = -2
	} else {
		qualityAdj = -4
	}
	i.Quality = max((i.Quality + qualityAdj), minQuality)
	i.SellIn--
}

// UpdateQuality ages the item by a day, and updates the Quality of the item according to the Requirements Specification.
func UpdateQuality(items ...*Item) {
	for _, item := range items {
		var u Updater
		category := itemCategory(item.Name)
		switch category {
		case normalCategory:
			u = normal{item}
		case agedBrieCategory:
			u = agedBrie{item}
		case backstageCategory:
			u = backstage{item}
		case sulfurasCategory:
			u = sulfuras{item}
		case conjuredCategory:
			u = conjured{item}
		}
		u.Update()
	}
}
