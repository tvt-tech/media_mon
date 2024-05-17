//go:build linux && !mipsle
// +build linux,!mipsle

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
	as_service := flag.Bool("s", false, "Start service")
	quiet := flag.Bool("q", false, "Quiet")
	eject := flag.Bool("e", false, "Eject device")
	eject_all := flag.Bool("A", false, "Eject all matched Archer devices")
	flag.Parse()
	if *debug {
		Logger.SetLevel(logrus.DebugLevel)
		Logger.Debug("Debug mode")
	}

	if *as_service {
		if !*quiet {
			go utils.RunTray()
		}
		utils.StartService()
		return
	}

	if *list {
		utils.PrintDrives()
		return
	}

	if *eject {
		if directory := flag.Arg(0); directory != "" {
			err := utils.EjectDrive(directory)
			if err != nil {
				Logger.Error(err)
			}
		} else if *eject_all {
			if drives, err := utils.Detect(); err == nil {
				for _, drive := range drives {
					if utils.FindMatchedArcherDevice(drive) {
						utils.EjectDrive(drive)
					}
				}
			}
		} else {
			Logger.Warn("No eject options specified")
		}
		return
	}

	if directory := flag.Arg(0); directory != "" {
		info, err := os.Stat(directory)
		if err == nil {
			if info.IsDir() {
				if err := utils.ArchiveNotAllowed(directory); err != nil {
					Logger.Error(err)
				}
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
				if err := utils.ArchiveNotAllowed(drive); err != nil {
					Logger.Error(err)
				}
			}
		}
	}
}
