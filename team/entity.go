package team

type core struct {
	Id         int
	TitleFull  string
	TitleShort string
}

type Team struct {
	core
}

func (t Team) Id() int {
	return t.core.Id
}

func (t *Team) SetId(id int) {
	t.core.Id = id
}

func (t Team) TitleFull() string {
	return t.core.TitleFull
}

func (t Team) TitleShort() string {
	return t.core.TitleShort
}

func New(id int, titleFull, titleShort string) Team {
	return Team{
		core{
			Id:         id,
			TitleFull:  titleFull,
			TitleShort: titleShort,
		},
	}
}
