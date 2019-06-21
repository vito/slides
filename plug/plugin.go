package main

import (
	_ "github.com/concourse/docs/go/chromastyle"

	"github.com/vito/booklit"
	"github.com/vito/booklit/chroma"
)

func init() {
	booklit.RegisterPlugin("slides", NewPlugin)
}

type Plugin struct {
	section *booklit.Section
	chroma  chroma.Plugin
}

func NewPlugin(section *booklit.Section) booklit.Plugin {
	return &Plugin{
		section: section,
		chroma:  chroma.NewPlugin(section).(chroma.Plugin),
	}
}

func (p Plugin) TitleSlide(title booklit.Content, subtitle ...booklit.Content) booklit.Content {
	var style booklit.Style = "title-slide"
	partials := booklit.Partials{}

	if len(subtitle) != 0 {
		style = "subtitle-slide"
		partials["Subtitle"] = subtitle[0]
	}

	return booklit.Styled{
		Style:    style,
		Content:  title,
		Block:    true,
		Partials: partials,
	}
}

func (p Plugin) Slide(title, body booklit.Content) booklit.Content {
	return booklit.Styled{
		Style:   "slide",
		Content: body,
		Block:   true,
		Partials: booklit.Partials{
			"Title": title,
		},
	}
}

func (p Plugin) Detail(title, body booklit.Content) booklit.Content {
	return booklit.Styled{
		Style:   "detail-slide",
		Content: body,
		Block:   true,
		Partials: booklit.Partials{
			"Title": title,
		},
	}
}

func (p Plugin) Codeblock(language string, code booklit.Content) (booklit.Content, error) {
	return p.chroma.Syntax(language, code, "concourseci")
}

func (p Plugin) TitledCodeblock(title booklit.Content, language string, code booklit.Content) (booklit.Content, error) {
	codeblock, err := p.Codeblock(language, code)
	if err != nil {
		return nil, err
	}

	return booklit.Styled{
		Style: "titled-codeblock",
		Block: true,

		Content: codeblock,

		Partials: booklit.Partials{
			"Title": booklit.Styled{
				Style:   booklit.StyleVerbatim,
				Content: title,
			},
		},
	}, nil
}

func (p Plugin) Youtube(id booklit.Content) booklit.Content {
	return booklit.Styled{
		Style:   "youtube-embed",
		Block:   true,
		Content: id,
	}
}
