"use client"

import { useRouter } from "next/navigation"
import { useTheme } from "next-themes"
import { useEffect, useRef, useState } from "react"
import { Lang } from "@/components/lang"
import { Logo } from "@/components/logo"
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from "@/components/ui/accordion"
import { Button } from "@/components/ui/button"
import { Carousel, type CarouselApi, CarouselContent, CarouselItem } from "@/components/ui/carousel"

export default function HomePage() {
  const [activeSection, setActiveSection] = useState(0)
  const [mobileApi, setMobileApi] = useState<CarouselApi>()
  const [mobileCurrent, setMobileCurrent] = useState(0)
  const scrollContainerRef = useRef<HTMLDivElement>(null)
  const router = useRouter()
  const { theme, resolvedTheme } = useTheme()
  const [mounted, setMounted] = useState(false)

  useEffect(() => {
    setMounted(true)
  }, [])

  // Mobile carousel API
  useEffect(() => {
    if (!mobileApi) return

    setMobileCurrent(mobileApi.selectedScrollSnap())

    mobileApi.on("select", () => {
      setMobileCurrent(mobileApi.selectedScrollSnap())
    })
  }, [mobileApi])

  useEffect(() => {
    const handleKeyDown = (event: KeyboardEvent) => {
      const sections = document.querySelectorAll(".snap-section")
      const currentSection = Array.from(sections).findIndex((section) => {
        const rect = section.getBoundingClientRect()
        return rect.top >= -100 && rect.top <= 100
      })

      if (event.key === "ArrowDown" && currentSection < sections.length - 1) {
        event.preventDefault()
        sections[currentSection + 1].scrollIntoView({ behavior: "smooth" })
      } else if (event.key === "ArrowUp" && currentSection > 0) {
        event.preventDefault()
        sections[currentSection - 1].scrollIntoView({ behavior: "smooth" })
      }
    }

    const handleScroll = () => {
      if (!scrollContainerRef.current) return

      const sections = document.querySelectorAll(".snap-section")
      const containerRect = scrollContainerRef.current.getBoundingClientRect()
      const containerCenter = containerRect.top + containerRect.height / 2

      let currentSection = 0
      let minDistance = Number.POSITIVE_INFINITY

      sections.forEach((section, index) => {
        const rect = section.getBoundingClientRect()
        const sectionCenter = rect.top + rect.height / 2
        const distance = Math.abs(sectionCenter - containerCenter)

        if (distance < minDistance) {
          minDistance = distance
          currentSection = index
        }
      })

      setActiveSection(currentSection)
    }

    let scrollTimeout: NodeJS.Timeout
    const throttledHandleScroll = () => {
      clearTimeout(scrollTimeout)
      scrollTimeout = setTimeout(handleScroll, 50)
    }

    const scrollContainer = scrollContainerRef.current

    window.addEventListener("keydown", handleKeyDown)
    if (scrollContainer) {
      scrollContainer.addEventListener("scroll", throttledHandleScroll, { passive: true })
    }

    setTimeout(handleScroll, 200)

    return () => {
      window.removeEventListener("keydown", handleKeyDown)
      if (scrollContainer) {
        scrollContainer.removeEventListener("scroll", throttledHandleScroll)
      }
      clearTimeout(scrollTimeout)
    }
  }, [])

  const scrollToSection = (index: number) => {
    const sections = document.querySelectorAll(".snap-section")
    if (sections[index]) {
      sections[index].scrollIntoView({ behavior: "smooth" })
    }
  }

  return (
    <>
      {/* Mobile Carousel Layout */}
      <div className="lg:hidden h-screen bg-foreground dark:bg-foreground text-black relative overflow-hidden">
        {/* Mode indicator in top right corner */}
        {mounted && (
          <div className="fixed top-4 right-4 z-50 bg-white border border-black px-2 py-1 text-xs font-light text-black">
            {theme === "system" ? `Auto (${resolvedTheme})` : theme}
          </div>
        )}

        {/* Carousel Container */}
        <Carousel
          className="h-full w-full mb-12"
          setApi={setMobileApi}
          opts={{
            align: "start",
            loop: false,
            dragFree: false,
          }}
        >
          <CarouselContent className="h-full -ml-0 flex">
            {/* Section 1: Main content */}
            <CarouselItem className="h-full pl-0">
              <div className="w-full h-full p-3 pb-16">
                <div className="w-full h-full flex flex-col justify-between bg-background rounded-xl p-4 overflow-hidden">
                  {/* Main content */}
                  <div className="text-center flex-1 flex flex-col justify-center">
                    <div className="flex items-center justify-center space-x-3 mb-4">
                      <Logo size="medium" autoStart={true} />
                      <span className="text-5xl font-thin text-foreground">4nuel</span>
                    </div>
                    <h1 className="text-lg font-light mb-3 text-foreground">Lead Software Engineer</h1>
                    <p className="text-xs text-foreground mb-4 font-light leading-relaxed px-2">
                      Specializing in AI/ML deep tech architecture. Proficient in C, Rust, Go, JS/TS, Python, and Zig.
                      Let's build cutting-edge solutions together.
                    </p>
                  </div>

                  {/* Footer at bottom */}
                  <div className="flex items-end justify-between w-full pb-4">
                    {/* Legal links on the left */}
                    <div className="flex flex-row gap-3">
                      <Button
                        type="button"
                        onClick={() => router.push("/imprint")}
                        className="text-foreground hover:text-foreground text-xs font-light cursor-pointer transition-all p-0 h-auto rounded-none shadow-none border-0 bg-transparent hover:bg-transparent hover:[text-shadow:_0_0_10px_rgba(0,0,0,0.3)] dark:hover:[text-shadow:_0_0_10px_rgba(255,255,255,0.5)] active:[text-shadow:_0_0_15px_rgba(0,0,0,0.5)] dark:active:[text-shadow:_0_0_15px_rgba(255,255,255,0.7)]"
                      >
                        Imprint
                      </Button>
                      <Button
                        type="button"
                        onClick={() => router.push("/gdpr")}
                        className="text-foreground hover:text-foreground text-xs font-light cursor-pointer transition-all p-0 h-auto rounded-none shadow-none border-0 bg-transparent hover:bg-transparent hover:[text-shadow:_0_0_10px_rgba(0,0,0,0.3)] dark:hover:[text-shadow:_0_0_10px_rgba(255,255,255,0.5)] active:[text-shadow:_0_0_15px_rgba(0,0,0,0.5)] dark:active:[text-shadow:_0_0_15px_rgba(255,255,255,0.7)]"
                      >
                        GDPR
                      </Button>
                      <Button
                        type="button"
                        onClick={() => router.push("/terms")}
                        className="text-foreground hover:text-foreground text-xs font-light cursor-pointer transition-all p-0 h-auto rounded-none shadow-none border-0 bg-transparent hover:bg-transparent hover:[text-shadow:_0_0_10px_rgba(0,0,0,0.3)] dark:hover:[text-shadow:_0_0_10px_rgba(255,255,255,0.5)] active:[text-shadow:_0_0_15px_rgba(0,0,0,0.5)] dark:active:[text-shadow:_0_0_15px_rgba(255,255,255,0.7)]"
                      >
                        Terms of Service
                      </Button>
                    </div>

                    {/* Theme and Language switchers on the right */}
                    <div className="flex flex-row gap-2">
                      <Lang type="mode" />
                      <Lang type="language" />
                    </div>
                  </div>
                </div>
              </div>
            </CarouselItem>

            {/* Section 2: Offer */}
            <CarouselItem className="h-full pl-0">
              <div className="w-full h-full p-3 pb-16">
                <div className="w-full h-full flex flex-col justify-center bg-background rounded-xl p-4 overflow-hidden">
                  <div className="max-w-2xl mx-auto text-center px-4">
                    <h2 className="text-lg font-light mb-3 text-foreground">Professional Cloud Browser Solutions</h2>
                    <p className="text-xs text-foreground mb-4 font-light leading-relaxed">
                      We provide cutting-edge cloud browser technology for businesses and developers. Scale your
                      operations with our reliable, secure, and high-performance browser infrastructure.
                    </p>
                    <Button
                      type="button"
                      className="bg-foreground text-background border-0 outline-none cursor-pointer text-xs px-3 py-2 font-light transition-colors h-auto rounded-none shadow-none hover:bg-foreground"
                      onMouseEnter={(e) => {
                        e.currentTarget.classList.add("opacity-80")
                      }}
                      onMouseLeave={(e) => {
                        e.currentTarget.classList.remove("opacity-80")
                      }}
                    >
                      Get in Touch
                    </Button>
                  </div>
                </div>
              </div>
            </CarouselItem>

            {/* Section 3: Pricing */}
            <CarouselItem className="h-full pl-0">
              <div className="w-full h-full p-3 pb-16">
                <div className="w-full h-full flex flex-col justify-center bg-background rounded-xl p-4 overflow-hidden">
                  <div className="w-full px-4">
                    <div className="text-center mb-6">
                      <h2 className="text-lg font-light mb-2 text-foreground">Pricing</h2>
                      <p className="text-xs text-foreground max-w-xl mx-auto font-light">
                        Choose the plan that best fits your business needs
                      </p>
                    </div>
                    <div className="grid grid-cols-1 gap-4 max-w-md mx-auto">
                      <div className="p-4 text-center border border-foreground">
                        <h3 className="text-lg font-light mb-2 text-foreground">Hourly</h3>
                        <div className="mb-2">
                          <span className="text-xl font-light text-foreground">$100</span>
                          <span className="text-xs text-foreground">/hour</span>
                        </div>
                        <p className="text-xs font-light text-foreground">Perfect for quick tasks</p>
                      </div>
                      <div className="border border-foreground p-4 text-center bg-foreground text-background">
                        <h3 className="text-lg font-light mb-2">Daily</h3>
                        <div className="mb-2">
                          <span className="text-xl font-light">$800</span>
                          <span className="text-xs">/day</span>
                        </div>
                        <p className="text-xs font-light">Ideal for ongoing projects</p>
                      </div>
                      <div className="p-4 text-center border border-foreground">
                        <h3 className="text-lg font-light mb-2 text-foreground">Startups</h3>
                        <div className="mb-2">
                          <span className="text-lg font-light text-foreground">Co-founder</span>
                        </div>
                        <p className="text-xs font-light text-foreground">Equity participation</p>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </CarouselItem>

            {/* Section 4: Q&A */}
            <CarouselItem className="h-full pl-0">
              <div className="w-full h-full p-3 pb-16">
                <div className="w-full h-full flex flex-col justify-center bg-background rounded-xl p-4 overflow-y-auto">
                  <div className="w-full px-4 max-w-2xl mx-auto">
                    <div className="text-center mb-4">
                      <h2 className="text-lg font-light mb-2 text-foreground">Questions</h2>
                      <p className="text-xs text-foreground font-light">
                        Everything you need to know about working together
                      </p>
                    </div>
                    <Accordion type="single" collapsible className="border border-foreground">
                      {[
                        {
                          question: "What types of services do I offer?",
                          answer:
                            "I specialize in full-stack web development, API design and implementation, cloud architecture, database optimization, and custom software solutions.",
                        },
                        {
                          question: "How do I approach new projects?",
                          answer:
                            "I begin with a thorough discovery phase to understand your requirements, followed by a detailed project proposal with clear milestones.",
                        },
                        {
                          question: "What is my pricing structure?",
                          answer:
                            "I offer flexible pricing models including hourly rates for short-term work, daily rates for ongoing projects, and equity-based partnerships for startups.",
                        },
                        {
                          question: "What technologies do I work with?",
                          answer:
                            "I'm proficient in modern web technologies including React, Next.js, Node.js, TypeScript, Python, and various databases.",
                        },
                      ].map((faq, index) => (
                        <AccordionItem
                          key={index}
                          value={`item-${index}`}
                          className="px-4 border-b border-foreground last:border-b-0"
                        >
                          <AccordionTrigger className="text-left text-sm font-light hover:text-foreground py-3 hover:no-underline cursor-pointer text-foreground">
                            {faq.question}
                          </AccordionTrigger>
                          <AccordionContent className="text-foreground pb-3 text-xs leading-relaxed font-light">
                            {faq.answer}
                          </AccordionContent>
                        </AccordionItem>
                      ))}
                    </Accordion>
                  </div>
                </div>
              </div>
            </CarouselItem>
          </CarouselContent>
        </Carousel>

        {/* Section Indicators */}
        <div className="absolute bottom-6 left-1/2 -translate-x-1/2 z-50 flex gap-1">
          {[0, 1, 2, 3].map((index) => (
            <button
              key={index}
              onClick={() => mobileApi?.scrollTo(index)}
              className={`w-3 h-3 transition-all duration-300 touch-manipulation cursor-pointer border rounded-none shadow-none ${
                mobileCurrent === index
                  ? "bg-white border-white dark:bg-black dark:border-black shadow-lg w-10 h-3"
                  : "bg-transparent border-white dark:border-black"
              }`}
              aria-label={`Go to section ${index + 1}`}
            />
          ))}
        </div>
      </div>

      {/* Desktop Layout */}
      <div className="hidden lg:grid min-h-screen bg-foreground dark:bg-foreground text-black grid-cols-5 selectable-content relative">
        {/* Mode indicator in top right corner */}
        {mounted && (
          <div className="fixed top-4 right-4 z-50 bg-white border border-black px-3 py-1 text-xs font-light text-black">
            {theme === "system" ? `Auto (${resolvedTheme})` : theme}
          </div>
        )}

        {/* Left Column - Fixed - 2/5 width on desktop */}
        <div className="h-screen p-6 lg:pl-10 lg:pr-3 col-span-2 bg-foreground dark:bg-foreground">
          <div className="p-8 w-full h-full flex flex-col justify-between bg-background rounded-2xl">
            {/* Main content */}
            <div className="text-center flex-1 flex flex-col justify-center">
              <div className="flex items-center justify-center space-x-4 mb-8">
                <Logo size="medium" autoStart={true} />
                <span className="text-4xl md:text-6xl lg:text-8xl xl:text-9xl font-thin text-foreground">4nuel</span>
              </div>
              <h1 className="text-xl md:text-2xl lg:text-3xl xl:text-4xl font-light mb-4 lg:mb-6 text-foreground">
                Lead Software Engineer
              </h1>
              <p className="text-xs md:text-sm lg:text-base xl:text-lg text-foreground mb-6 lg:mb-8 font-light leading-relaxed">
                Specializing in AI/ML deep tech architecture. Proficient in C, Rust, Go, JS/TS, Python, and Zig. Let's
                build cutting-edge solutions together.
              </p>
            </div>

            {/* Footer at bottom */}
            <div className="flex items-end justify-between w-full pb-0">
              {/* Legal links on the left */}
              <div className="flex flex-row gap-2 sm:gap-4 md:gap-6 lg:gap-8">
                <Button
                  type="button"
                  onClick={() => router.push("/imprint")}
                  className="text-foreground hover:text-foreground text-xs sm:text-sm font-light cursor-pointer transition-all p-0 h-auto rounded-none shadow-none border-0 bg-transparent hover:bg-transparent hover:[text-shadow:_0_0_10px_rgba(0,0,0,0.3)] dark:hover:[text-shadow:_0_0_10px_rgba(255,255,255,0.5)] active:[text-shadow:_0_0_15px_rgba(0,0,0,0.5)] dark:active:[text-shadow:_0_0_15px_rgba(255,255,255,0.7)]"
                >
                  Imprint
                </Button>
                <Button
                  type="button"
                  onClick={() => router.push("/gdpr")}
                  className="text-foreground hover:text-foreground text-xs sm:text-sm font-light cursor-pointer transition-all p-0 h-auto rounded-none shadow-none border-0 bg-transparent hover:bg-transparent hover:[text-shadow:_0_0_10px_rgba(0,0,0,0.3)] dark:hover:[text-shadow:_0_0_10px_rgba(255,255,255,0.5)] active:[text-shadow:_0_0_15px_rgba(0,0,0,0.5)] dark:active:[text-shadow:_0_0_15px_rgba(255,255,255,0.7)]"
                >
                  GDPR
                </Button>
                <Button
                  type="button"
                  onClick={() => router.push("/terms")}
                  className="text-foreground hover:text-foreground text-xs sm:text-sm font-light cursor-pointer transition-all p-0 h-auto rounded-none shadow-none border-0 bg-transparent hover:bg-transparent hover:[text-shadow:_0_0_10px_rgba(0,0,0,0.3)] dark:hover:[text-shadow:_0_0_10px_rgba(255,255,255,0.5)] active:[text-shadow:_0_0_15px_rgba(0,0,0,0.5)] dark:active:[text-shadow:_0_0_15px_rgba(255,255,255,0.7)]"
                >
                  Terms of Service
                </Button>
              </div>

              {/* Theme and Language switchers on the right */}
              <div className="flex flex-row gap-1 sm:gap-2 md:gap-3 lg:gap-4">
                <Lang type="mode" />
                <Lang type="language" />
              </div>
            </div>
          </div>
        </div>

        {/* Right Column - Scrollable - 3/5 width on desktop */}
        <div className="h-screen relative col-span-3 p-6 lg:pl-3 lg:pr-10 bg-foreground dark:bg-foreground">
          <div className="h-full bg-background rounded-2xl overflow-hidden">
            <div
              ref={scrollContainerRef}
              className="h-full overflow-y-scroll snap-y snap-mandatory scrollbar-hide bg-background rounded-2xl"
            >
              <main className="w-full bg-background">
                <div
                  id="offer"
                  className="snap-section snap-start h-screen flex items-center justify-center bg-background"
                >
                  <div className="max-w-2xl mx-auto text-center px-4 lg:px-8">
                    <h2 className="text-xl md:text-2xl lg:text-3xl xl:text-4xl font-light mb-4 lg:mb-6 text-foreground">
                      Professional Cloud Browser Solutions
                    </h2>
                    <p className="text-xs md:text-sm lg:text-base xl:text-lg text-foreground mb-6 lg:mb-8 font-light leading-relaxed">
                      We provide cutting-edge cloud browser technology for businesses and developers. Scale your
                      operations with our reliable, secure, and high-performance browser infrastructure.
                    </p>
                    <Button
                      type="button"
                      className="bg-foreground text-background border-0 outline-none cursor-pointer text-xs md:text-sm lg:text-base px-4 lg:px-6 py-2 lg:py-3 font-light transition-colors h-auto rounded-none shadow-none hover:bg-foreground"
                      onMouseEnter={(e) => {
                        e.currentTarget.classList.add("opacity-80")
                      }}
                      onMouseLeave={(e) => {
                        e.currentTarget.classList.remove("opacity-80")
                      }}
                    >
                      Get in Touch
                    </Button>
                  </div>
                </div>

                <div
                  id="pricing"
                  className="snap-section snap-start h-screen flex items-center justify-center bg-background"
                >
                  <div className="w-full px-4 lg:px-8">
                    <div className="text-center mb-6 lg:mb-8">
                      <h2 className="text-xl md:text-2xl lg:text-3xl xl:text-4xl font-light mb-3 lg:mb-4 text-foreground">
                        Pricing
                      </h2>
                      <p className="text-xs md:text-sm lg:text-base xl:text-lg text-foreground max-w-xl mx-auto font-light">
                        Choose the plan that best fits your business needs
                      </p>
                    </div>
                    <div className="grid grid-cols-1 md:grid-cols-3 gap-4 max-w-4xl mx-auto">
                      <div className="p-6 text-center border border-foreground">
                        <h3 className="text-xl font-light mb-2 text-foreground">Hourly</h3>
                        <div className="mb-4">
                          <span className="text-2xl font-light text-foreground">$100</span>
                          <span className="text-sm text-foreground">/hour</span>
                        </div>
                        <p className="text-sm font-light text-foreground">Perfect for quick tasks</p>
                      </div>
                      <div className="border border-foreground p-6 text-center bg-foreground text-background">
                        <h3 className="text-xl font-light mb-2">Daily</h3>
                        <div className="mb-4">
                          <span className="text-2xl font-light">$800</span>
                          <span className="text-sm">/day</span>
                        </div>
                        <p className="text-sm font-light">Ideal for ongoing projects</p>
                      </div>
                      <div className="p-6 text-center border border-foreground">
                        <h3 className="text-xl font-light mb-2 text-foreground">Startups</h3>
                        <div className="mb-4">
                          <span className="text-xl font-light text-foreground">Co-founder</span>
                        </div>
                        <p className="text-sm font-light text-foreground">Equity participation</p>
                      </div>
                    </div>
                  </div>
                </div>

                <div
                  id="qa"
                  className="snap-section snap-start h-screen flex items-center justify-center bg-background"
                >
                  <div className="w-full px-4 lg:px-8 max-w-2xl mx-auto">
                    <div className="text-center mb-4 lg:mb-6">
                      <h2 className="text-xl md:text-2xl lg:text-3xl xl:text-4xl font-light mb-3 lg:mb-4 text-foreground">
                        Questions
                      </h2>
                      <p className="text-xs md:text-sm lg:text-base xl:text-lg text-foreground font-light">
                        Everything you need to know about working together
                      </p>
                    </div>
                    <Accordion type="single" collapsible className="border border-foreground">
                      {[
                        {
                          question: "What types of services do I offer?",
                          answer:
                            "I specialize in full-stack web development, API design and implementation, cloud architecture, database optimization, and custom software solutions.",
                        },
                        {
                          question: "How do I approach new projects?",
                          answer:
                            "I begin with a thorough discovery phase to understand your requirements, followed by a detailed project proposal with clear milestones.",
                        },
                        {
                          question: "What is my pricing structure?",
                          answer:
                            "I offer flexible pricing models including hourly rates for short-term work, daily rates for ongoing projects, and equity-based partnerships for startups.",
                        },
                        {
                          question: "What technologies do I work with?",
                          answer:
                            "I'm proficient in modern web technologies including React, Next.js, Node.js, TypeScript, Python, and various databases.",
                        },
                      ].map((faq, index, array) => (
                        <AccordionItem
                          key={index}
                          value={`item-${index}`}
                          className="px-4 border-b border-foreground last:border-b-0"
                        >
                          <AccordionTrigger className="text-left text-sm md:text-base font-light hover:text-foreground py-3 hover:no-underline cursor-pointer text-foreground">
                            {faq.question}
                          </AccordionTrigger>
                          <AccordionContent className="text-foreground pb-3 text-xs md:text-sm leading-relaxed font-light">
                            {faq.answer}
                          </AccordionContent>
                        </AccordionItem>
                      ))}
                    </Accordion>
                  </div>
                </div>
              </main>
            </div>

            {/* Section Navigation - Right Column - Hidden on mobile */}
            <div className="absolute right-2 lg:right-4 top-1/2 -translate-y-1/2 z-40 flex-col gap-1 hidden lg:flex">
              {[0, 1, 2].map((index) => (
                <button
                  type="button"
                  key={index}
                  onClick={() => scrollToSection(index)}
                  className={`w-3 h-3 transition-all duration-300 cursor-pointer border rounded-none shadow-none ${
                    activeSection === index
                      ? "bg-white border-white dark:bg-black dark:border-black shadow-lg w-3 h-10"
                      : "bg-transparent border-white dark:border-black hover:bg-transparent"
                  }`}
                  aria-label={`Go to section ${index + 1}`}
                />
              ))}
            </div>
          </div>
        </div>
      </div>
    </>
  )
}
