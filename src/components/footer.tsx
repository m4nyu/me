"use client"

import { useRouter } from "next/navigation"
import { Lang } from "@/components/lang"
import { Logo } from "@/components/logo"
import { Button } from "@/components/ui/button"

export function Footer() {
  const router = useRouter()

  return (
    <div className="h-screen bg-background text-foreground relative overflow-hidden">
      <div className="absolute bottom-0 left-0 right-0 w-full px-8 sm:px-12 md:px-16 lg:px-20 z-20">
        <div className="max-w-7xl mx-auto relative">
          <div className="absolute -top-16 sm:-top-12 left-0 flex flex-row gap-3 sm:gap-8 md:gap-16 lg:gap-24">
            <Button
              type="button"
              onClick={() => router.push("/imprint")}
              className="text-foreground hover:text-foreground text-xs sm:text-sm font-light cursor-pointer transition-colors p-0 h-auto rounded-none shadow-none border-0 bg-transparent hover:bg-transparent"
            >
              Imprint
            </Button>
            <Button
              type="button"
              onClick={() => router.push("/gdpr")}
              className="text-foreground hover:text-foreground text-xs sm:text-sm font-light cursor-pointer transition-colors p-0 h-auto rounded-none shadow-none border-0 bg-transparent hover:bg-transparent"
            >
              GDPR
            </Button>
            <Button
              type="button"
              onClick={() => router.push("/terms")}
              className="text-foreground hover:text-foreground text-xs sm:text-sm font-light cursor-pointer transition-colors p-0 h-auto rounded-none shadow-none border-0 bg-transparent hover:bg-transparent"
            >
              Terms of Service
            </Button>
          </div>

          <div className="absolute -top-16 sm:-top-12 right-0 flex flex-row gap-3 sm:gap-4 md:gap-8 lg:gap-12">
            <Lang type="language" />
            <Lang type="mode" />
          </div>

          <div className="flex items-end justify-center space-x-1 md:space-x-2 lg:space-x-3 pb-0">
            <Logo size="medium" autoStart={false} />

            <span className="text-foreground font-thin text-6xl sm:text-7xl md:text-8xl lg:text-9xl xl:text-[12rem] leading-none">
              4nuel
            </span>
          </div>
        </div>
      </div>
    </div>
  )
}
