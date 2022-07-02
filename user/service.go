package user

type userService struct {
	repo Repository
}

type Service interface {
	GetTotalUser() (total int64, err error)
	Create(user User) (err error)
}

func NewUserService(repo Repository) Service {
	return &userService{
		repo: repo,
	}
}

func (u userService) GetTotalUser() (total int64, err error) {
	return u.repo.GetTotalUser()
}

func (u userService) Create(user User) (err error) {
	return u.repo.Create(user)
}
