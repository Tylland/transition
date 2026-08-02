package specs

import (
	"fmt"
	"strings"
	"testing"
)

type UserStory struct {
	AsA     string
	IWantTo string
	SoThat  string
}

type Story struct {
	UserStory UserStory
	tags      map[string]string
	reported  bool
}

func NewStory(story UserStory) *Story {
	return &Story{UserStory: story, tags: make(map[string]string)}
}

func (f *Story) Scenario(t *testing.T) *Scenario {
	return NewScenario(f, t)
}

func (f *Story) isVowel(str string) bool {
	runes := []rune(strings.ToLower(str))

	return len(runes) > 0 && (runes[0] == 'a' || runes[0] == 'e' || runes[0] == 'i' || runes[0] == 'o' || runes[0] == 'u' || runes[0] == 'å' || runes[0] == 'ä' || runes[0] == 'ö')
}

func (f *Story) AsAPrefix() string {
	if f.isVowel(f.UserStory.AsA) {
		return "an"
	} else {
		return "a"
	}
}

func (f *Story) String() string {
	return fmt.Sprintf("As %s %s\nI want to %s\nSo that %s\n", f.AsAPrefix(), f.UserStory.AsA, f.UserStory.IWantTo, f.UserStory.SoThat)
}

func (f *Story) WithTags(tags ...string) *Story {
	for i := 0; i < len(tags); i += 2 {
		if len(tags)-i > 1 {
			f.tags[tags[i]] = tags[i+1]
		}
	}

	return f
}
