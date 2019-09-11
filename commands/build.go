package commands

import (
	"golb/saving"
	"golb/taxonomy"
)

func Build() {
	pr := taxonomy.NewPageRepository()
	savingService := saving.NewService(pr)
	savingService.SaveAll()
	savingService.SaveAssets()
}
