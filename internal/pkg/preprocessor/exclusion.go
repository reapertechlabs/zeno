package preprocessor

import (
	"regexp"

	"github.com/reapertechlabs/zeno/pkg/models"
)

func matchRegexExclusion(ExclusionRegexes []*regexp.Regexp, item *models.Item) bool {
	for _, exclusion := range ExclusionRegexes {
		if exclusion.MatchString(item.GetURL().String()) {
			return true
		}
	}

	return false
}
