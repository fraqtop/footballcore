package competition

type core struct {
	Id    int
	Title string
}

type Competition struct {
	core
}

func New(id int, title string) Competition {
	return Competition{
		core{
			Id:    id,
			Title: title,
		},
	}
}

func (this Competition) Id() int {
	return this.core.Id
}

func (this Competition) Title() string {
	return this.core.Title
}

func (this Competition) IsValid() bool {
	return len(this.core.Title) > 0
}
