package team

type core struct {
	Id         int
	TitleFull  string
	TitleShort string
}

type Team struct {
	core
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

func (this Team) Id() int {
	return this.core.Id
}

func (this *Team) SetId(id int) {
	this.core.Id = id
}

func (this Team) TitleFull() string {
	return this.core.TitleFull
}

func (this Team) TitleShort() string {
	return this.core.TitleShort
}

func (this Team) IsValid() bool {
	return len(this.core.TitleShort) > 0 && len(this.core.TitleFull) > 0
}

