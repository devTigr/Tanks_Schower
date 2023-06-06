package controller

import (
	"ProjectZero/models"
	"ProjectZero/view"
	"log"
)

// Run does the running of the models application
func Run(enablePersistence bool) {
	if enablePersistence {
		models.EnableFilePersistence()
	} else {
		models.DisableFilePersistence()
	}

	err := models.Initialize()
	checkAndHandleErrorWithTermination(err)

	for true {
		executeCommand()
	}
}

func checkAndHandleErrorWithTermination(err error) {
	if err != nil {
		view.PrintError(err)
		log.Fatal(err)
	}
}

func checkAndHandleErrorWithoutTermination(err error) {
	if err != nil {
		view.PrintMessage("The following error occurred:")
		view.PrintError(err)
		view.PrintMessage("Press c to continue!")
	}
}

func executeCommand() {
	view.RunScreen()
	//view.Clear()
	//view.UpdateInfo("nei")
	//models.GetNewInput()
	//view.GlobalLable = "nei2"
	//view.CheckQuitCommand()
	executeCommand()
}
