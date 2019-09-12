package commands

type settingsGetter interface {
	Get(name string) (v string)
}
