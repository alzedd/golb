package fsutils

import "github.com/alzedd/golb/settings"

type settingsGetter interface {
	Get(name string) (v string)
}

var settingsgetter settingsGetter

func init() {
	settingsgetter = settings.Instance()
}
