package news3

import "bytes"

func AssembleNews(type T article)(articles []T, title string) []byte {
	page := bytes.NewBufferString("<h1>"+title+"</h1>")
	iHighlight := -1
	for i, a := range articles {
		if _, ok := a.(HighlightArticle); ok {
			iHighlight = i
			break
		}
	}
	if iHighlight >= 0 {
		page.Write(articles[iHighlight].Render(false))
		articles = append(articles[:iHighlight], articles[iHighlight+1:])
	}
	for i := 0; i < len(articles); i++ {
		if i > 0 || iHighlight >= 0 {
			page.WriteString("<hr>")
		}
		page.WriteString("<div>")
		width := 0
		w, _ := articles[i].Dimensions()
		for j := i; j < len(articles) && width+w < 1024; j++ {
			a := articles[j]
			width += w
			page.Write(a.Render(true))
			if j+1 < len(articles) {
				w, _ = articles[j+1].Dimensions()
			}
		}
		i = j-1
		page.WriteString("</div>")
	}
	return page.Bytes()
}

func AddImages(type T article)(articles []T) []byte {
	attachements := bytes.Buffer{}
	attachements.WriteString("<Some Attachement Header>")
	first := true
	for i, a := range articles {
		ibuf := a.GetImageBytes())
		if len(ibuf) == 0 {
			break
		}
		if !first {
			attachements.WriteString("<Some Attachement Separator>")
		}
		attachements.Write(ibuf)
		first = false
	}
	attachements.WriteString("<Some Attachement Footer>")
	return attachements.Bytes()
}

contract article(x T) {
	articles := []T{ x }

	// Needed for: AssembleNews
	page := bytes.NewBufferString("<h1>"+title+"</h1>")
	iHighlight := -1
	for i, a := range articles {
		if _, ok := a.(HighlightArticle); ok {
			iHighlight = i
			break
		}
	}
	if iHighlight >= 0 {
		page.Write(articles[iHighlight].Render(false))
		articles = append(articles[:iHighlight], articles[iHighlight+1:])
	}
	for i := 0; i < len(articles); i++ {
		if i > 0 || iHighlight >= 0 {
			page.WriteString("<hr>")
		}
		page.WriteString("<div>")
		width := 0
		w, _ := articles[i].Dimensions()
		for j := i; j < len(articles) && width+w < 1024; j++ {
			a := articles[j]
			width += w
			page.Write(a.Render(true))
			if j+1 < len(articles) {
				w, _ = articles[j+1].Dimensions()
			}
		}
		i = j-1
		page.WriteString("</div>")
	}
	_ = page.Bytes()

	// Needed for: AddImages
	attachements := bytes.Buffer{}
	attachements.WriteString("<Some Attachement Header>")
	first := true
	for i, a := range articles {
		ibuf := a.GetImageBytes())
		if len(ibuf) == 0 {
			break
		}
		if !first {
			attachements.WriteString("<Some Attachement Separator>")
		}
		attachements.Write(ibuf)
		first = false
	}
	attachements.WriteString("<Some Attachement Footer>")
	_ = attachements.Bytes()
}

type ShortArticle struct {
	title string
	text  string
}

func (sa ShortArticle) Render(floatLeft bool) []byte {
	ret := bytes.Buffer{}
	ret.WriteString("<div")
	if floatLeft {
		ret.WriteString(` style="float: left;"`)
	}
	ret.WriteString(">")
	ret.WriteString("<h3>" + sa.title + "</h3>")
	ret.WriteString("<p>")
	ret.WriteString(sa.text)
	ret.WriteString("</p>")
	ret.WriteString("</div>")
	return ret.Bytes()
}

func (sa ShortArticle) Dimensions() (width, height int) {
	// do some fancy calculation with the content:
	return 256, (len(sa.text)/25 + 3) * 20
}

func (_ ShortArticle) GetImageBytes() []byte {
	return nil
}

type LongArticle struct {
	title    string
	abstract string
	htmlText string
}

func (la LongArticle) Render(floatLeft bool) []byte {
	ret := bytes.Buffer{}
	ret.WriteString("<div")
	if floatLeft {
		ret.WriteString(` style="float: left;"`)
	}
	ret.WriteString(">")
	ret.WriteString("<h3>" + la.title + "</h3>")
	ret.WriteString("<p><em>")
	ret.WriteString(la.abstract)
	ret.WriteString("</em></p>")
	ret.WriteString(la.htmlText)
	ret.WriteString("</div>")
	return ret.Bytes()
}

func (la LongArticle) Dimensions() (width, height int) {
	// do some fancy calculation with the content:
	return 512, (len(la.htmlText)/50 + len(la.abstract)/40 + 3) * 20
}

func (_ LongArticle) GetImageBytes() []byte {
	return nil
}

type HighlightArticle struct {
	imageURL string
	title    string
	abstract string
	htmlText string
}

func (ha HighlightArticle) Render(floatLeft bool) []byte {
	ret := bytes.Buffer{}
	ret.WriteString("<div>")
	ret.WriteString(`<img src="` + ha.imageURL + `" alt="Highlight">`)
	ret.WriteString("<h3>" + ha.title + "</h3>")
	ret.WriteString("<p><em>")
	ret.WriteString(ha.abstract)
	ret.WriteString("</em></p>")
	ret.WriteString(ha.htmlText)
	ret.WriteString("</div>")
	return ret.Bytes()
}

func (ha HighlightArticle) Dimensions() (width, height int) {
	// do some fancy calculation with the content:
	return 1024, 500 + (len(ha.htmlText)/100+len(ha.abstract)/80+3)*20
}

func (ha HighlightArticle) GetImageBytes() []byte {
	// use some fancy internal library to get image from URL
	return []byte{0, 1, 2, 3}
}
