package counter

type Repository interface {
	Increment(key string) (int64, error)
	Get(key string) (int64, error)
}
