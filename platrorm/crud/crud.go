package crud

type Item struct {
	Name    string `json:name`
	Surname string `json:surname`
}

type AllItems struct {
	Items []Item
}

func New() *AllItems {
	return &AllItems{
		Items: []Item{},
	}
}

func (a *AllItems) Add(i Item) {
	a.Items = append(a.Items, i)
}

func (a *AllItems) List() []Item {
	return a.Items
}
