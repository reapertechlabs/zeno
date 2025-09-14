package sitespecific

import (
	"net/http"

	"github.com/reapertechlabs/zeno/internal/pkg/preprocessor/sitespecific/npr"
	"github.com/reapertechlabs/zeno/internal/pkg/preprocessor/sitespecific/reddit"
	"github.com/reapertechlabs/zeno/internal/pkg/preprocessor/sitespecific/tiktok"
	"github.com/reapertechlabs/zeno/internal/pkg/preprocessor/sitespecific/truthsocial"
	"github.com/reapertechlabs/zeno/pkg/models"
)

type Preprocessor interface {
	Match(*models.URL) bool
	Apply(*http.Request)
}

var preprocessors = []Preprocessor{
	npr.NPRPreprocessor{},
	reddit.RedditPreprocessor{},
	tiktok.TikTokPreprocessor{},
	truthsocial.TruthsocialStatusPreprocessor{},
	truthsocial.TruthsocialAccountsPreprocessor{},
}

// Apply the first matching preprocessor.
func RunPreprocessors(URL *models.URL, req *http.Request) {
	for _, p := range preprocessors {
		if p.Match(URL) {
			p.Apply(req)
			break
		}
	}
}
