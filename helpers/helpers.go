package helpers

import (
	"errors"
	"strings"
)

// GetAlbumIDFromURL gets the album ID from a spotify url
func GetAlbumIDFromURL(URL string) (string, error) {
	splitURL := strings.Split(URL, "/")

	indexBeforeID, err := indexOfStringElement("album", splitURL)

	if err != nil {
		return "", errors.New("Album text not found in URL")
	}

	// Get album ID by referring to the index after "album" text in url
	// and remove any query strings attached
	albumID := strings.Split(splitURL[indexBeforeID+1], "?")[0]

	return albumID, nil
}

func indexOfStringElement(element string, data []string) (int, error) {
	for i, v := range data {
		if strings.Compare(element, v) == 0 {
			return i, nil
		}
	}

	return -1, errors.New("Element not found")
}
