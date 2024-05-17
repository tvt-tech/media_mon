package main

import (
	"flag"
	"os"

	"github.com/tvt-tech/usb-file-filter/utils"

	"github.com/sirupsen/logrus"
)

var Logger = utils.Logger // Assign utils.Logger to Log

func init() {

}

func main() {
	debug := flag.Bool("d", false, "Debug")
	list := flag.Bool("l", false, "List devices")
	flag.Parse()

	if *debug {
		Logger.SetLevel(logrus.DebugLevel)
		Logger.Debug("Debug mode")
	}

	if *list {
		utils.PrintDrives()
		return
	}

	if directory := flag.Arg(0); directory != "" {
		info, err := os.Stat(directory)
		if err == nil {
			if info.IsDir() {
				utils.RenameNotAllowed(directory)
			}
		} else {
			Logger.Error("Wrong path")
		}
	} else {
		if Logger.GetLevel() == logrus.DebugLevel {
			utils.PrintDrives()
		}
		if drives, err := utils.Detect(); err == nil {
			for _, drive := range drives {
				utils.RenameNotAllowed(drive)
			}
		}
	}
}
