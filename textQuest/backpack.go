package main


type Backpack struct {
	Items map[string]*Item
	worn  bool
}

func (b *Backpack) AddItem(name string, item *Item) {
	b.Items[name] = item
}

func (b *Backpack) HasItem(name string) bool {
	_, exists := b.Items[name]
	return exists
}

func (b *Backpack) HasItems() bool {
	return len(b.Items) > 0
}

func (b *Backpack) Wear() {
	b.worn = true
}

func (b *Backpack) IsWorn() bool {
	return b.worn
}
