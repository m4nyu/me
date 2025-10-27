package pages

import (
	"engineer/src/app/internal/components"
	"engineer/src/app/internal/i18n"
	"fmt"

	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type HomePageProps struct {
	Theme       string
	CurrentLang string
}

func HomePage(props HomePageProps) g.Node {
	faqItems := []components.AccordionItem{
		{
			Question: i18n.T(props.CurrentLang, "faq_q1"),
			Answer:   i18n.T(props.CurrentLang, "faq_a1"),
		},
		{
			Question: i18n.T(props.CurrentLang, "faq_q2"),
			Answer:   i18n.T(props.CurrentLang, "faq_a2"),
		},
		{
			Question: i18n.T(props.CurrentLang, "faq_q3"),
			Answer:   i18n.T(props.CurrentLang, "faq_a3"),
		},
		{
			Question: i18n.T(props.CurrentLang, "faq_q4"),
			Answer:   i18n.T(props.CurrentLang, "faq_a4"),
		},
		{
			Question: i18n.T(props.CurrentLang, "faq_q5"),
			Answer:   i18n.T(props.CurrentLang, "faq_a5"),
		},
		{
			Question: i18n.T(props.CurrentLang, "faq_q6"),
			Answer:   i18n.T(props.CurrentLang, "faq_a6"),
		},
		{
			Question: i18n.T(props.CurrentLang, "faq_q7"),
			Answer:   i18n.T(props.CurrentLang, "faq_a7"),
		},
		{
			Question: i18n.T(props.CurrentLang, "faq_q8"),
			Answer:   i18n.T(props.CurrentLang, "faq_a8"),
		},
		{
			Question: i18n.T(props.CurrentLang, "faq_q9"),
			Answer:   i18n.T(props.CurrentLang, "faq_a9"),
		},
		{
			Question: i18n.T(props.CurrentLang, "faq_q10"),
			Answer:   i18n.T(props.CurrentLang, "faq_a10"),
		},
	}

	return g.Group([]g.Node{
		g.El("style", g.Raw(`
			/* Scrollbar hiding */
			.scrollbar-hide {
				-ms-overflow-style: none;
				scrollbar-width: none;
			}

			.scrollbar-hide::-webkit-scrollbar,
			#mobile-swipe-container::-webkit-scrollbar {
				display: none;
			}

			/* Selectable content */
			.selectable-content,
			.selectable-content * {
				user-select: text !important;
				-webkit-user-select: text !important;
				-moz-user-select: text !important;
				-ms-user-select: text !important;
			}

			.selectable-content::selection {
				background-color: #000000;
				color: #ffffff;
			}

			.selectable-content ::-moz-selection {
				background-color: #000000;
				color: #ffffff;
			}

			.dark .selectable-content::selection {
				background-color: #ffffff;
				color: #000000;
			}

			.dark .selectable-content ::-moz-selection {
				background-color: #ffffff;
				color: #000000;
			}
		`)),
		MobileCarouselLayout(props, faqItems),
		DesktopLayout(props, faqItems),
		components.LanguageDrawer(props.CurrentLang),
	})
}

func MobileCarouselLayout(props HomePageProps, faqItems []components.AccordionItem) g.Node {
	return Div(
		Class("lg:hidden h-screen bg-foreground dark:bg-foreground text-black relative overflow-hidden"),
		ID("mobile-carousel"),
		Div(
			Class("relative overflow-hidden h-full w-full mb-12"),
			Input(Type("radio"), ID("slide1"), Name("mobile-carousel"), Checked(), g.Attr("style", "display: none;")),
			Input(Type("radio"), ID("slide2"), Name("mobile-carousel"), g.Attr("style", "display: none;")),
			Input(Type("radio"), ID("slide3"), Name("mobile-carousel"), g.Attr("style", "display: none;")),
			Div(
				Class("carousel-slides flex h-full transition-transform duration-300 ease-in-out"),
				MobileSection1Swiper(props),
				MobileSection3Swiper(props),
				MobileSection4Swiper(props, faqItems),
			),
			Div(
				Class("carousel-indicators absolute bottom-2 left-1/2 -translate-x-1/2 z-50 flex gap-1"),
				g.Group([]g.Node{
					CarouselDot(1),
					CarouselDot(2),
					CarouselDot(3),
				}),
			),
		),
		g.El("style", g.Raw(`
			/* Radio button controls */
			#slide1:checked ~ .carousel-slides {
				transform: translateX(0%);
			}

			#slide2:checked ~ .carousel-slides {
				transform: translateX(-100%);
			}

			#slide3:checked ~ .carousel-slides {
				transform: translateX(-200%);
			}

			/* Active indicator styling */
			#slide1:checked ~ .carousel-indicators .indicator:nth-child(1),
			#slide2:checked ~ .carousel-indicators .indicator:nth-child(2),
			#slide3:checked ~ .carousel-indicators .indicator:nth-child(3) {
				background-color: white;
				border-color: white;
				box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
				width: 2.5rem;
				height: 0.75rem;
			}

			.dark #slide1:checked ~ .carousel-indicators .indicator:nth-child(1),
			.dark #slide2:checked ~ .carousel-indicators .indicator:nth-child(2),
			.dark #slide3:checked ~ .carousel-indicators .indicator:nth-child(3) {
				background-color: black;
				border-color: black;
			}
		`)),
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

					const isHorizontalSwipe = Math.abs(diffX) > Math.abs(diffY);
					const isSignificantDistance = Math.abs(diffX) > 15;
					const isFastSwipe = velocity > 0.3 && Math.abs(diffX) > 5;

					if (isHorizontalSwipe && (isSignificantDistance || isFastSwipe)) {
						const currentSlide = getCurrentSlide();
						if (diffX > 0) {
							goToSlide(currentSlide + 1);
						} else {
							goToSlide(currentSlide - 1);
						}
					}
				}, { passive: true });

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

					const isHorizontalSwipe = Math.abs(diffX) > Math.abs(diffY);
					const isSignificantDistance = Math.abs(diffX) > 15;
					const isFastSwipe = velocity > 0.3 && Math.abs(diffX) > 5;

					if (isHorizontalSwipe && (isSignificantDistance || isFastSwipe)) {
						const currentSlide = getCurrentSlide();
						if (diffX > 0) {
							goToSlide(currentSlide + 1);
						} else {
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

func CarouselDot(index int) g.Node {
	return g.El("label",
		g.Attr("for", fmt.Sprintf("slide%d", index)),
		Class("indicator w-3 h-3 transition-all duration-300 touch-manipulation cursor-pointer border rounded-none bg-transparent border-white dark:border-black inline-block"),
		g.Attr("aria-label", fmt.Sprintf("Go to section %d", index)),
	)
}

func MobileSection1Swiper(props HomePageProps) g.Node {
	return Div(
		Class("min-w-full flex-shrink-0 h-full w-full"),
		Div(
			Class("w-full h-full p-3 pb-8"),
			Div(
				Class("w-full h-full flex flex-col justify-between bg-background rounded-xl p-4 overflow-hidden"),
				Div(
					Class("text-center flex-1 flex flex-col justify-center"),
					Div(
						Class("flex items-center justify-center space-x-3 mb-4"),
						components.Logo("medium"),
						Span(Class("text-5xl font-thin text-foreground"), g.Text("4nuel")),
					),
					H1(Class("text-lg font-light mb-3 text-foreground"), g.Text(i18n.T(props.CurrentLang, "tagline"))),
					P(
						Class("text-xs text-foreground mb-4 font-light leading-relaxed px-2"),
						g.Text(i18n.T(props.CurrentLang, "intro_desc")),
					),
				),
				Div(
					Class("flex items-end justify-center w-full pb-4"),
					Div(
						Class("flex flex-row gap-2"),
						components.ThemeSwitcher(props.Theme),
						components.LanguageSwitcher(props.CurrentLang),
					),
				),
			),
		),
	)
}

func MobileSection3Swiper(props HomePageProps) g.Node {
	return Div(
		Class("min-w-full flex-shrink-0 h-full w-full"),
		Div(
			Class("w-full h-full p-3 pb-8 overflow-hidden"),
			Div(
				Class("w-full h-full flex flex-col justify-center bg-background rounded-xl p-4"),
				Div(
					Class("w-full px-2"),
					Div(
						Class("text-center mb-4"),
						H2(Class("text-base font-light mb-1 text-foreground"), g.Text(i18n.T(props.CurrentLang, "core_strengths"))),
						P(Class("text-[10px] text-foreground max-w-xl mx-auto font-light"), g.Text(i18n.T(props.CurrentLang, "core_strengths_subtitle"))),
					),
					Div(
						Class("grid grid-cols-1 gap-2 max-w-md mx-auto"),
						MobilePricingCard(i18n.T(props.CurrentLang, "services_title"), i18n.T(props.CurrentLang, "services_lang"), "", i18n.T(props.CurrentLang, "services_desc"), false),
						MobilePricingCard(i18n.T(props.CurrentLang, "web_title"), i18n.T(props.CurrentLang, "web_lang"), "", i18n.T(props.CurrentLang, "web_desc"), false),
						MobilePricingCard(i18n.T(props.CurrentLang, "aiml_title"), i18n.T(props.CurrentLang, "aiml_lang"), "", i18n.T(props.CurrentLang, "aiml_desc"), false),
					),
				),
			),
		),
	)
}

func MobileSection4Swiper(props HomePageProps, faqItems []components.AccordionItem) g.Node {
	return Div(
		Class("min-w-full flex-shrink-0 h-full w-full"),
		Div(
			Class("w-full h-full p-3 pb-8"),
			Div(
				Class("w-full h-full flex flex-col justify-center bg-background rounded-xl p-4 overflow-y-auto"),
				Div(
					Class("w-full px-4 max-w-2xl mx-auto"),
					Div(
						Class("text-center mb-4"),
						H2(Class("text-lg font-light mb-2 text-foreground"), g.Text(i18n.T(props.CurrentLang, "opinions"))),
						P(Class("text-xs text-foreground font-light"), g.Text(i18n.T(props.CurrentLang, "opinions_subtitle"))),
					),
					components.Accordion("border border-foreground", faqItems),
				),
			),
		),
	)
}

func DesktopLayout(props HomePageProps, faqItems []components.AccordionItem) g.Node {
	return Div(
		Class("hidden lg:grid min-h-screen bg-foreground dark:bg-foreground text-black grid-cols-5 selectable-content relative"),
		DesktopLeftColumn(props),
		DesktopRightColumn(props, faqItems),
	)
}

func DesktopLeftColumn(props HomePageProps) g.Node {
	return Div(
		Class("h-screen p-6 lg:pl-10 lg:pr-3 col-span-2 bg-foreground dark:bg-foreground"),
		Div(
			Class("p-8 w-full h-full flex flex-col justify-between bg-background rounded-2xl"),
			Div(
				Class("text-center flex-1 flex flex-col justify-center"),
				Div(
					Class("flex items-center justify-center space-x-4 mb-8"),
					components.Logo("medium"),
					Span(Class("text-4xl md:text-6xl lg:text-8xl xl:text-9xl font-thin text-foreground"), g.Text("4nuel")),
				),
				H1(
					Class("text-xl md:text-2xl lg:text-3xl xl:text-4xl font-light mb-4 lg:mb-6 text-foreground"),
					g.Text(i18n.T(props.CurrentLang, "tagline")),
				),
				P(
					Class("text-xs md:text-sm lg:text-base xl:text-lg text-foreground mb-6 lg:mb-8 font-light leading-relaxed"),
					g.Text(i18n.T(props.CurrentLang, "intro_desc")),
				),
			),
			Div(
				Class("flex items-end justify-center w-full pb-0"),
				Div(
					Class("flex flex-row gap-1 sm:gap-2 md:gap-3 lg:gap-4"),
					components.ThemeSwitcher(props.Theme),
					components.LanguageSwitcher(props.CurrentLang),
				),
			),
		),
	)
}

func DesktopRightColumn(props HomePageProps, faqItems []components.AccordionItem) g.Node {
	return Div(
		Class("h-screen relative col-span-3 p-6 lg:pl-3 lg:pr-10 bg-foreground dark:bg-foreground"),
		Div(
			Class("h-full bg-background rounded-2xl overflow-hidden"),
			Div(
				ID("scroll-container"),
				Class("h-full overflow-y-scroll snap-y snap-mandatory scrollbar-hide bg-background rounded-2xl"),
				Main(
					Class("w-full bg-background"),
					DesktopSkillsSection(props),
					DesktopQASection(props, faqItems),
				),
			),
			Div(
				Class("absolute right-2 lg:right-4 top-1/2 -translate-y-1/2 z-40 flex-col gap-1 hidden lg:flex"),
				g.Group([]g.Node{
					DesktopSectionDot(0),
					DesktopSectionDot(1),
				}),
			),
			g.El("style", g.Raw(`
				.desktop-nav-dot.active {
					background-color: white;
					border-color: white;
					box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
					width: 0.75rem;
					height: 2.5rem;
				}

				.dark .desktop-nav-dot.active {
					background-color: black;
					border-color: black;
				}
			`)),
			g.El("script", g.Raw(`
				(function() {
					const container = document.getElementById('scroll-container');
					const sections = container.querySelectorAll('.snap-start');
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

func DesktopSkillsSection(props HomePageProps) g.Node {
	return Div(
		ID("skills"),
		Class("snap-start snap-always min-h-screen h-screen flex items-center justify-center bg-background"),
		Div(
			Class("w-full px-4 lg:px-8"),
			Div(
				Class("text-center mb-6 lg:mb-8"),
				H2(Class("text-xl md:text-2xl lg:text-3xl xl:text-4xl font-light mb-3 lg:mb-4 text-foreground"), g.Text(i18n.T(props.CurrentLang, "core_strengths"))),
				P(Class("text-xs md:text-sm lg:text-base xl:text-lg text-foreground max-w-xl mx-auto font-light"), g.Text(i18n.T(props.CurrentLang, "core_strengths_subtitle"))),
			),
			Div(
				Class("grid grid-cols-1 md:grid-cols-3 gap-4 max-w-4xl mx-auto"),
				PricingCard(i18n.T(props.CurrentLang, "services_title"), i18n.T(props.CurrentLang, "services_lang"), "", i18n.T(props.CurrentLang, "services_desc"), false),
				PricingCard(i18n.T(props.CurrentLang, "web_title"), i18n.T(props.CurrentLang, "web_lang"), "", i18n.T(props.CurrentLang, "web_desc"), false),
				PricingCard(i18n.T(props.CurrentLang, "aiml_title"), i18n.T(props.CurrentLang, "aiml_lang"), "", i18n.T(props.CurrentLang, "aiml_desc"), false),
			),
		),
	)
}

func DesktopQASection(props HomePageProps, faqItems []components.AccordionItem) g.Node {
	return Div(
		ID("qa"),
		Class("snap-start snap-always min-h-screen h-screen flex items-center justify-center bg-background"),
		Div(
			Class("w-full px-4 lg:px-8 max-w-2xl mx-auto"),
			Div(
				Class("text-center mb-4 lg:mb-6"),
				H2(Class("text-xl md:text-2xl lg:text-3xl xl:text-4xl font-light mb-3 lg:mb-4 text-foreground"), g.Text(i18n.T(props.CurrentLang, "opinions"))),
				P(Class("text-xs md:text-sm lg:text-base xl:text-lg text-foreground font-light"), g.Text(i18n.T(props.CurrentLang, "opinions_subtitle"))),
			),
			components.Accordion("border border-foreground", faqItems),
		),
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
