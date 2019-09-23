package commands

import (
	"github.com/alzedd/golb/saving"
	"github.com/alzedd/golb/taxonomy"
)

func Build() {
	pr := taxonomy.NewPageRepository()
	savingService := saving.NewService(pr)
	savingService.SaveAll()
	savingService.SaveAssets()
}
