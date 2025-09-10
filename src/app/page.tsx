"use client"

import { useEffect, useRef, useState } from "react"
import { Footer } from "@/components/footer"
import { Header } from "@/components/header"
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from "@/components/ui/accordion"
import { Mask } from "@/components/ui/animations/mask"
import { Card, CardHeader } from "@/components/ui/card"
import { Structure } from "@/components/structure"
import { Button } from "@/components/ui/button"
import {
  Carousel,
  CarouselContent,
  CarouselItem,
  CarouselApi,
} from "@/components/ui/carousel"

export default function HomePage() {
  const [activeSection, setActiveSection] = useState(0)
  const [api, setApi] = useState<CarouselApi>()
  const [current, setCurrent] = useState(0)
  const [count, setCount] = useState(0)
  const scrollContainerRef = useRef<HTMLDivElement>(null)

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

  useEffect(() => {
    if (!api) return

    setCount(api.scrollSnapList().length)
    setCurrent(api.selectedScrollSnap() + 1)

    api.on("select", () => {
      setCurrent(api.selectedScrollSnap() + 1)
    })
  }, [api])

  const scrollToSection = (index: number) => {
    const sections = document.querySelectorAll(".snap-section")
    if (sections[index]) {
      sections[index].scrollIntoView({ behavior: "smooth" })
    }
  }

  return (
    <div className="min-h-screen bg-background text-foreground">
      <Header activeSection={activeSection} />

      <div ref={scrollContainerRef} className="h-screen overflow-y-scroll snap-y snap-mandatory scrollbar-hide">
        <main className="w-full">
          <div id="offer" className="snap-section snap-start">
            <Structure>
              <div className="relative">
                <Mask
                  revealText={
                    <div className="max-w-4xl mx-auto text-center pt-32 px-4">
                      <h1 className="text-3xl sm:text-4xl md:text-6xl font-light mb-8 text-balance text-foreground">
                        Professional Cloud Browser Solutions
                      </h1>
                      <p className="text-lg sm:text-xl text-foreground mb-12 max-w-2xl mx-auto text-pretty font-light">
                        We provide cutting-edge cloud browser technology for businesses and developers. Scale your
                        operations with our reliable, secure, and high-performance browser infrastructure.
                      </p>
                      <div className="flex justify-center">
                        <Button
                          type="button"
                          className="bg-foreground text-background border-0 outline-none cursor-pointer text-base sm:text-lg px-6 sm:px-8 py-3 sm:py-4 font-light transition-colors h-auto rounded-none shadow-none hover:bg-foreground"
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
                  }
                >
                  <div className="max-w-4xl mx-auto text-center pt-32 px-4">
                    <h1 className="text-3xl sm:text-4xl md:text-6xl font-light mb-8 text-balance text-background">
                      Advanced Browser Infrastructure
                    </h1>
                    <p className="text-lg sm:text-xl text-background mb-12 max-w-2xl mx-auto text-pretty font-light">
                      Experience the future of cloud computing with our revolutionary browser technology that adapts to
                      your needs.
                    </p>
                    <div className="flex justify-center">
                      <Button
                        type="button"
                        className="bg-background text-foreground border-0 outline-none cursor-pointer text-base sm:text-lg px-6 sm:px-8 py-3 sm:py-4 font-light transition-colors h-auto rounded-none shadow-none hover:bg-background"
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
                </Mask>
              </div>
            </Structure>
          </div>

          <div id="pricing" className="snap-section snap-start">
            <Structure hasNav>
              <section className="h-full bg-background text-foreground flex flex-col justify-center px-6 py-8">
                <div className="max-w-6xl w-full mx-auto">
                  <div className="text-center mb-8 md:mb-12">
                    <h2 className="text-3xl md:text-5xl font-light mb-4 md:mb-6">Pricing</h2>
                    <p className="text-sm md:text-xl text-foreground max-w-2xl mx-auto font-light">
                      Choose the plan that best fits your business needs
                    </p>
                  </div>
                  <Carousel 
                    className="flex-1 md:hidden"
                    setApi={setApi}
                    opts={{
                      align: "center",
                      loop: false,
                      dragFree: false,
                      containScroll: "trimSnaps",
                      slidesToScroll: 1,
                    }}
                  >
                    <CarouselContent className="-ml-2 md:-ml-4">
                      {[
                      {
                        name: "Hourly",
                        price: "$100",
                        period: "/hour",
                        description: "Perfect for quick tasks and testing",
                        features: ["Pay as you go", "Instant setup", "Basic support", "No commitments"],
                      },
                      {
                        name: "Daily",
                        price: "$800",
                        period: "/day",
                        description: "Ideal for ongoing projects",
                        features: [
                          "Full day access",
                          "Priority support",
                          "Advanced features",
                          "24/7 availability",
                          "Custom configurations",
                        ],
                        popular: true,
                      },
                      {
                        name: "Startups",
                        price: "Co-founder",
                        period: "",
                        description: "Equity participation and ownership stake",
                        features: [
                          "Equity-based partnership",
                          "Active participation",
                          "Long-term commitment",
                          "Shared ownership",
                          "Strategic involvement",
                        ],
                      },
                      ].map((plan, index) => (
                        <CarouselItem key={index} className="pl-2 md:pl-4">
                          <div className="p-1">
                            <Card
                          className={`relative w-full max-w-sm aspect-square flex flex-col justify-center rounded-none ${
                            plan.popular 
                              ? "bg-foreground border-2 border-foreground shadow-2xl transform scale-105" 
                              : "bg-background border-2 border-foreground"
                          }`}
                        >
                          {plan.popular && (
                            <div className="absolute top-0 left-1/2 transform -translate-x-1/2 z-10">
                              <div className="bg-background text-foreground px-6 py-2 text-sm font-medium border border-foreground">
                                MOST POPULAR
                              </div>
                            </div>
                          )}
                          <CardHeader className="text-center pb-4 flex-1 flex flex-col justify-center pt-12">
                            <h3 className={`text-2xl font-light mb-2 ${
                              plan.popular ? "text-background" : "text-foreground"
                            }`}>{plan.name}</h3>
                            <div className="mb-4">
                              <span className={`text-4xl font-light ${
                                plan.popular ? "text-background" : "text-foreground"
                              }`}>{plan.price}</span>
                              <span className={`text-lg font-light ${
                                plan.popular ? "text-background" : "text-foreground"
                              }`}>{plan.period}</span>
                            </div>
                            <p className={`font-light text-sm mb-6 ${
                              plan.popular ? "text-background" : "text-foreground"
                            }`}>{plan.description}</p>
                            <ul className="space-y-2">
                              {plan.features.map((feature, featureIndex) => (
                                <li
                                  key={featureIndex}
                                  className={`flex items-start font-light text-sm justify-center ${
                                    plan.popular ? "text-background" : "text-foreground"
                                  }`}
                                >
                                  <div className={`w-1 h-1 mt-2 mr-3 flex-shrink-0 ${
                                    plan.popular ? "bg-background" : "bg-foreground"
                                  }`}></div>
                                  {feature}
                                </li>
                              ))}
                            </ul>
                            </CardHeader>
                          </Card>
                        </div>
                      </CarouselItem>
                    ))}
                  </CarouselContent>
                </Carousel>
                
                <div className="flex justify-center mt-4 gap-2 md:hidden">
                  {Array.from({ length: count }, (_, index) => (
                    <Button
                      key={index}
                      type="button"
                      onClick={() => api?.scrollTo(index)}
                      className={`w-2 h-2 transition-all p-0 rounded-none shadow-none border-0 ${
                        index + 1 === current
                          ? "bg-foreground"
                          : "bg-foreground opacity-30 hover:opacity-60"
                      }`}
                      variant="ghost"
                      size="icon"
                      aria-label={`Go to slide ${index + 1}`}
                    />
                  ))}
                </div>
                  
                <div className="hidden md:grid md:grid-cols-3 gap-8">
                    {[
                      {
                        name: "Hourly",
                        price: "$100",
                        period: "/hour",
                        description: "Perfect for quick tasks and testing",
                        features: ["Pay as you go", "Instant setup", "Basic support", "No commitments"],
                      },
                      {
                        name: "Daily",
                        price: "$800",
                        period: "/day",
                        description: "Ideal for ongoing projects",
                        features: [
                          "Full day access",
                          "Priority support",
                          "Advanced features",
                          "24/7 availability",
                          "Custom configurations",
                        ],
                        popular: true,
                      },
                      {
                        name: "Startups",
                        price: "Co-founder",
                        period: "",
                        description: "Equity participation and ownership stake",
                        features: [
                          "Equity-based partnership",
                          "Active participation",
                          "Long-term commitment",
                          "Shared ownership",
                          "Strategic involvement",
                        ],
                      },
                    ].map((plan, index) => (
                      <Card
                        key={index}
                        className={`relative aspect-square flex flex-col justify-center rounded-none ${
                          plan.popular 
                            ? "bg-foreground border-2 border-foreground shadow-2xl transform scale-105" 
                            : "bg-background border-2 border-foreground"
                        }`}
                      >
                        {plan.popular && (
                          <div className="absolute -top-4 left-1/2 transform -translate-x-1/2 z-10">
                            <div className="bg-background text-foreground px-6 py-2 text-sm font-medium border border-foreground">
                              MOST POPULAR
                            </div>
                          </div>
                        )}
                        <CardHeader className="text-center pb-4 flex-1 flex flex-col justify-center">
                          <h3 className={`text-2xl font-light mb-2 ${
                            plan.popular ? "text-background" : "text-foreground"
                          }`}>{plan.name}</h3>
                          <div className="mb-4">
                            <span className={`text-4xl font-light ${
                              plan.popular ? "text-background" : "text-foreground"
                            }`}>{plan.price}</span>
                            <span className={`text-lg font-light ${
                              plan.popular ? "text-background" : "text-foreground"
                            }`}>{plan.period}</span>
                          </div>
                          <p className={`font-light text-sm mb-6 ${
                            plan.popular ? "text-background" : "text-foreground"
                          }`}>{plan.description}</p>
                          <ul className="space-y-2">
                            {plan.features.map((feature, featureIndex) => (
                              <li
                                key={featureIndex}
                                className={`flex items-start font-light text-sm justify-center ${
                                  plan.popular ? "text-background" : "text-foreground"
                                }`}
                              >
                                <div className={`w-1 h-1 mt-2 mr-3 flex-shrink-0 ${
                                  plan.popular ? "bg-background" : "bg-foreground"
                                }`}></div>
                                {feature}
                              </li>
                            ))}
                          </ul>
                        </CardHeader>
                      </Card>
                    ))}
                  </div>
                </div>
              </section>
            </Structure>
          </div>

          <div id="qa" className="snap-section snap-start">
            <Structure hasNav>
              <div className="w-full bg-background text-foreground" style={{ paddingBottom: "50px" }}>
                <div className="w-full max-w-4xl mx-auto px-6">
                  <div className="text-center mb-8">
                    <h2 className="text-3xl md:text-4xl font-light mb-4">Questions</h2>
                    <p className="text-sm md:text-base text-foreground max-w-2xl mx-auto font-light">
                      Everything you need to know about working together
                    </p>
                  </div>
                  <div className="w-full">
                    <Accordion type="single" collapsible className="border border-foreground">
                      {[
                        {
                          question: "What types of services do I offer?",
                          answer:
                            "I specialize in full-stack web development, API design and implementation, cloud architecture, database optimization, and custom software solutions. Whether you need a new application built from scratch, legacy code modernization, or technical consulting, I can help bring your project to life.",
                        },
                        {
                          question: "How do I approach new projects?",
                          answer:
                            "I begin with a thorough discovery phase to understand your requirements, followed by a detailed project proposal with clear milestones. I work in iterative sprints with regular check-ins, ensuring transparency and flexibility throughout the development process. Timelines are estimated based on project complexity and agreed upon before starting.",
                        },
                        {
                          question: "What is my pricing structure?",
                          answer:
                            "I offer flexible pricing models including hourly rates for short-term work, daily rates for ongoing projects, and equity-based partnerships for startups. Each project is unique, and I provide customized quotes based on scope, complexity, and timeline. Contact me for a detailed estimate tailored to your specific needs.",
                        },
                        {
                          question: "What technologies do I work with?",
                          answer:
                            "I'm proficient in modern web technologies including React, Next.js, Node.js, TypeScript, Python, and various databases. I also have experience with cloud platforms like AWS and GCP, containerization with Docker, and CI/CD pipelines. I stay current with industry trends and can adapt to your existing tech stack.",
                        },
                        {
                          question: "How do I ensure code quality?",
                          answer:
                            "I follow industry best practices including test-driven development, code reviews, and comprehensive documentation. All code is version-controlled, well-commented, and built with scalability in mind. I maintain clear communication throughout the project and provide post-launch support to ensure smooth deployment.",
                        },
                      ].map((faq, index, array) => (
                        <AccordionItem key={index} value={`item-${index}`} className={`px-4 ${index < array.length - 1 ? 'border-b border-foreground' : ''}`}>
                          <AccordionTrigger className="text-left text-sm md:text-base font-light hover:text-foreground py-3 hover:no-underline cursor-pointer">
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
              </div>
            </Structure>
          </div>

          <div id="footer" className="snap-section snap-start">
            <Footer />
          </div>
        </main>
      </div>

      <div className="fixed right-2 sm:right-3 lg:right-4 top-1/2 -translate-y-1/2 z-40 md:hidden flex flex-col gap-2">
        {[0, 1, 2, 3].map((index) => (
          <Button
            type="button"
            key={index}
            onClick={() => scrollToSection(index)}
            className={`w-2 sm:w-3 h-8 sm:h-10 lg:h-12 transition-all duration-300 cursor-pointer border p-0 rounded-none shadow-none ${
              activeSection === index
                ? "bg-foreground border-foreground shadow-lg"
                : "bg-transparent border-foreground hover:border-foreground hover:bg-muted"
            }`}
            variant="ghost"
            aria-label={`Go to section ${index + 1}`}
          />
        ))}
      </div>
    </div>
  )
}
