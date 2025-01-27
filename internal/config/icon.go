package config

import (
	"errors"
	"os"
	"path"

	"github.com/adrg/xdg"
)

func RegisterIcon(icon string) (string, error) {
	iconsFolder := xdg.Home + "/.local/share/icons/hicolor/48x48/apps/"
	iconPath := path.Join(iconsFolder, "bills.svg")

	_, err := os.Stat(iconsFolder)

	if errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(iconsFolder, os.ModePerm)
	}

	if err != nil {
		return "", err
	}

	file, err := os.Create(iconPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.WriteString(icon)
	if err != nil {
		return "", err
	}

	return iconPath, nil
}
