package competition

type core struct {
	Id    int
	Title string
}

type Competition struct {
	core
}

func (c Competition) Id() int {
	return c.core.Id
}

func (c Competition) Title() string {
	return c.core.Title
}

func New(id int, title string) Competition {
	return Competition{
		core{
			Id: id,
			Title: title,
		},
	}
}
