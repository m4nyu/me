"use client"

import { Logo } from "@/components/logo"
import { Picker } from "@/components/picker"

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
      0: 'offer',
      1: 'pricing', 
      2: 'qa'
    }
    return sectionMap[activeSection] === sectionId
  }

  return (
    <header className="fixed top-0 left-0 right-0 z-50 p-2 md:p-6">
      <nav className="bg-black backdrop-blur-md shadow-lg border-4 border-white hidden md:block max-w-7xl mx-auto">
        <div className="flex items-center px-8 py-6 w-full">
          <div className="flex items-center space-x-1">
            <Logo size="small" autoStart={true} />
            <span className="text-white font-light text-xl">4nuel</span>
          </div>

          <div className="flex-1 flex justify-center">
            <div className="flex items-center space-x-8">
              <button
                onClick={() => scrollToSection("offer")}
                className={`${
                  getActiveNavItem("offer") ? "text-white font-bold" : "text-gray-300 hover:text-white font-light"
                } transition-all text-base cursor-pointer`}
              >
                Offer
              </button>
              <button
                onClick={() => scrollToSection("pricing")}
                className={`${
                  getActiveNavItem("pricing") ? "text-white font-bold" : "text-gray-300 hover:text-white font-light"
                } transition-all text-base cursor-pointer`}
              >
                Pricing
              </button>
              <button
                onClick={() => scrollToSection("qa")}
                className={`${
                  getActiveNavItem("qa") ? "text-white font-bold" : "text-gray-300 hover:text-white font-light"
                } transition-all text-base cursor-pointer`}
              >
                Q&A
              </button>
            </div>
          </div>

          <div className="flex items-center space-x-3">
            <button
              className="cursor-pointer text-base px-6 py-3 font-light transition-colors"
              style={{
                backgroundColor: "#ffffff",
                color: "#000000",
                border: "none",
                outline: "none",
              }}
              onMouseEnter={(e) => {
                e.currentTarget.style.backgroundColor = "#f3f4f6"
              }}
              onMouseLeave={(e) => {
                e.currentTarget.style.backgroundColor = "#ffffff"
              }}
            >
              Get in Touch
            </button>
          </div>
        </div>
      </nav>

      <nav className="bg-black backdrop-blur-md shadow-lg border-2 border-white md:hidden max-w-full mx-auto">
        <div className="flex items-center justify-between px-6 py-4 w-full">
          <div className="flex items-center space-x-1">
            <Logo size="small" autoStart={true} />
            <span className="text-white font-light text-base">4nuel</span>
          </div>

          <button
            className="cursor-pointer text-sm px-4 py-2 font-light transition-colors"
            style={{
              backgroundColor: "#ffffff",
              color: "#000000",
              border: "none",
              outline: "none",
            }}
            onMouseEnter={(e) => {
              e.currentTarget.style.backgroundColor = "#f3f4f6"
            }}
            onMouseLeave={(e) => {
              e.currentTarget.style.backgroundColor = "#ffffff"
            }}
          >
            Get in Touch
          </button>
        </div>
      </nav>
    </header>
  )
}
