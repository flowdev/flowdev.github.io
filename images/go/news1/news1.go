package news1

import "bytes"

func AssembleNews(type T article)(articles []T) []byte {
	page := bytes.NewBufferString("<h1>News</h1>")
	for i, a := range articles {
		if i > 0 {
			page.WriteString("<hr>")
		}
		page.Write(a.Render())
	}
	return page.Bytes()
}

contract article(x T) {
	articles := []T{ x }
	page := bytes.NewBufferString("<h1>News</h1>")
	for i, a := range articles {
		if i > 0 {
			page.WriteString("<hr>")
		}
		page.Write(a.Render())
	}
	_ = page.Bytes()
}

type ShortArticle struct {
	title string
	text  string
}

func (sa *ShortArticle) Render() []byte {
	ret := bytes.Buffer{}
	ret.WriteString("<h3>"+sa.title+"</h3>")
	ret.WriteString("<p>")
	ret.WriteString(sa.text)
	ret.WriteString("</p>")
	return ret.Bytes()
}

type LongArticle struct {
	title    string
	abstract string
	htmlText string
}

func (la *LongArticle) Render() []byte {
	ret := bytes.Buffer{}
	ret.WriteString("<h3>"+la.title+"</h3>")
	ret.WriteString("<p><em>")
	ret.WriteString(la.abstract)
	ret.WriteString("</em></p>")
	ret.WriteString(la.htmlText)
	return ret.Bytes()
}
