package staticpageservice

import (
	"metadatasvc/gen/go/staticpagepb"
	"shared/models/staticpagemodel"
)

var updateSetters = map[string]func(*staticpagemodel.StaticPage, *staticpagepb.StaticPage){
	"title": func(m *staticpagemodel.StaticPage, p *staticpagepb.StaticPage) {
		m.Title = p.Title
	},
	"slug": func(m *staticpagemodel.StaticPage, p *staticpagepb.StaticPage) {
		m.Slug = p.Slug
	},
}

func applyUpdates(pageModel *staticpagemodel.StaticPage, req *staticpagepb.UpdateRequest) *staticpagemodel.StaticPage {
	page := req.Page
	for _, f := range req.Fields {
		if setter, ok := updateSetters[f]; ok {
			setter(pageModel, page)
		}
	}
	return pageModel
}
