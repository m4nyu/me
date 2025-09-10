"use client"

import { Logo } from "@/components/logo"
import { Button } from "@/components/ui/button"

interface HeaderProps {
  activeSection?: number
}

export function Header({ activeSection = 0 }: HeaderProps) {
  const scrollToSection = (sectionId: string) => {
    const section = document.getElementById(sectionId)
    if (section) {
      section.scrollIntoView({ behavior: "smooth" })
    }
  }

  const getActiveNavItem = (sectionId: string) => {
    const sectionMap: Record<number, string> = {
      0: "offer",
      1: "pricing",
      2: "qa",
    }
    return sectionMap[activeSection] === sectionId
  }

  return (
    <header className="fixed top-0 left-0 right-0 z-50 p-2 md:p-6">
      <nav className="bg-background backdrop-blur-md shadow-lg border-4 border-foreground hidden md:block max-w-7xl mx-auto">
        <div className="flex items-center px-8 py-6 w-full">
          <div className="flex items-center space-x-1">
            <Logo size="small" autoStart={true} />
            <span className="text-foreground font-light text-xl">4nuel</span>
          </div>

          <div className="flex-1 flex justify-center">
            <div className="flex items-center space-x-8">
              <Button
                type="button"
                onClick={() => scrollToSection("offer")}
                className={`${
                  getActiveNavItem("offer")
                    ? "text-foreground font-bold"
                    : "text-foreground hover:text-foreground font-light"
                } transition-all text-base cursor-pointer p-0 h-auto rounded-none shadow-none border-0 bg-transparent hover:bg-transparent`}
              >
                Offer
              </Button>
              <Button
                type="button"
                onClick={() => scrollToSection("pricing")}
                className={`${
                  getActiveNavItem("pricing")
                    ? "text-foreground font-bold"
                    : "text-foreground hover:text-foreground font-light"
                } transition-all text-base cursor-pointer p-0 h-auto rounded-none shadow-none border-0 bg-transparent hover:bg-transparent`}
              >
                Pricing
              </Button>
              <Button
                type="button"
                onClick={() => scrollToSection("qa")}
                className={`${
                  getActiveNavItem("qa")
                    ? "text-foreground font-bold"
                    : "text-foreground hover:text-foreground font-light"
                } transition-all text-base cursor-pointer p-0 h-auto rounded-none shadow-none border-0 bg-transparent hover:bg-transparent`}
              >
                Q&A
              </Button>
            </div>
          </div>

          <div className="flex items-center space-x-3">
            <Button
              type="button"
              className="bg-foreground text-background border-0 outline-none cursor-pointer text-base px-6 py-3 font-light transition-colors h-auto rounded-none shadow-none hover:bg-foreground"
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
      </nav>

      <nav className="bg-background backdrop-blur-md shadow-lg border-2 border-foreground md:hidden max-w-full mx-auto">
        <div className="flex items-center justify-between px-6 py-4 w-full">
          <div className="flex items-center space-x-1">
            <Logo size="small" autoStart={true} />
            <span className="text-foreground font-light text-base">4nuel</span>
          </div>

          <Button
            type="button"
            className="bg-foreground text-background border-0 outline-none cursor-pointer text-sm px-4 py-2 font-light transition-colors h-auto rounded-none shadow-none hover:bg-foreground"
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
      </nav>
    </header>
  )
}
