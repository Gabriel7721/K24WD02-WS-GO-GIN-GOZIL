package room

type Controller struct {
	Repo *Repository
}

func NewController(r *Repository) *Controller {
	return &Controller{Repo: r}
}
