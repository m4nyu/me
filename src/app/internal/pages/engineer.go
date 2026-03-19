package pages

import (
	"engineer/src/app/internal/i18n"

	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func EngineerPage(props HomePageProps) g.Node {
	lang := props.CurrentLang

	return Doctype(
		HTML(
			Lang("en"),
			Head(
				Meta(Charset("utf-8")),
				Meta(Name("viewport"), Content("width=device-width, initial-scale=1.0")),
				TitleEl(g.Text("Manuel — Engineer Mode")),
			),
			Body(
				Header(
					Nav(
						A(Href("/"), g.Text("Exit Engineer Mode")),
					),
					g.El("hr"),
				),
				Main(
					H1(g.Text("Manuel")),
					P(g.El("em", g.Text(i18n.T(lang, "tagline")))),
					P(g.Text(i18n.T(lang, "intro_desc"))),
					g.El("hr"),

					H2(g.Text(i18n.T(lang, "core_strengths"))),
					P(g.Text(i18n.T(lang, "core_strengths_subtitle"))),

					H3(g.Text(i18n.T(lang, "services_title")+" — "+i18n.T(lang, "services_lang"))),
					P(g.Text(i18n.T(lang, "services_desc"))),

					H3(g.Text(i18n.T(lang, "web_title")+" — "+i18n.T(lang, "web_lang"))),
					P(g.Text(i18n.T(lang, "web_desc"))),

					H3(g.Text(i18n.T(lang, "aiml_title")+" — "+i18n.T(lang, "aiml_lang"))),
					P(g.Text(i18n.T(lang, "aiml_desc"))),

					g.El("hr"),

					H2(g.Text(i18n.T(lang, "opinions"))),
					P(g.Text(i18n.T(lang, "opinions_subtitle"))),

					Dl(
						faqEntry(lang, "faq_q1", "faq_a1"),
						faqEntry(lang, "faq_q2", "faq_a2"),
						faqEntry(lang, "faq_q3", "faq_a3"),
						faqEntry(lang, "faq_q4", "faq_a4"),
						faqEntry(lang, "faq_q5", "faq_a5"),
						faqEntry(lang, "faq_q6", "faq_a6"),
						faqEntry(lang, "faq_q7", "faq_a7"),
						faqEntry(lang, "faq_q8", "faq_a8"),
						faqEntry(lang, "faq_q9", "faq_a9"),
						faqEntry(lang, "faq_q10", "faq_a10"),
					),

					g.El("hr"),
				),
				Footer(
					P(
						A(Href("https://m4nuel.com"), g.Text("m4nuel.com")),
					),
				),
			),
		),
	)
}

func faqEntry(lang, qKey, aKey string) g.Node {
	return g.Group([]g.Node{
		g.El("dt", g.El("strong", g.Text(i18n.T(lang, qKey)))),
		g.El("dd", g.Text(i18n.T(lang, aKey))),
	})
}
