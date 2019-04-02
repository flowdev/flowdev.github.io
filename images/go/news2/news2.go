package news2

import "bytes"

func AssembleNews(type T article)(articles []T) []byte {
	page := bytes.NewBufferString("<h1>News</h1>")
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

contract article(x T) {
	articles := []T{ x }
	page := bytes.NewBufferString("<h1>News</h1>")
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
	ret.WriteString("<p>")
	ret.WriteString(la.htmlText)
	ret.WriteString("</p>")
	ret.WriteString("</div>")
	return ret.Bytes()
}

func (la LongArticle) Dimensions() (width, height int) {
	// do some fancy calculation with the content:
	return 512, (len(la.htmlText)/50 + len(la.abstract)/40 + 3) * 20
}
