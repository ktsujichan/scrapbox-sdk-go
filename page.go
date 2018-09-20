package scrapbox

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

type Page struct {
	ID            string          `json:"id"`
	Title         string          `json:"title"`
	TitleLc       *string         `json:"titleLc,omitempty"`
	Image         string          `json:"image"`
	Descriptions  []string        `json:"descriptions"`
	User          interface{}     `json:"user,omitempty"`
	Pin           uint            `json:"pin,omitempty"`
	Views         uint            `json:"views,omitempty"`
	Point         uint            `json:"point,omitempty"`
	Linked        uint            `json:"linked,omitempty"`
	CommitID      string          `json:"commitId,omitempty"`
	Created       uint            `json:"created,omitempty"`
	Updated       uint            `json:"updated"`
	Accessed      uint            `json:"accessed"`
	Persistent    bool            `json:"persistent,omitempty"`
	Lines         []Line          `json:"lines,omitempty"`
	Links         []string        `json:"links,omitempty"`
	LinksLc       *[]string       `json:"linksLc,omitempty"`
	Icons         map[string]uint `json:"icons,omitempty"`
	RelatedPages  *RelatedPages   `json:"relatedPages,omitempty"`
	Collaborators []User          `json:"collaborators,omitempty"`
	LastAccessed  interface{}     `json:"lastAccessed,omitempty"`
}

type Pages struct {
	Skip  uint   `json:"skip"`
	Limit uint   `json:"limit"`
	Count uint   `json:"count"`
	Pages []Page `json:"pages"`
}

type RelatedPages struct {
	Links1Hop []Page        `json:"links1hop"`
	Links2Hop []Page        `json:"links2hop"`
	Icons1Hop []interface{} `json:"icons1hop"`
}

type ListPagesOptions struct {
	Skip  uint `url:"skip,omitempty"`
	Limit uint `url:"limit,omitempty"`
}

func (c *Client) ListPages(ctx context.Context, projectName string, opt *ListPagesOptions) (*Pages, error) {
	res, err := c.get(ctx, fmt.Sprintf("/pages/%s", projectName), opt)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var pages Pages
	if err := decodeBody(res, &pages); err != nil {
		return nil, errors.WithStack(err)
	}
	return &pages, nil
}

func (c *Client) GetPage(ctx context.Context, projectName, pageTitle string) (*Page, error) {
	res, err := c.get(ctx, fmt.Sprintf("/pages/%s/%s", projectName, pageTitle), nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var page Page
	if err := decodeBody(res, &page); err != nil {
		return nil, errors.WithStack(err)
	}
	return &page, nil
}
