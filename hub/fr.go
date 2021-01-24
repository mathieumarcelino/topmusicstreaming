package hubfr

import (
	collectorfr "topmusicstreaming/collector"
	"topmusicstreaming/sorter"
)

func Hub_FR() {
	allInfosSpotify := collectorfr.Spotify()
	allInfosAppleMusic := collectorfr.AppleMusic()
	allInfosDeezer := collectorfr.Deezer()

	sorter.Sorter(allInfosSpotify, allInfosAppleMusic, allInfosDeezer, "fr")
}