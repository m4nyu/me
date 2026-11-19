// Package seo renders search and AI-discoverability metadata for m4nuel.net:
// per-language meta tags, hreflang alternates, Open Graph and JSON-LD, plus
// robots.txt, sitemap.xml and llms.txt.
package seo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"engineer/src/app/internal/i18n"

	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

const (
	SiteURL = "https://m4nuel.net"
	ogImage = SiteURL + "/static/og.png"
	author  = "Manuel Szedlak"
)

// Languages served at /<code>, in the order they appear in the sitemap.
var Languages = []string{"en", "de", "fr", "es", "it", "pt", "nl", "ru", "ja", "ko", "zh", "ar"}

func pageURL(lang string) string { return SiteURL + "/" + strings.ToLower(lang) }

// Tags returns the meta tags for the landing page in lang (e.g. "EN").
func Tags(lang string) []g.Node {
	title := i18n.T(lang, "title") + ": " + i18n.T(lang, "tagline")
	desc := i18n.T(lang, "intro_desc")
	url := pageURL(lang)

	nodes := []g.Node{
		Meta(Name("description"), Content(desc)),
		Link(Rel("canonical"), Href(url)),
	}
	for _, l := range Languages {
		nodes = append(nodes, Link(Rel("alternate"), g.Attr("hreflang", l), Href(pageURL(l))))
	}
	nodes = append(nodes,
		Link(Rel("alternate"), g.Attr("hreflang", "x-default"), Href(pageURL("en"))),
		Meta(g.Attr("property", "og:type"), Content("website")),
		Meta(g.Attr("property", "og:site_name"), Content("4nuel")),
		Meta(g.Attr("property", "og:title"), Content(title)),
		Meta(g.Attr("property", "og:description"), Content(desc)),
		Meta(g.Attr("property", "og:url"), Content(url)),
		Meta(g.Attr("property", "og:image"), Content(ogImage)),
		Meta(g.Attr("property", "og:image:width"), Content("1200")),
		Meta(g.Attr("property", "og:image:height"), Content("630")),
		Meta(g.Attr("property", "og:image:alt"), Content("4nuel: I am unopinionated, therefore I am!")),
		Meta(g.Attr("property", "og:locale"), Content(strings.ToLower(lang))),
		Meta(Name("twitter:card"), Content("summary_large_image")),
		Meta(Name("twitter:creator"), Content("@ManuelSzedlak")),
		Meta(Name("twitter:title"), Content(title)),
		Meta(Name("twitter:description"), Content(desc)),
		Meta(Name("twitter:image"), Content(ogImage)),
		structuredData(lang, desc),
	)
	return nodes
}

func structuredData(lang, desc string) g.Node {
	data := []map[string]any{
		{
			"@context":      "https://schema.org",
			"@type":         "Person",
			"name":          author,
			"alternateName": "4nuel",
			"url":           SiteURL,
			"image":         ogImage,
			"description":   desc,
			"jobTitle":      "Founder & Engineer",
			"knowsAbout":    []string{"Distributed systems", "AI/ML training and inference", "Service architecture", "Rust", "Go"},
			"sameAs": []string{
				"https://m4nuel.com",
				"https://m4nuel.blog",
				"https://github.com/m4nyu",
				"https://www.linkedin.com/in/manuel-szedlak",
				"https://x.com/ManuelSzedlak",
			},
		},
		{
			"@context":   "https://schema.org",
			"@type":      "WebSite",
			"name":       "4nuel",
			"url":        SiteURL,
			"inLanguage": strings.ToLower(lang),
		},
	}
	b, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	// json.Marshal escapes <, > and & as \u00XX, so the text cannot close the tag.
	return Script(Type("application/ld+json"), g.Raw(string(b)))
}

// Robots allows every crawler, AI agents included, and points to the sitemap.
func Robots(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "User-agent: *\nAllow: /\n\nSitemap: %s/sitemap.xml\n", SiteURL)
}

// Sitemap lists every language page with its hreflang alternates.
func Sitemap(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/xml; charset=utf-8")
	var b strings.Builder
	b.WriteString(`<?xml version="1.0" encoding="UTF-8"?>` + "\n")
	b.WriteString(`<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:xhtml="http://www.w3.org/1999/xhtml">` + "\n")
	for _, l := range Languages {
		fmt.Fprintf(&b, "  <url>\n    <loc>%s</loc>\n", pageURL(l))
		for _, alt := range Languages {
			fmt.Fprintf(&b, "    <xhtml:link rel=\"alternate\" hreflang=\"%s\" href=\"%s\"/>\n", alt, pageURL(alt))
		}
		b.WriteString("  </url>\n")
	}
	b.WriteString("</urlset>\n")
	_, _ = w.Write([]byte(b.String()))
}

// LLMs serves /llms.txt: a plain summary for AI answer engines.
func LLMs(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, `# 4nuel (Manuel Szedlak)

> %s

%s

## Pages

- [m4nuel.net](%s): profile in %d languages (%s)

## Elsewhere

- [m4nuel.blog](https://m4nuel.blog): notes on software, with its own llms.txt
- [m4nuel.com](https://m4nuel.com): links and contact
- [GitHub](https://github.com/m4nyu)
- [LinkedIn](https://www.linkedin.com/in/manuel-szedlak)
`, i18n.T("EN", "tagline"), i18n.T("EN", "intro_desc"), pageURL("en"), len(Languages), strings.Join(Languages, ", "))
}
