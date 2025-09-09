"use client"

import { AnimatedLogo } from "@/components/animated-logo"
import { LanguageSelector } from "@/components/language-selector"
import { ThemeSelector } from "@/components/theme-selector"

export function Footer() {
  return (
    <div className="h-screen bg-black text-white relative overflow-hidden">
      {/* Legal links and settings positioned directly over the logo and text */}
      <div className="absolute bottom-0 left-0 right-0 w-full px-8 sm:px-12 md:px-16 lg:px-20 z-20">
        <div className="max-w-7xl mx-auto relative">
          {/* Legal links positioned all the way to the left */}
          <div className="absolute -top-16 sm:-top-12 left-0 flex flex-row gap-3 sm:gap-8 md:gap-16 lg:gap-24">
            <a
              href="#"
              className="text-gray-400 hover:text-white text-xs sm:text-sm font-light cursor-pointer transition-colors"
            >
              Imprint
            </a>
            <a
              href="#"
              className="text-gray-400 hover:text-white text-xs sm:text-sm font-light cursor-pointer transition-colors"
            >
              GDPR
            </a>
            <a
              href="#"
              className="text-gray-400 hover:text-white text-xs sm:text-sm font-light cursor-pointer transition-colors"
            >
              Terms of Service
            </a>
          </div>

          {/* Theme and Language switchers positioned all the way to the right */}
          <div className="absolute -top-16 sm:-top-12 right-0 flex flex-row gap-3 sm:gap-4 md:gap-8 lg:gap-12">
            {/* Theme selector */}
            <ThemeSelector />

            {/* Language switcher */}
            <LanguageSelector />
          </div>

          {/* Centered logo and text container positioned at the very bottom */}
          <div className="flex items-end justify-center space-x-1 md:space-x-2 lg:space-x-3 pb-0">
            {/* Logo with matrix animation */}
            <AnimatedLogo size="medium" autoStart={false} />

            <span className="text-white font-thin text-6xl sm:text-7xl md:text-8xl lg:text-9xl xl:text-[12rem] leading-none">
              4nuel
            </span>
          </div>
        </div>
      </div>
    </div>
  )
}
