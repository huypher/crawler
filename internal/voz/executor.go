package voz

import (
	"context"
	"errors"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/huypher/crawler/internal/websocket"

	"github.com/huypher/crawler/internal/cache"

	"github.com/huypher/crawler/internal/pkg/container"

	"github.com/huypher/crawler/internal/pkg/utils"

	"golang.org/x/net/html"

	"github.com/gocolly/colly"
)

const (
	originalUrl = "https://voz.vn"

	cmtCacheKey      = "voz_processed_cache_key_%s"
	fieldLastPageIdx = "voz_processed_last_page_idx"
	fieldLastCmtIdx  = "voz_processed_last_cmt_idx"
	url              = "https://voz.vn/t/bong-da-viet-nam-nam-2021.206655/"
)

var (
	lastPateIndex = 0
	lasCmtIndex   = 0
)

type Executor interface {
	Do()
	GetVisiblePages(ctx context.Context, url string) (container.Map, []int)
	GetExpected(ctx context.Context, pages container.Map, pageIndex []int) (container.Map, []int, error)
	GetCmts(ctx context.Context, pages container.Map, pageIndex []int) (map[int][]Cmt, error)
}

type Broadcast interface {
	Broadcast(message *websocket.Message)
}

type executor struct {
	cache     cache.Cache
	broadcast Broadcast
}

func NewVozExecutor(cache cache.Cache, broadcast Broadcast) *executor {
	return &executor{
		cache:     cache,
		broadcast: broadcast,
	}
}

func (e *executor) Do() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "post_id", 206655)
	pages, pageIndex := e.GetVisiblePages(ctx, url)
	if len(pages) == 0 {
		return
	}

	expectedPages, expectedPageIndex, err := e.GetExpected(ctx, pages, pageIndex)
	if err != nil {
		return
	}

	cmts, err := e.GetCmts(ctx, expectedPages, expectedPageIndex)
	if err != nil {
		return
	}

	data := make([]string, 0)
	for _, cmt := range cmts {
		for _, ele := range cmt {
			data = append(data, ele.Content)
		}
	}

	fmt.Println("Sending to client...")
	for _, body := range data {
		e.broadcast.Broadcast(&websocket.Message{
			Type: "data",
			Body: body,
		})
		time.Sleep(2 * time.Second)
	}
}

func (e *executor) GetVisiblePages(ctx context.Context, url string) (container.Map, []int) {
	c := colly.NewCollector()

	selector := ".block.block--messages .block-outer.block-outer--after .block-outer-main .pageNav-main"
	pages := make(container.Map)
	pageIndex := make([]int, 0)
	c.OnHTML(selector, func(e *colly.HTMLElement) {
		e.ForEach("li", func(idx int, e1 *colly.HTMLElement) {
			pages[e1.Text] = e1.ChildAttr("a[href]", "href")
			if n, err := strconv.Atoi(e1.Text); err == nil {
				pageIndex = append(pageIndex, n)
			}
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visit page... %s\n", r.URL)
	})

	c.Visit(url)

	if len(pageIndex) == 0 {
		return container.Map{}, []int{}
	}

	sort.Ints(pageIndex)

	for idx, link := range pages {
		pages[idx] = fmt.Sprintf("%s%s", originalUrl, link)
	}

	return pages, pageIndex
}

type Cmt struct {
	Index    int
	MainAttr []html.Attribute
	Content  string
	Metadata container.Map
}

func (e *executor) GetCmts(ctx context.Context, pages container.Map, pageIndex []int) (map[int][]Cmt, error) {
	//var cacheKey string
	//val := ctx.Value("post_id")
	//if postID, ok := val.(string); ok {
	//	cacheKey = fmt.Sprintf(cmtCacheKey, postID)
	//}

	//cacheRes, err := e.cache.HMGetInt(ctx, cacheKey, fieldLastPageIdx, fieldLastCmtIdx)
	//if err != nil {
	//	return map[int][]Cmt{}, err
	//}

	newPages := make(container.Map)
	newPageIndex := make([]int, 0)
	if lastPateIndex != 0 && lasCmtIndex != 0 {
		newPageIndex = utils.GreaterOrEqualInt(pageIndex, lastPateIndex)
		newPages = pages.Include(utils.IntsToStrings(newPageIndex))
		sort.Ints(newPageIndex)
	} else {
		newPages = pages
		newPageIndex = pageIndex
	}

	c := colly.NewCollector()

	res := make(map[int][]Cmt)
	for _, idx := range newPageIndex {
		url, ok := newPages.GetString(strconv.Itoa(idx))
		if !ok {
			continue
		}

		cmts := make([]Cmt, 0)
		metadata := make(container.Map)
		c.OnHTML(".block-body.js-replyNewMessageContainer article.message.message--post.js-post.js-inlineModContainer", func(e *colly.HTMLElement) {
			htmlContent, err := e.DOM.Html()
			if err == nil {

				var index int
				re := regexp.MustCompile(`^#.*`)
				e.ForEachWithBreak("a[rel=nofollow]", func(i int, element *colly.HTMLElement) bool {
					text := strings.TrimSpace(element.Text)
					if re.MatchString(text) {
						index = func(index string) int {
							if n, err := strconv.Atoi(index); err == nil {
								return n
							}
							return 0
						}(utils.RemoveAllChars(text, []string{"#", ","}))
						return false
					}
					return true
				})
				if index <= lasCmtIndex {
					return
				}

				nodes := e.DOM.Nodes

				var mainAttr []html.Attribute
				if len(nodes) > 0 {
					mainAttr = nodes[len(nodes)-1].Attr
				}

				metadata = container.Map{}
				metadata.AppendSliceString("img", utils.RemoveDupString(e.ChildAttrs("img", "src")))
				metadata.AppendSliceString("src_set", utils.RemoveDupString(e.ChildAttrs("img", "srcset")))
				metadata.AppendSliceString("a_link", utils.RemoveDupString(e.ChildAttrs("a", "href")))

				fullContent := completeCmt(htmlContent, mainAttr, metadata)

				cmts = append(cmts, Cmt{
					Index:    index,
					MainAttr: mainAttr,
					Content:  fullContent,
					Metadata: metadata,
				})
			} else {
				log.Printf("err=%v", err)
			}
		})

		c.OnRequest(func(r *colly.Request) {
			fmt.Printf("Visit page... %s\n", r.URL)
		})

		c.Visit(url)

		if len(cmts) > 0 {
			res[idx] = cmts

			lastPateIndex = newPageIndex[len(newPageIndex)-1]
			lasCmtIndex = cmts[len(cmts)-1].Index
			//_, err := e.cache.HMSet(ctx, cacheKey, fieldLastPageIdx, strconv.Itoa(lastPateIndex), fieldLastCmtIdx, strconv.Itoa(lasCmtIndex))
			//if err != nil {
			//	log.Printf("cache err=%v", err)
			//}
			//e.cache.Expire(ctx, cacheKey, 10*time.Minute)
		}
	}

	if len(res) > 0 {
		return res, nil
	}

	return map[int][]Cmt{}, errors.New("data is nil")
}

func (e *executor) GetExpected(ctx context.Context, pages container.Map, pageIndex []int) (container.Map, []int, error) {
	lastPageIndex := pageIndex[len(pageIndex)-1]
	if lastPageIndex == 1 {
		return pages.Include([]string{strconv.Itoa(lastPageIndex)}), []int{lastPageIndex}, nil
	}

	lastPage, ok := pages.GetString(strconv.Itoa(lastPageIndex))
	if !ok {
		return container.Map{}, []int{}, nil
	}

	p := strings.Replace(lastPage, fmt.Sprintf("page-%s", strconv.Itoa(lastPageIndex)), fmt.Sprintf("page-%s", strconv.Itoa(lastPageIndex-1)), 1)

	pages.Add(strconv.Itoa(lastPageIndex-1), p)

	return pages.Include([]string{strconv.Itoa(lastPageIndex - 1), strconv.Itoa(lastPageIndex)}), []int{lastPageIndex - 1, lastPageIndex}, nil
}

func completeCmt(cmt string, attrs []html.Attribute, partialLinks container.Map) string {
	s1 := ""

	for _, a := range attrs {
		s1 += a.Key + `="` + a.Val + `" `
	}

	return `<article ` + s1 + `>` + makeFullLink(cmt, partialLinks) + `</article>`
}

func makeFullLink(cmt string, metadata container.Map) string {
	re := regexp.MustCompile(`^https:\/\/.*`)

	f := func(p container.Map, key string) []string {
		if data, ok := p.GetSliceString(key); ok {
			return data
		}
		return []string{}
	}

	srcSets := utils.RemoveDupString(f(metadata, "src_set"))

	remainMap := metadata.Exclude([]string{"src_set"})

	for _, l := range srcSets {
		splitted := strings.Split(l, ",")
		if len(splitted) == 1 {
			cmt = strings.Replace(cmt, l, originalUrl+l, -1)
		}

		temp := make([]string, len(splitted))
		for idx, s := range splitted {
			if idx == 0 {
				temp[idx] = s
				continue
			}
			temp[idx] = originalUrl + strings.TrimSpace(s)
		}
		cmt = strings.Replace(cmt, l, strings.Join(temp, ","), -1)
	}

	for k := range remainMap {
		links := utils.RemoveDupString(f(remainMap, k))
		for _, l := range links {
			if re.MatchString(l) {
				continue
			}
			cmt = strings.Replace(cmt, l, originalUrl+l, -1)
		}
	}

	return cmt
}
