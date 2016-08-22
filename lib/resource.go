package lib

type Resources interface {
	Update(updater Updater) error
}
