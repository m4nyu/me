package main

import (
	"engineer/src/internal/components"
	"engineer/src/internal/handlers"
	"engineer/src/internal/middleware"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func main() {
	r := chi.NewRouter()

	// Middleware
	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)
	r.Use(middleware.CookieMiddleware)

	// Static files
	fileServer := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	// Cookie handlers
	r.Post("/api/theme", handlers.SetTheme)
	r.Post("/api/lang", handlers.SetLanguage)

	// Routes
	r.Get("/", handleHome)
	r.Get("/imprint", handleImprint)
	r.Get("/gdpr", handleGDPR)
	r.Get("/terms", handleTerms)

	port := ":3000"
	fmt.Printf("Server starting on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	theme := middleware.GetTheme(r)
	lang := middleware.GetLang(r)

	page := components.BaseLayout("4nuel", theme,
		components.HomePage(components.HomePageProps{
			Theme:       theme,
			CurrentLang: lang,
		}),
	)

	w.Header().Set("Content-Type", "text/html")
	_ = page.Render(w)
}

func handleImprint(w http.ResponseWriter, r *http.Request) {
	theme := middleware.GetTheme(r)

	page := components.BaseLayout("Imprint", theme,
		Div(Class("min-h-screen bg-background text-foreground"),
			g.El("header",
				Class("p-6"),
				A(
					Href("/"),
					Class("flex items-center gap-2 text-foreground hover:opacity-80 transition-opacity cursor-pointer p-0 h-auto rounded-none shadow-none border-0 bg-transparent hover:bg-transparent"),
					components.ArrowLeftIcon(),
					Span(Class("text-sm font-light"), g.Text("Back")),
				),
			),
			g.El("main",
				Class("max-w-4xl mx-auto px-6 pb-12 legal-page-content"),
				Div(Class("mb-12"),
					H1(Class("text-4xl font-light mb-8"), g.Text("Imprint")),
					Div(Class("space-y-8 font-light"),
						g.El("section",
							H2(Class("text-xl font-medium mb-4"), g.Text("Information according to § 5 TMG")),
							Div(Class("space-y-2"),
								P(g.Text("Manuel [Your Last Name]")),
								P(g.Text("[Your Street and Number]")),
								P(g.Text("[Your Postal Code and City]")),
								P(g.Text("[Your Country]")),
							),
						),
						g.El("section",
							H2(Class("text-xl font-medium mb-4"), g.Text("Contact")),
							Div(Class("space-y-2"),
								P(g.Text("Email: [your.email@example.com]")),
								P(g.Text("Phone: [your phone number]")),
							),
						),
						g.El("section",
							H2(Class("text-xl font-medium mb-4"), g.Text("VAT ID")),
							P(g.Text("Sales tax identification number according to § 27 a of the Sales Tax Law:")),
							P(g.Text("[Your VAT ID if applicable]")),
						),
						g.El("section",
							H2(Class("text-xl font-medium mb-4"), g.Text("Responsible for the content according to § 55, para. 2 RStV")),
							Div(Class("space-y-2"),
								P(g.Text("Manuel [Your Last Name]")),
								P(g.Text("[Your Street and Number]")),
								P(g.Text("[Your Postal Code and City]")),
							),
						),
						g.El("section",
							H2(Class("text-xl font-medium mb-4"), g.Text("Disclaimer")),
							Div(Class("space-y-4"),
								Div(
									H3(Class("font-medium mb-2"), g.Text("Liability for Contents")),
									P(Class("text-sm leading-relaxed"), g.Text("As service providers, we are liable for own contents of these websites according to Sec. 7, para. 1 of the TMG (Telemediengesetz – Tele Media Act by German law). However, according to Sec. 8 to 10 of the TMG, we as service providers are not under obligation to monitor external information provided or stored on our website. Once we have become aware of a specific infringement of law, we will immediately remove the content in question. Any liability concerning this matter can only be assumed from the point in time at which the infringement becomes known to us.")),
								),
								Div(
									H3(Class("font-medium mb-2"), g.Text("Liability for Links")),
									P(Class("text-sm leading-relaxed"), g.Text("Our website contains links to the websites of third parties (\"external links\"). As the contents of these websites are not under our control, we cannot assume any liability for such external content. In all cases, the provider of information of the linked websites is liable for the content and accuracy of the information provided. At the point in time when the links were placed, no infringements of the law were recognisable to us. As soon as an infringement of the law becomes known to us, we will immediately remove the link in question.")),
								),
								Div(
									H3(Class("font-medium mb-2"), g.Text("Copyright")),
									P(Class("text-sm leading-relaxed"), g.Text("The content and works published on this website are governed by the copyright laws of Germany. Any duplication, processing, distribution or any form of utilisation beyond the scope of copyright law shall require the prior written consent of the author or authors in question.")),
								),
							),
						),
					),
				),
			),
		),
	)

	w.Header().Set("Content-Type", "text/html")
	_ = page.Render(w)
}

func handleGDPR(w http.ResponseWriter, r *http.Request) {
	theme := middleware.GetTheme(r)

	page := components.BaseLayout("GDPR", theme,
		Div(Class("min-h-screen bg-background text-foreground"),
			g.El("header",
				Class("p-6"),
				A(
					Href("/"),
					Class("flex items-center gap-2 text-foreground hover:opacity-80 transition-opacity cursor-pointer p-0 h-auto rounded-none shadow-none border-0 bg-transparent hover:bg-transparent"),
					components.ArrowLeftIcon(),
					Span(Class("text-sm font-light"), g.Text("Back")),
				),
			),
			g.El("main",
				Class("max-w-4xl mx-auto px-6 pb-12 legal-page-content"),
				Div(Class("mb-12"),
					H1(Class("text-4xl font-light mb-8"), g.Text("Privacy Policy (GDPR)")),
					Div(Class("space-y-8 font-light"),
						g.El("section",
							H2(Class("text-xl font-medium mb-4"), g.Text("1. Data Protection Overview")),
							P(Class("text-sm leading-relaxed mb-4"), g.Text("The following gives a simple overview of what happens to your personal information when you visit our website. Personal information is any data with which you could be personally identified.")),
						),
						g.El("section",
							H2(Class("text-xl font-medium mb-4"), g.Text("2. Data Controller")),
							Div(Class("space-y-2 text-sm"),
								P(g.Text("The party responsible for processing data on this website is:")),
								Div(Class("mt-2"),
									P(g.Text("Manuel [Your Last Name]")),
									P(g.Text("[Your Street and Number]")),
									P(g.Text("[Your Postal Code and City]")),
									P(g.Text("Email: [your.email@example.com]")),
								),
							),
						),
						g.El("section",
							H2(Class("text-xl font-medium mb-4"), g.Text("3. How We Collect Your Data")),
							Div(Class("space-y-4 text-sm"),
								Div(
									H3(Class("font-medium mb-2"), g.Text("Contact Forms")),
									P(Class("leading-relaxed"), g.Text("If you submit data to us via a contact form, we collect the data entered in the form, including the contact details you provide, to answer your question and any follow-up questions.")),
								),
								Div(
									H3(Class("font-medium mb-2"), g.Text("Server Log Files")),
									P(Class("leading-relaxed"), g.Text("The website provider automatically collects information in so-called server log files, which your browser automatically transmits to us. This information comprises:")),
									g.El("ul", Class("list-disc list-inside mt-2 space-y-1 ml-4"),
										g.El("li", g.Text("Browser type and browser version")),
										g.El("li", g.Text("Operating system used")),
										g.El("li", g.Text("Referrer URL")),
										g.El("li", g.Text("Host name of the accessing computer")),
										g.El("li", g.Text("Time of the server request")),
										g.El("li", g.Text("IP address")),
									),
								),
							),
						),
						g.El("section",
							H2(Class("text-xl font-medium mb-4"), g.Text("4. Your Rights")),
							Div(Class("text-sm space-y-2"),
								P(g.Text("You always have the right to request information about your stored data, its origin, its recipients, and the purpose of its collection at no charge. You also have the right to request that it be corrected, blocked, or deleted.")),
								P(g.Text("You have the following rights under the GDPR:")),
								g.El("ul", Class("list-disc list-inside mt-2 space-y-1 ml-4"),
									g.El("li", g.Text("Right to information (Article 15 GDPR)")),
									g.El("li", g.Text("Right to rectification (Article 16 GDPR)")),
									g.El("li", g.Text("Right to erasure (Article 17 GDPR)")),
									g.El("li", g.Text("Right to restrict processing (Article 18 GDPR)")),
									g.El("li", g.Text("Right to data portability (Article 20 GDPR)")),
									g.El("li", g.Text("Right to object (Article 21 GDPR)")),
									g.El("li", g.Text("Right to withdraw consent (Article 7 GDPR)")),
								),
							),
						),
						g.El("section",
							H2(Class("text-xl font-medium mb-4"), g.Text("5. Analytics and Third-Party Tools")),
							P(Class("text-sm leading-relaxed"), g.Text("Currently, this website does not use analytics tools or third-party tracking services. Should this change in the future, this privacy policy will be updated accordingly.")),
						),
						g.El("section",
							H2(Class("text-xl font-medium mb-4"), g.Text("6. SSL/TLS Encryption")),
							P(Class("text-sm leading-relaxed"), g.Text("This site uses SSL/TLS encryption for security reasons and to protect the transmission of confidential content, such as the inquiries you send to us as the site operator. You can recognize an encrypted connection by the fact that the address line of the browser changes from \"http://\" to \"https://\".")),
						),
						g.El("section",
							H2(Class("text-xl font-medium mb-4"), g.Text("7. Data Retention")),
							P(Class("text-sm leading-relaxed"), g.Text("We store your data only as long as necessary to provide our services or as required by law. Contact form data is typically deleted after your inquiry has been resolved, unless you have agreed to further communication.")),
						),
						g.El("section",
							H2(Class("text-xl font-medium mb-4"), g.Text("8. Contact")),
							P(Class("text-sm leading-relaxed"), g.Text("If you have questions about this privacy policy or our data processing practices, please contact us at [your.email@example.com].")),
						),
						g.El("section", Class("text-xs opacity-75"),
							P(g.Text("Last updated: [Date]")),
						),
					),
				),
			),
		),
	)

	w.Header().Set("Content-Type", "text/html")
	_ = page.Render(w)
}

func handleTerms(w http.ResponseWriter, r *http.Request) {
	theme := middleware.GetTheme(r)

	page := components.BaseLayout("Terms", theme,
		Div(Class("min-h-screen bg-background text-foreground"),
			g.El("header",
				Class("p-6"),
				A(
					Href("/"),
					Class("flex items-center gap-2 text-foreground hover:opacity-80 transition-opacity cursor-pointer p-0 h-auto rounded-none shadow-none border-0 bg-transparent hover:bg-transparent"),
					components.ArrowLeftIcon(),
					Span(Class("text-sm font-light"), g.Text("Back")),
				),
			),
			g.El("main",
				Class("max-w-4xl mx-auto px-6 pb-12 legal-page-content"),
				Div(Class("mb-12"),
					H1(Class("text-4xl font-light mb-8"), g.Text("Terms of Service")),
					Div(Class("space-y-8 font-light"),
						g.El("section",
							H2(Class("text-xl font-medium mb-4"), g.Text("1. Acceptance of Terms")),
							P(Class("text-sm leading-relaxed"), g.Text("By accessing and using this website and engaging our services, you accept and agree to be bound by the terms and provision of this agreement. If you do not agree to abide by the above, please do not use this service.")),
						),
						g.El("section",
							H2(Class("text-xl font-medium mb-4"), g.Text("2. Services")),
							Div(Class("text-sm space-y-4"),
								P(Class("leading-relaxed"), g.Text("We provide professional software development and consulting services including but not limited to:")),
								g.El("ul", Class("list-disc list-inside space-y-1 ml-4"),
									g.El("li", g.Text("Full-stack web development")),
									g.El("li", g.Text("API design and implementation")),
									g.El("li", g.Text("Cloud architecture consulting")),
									g.El("li", g.Text("Database optimization")),
									g.El("li", g.Text("Custom software solutions")),
									g.El("li", g.Text("Technical consulting")),
								),
							),
						),
						g.El("section",
							H2(Class("text-xl font-medium mb-4"), g.Text("3. Pricing and Payment")),
							Div(Class("text-sm space-y-2"),
								P(Class("leading-relaxed"), g.Text("Service pricing is determined on a project-by-project basis and may include:")),
								g.El("ul", Class("list-disc list-inside space-y-1 ml-4"),
									g.El("li", g.Text("Hourly rates for short-term engagements")),
									g.El("li", g.Text("Daily rates for ongoing projects")),
									g.El("li", g.Text("Fixed project pricing for defined scope work")),
									g.El("li", g.Text("Equity partnerships for qualifying startup ventures")),
								),
								P(Class("leading-relaxed mt-4"), g.Text("Payment terms will be specified in individual project agreements. Generally, payment is due within 30 days of invoice date unless otherwise agreed upon in writing.")),
							),
						),
						g.El("section",
							H2(Class("text-xl font-medium mb-4"), g.Text("4. Project Scope and Changes")),
							Div(Class("text-sm space-y-2"),
								P(Class("leading-relaxed"), g.Text("All projects begin with a discovery phase to establish clear requirements and deliverables. Any changes to the agreed-upon scope may result in additional charges and timeline adjustments.")),
								P(Class("leading-relaxed"), g.Text("We work in iterative phases with regular check-ins to ensure alignment and allow for necessary adjustments throughout the development process.")),
							),
						),
						g.El("section",
							H2(Class("text-xl font-medium mb-4"), g.Text("5. Intellectual Property")),
							Div(Class("text-sm space-y-2"),
								P(Class("leading-relaxed"), g.Text("Upon full payment for services, all custom code and deliverables created specifically for your project will be transferred to you. However, we retain rights to:")),
								g.El("ul", Class("list-disc list-inside space-y-1 ml-4"),
									g.El("li", g.Text("General methodologies and techniques")),
									g.El("li", g.Text("Reusable code components and libraries")),
									g.El("li", g.Text("Pre-existing intellectual property")),
								),
							),
						),
						g.El("section",
							H2(Class("text-xl font-medium mb-4"), g.Text("6. Confidentiality")),
							P(Class("text-sm leading-relaxed"), g.Text("We maintain strict confidentiality regarding all client information, project details, and business data. We will not disclose any confidential information to third parties without explicit written consent, except as required by law.")),
						),
						g.El("section",
							H2(Class("text-xl font-medium mb-4"), g.Text("7. Warranties and Disclaimers")),
							Div(Class("text-sm space-y-2"),
								P(Class("leading-relaxed"), g.Text("We provide services using industry best practices and current technologies. However, software development involves inherent risks and complexities. While we strive for error-free delivery:")),
								g.El("ul", Class("list-disc list-inside space-y-1 ml-4"),
									g.El("li", g.Text("We do not guarantee that software will be completely error-free")),
									g.El("li", g.Text("We provide reasonable support for bug fixes identified within 30 days of delivery")),
									g.El("li", g.Text("Client is responsible for thorough testing before production deployment")),
								),
							),
						),
						g.El("section",
							H2(Class("text-xl font-medium mb-4"), g.Text("8. Limitation of Liability")),
							P(Class("text-sm leading-relaxed"), g.Text("Our liability for any claims arising from our services shall be limited to the amount paid for the specific project or service in question. We shall not be liable for any indirect, incidental, special, or consequential damages.")),
						),
						g.El("section",
							H2(Class("text-xl font-medium mb-4"), g.Text("9. Termination")),
							P(Class("text-sm leading-relaxed"), g.Text("Either party may terminate a project agreement with written notice. In such cases, payment will be due for all work completed up to the termination date. Any deliverables completed will be transferred upon payment.")),
						),
						g.El("section",
							H2(Class("text-xl font-medium mb-4"), g.Text("10. Governing Law")),
							P(Class("text-sm leading-relaxed"), g.Text("These terms shall be governed by and construed in accordance with the laws of [Your Jurisdiction], without regard to its conflict of law provisions.")),
						),
						g.El("section",
							H2(Class("text-xl font-medium mb-4"), g.Text("11. Changes to Terms")),
							P(Class("text-sm leading-relaxed"), g.Text("We reserve the right to modify these terms at any time. Changes will be effective immediately upon posting to this website. Your continued use of our services constitutes acceptance of any changes.")),
						),
						g.El("section",
							H2(Class("text-xl font-medium mb-4"), g.Text("12. Contact Information")),
							P(Class("text-sm leading-relaxed"), g.Text("For questions about these Terms of Service, please contact us at [your.email@example.com].")),
						),
						g.El("section", Class("text-xs opacity-75"),
							P(g.Text("Last updated: [Date]")),
						),
					),
				),
			),
		),
	)

	w.Header().Set("Content-Type", "text/html")
	_ = page.Render(w)
}