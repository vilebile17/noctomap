package main

import (
	"encoding/json"
	"os"
)

type Theme struct {
	WallpaperPath      string `json:"wallpaper_path"`
	ColourSchemeSource string `json:"colour_scheme_source"`
	ColourSchemeName   string `json:"colour_scheme_name"`
}

func getConfiguration() ([]Theme, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return []Theme{}, err
	}

	data, err := os.ReadFile(homedir + "/.noctomap.json")
	if err != nil {
		return []Theme{}, err
	}

	var themes struct {
		Themes []Theme `json:"themes"`
	}
	if err = json.Unmarshal(data, &themes); err != nil {
		return []Theme{}, err
	}

	return themes.Themes, nil
}
