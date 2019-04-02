---
layout: post
---

TL;DR

* Generics is difficult to get right.
* The officially proposed contracts are very flexible and can be used in many
  ways.
* This can lead to very ugly copy/paste programming.
* This is especially true since now are more and more pragmatic programmers
  becoming Gophers.
* Generics for containers can be achieved with interfaces, too.
* Generics for (un-)marshaling and templating seem to be more difficult.

## Disclaimer

* I am writing this as a private person and Gopher and neither as an employee
  of the company I am currently working for nor as co-organizer of the GDG
  Berlin Golang meetup.
* I am not saying that generics themself are a bad idea.
* The presented example is a bit silly and a worst case scenario for its size.
  But I have seen very similar things happen in real live in the Java world.

## What Is Generics Good For?

Suppose you have to write and maintain a high performance container data
structure (e.g. some kind of tree) for all kinds of integer numbers (uint8,
uint16, uint32, uint64, int8, int16, int32 and int64).

Since the performance matters you can't use `interface{}`. And maintaining 8
copies of essentially the same code that only differ in the type annotations
quickly becomes a nightmare.

Thus Go currently has got no clean way to handle this nicely.
The best approach usually is to generate the code for all 8 types from a single
source.

## What Are Go Contracts?

The simplest example given at the
[official proposal site](https://go.googlesource.com/proposal/+/master/design/go2draft-contracts.md)
is this:
```go
func Stringify(type T stringer)(s []T) (ret []string) {
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return ret
}

contract stringer(x T) {
	var s string = x.String()
}
```

The `stringer` contract ensures that every instance of that type has got a
no-arguments method called `String` that returns a `string`.

All that contracts are good for is ensuring properties of types.
In this particular case it could (and should) be done simpler with the
[Stringer](https://golang.org/pkg/fmt/#Stringer) interface.

But contracts are much more powerful because they can be used to ensure any
property of a type that can be expressed by a usage in Go code.
This allows for example to ensure the existance of operators as is needed for
almost all kinds of data containers.
The two things that are not allowed in contracts are:
* Contracts may not have normal parameters or return parameters. It follows
  that they may not use `return` statements.
* Contracts may not refer to any name defined in the current package. This
  restriction exists to make it harder to accidentally change the meaning of a
  contract.

So Go contracts are good enough for the generics example above whereas
interfaces are not.

## Adding A Significant Feature Like Generics To Go Changes The Language For Good

![Rune Stone](/images/runeStone.jpg)

This is a very fundamental and long term decision!
Python is still strugling to get version 3 fully adopted and it started more
than 10 years ago (2008).
Go 'contracts' is still the official proposal.
All [feedback](https://github.com/golang/go/wiki/Go2GenericsFeedback) so far
didn't change that.
Once generics is widely used it will be impossible to remove it without
effectively creating a language fork with new library ecosystem and community.
Binary compatibility might mitigate that a bit but not fully and it is a
crutch.

## The Go Community Changes

There are lot's of new Gophers coming to the community.
The last couple of years the Go community has been roughly doubling every year.
The Go community is now well beyond just innovators and early adopters but well
into the early majority.

![adoption curve](/images/adoption-curve.jpg)

As great as this is it is changing the culture of the Go community, too.
Please don't get me wrong, these are nice people and a valueable addition to
the Go community. It is a sign that the Go community is becoming more mature.
But they are often less interested in the finer technical details of the tools
they use.  They are pragmatic and use a programming language to get a job done.
So they use the programming language in simple and pragmatic ways.  Often
without caring about the technical details.

I generally like that a lot and pragmatic programmers can very well be better
programmers. They often know the problem domain very well and it is better to
solve the right problem in a simple, straight forward way than elegantly
solving the wrong problem.  Or even solving the (right or wrong) problem in an
overengineered way.

## Let's Solve A Small Problem

Now I would like to demonstrate how Go contracts will be used in such a setting.
I understand that this specific example is *not realistic*. But I have seen
similar things happen already in the Java world.
So I speak from experiance and don't just express my nightmares here.

This example goes through 3 stages with problem definition, main function
implementation and contract implementation in each stage.

### Problem Definition For First Stage

* Articles should be assembled into a news page.
* Short articles with `title` and `text`.
* Long articles with additional `abstract`.
* All articles know how to render themself.

### Main Method Implementation For First Stage

This is a simple implementation of the `AssembleNews` function:
```go
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
```
The first line shows the expected signature
`func AssembleNews(type T article)(articles []T) []byte`.
The article is really used in Line 11: `page.Write(a.Render())`
So the `Render` method is the interesting part for the contract.

### Go Contract Implementation For First Stage

For the contract implementation we need just an example usage of the type.
As a pragmatic programmer we would like to use the most successful and often
used pattern in programming: **copy & paste**
Usually we would google such a usage and copy it from the internet (stack overflow).
But as this is an internal type no usage available ... 
except the one we just wrote!

You might cringe now but
I have seen very similar lines of thinking happen in real live in the Java world.

So consequently the contract looks like this:
```go
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
```
So the contract has got the expected signature. As we aren't allowed to have
normal parameters but need a slice of articles as input we just create one
with: `articles := []T{ x }`
Finally we have to be careful not to use `return` statements. but we can
replace them with simple assignments:
`return page.Bytes()` becomes `_ = page.Bytes()`

So these small formal changes are all we need to turn our usage ot the type
into a contract.

So let's evaluate this solution in the context of being a pragmatic person.
The upsides are:
1. I don't have to understand what a contract really is or how it works.
1. I just have to follow a small set of simple and formal rules.
1. I exactly know that my real usage is supported.

The only downside is readability. But not for me since I know about `Render`.
I only feel the immediate upsides and don't feel the long term downside at all.
The first upside can't be valued high enough because there are many people who
simply don't understand abstract concepts. They just remember and reproduce
some important use cases.

The full source code for this stage (all in one file) can be found in
[news1.go](/images/go/news1/news1.go).

### Problem Definition For Second Stage

Our solution is accepted by the PO but marketing rejects it.
They want highlight articles with an image and better layout.
The articles have to support a `Dimensions` method that returns width and
height of the rendered article in pixels.

### Main Method Implementation For Second Stage

Our `AssembleNews` function had to become a bit more complex to accommodate the
changes:
```go
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
```

The signature stays the same.
And the first 12 lines inside the function are used to find and handle the
hightlight article.
Please note that this doesn't need any change to the contract itself.

In the outer `for` loop the articles are rendered in rows (inner `for` loop).
We use the universal wisdom of the marketing department that assures us the
modern screens are 1024 pixel wide.
Highlight articles fill the full width, others create nice columns.

### Go Contract Implementation For Second Stage

For the contract implementation we have got an working pattern and want to see
how well this holds up.

```go
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
```

Phew, that went well! The formal changes we had to do for the first version can
stay exactly the same and we don't need any additional ones.
So the copy and pasting is really painless.
Now we are sure we are on the right track.

The full source code for this stage (all in one file) can be found in
[news2.go](/images/go/news2/news2.go).

### Problem Definition For Third Stage

Our solution is a success in the market!
Now people want the great news in their inbox.
As HTML is supported in emails this shouldn't be too hard.
We just have to add the images as attachments to the emails.
So we need a new `GetImageBytes` method for that.

### Main Method Implementation For Third Stage

Thankfully the big `AssembleNews` function doesn't have to change at all.
The big change is a new `AddImages` function that is calling a `GetImageBytes`
on the articles.
```go
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
```

This is quite straight forward and I am sparing you the details of the relevant
RFC.


### Go Contract Implementation For Last Stage

As the copy and paste model has been so successful so far we will of course
stick to it.  We could invent a second contract for this but as we aren't too
big a fan of those and we like to keep this difficult stuff in one place we
just add it to the existing one.
As good programmers we are adding nice comments of course:
```go
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
```

The first part of the contract didn't change at all (except for the added
comment) and we simply added the content of the `AddImages` function.
In the last line we have to replace `return` with `_ =` again but that isn't
too difficult.
We feel good with this pragmatic style and ready for the years to come.

The full source code for this stage (all in one file) can be found in
[news3.go](/images/go/news3/news3.go).

## Criticism Of The Last Contract

As contrast this is the minimal contract that we should have created:
```go
contract article(a T) {
	var html []byte = a.Render(false)
	var w, h int = a.Dimensions()
	var buf []byte = a.GetImageBytes()
}
```

And we should really use an interface instead:
```go
type Article interface {
        Render(floatLeft bool) string
        Dimensions() (width, height int)
        GetImageBytes() []byte
}
```

It has the same number of lines as the contract but in my humble opinion
declarations are inherently easier to read than definitions.
