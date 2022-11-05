package slackemoji

import (
	"fmt"
	"strings"
	"sync"
)

var (
	cacheLock sync.Mutex
	cache     = map[string]map[string]string{}
)

func FillTeamCache(teamID string, emoji map[string]string) {
	cacheLock.Lock()
	defer cacheLock.Unlock()

	cache[teamID] = emoji
}

func CustomURL(teamID, colonSyntax string, fillCache func() (map[string]string, error)) (string, error) {
	emojiName := strings.Trim(colonSyntax, ":")

	cacheLock.Lock()
	defer cacheLock.Unlock()

	visited := map[string]bool{}
	for {
		if visited[emojiName] || emojiName == "" {
			// circular alias loop
			return "", nil
		}

		teamCache, ok := cache[teamID]
		if !ok {
			// team not in cache
			var err error
			teamCache, err = fillCache()
			if err != nil {
				return "", fmt.Errorf("unable to fill cache: %w", err)
			}

			cache[teamID] = teamCache
		}

		url, ok := teamCache[emojiName]
		if !ok {
			// emoji not in cache
			return "", nil
		}

		if !strings.HasPrefix(url, "alias:") {
			return url, nil
		}

		visited[emojiName] = true
		emojiName = strings.TrimPrefix(url, "alias:")
	}
}
