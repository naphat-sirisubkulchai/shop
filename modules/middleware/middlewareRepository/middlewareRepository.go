package middlewarerepository

type(
	MiddlewareRepositoryService interface{}

	middlewareRepository struct{}
)

func NewMiddlewareRepository() MiddlewareRepositoryService {
	return &middlewareRepository{}
}