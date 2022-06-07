package team

type Team interface {
	Id() int
	SetId(id int)
	TitleFull() string
	TitleShort() string
}

type team struct {
	id         int
	fullTitle  string
	shortTitle string
}

func (t team) Id() int {
	return t.id
}

func (t *team) SetId(id int) {
	t.id = id
}

func (t team) TitleFull() string {
	return t.fullTitle
}

func (t team) TitleShort() string {
	return t.shortTitle
}

func New(id int, fullTitle, shortTitle string) Team {
	return &team{
		id:         id,
		shortTitle: shortTitle,
		fullTitle:  fullTitle,
	}
}
