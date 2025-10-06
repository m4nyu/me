package components

import (
	"fmt"

	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type HomePageProps struct {
	Theme        string
	CurrentLang  string
	OnNavigate   func(string) string
}

// HomePage creates the complete homepage with mobile carousel and desktop scroll layout
func HomePage(props HomePageProps) g.Node {
	faqItems := []AccordionItem{
		{
			Question: "What's the biggest myth in software engineering?",
			Answer:   "That there's always a 'right' technology or architecture. Every choice is a trade-off. The best solution depends on your constraints, team, and business context. Being unopinionated means choosing the right tool for the job, not the trendy one.",
		},
		{
			Question: "What's your take on AI/ML in production?",
			Answer:   "Most companies don't need custom models—they need better data pipelines and inference infrastructure. Focus on deployment, monitoring, and iteration speed. The model is often the easy part; productionizing it is where the real engineering happens.",
		},
		{
			Question: "How do you approach technical debt?",
			Answer:   "Technical debt is a financial metaphor—treat it like one. Some debt is strategic. The key is knowing the interest rate: high-traffic code with mounting complexity needs immediate attention. Low-touch systems? Let them be until they need to change.",
		},
		{
			Question: "What's the most underrated skill?",
			Answer:   "Deleting code. The best engineers know what NOT to build. Every line of code is a liability—it needs to be tested, maintained, and understood. The most elegant solution is often the one that requires no code at all.",
		},
		{
			Question: "Testing philosophy?",
			Answer:   "Test the contracts, not the implementation. Focus on integration tests over unit tests for business logic. Perfect coverage is a vanity metric. What matters is confidence in your ability to ship without breaking things.",
		},
		{
			Question: "What drives your technical decisions?",
			Answer:   "Business value first, technical elegance second. The best architecture is the one that ships on time and scales when needed. Optimize for iteration speed early, performance later. Premature optimization kills more projects than premature scaling.",
		},
		{
			Question: "Thoughts on the future of programming?",
			Answer:   "AI will amplify good engineers and expose bad ones. The fundamentals—algorithms, systems thinking, trade-off analysis—matter more than ever. Tools change, but problem-solving is timeless. Focus on understanding deeply, not learning frameworks shallowly.",
		},
		{
			Question: "Static or dynamic typing?",
			Answer:   "Use types where they add value: at system boundaries, in critical paths, and for team coordination. Don't dogmatically type everything. Python with type hints beats TypeScript with 'any' everywhere. The goal is clarity, not ceremony.",
		},
		{
			Question: "Database choices?",
			Answer:   "Postgres for 90% of use cases. Start simple, scale vertically, then horizontally. Most 'scale' problems are actually query optimization problems. NoSQL for specific workloads where relational doesn't fit. Multi-database architectures are organizational complexity you probably don't need.",
		},
		{
			Question: "Performance vs. readability?",
			Answer:   "Readable code first, fast code when you measure a problem. Premature optimization is evil, but so is ignoring algorithmic complexity. Profile before you optimize. Most performance issues are at architectural level—choose the right algorithm, not clever micro-optimizations.",
		},
	}

	return g.Group([]g.Node{
		// Mobile Carousel Layout
		MobileCarouselLayout(props, faqItems),
		// Desktop Layout
		DesktopLayout(props, faqItems),
		// Language Drawer (rendered at body level to avoid overflow issues)
		LanguageDrawer(props.CurrentLang),
	})
}

// MobileCarouselLayout creates the mobile carousel view with swipe support
func MobileCarouselLayout(props HomePageProps, faqItems []AccordionItem) g.Node {
	return Div(
		Class("lg:hidden h-screen bg-foreground dark:bg-foreground text-black relative overflow-hidden"),
		ID("mobile-carousel"),
		// Carousel Container
		Div(
			Class("carousel-container h-full w-full mb-12"),
			// Hidden radio buttons for carousel control
			Input(Type("radio"), ID("slide1"), Name("mobile-carousel"), Checked(), g.Attr("style", "display: none;")),
			Input(Type("radio"), ID("slide2"), Name("mobile-carousel"), g.Attr("style", "display: none;")),
			Input(Type("radio"), ID("slide3"), Name("mobile-carousel"), g.Attr("style", "display: none;")),
			// Slides wrapper
			Div(
				Class("carousel-slides h-full"),
				// Section 1: Main content
				MobileSection1Swiper(props),
				// Section 2: Core Strengths
				MobileSection3Swiper(),
				// Section 3: Opinions
				MobileSection4Swiper(faqItems),
			),
			// Section Indicators
			Div(
				Class("carousel-indicators absolute bottom-2 left-1/2 -translate-x-1/2 z-50 flex gap-1"),
				g.Group([]g.Node{
					CarouselDot(1),
					CarouselDot(2),
					CarouselDot(3),
				}),
			),
		),
		// JavaScript for swipe detection
		g.El("script", g.Raw(`
			(function() {
				const carousel = document.getElementById('mobile-carousel');
				if (!carousel) return;

				let startX = 0;
				let startY = 0;
				let startTime = 0;
				let isDragging = false;

				function getCurrentSlide() {
					if (document.getElementById('slide1').checked) return 1;
					if (document.getElementById('slide2').checked) return 2;
					if (document.getElementById('slide3').checked) return 3;
					return 1;
				}

				function goToSlide(slideNum) {
					if (slideNum >= 1 && slideNum <= 3) {
						document.getElementById('slide' + slideNum).checked = true;
					}
				}

				// Touch events
				carousel.addEventListener('touchstart', function(e) {
					startX = e.touches[0].clientX;
					startY = e.touches[0].clientY;
					startTime = Date.now();
					isDragging = true;
				}, { passive: true });

				carousel.addEventListener('touchmove', function(e) {
					if (!isDragging) return;
				}, { passive: true });

				carousel.addEventListener('touchend', function(e) {
					if (!isDragging) return;
					isDragging = false;

					const endX = e.changedTouches[0].clientX;
					const endY = e.changedTouches[0].clientY;
					const endTime = Date.now();
					const diffX = startX - endX;
					const diffY = startY - endY;
					const timeDiff = endTime - startTime;
					const velocity = Math.abs(diffX) / timeDiff;

					// Trigger on: horizontal swipe (15px+) OR fast swipe (velocity > 0.3)
					const isHorizontalSwipe = Math.abs(diffX) > Math.abs(diffY);
					const isSignificantDistance = Math.abs(diffX) > 15;
					const isFastSwipe = velocity > 0.3 && Math.abs(diffX) > 5;

					if (isHorizontalSwipe && (isSignificantDistance || isFastSwipe)) {
						const currentSlide = getCurrentSlide();
						if (diffX > 0) {
							// Swipe left - next slide
							goToSlide(currentSlide + 1);
						} else {
							// Swipe right - previous slide
							goToSlide(currentSlide - 1);
						}
					}
				}, { passive: true });

				// Mouse events for desktop testing
				carousel.addEventListener('mousedown', function(e) {
					startX = e.clientX;
					startY = e.clientY;
					startTime = Date.now();
					isDragging = true;
					e.preventDefault();
				});

				carousel.addEventListener('mousemove', function(e) {
					if (!isDragging) return;
					e.preventDefault();
				});

				carousel.addEventListener('mouseup', function(e) {
					if (!isDragging) return;
					isDragging = false;

					const endX = e.clientX;
					const endY = e.clientY;
					const endTime = Date.now();
					const diffX = startX - endX;
					const diffY = startY - endY;
					const timeDiff = endTime - startTime;
					const velocity = Math.abs(diffX) / timeDiff;

					// Trigger on: horizontal swipe (15px+) OR fast swipe (velocity > 0.3)
					const isHorizontalSwipe = Math.abs(diffX) > Math.abs(diffY);
					const isSignificantDistance = Math.abs(diffX) > 15;
					const isFastSwipe = velocity > 0.3 && Math.abs(diffX) > 5;

					if (isHorizontalSwipe && (isSignificantDistance || isFastSwipe)) {
						const currentSlide = getCurrentSlide();
						if (diffX > 0) {
							// Swipe left - next slide
							goToSlide(currentSlide + 1);
						} else {
							// Swipe right - previous slide
							goToSlide(currentSlide - 1);
						}
					}
				});

				carousel.addEventListener('mouseleave', function() {
					isDragging = false;
				});
			})();
		`)),
	)
}

// CarouselDot creates a navigation label for the carousel (pure CSS)
func CarouselDot(index int) g.Node {
	return g.El("label",
		g.Attr("for", fmt.Sprintf("slide%d", index)),
		Class("indicator w-3 h-3 transition-all duration-300 touch-manipulation cursor-pointer border rounded-none bg-transparent border-white dark:border-black inline-block"),
		g.Attr("aria-label", fmt.Sprintf("Go to section %d", index)),
	)
}

// Mobile Sections (Carousel Slides)
func MobileSection1Swiper(props HomePageProps) g.Node {
	return Div(
		Class("carousel-slide h-full w-full"),
		Div(
			Class("w-full h-full p-3 pb-8"),
			Div(
				Class("w-full h-full flex flex-col justify-between bg-background rounded-xl p-4 overflow-hidden"),
				// Main content
				Div(
					Class("text-center flex-1 flex flex-col justify-center"),
					Div(
						Class("flex items-center justify-center space-x-3 mb-4"),
						Logo("medium"),
						Span(Class("text-5xl font-thin text-foreground"), g.Text("4nuel")),
					),
					H1(Class("text-lg font-light mb-3 text-foreground"), g.Text("I am unopinionated, therefore I am")),
					P(
						Class("text-xs text-foreground mb-4 font-light leading-relaxed px-2"),
						g.Text("I take systems from SOA to serverless, monolithic to microservices. From MVP to production and scaling. I specialize in AI/ML training and inference, distributed systems, and service architecture."),
					),
				),
				// Footer
				Div(
					Class("flex items-end justify-center w-full pb-4"),
					// Theme and Language switchers
					Div(
						Class("flex flex-row gap-2"),
						ThemeSwitcher(props.Theme),
						LanguageSwitcher(props.CurrentLang),
					),
				),
			),
		),
	)
}

func MobileSection2Swiper() g.Node {
	return Div(
		Class("carousel-slide h-full w-full"),
		Div(
			Class("w-full h-full p-3 pb-8"),
			Div(
				Class("w-full h-full flex flex-col justify-center bg-background rounded-xl p-4 overflow-hidden"),
				Div(
					Class("max-w-2xl mx-auto text-center px-4"),
					H2(Class("text-lg font-light mb-3 text-foreground"), g.Text("I am unopinionated, therefore I am")),
					P(
						Class("text-xs text-foreground mb-4 font-light leading-relaxed"),
						g.Text("I take systems from SOA to serverless, monolithic to microservices. From MVP to production and scaling. I specialize in AI/ML training and inference, distributed systems, and service architecture."),
					),
				),
			),
		),
	)
}

func MobileSection3Swiper() g.Node {
	return Div(
		Class("carousel-slide h-full w-full"),
		Div(
			Class("w-full h-full p-3 pb-8 overflow-hidden"),
			Div(
				Class("w-full h-full flex flex-col justify-center bg-background rounded-xl p-4"),
				Div(
					Class("w-full px-2"),
					Div(
						Class("text-center mb-4"),
						H2(Class("text-base font-light mb-1 text-foreground"), g.Text("Core Strengths")),
						P(Class("text-[10px] text-foreground max-w-xl mx-auto font-light"), g.Text("Technologies and domains I excel in")),
					),
					Div(
						Class("grid grid-cols-1 gap-2 max-w-md mx-auto"),
						MobilePricingCard("Services", "C, Rust, Zig", "", "Performance-critical services", false),
						MobilePricingCard("Web", "Go, JS/TS", "", "Pulumi, Terraform | Cloud & IaC", false),
						MobilePricingCard("AI/ML", "Python", "", "ML, RL, SNNs, Liquid NNs", false),
					),
				),
			),
		),
	)
}

func MobileSection4Swiper(faqItems []AccordionItem) g.Node {
	return Div(
		Class("carousel-slide h-full w-full"),
		Div(
			Class("w-full h-full p-3 pb-8"),
			Div(
				Class("w-full h-full flex flex-col justify-center bg-background rounded-xl p-4 overflow-y-auto"),
				Div(
					Class("w-full px-4 max-w-2xl mx-auto"),
					Div(
						Class("text-center mb-4"),
						H2(Class("text-lg font-light mb-2 text-foreground"), g.Text("Opinions")),
						P(Class("text-xs text-foreground font-light"), g.Text("Thoughts on engineering, architecture, and technology")),
					),
					Accordion("border border-foreground", faqItems),
				),
			),
		),
	)
}

// Desktop Layout
func DesktopLayout(props HomePageProps, faqItems []AccordionItem) g.Node {
	return Div(
		Class("hidden lg:grid min-h-screen bg-foreground dark:bg-foreground text-black grid-cols-5 selectable-content relative"),
		// Left Column - Fixed
		DesktopLeftColumn(props),
		// Right Column - Scrollable
		DesktopRightColumn(faqItems),
	)
}

func DesktopLeftColumn(props HomePageProps) g.Node {
	return Div(
		Class("h-screen p-6 lg:pl-10 lg:pr-3 col-span-2 bg-foreground dark:bg-foreground"),
		Div(
			Class("p-8 w-full h-full flex flex-col justify-between bg-background rounded-2xl"),
			// Main content
			Div(
				Class("text-center flex-1 flex flex-col justify-center"),
				Div(
					Class("flex items-center justify-center space-x-4 mb-8"),
					Logo("medium"),
					Span(Class("text-4xl md:text-6xl lg:text-8xl xl:text-9xl font-thin text-foreground"), g.Text("4nuel")),
				),
				H1(
					Class("text-xl md:text-2xl lg:text-3xl xl:text-4xl font-light mb-4 lg:mb-6 text-foreground"),
					g.Text("I am unopinionated, therefore I am"),
				),
				P(
					Class("text-xs md:text-sm lg:text-base xl:text-lg text-foreground mb-6 lg:mb-8 font-light leading-relaxed"),
					g.Text("I take systems from SOA to serverless, monolithic to microservices. From MVP to production and scaling. I specialize in AI/ML training and inference, distributed systems, and service architecture."),
				),
			),
			// Footer
			Div(
				Class("flex items-end justify-center w-full pb-0"),
				// Theme and Language switchers
				Div(
					Class("flex flex-row gap-1 sm:gap-2 md:gap-3 lg:gap-4"),
					ThemeSwitcher(props.Theme),
					LanguageSwitcher(props.CurrentLang),
				),
			),
		),
	)
}

func DesktopRightColumn(faqItems []AccordionItem) g.Node {
	return Div(
		Class("h-screen relative col-span-3 p-6 lg:pl-3 lg:pr-10 bg-foreground dark:bg-foreground"),
		Div(
			Class("h-full bg-background rounded-2xl overflow-hidden"),
			// Scrollable content
			Div(
				ID("scroll-container"),
				Class("h-full overflow-y-scroll snap-y snap-mandatory scrollbar-hide bg-background rounded-2xl"),
				Main(
					Class("w-full bg-background"),
					// Section 1: Core Strengths
					DesktopSkillsSection(),
					// Section 2: Opinions
					DesktopQASection(faqItems),
				),
			),
			// Section Navigation Dots
			Div(
				Class("absolute right-2 lg:right-4 top-1/2 -translate-y-1/2 z-40 flex-col gap-1 hidden lg:flex"),
				g.Group([]g.Node{
					DesktopSectionDot(0),
					DesktopSectionDot(1),
				}),
			),
			// JavaScript for scroll-based indicator highlighting
			g.El("script", g.Raw(`
				(function() {
					const container = document.getElementById('scroll-container');
					const sections = container.querySelectorAll('.snap-section');
					const dots = document.querySelectorAll('.desktop-nav-dot');

					function updateActiveIndicator() {
						const scrollTop = container.scrollTop;
						const containerHeight = container.clientHeight;

						let activeIndex = 0;
						sections.forEach((section, index) => {
							const sectionTop = section.offsetTop;
							const sectionHeight = section.offsetHeight;

							if (scrollTop >= sectionTop - containerHeight / 2) {
								activeIndex = index;
							}
						});

						dots.forEach((dot, index) => {
							if (index === activeIndex) {
								dot.classList.add('active');
							} else {
								dot.classList.remove('active');
							}
						});
					}

					// Add click handlers to dots
					dots.forEach((dot) => {
						dot.addEventListener('click', function() {
							const sectionId = this.getAttribute('data-section');
							const section = document.getElementById(sectionId);
							if (section) {
								section.scrollIntoView({ behavior: 'smooth', block: 'start' });
							}
						});
					});

					container.addEventListener('scroll', updateActiveIndicator);
					updateActiveIndicator();
				})();
			`)),
		),
	)
}

func DesktopSectionDot(index int) g.Node {
	sectionID := ""

	switch index {
	case 0:
		sectionID = "skills"
	case 1:
		sectionID = "qa"
	}

	return Div(
		Class("desktop-nav-dot w-3 h-3 transition-all duration-300 cursor-pointer border rounded-none bg-transparent border-white dark:border-black hover:bg-transparent block"),
		g.Attr("data-section", sectionID),
		g.Attr("aria-label", fmt.Sprintf("Go to section %d", index+1)),
	)
}

func DesktopSection(id, title, description string, withButton bool) g.Node {
	children := []g.Node{
		H2(
			Class("text-xl md:text-2xl lg:text-3xl xl:text-4xl font-light mb-4 lg:mb-6 text-foreground"),
			g.Text(title),
		),
		P(
			Class("text-xs md:text-sm lg:text-base xl:text-lg text-foreground mb-6 lg:mb-8 font-light leading-relaxed"),
			g.Text(description),
		),
	}

	if withButton {
		children = append(children, Button(
			Type("button"),
			Class("bg-foreground text-background border-0 outline-none cursor-pointer text-xs md:text-sm lg:text-base px-4 lg:px-6 py-2 lg:py-3 font-light transition-colors h-auto rounded-none shadow-none hover:bg-foreground hover:opacity-80"),
			g.Text("Get in Touch"),
		))
	}

	return Div(
		ID(id),
		Class("snap-section snap-start h-screen flex items-center justify-center bg-background"),
		Div(
			Class("max-w-2xl mx-auto text-center px-4 lg:px-8"),
			g.Group(children),
		),
	)
}

func DesktopSkillsSection() g.Node {
	return Div(
		ID("skills"),
		Class("snap-section snap-start h-screen flex items-center justify-center bg-background"),
		Div(
			Class("w-full px-4 lg:px-8"),
			Div(
				Class("text-center mb-6 lg:mb-8"),
				H2(Class("text-xl md:text-2xl lg:text-3xl xl:text-4xl font-light mb-3 lg:mb-4 text-foreground"), g.Text("Core Strengths")),
				P(Class("text-xs md:text-sm lg:text-base xl:text-lg text-foreground max-w-xl mx-auto font-light"), g.Text("Technologies and domains I excel in")),
			),
			Div(
				Class("grid grid-cols-1 md:grid-cols-3 gap-4 max-w-4xl mx-auto"),
				PricingCard("Services", "C, Rust, Zig", "", "Performance-critical services", false),
				PricingCard("Web", "Go, JS/TS", "", "Pulumi, Terraform | Cloud & IaC", false),
				PricingCard("AI/ML", "Python", "", "ML, RL, SNNs, Liquid NNs", false),
			),
		),
	)
}

func DesktopQASection(faqItems []AccordionItem) g.Node {
	return Div(
		ID("qa"),
		Class("snap-section snap-start h-screen flex items-center justify-center bg-background"),
		Div(
			Class("w-full px-4 lg:px-8 max-w-2xl mx-auto"),
			Div(
				Class("text-center mb-4 lg:mb-6"),
				H2(Class("text-xl md:text-2xl lg:text-3xl xl:text-4xl font-light mb-3 lg:mb-4 text-foreground"), g.Text("Opinions")),
				P(Class("text-xs md:text-sm lg:text-base xl:text-lg text-foreground font-light"), g.Text("Thoughts on engineering, architecture, and technology")),
			),
			Accordion("border border-foreground", faqItems),
		),
	)
}

// Helper Components
func LegalLink(href, text string) g.Node {
	return A(
		Href(href),
		Class("text-foreground hover:text-foreground text-xs sm:text-sm font-light cursor-pointer transition-all p-0 h-auto rounded-none shadow-none border-0 bg-transparent hover:bg-transparent hover:[text-shadow:_0_0_10px_rgba(0,0,0,0.3)] dark:hover:[text-shadow:_0_0_10px_rgba(255,255,255,0.5)] active:[text-shadow:_0_0_15px_rgba(0,0,0,0.5)] dark:active:[text-shadow:_0_0_15px_rgba(255,255,255,0.7)]"),
		g.Text(text),
	)
}

func MobilePricingCard(title, price, period, description string, highlighted bool) g.Node {
	bgClass := "p-2 text-center border border-foreground"
	textClass := "text-foreground"

	if highlighted {
		bgClass = "border border-foreground p-2 text-center bg-foreground text-background"
		textClass = "text-background"
	}

	priceContent := g.Text(price)
	if period != "" {
		return Div(
			Class(bgClass),
			H3(Class("text-sm font-light mb-1 "+textClass), g.Text(title)),
			Div(
				Class("mb-1"),
				Span(Class("text-sm font-light "+textClass), g.Text(price)),
				Span(Class("text-[10px] "+textClass), g.Text(period)),
			),
			P(Class("text-[10px] font-light leading-tight "+textClass), g.Text(description)),
		)
	}

	return Div(
		Class(bgClass),
		H3(Class("text-sm font-light mb-1 "+textClass), g.Text(title)),
		Div(
			Class("mb-1"),
			Span(Class("text-sm font-light "+textClass), priceContent),
		),
		P(Class("text-[10px] font-light leading-tight "+textClass), g.Text(description)),
	)
}

func PricingCard(title, price, period, description string, highlighted bool) g.Node {
	bgClass := "p-6 text-center border border-foreground"
	textClass := "text-foreground"

	if highlighted {
		bgClass = "border border-foreground p-6 text-center bg-foreground text-background"
		textClass = "text-background"
	}

	priceContent := g.Text(price)
	if period != "" {
		return Div(
			Class(bgClass),
			H3(Class("text-xl font-light mb-2 "+textClass), g.Text(title)),
			Div(
				Class("mb-4"),
				Span(Class("text-2xl font-light "+textClass), g.Text(price)),
				Span(Class("text-sm "+textClass), g.Text(period)),
			),
			P(Class("text-sm font-light "+textClass), g.Text(description)),
		)
	}

	return Div(
		Class(bgClass),
		H3(Class("text-xl font-light mb-2 "+textClass), g.Text(title)),
		Div(
			Class("mb-4"),
			Span(Class("text-xl font-light "+textClass), priceContent),
		),
		P(Class("text-sm font-light "+textClass), g.Text(description)),
	)
}
