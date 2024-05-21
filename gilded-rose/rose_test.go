package rose

import (
	"testing"
)

var cases = []struct {
	name        string
	description string
	sellIn      int
	quality     int
	wantSellIn  int
	wantQuality int
}{
	{"normal", "product update pre sellIn",
		3, 8, 2, 7},
	{"normal2", "product update on sellIn",
		0, 3, -1, 1},
	{"normal3", "product update past sellIn",
		-1, 5, -2, 3},
	{"normal4", "product update with 0 quality",
		2, 0, 1, 0},

	{"Aged Brie", "pre sellIn with max quality",
		10, 50, 9, 50},
	{"Aged Brie", "on sellIn reaches max quality",
		0, 49, -1, 50},
	{"Aged Brie", "on sellIn with max quality",
		0, 50, -1, 50},
	{"Aged Brie", "past sellIn",
		-13, 10, -14, 12},
	{"Aged Brie", "past sellIn with max quality",
		-15, 50, -16, 50},

	{"Sulfuras, Hand of Ragnaros", "pre sellIn",
		1, 80, 1, 80},
	{"Sulfuras, Hand of Ragnaros", "past sellIn",
		-1, 80, -1, 80},

	{"Backstage passes to a TAFKAL80ETC concert", ">10 pre sellIn",
		15, 10, 14, 11},
	{"Backstage passes to a TAFKAL80ETC concert", ">10 pre sellIn max quality",
		11, 50, 10, 50},
	{"Backstage passes to a TAFKAL80ETC concert", "10>5 to sellIn #1",
		9, 10, 8, 12},
	{"Backstage passes to a TAFKAL80ETC concert", "10>5 to sellIn #2",
		6, 14, 5, 16},
	{"Backstage passes to a TAFKAL80ETC concert", "5>0 to sellIn #1",
		5, 9, 4, 12},
	{"Backstage passes to a TAFKAL80ETC concert", "5>0 to sellIn #2",
		1, 30, 0, 33},
	{"Backstage passes to a TAFKAL80ETC concert", "on sellIn",
		0, 49, -1, 0},
	{"Backstage passes to a TAFKAL80ETC concert", "past sellIn",
		-13, 17, -14, 0},

	{"Conjured Mana Cake", "pre sellIn",
		8, 12, 7, 10},
	{"Conjured Mana Cake", "pre sellIn zero Quality",
		3, 0, 2, 0},
	{"Conjured Mana Cake", "on sellIn",
		0, 5, -1, 1},
	{"Conjured Mana Cake", "on sellIn zero Quality",
		0, 0, -1, 0},
	{"Conjured Mana Cake", "past sellIn",
		-4, 9, -5, 5},
	{"Conjured Mana Cake", "past sellIn zero Quality",
		-15, 0, -16, 0},
}

func TestGildedRose(t *testing.T) {
	for _, c := range cases {
		got := Item{
			Name:    c.name,
			SellIn:  c.sellIn,
			Quality: c.quality,
		}
		UpdateQuality(&got)

		if got.SellIn != c.wantSellIn {
			t.Errorf("Unexpected SellIn: %s %s: %d, expected SellIn value: %d", c.name, c.description, got.SellIn, c.wantSellIn)
		}
		if got.Quality != c.wantQuality {
			t.Errorf("Unexpected Quality: %s %s: %d, expected Quality value: %d", c.name, c.description, got.Quality, c.wantQuality)
		}
	}
}
