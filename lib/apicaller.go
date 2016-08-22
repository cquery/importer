package lib

type APICaller interface {
	Call() (Resources, error)
}
