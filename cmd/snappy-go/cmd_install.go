package main

import (
	"os"

	"launchpad.net/snappy/snappy"
)

type cmdInstall struct {
}

func init() {
	var cmdInstallData cmdInstall
	_, _ = parser.AddCommand("install",
		"Install a snap package",
		"Install a snap package",
		&cmdInstallData)
}

func (x *cmdInstall) Execute(args []string) (err error) {
	if !isRoot() {
		return ErrRequiresRoot
	}

	err = snappy.Install(args)
	if err != nil {
		return err
	}
	// call show versions afterwards
	installed, err := snappy.ListInstalled()
	if err != nil {
		return err
	}

	showInstalledList(installed, os.Stdout)

	return nil
}
