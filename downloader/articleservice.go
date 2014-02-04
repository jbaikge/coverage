package downloader

import (
	"errors"
	"fmt"
	"github.com/300brand/coverage"
	"github.com/300brand/coverage/service"
	"github.com/300brand/logger"
	"net/url"
	"regexp"
	"strconv"
	"time"
)

type ArticleService struct {
}

var _ service.ArticleService = ArticleService{}

var (
	MaxMetaRefreshDelay = 60
	errMetaNotFound     = errors.New("No meta-refresh tag found")
	reMetaRefresh       = regexp.MustCompile(`(?i)<meta[^>]+http-equiv=["']?refresh["']?[^>]+>`)
	reMetaContent       = regexp.MustCompile(`(?i)content=['"](\d+)\s*;\s*([^'" ]+)['"]`)
)

func NewArticleService() ArticleService {
	return ArticleService{}
}

func (s ArticleService) Update(a *coverage.Article) error {
	return Article(a)
}

func Article(a *coverage.Article) error {
	a.Log.Service("downloader.ArticleService")
	r, err := Fetch(a.URL.String())
	if err != nil {
		return a.Log.Error(err)
	}
	a.LastCheck = time.Now()

	switch redirURL, err := metaRedirect(r.Body, a.URL); err {
	case errMetaNotFound:
	// Continue
	case nil:
		a.Log.Debug("Meta-redirected URL from [%s] to [%s]", a.URL.String(), redirURL)
		*a.URL = *redirURL
		return Article(a)
	default:
		logger.Warn.Printf("[P:%s] [F:%s] [A:%s] %s", a.PublicationId.Hex(), a.FeedId.Hex(), a.ID.Hex(), err)
	}

	a.Text.HTML = r.Body
	if a.URL.String() != r.RealURL {
		a.Log.Debug("Updating URL from [%s] to [%s]", a.URL.String(), r.RealURL)
		if a.URL, err = url.Parse(r.RealURL); err != nil {
			return a.Log.Error(err)
		}
	}
	a.Modified("HTML")
	return nil
}

func metaRedirect(body []byte, pageUrl *url.URL) (redirect *url.URL, err error) {
	tag := reMetaRefresh.Find(body)
	if tag == nil {
		err = errMetaNotFound
		return
	}

	content := reMetaContent.FindSubmatch(tag)
	if len(content) == 0 {
		err = fmt.Errorf("Improperly formatted meta refresh tag: %s", tag)
		return
	}

	delay, err := strconv.Atoi(string(content[1]))
	if err != nil {
		return
	}
	if delay > MaxMetaRefreshDelay {
		err = fmt.Errorf("Meta-refresh delay too large: %d > %d", delay, MaxMetaRefreshDelay)
		return
	}

	refUrl, err := url.Parse(string(content[2]))
	if err != nil {
		return
	}

	redirect = pageUrl.ResolveReference(refUrl)
	return
}
