package truthsocial

import (
	"regexp"

	"github.com/reapertechlabs/zeno/internal/pkg/postprocessor/extractor"
	"github.com/reapertechlabs/zeno/pkg/models"
)

var (
	postURLRegex       = regexp.MustCompile(`^https?:\/\/truthsocial\.com\/@[A-Za-z0-9_]+\/posts\/`)
	postIDRegex        = regexp.MustCompile(`^https?:\/\/truthsocial\.com\/@[A-Za-z0-9_]+\/posts\/(\d+)`)
	usernameRegex      = regexp.MustCompile(`^https?:\/\/truthsocial\.com\/@([^/]+)`)
	statusesRegex      = regexp.MustCompile(`^https?:\/\/truthsocial\.com\/api\/v1\/statuses\/\d+$`)
	accountLookupRegex = regexp.MustCompile(`^https?:\/\/truthsocial\.com\/api\/v1\/accounts\/lookup\?acct=[a-zA-Z0-9]+$`)
)

func NeedExtraction(URL *models.URL) bool {
	return IsStatusesURL(URL) || IsPostURL(URL)
}

func ExtractAssets(item *models.Item) (assets, outlinks []*models.URL, err error) {
	if IsStatusesURL(item.GetURL()) {
		truthsocialAssets, err := GenerateVideoURLsFromStatusesAPI(item.GetURL())
		if err != nil {
			return assets, outlinks, err
		}

		JSONAssets, outlinks, err := extractor.JSON(item.GetURL())
		if err != nil {
			return assets, outlinks, err
		}

		assets = append(truthsocialAssets, JSONAssets...)
	} else if IsPostURL(item.GetURL()) {
		truthsocialAssets, err := GeneratePostAssetsURLs(item.GetURL())
		if err != nil {
			return assets, outlinks, err
		}

		HTMLAssets, err := extractor.HTMLAssets(item)
		if err != nil {
			return assets, outlinks, err
		}

		assets = append(truthsocialAssets, HTMLAssets...)
	}

	return assets, outlinks, err
}
