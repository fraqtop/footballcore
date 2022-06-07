package competition

type Competition interface {
	Id() int
	Title() string
}

func (c competition) Id() int {
	return c.id
}

func (c competition) Title() string {
	return c.title
}

type competition struct {
	id    int
	title string
}

func New(id int, title string) Competition {
	return &competition{
		id:    id,
		title: title,
	}
}
