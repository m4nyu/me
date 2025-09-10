"use client"

import { Bot, Languages, Moon, Sun, X } from "lucide-react"
import { useTheme } from "next-themes"
import { useEffect, useState } from "react"
import { Dialog, DialogClose, DialogContent, DialogHeader, DialogTitle, DialogTrigger } from "@/components/ui/dialog"
import { Button } from "@/components/ui/button"

const languages = [
  { code: "EN", name: "English" },
  { code: "DE", name: "Deutsch" },
  { code: "FR", name: "Français" },
  { code: "ES", name: "Español" },
  { code: "IT", name: "Italiano" },
  { code: "PT", name: "Português" },
  { code: "NL", name: "Nederlands" },
  { code: "RU", name: "Русский" },
  { code: "JA", name: "日本語" },
  { code: "KO", name: "한국어" },
  { code: "ZH", name: "中文" },
  { code: "AR", name: "العربية" },
]

interface LangProps {
  type: "language" | "mode"
}

export function Lang({ type }: LangProps) {
  const [dialogOpen, setDialogOpen] = useState(false)
  const [selectedLang, setSelectedLang] = useState("EN")
  const { theme, setTheme } = useTheme()
  const [mounted, setMounted] = useState(false)

  useEffect(() => {
    setMounted(true)
  }, [])

  const selectLanguage = (lang: string) => {
    setSelectedLang(lang)
    setDialogOpen(false)
  }

  const updateFavicon = (isDark: boolean) => {
    const favicon = document.querySelector('link[rel="icon"]') as HTMLLinkElement
    if (favicon) {
      favicon.href = isDark ? "/favicon.svg" : "/favicon-light.svg"
    } else {
      const newFavicon = document.createElement("link")
      newFavicon.rel = "icon"
      newFavicon.href = isDark ? "/favicon.svg" : "/favicon-light.svg"
      document.head.appendChild(newFavicon)
    }
  }

  const cycleMode = () => {
    if (theme === "system") {
      setTheme("dark")
    } else if (theme === "dark") {
      setTheme("light")
    } else {
      setTheme("system")
    }
  }

  useEffect(() => {
    if (mounted) {
      const isDark =
        theme === "dark" || (theme === "system" && window.matchMedia("(prefers-color-scheme: dark)").matches)
      updateFavicon(isDark)
    }
  }, [theme, mounted, updateFavicon])

  useEffect(() => {
    if (mounted && theme === "system") {
      const mediaQuery = window.matchMedia("(prefers-color-scheme: dark)")
      const handleChange = (e: MediaQueryListEvent) => {
        updateFavicon(e.matches)
      }

      mediaQuery.addEventListener("change", handleChange)
      return () => mediaQuery.removeEventListener("change", handleChange)
    }
  }, [theme, mounted, updateFavicon])

  if (type === "mode") {
    if (!mounted) {
      return null
    }

    return (
      <Button
        type="button"
        onClick={cycleMode}
        className="text-foreground hover:text-foreground text-xs sm:text-sm font-light cursor-pointer transition-colors flex items-center gap-1 sm:gap-2 p-0 h-auto rounded-none shadow-none border-0 bg-transparent hover:bg-transparent"
        aria-label="Toggle dark/light/auto mode"
      >
        <div className="w-[14px] h-[14px] sm:w-4 sm:h-4 flex items-center justify-center">
          {theme === "dark" ? (
            <Moon size={14} className="sm:w-4 sm:h-4" />
          ) : theme === "light" ? (
            <Sun size={14} className="sm:w-4 sm:h-4" />
          ) : (
            <Bot size={14} className="sm:w-4 sm:h-4" />
          )}
        </div>
        <span className="hidden sm:inline capitalize w-8 text-left">{theme === "system" ? "Auto" : theme}</span>
      </Button>
    )
  }

  return (
    <Dialog open={dialogOpen} onOpenChange={setDialogOpen}>
      <DialogTrigger asChild>
        <Button
          type="button"
          className="text-foreground hover:text-foreground text-xs sm:text-sm font-light cursor-pointer transition-colors flex items-center gap-1 sm:gap-2 p-0 h-auto rounded-none shadow-none border-0 bg-transparent hover:bg-transparent"
        >
          <Languages size={14} className="sm:w-4 sm:h-4" />
          <span className="hidden sm:inline">{selectedLang}</span>
        </Button>
      </DialogTrigger>
      <DialogContent
        className="sm:max-w-[400px] bg-background border-foreground text-foreground max-h-[80vh] rounded-none p-0"
        showCloseButton={false}
      >
        <DialogClose asChild>
          <Button
            type="button"
            className="absolute top-4 right-4 z-50 cursor-pointer hover:opacity-70 transition-opacity p-1 h-auto rounded-none shadow-none border-0 bg-transparent hover:bg-transparent"
            aria-label="Close dialog"
          >
            <X className="w-4 h-4 text-foreground pointer-events-none" />
          </Button>
        </DialogClose>
        <div className="overflow-y-auto max-h-[80vh] p-4 pt-2">
          <DialogHeader className="flex flex-row items-center justify-between relative mt-2">
            <DialogTitle className="text-foreground font-light text-xl">Choose Language</DialogTitle>
          </DialogHeader>
          <div className="grid gap-6 py-4">
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-2">
              {languages.map((language) => (
                <Button
                  key={language.code}
                  type="button"
                  onClick={() => selectLanguage(language.code)}
                  className={`group relative border p-4 hover:bg-foreground transition-all duration-200 cursor-pointer flex items-center gap-3 h-auto rounded-none shadow-none ${
                    selectedLang === language.code ? "bg-foreground border-foreground" : "bg-transparent border-foreground"
                  }`}
                >
                  <div
                    className={`text-sm font-bold min-w-[1.5rem] text-center transition-colors ${
                      selectedLang === language.code ? "text-background" : "text-foreground group-hover:text-background"
                    }`}
                  >
                    {language.code}
                  </div>
                  <div
                    className={`text-[10px] font-medium transition-colors whitespace-nowrap ${
                      selectedLang === language.code ? "text-background" : "text-foreground group-hover:text-background"
                    }`}
                  >
                    {language.name}
                  </div>
                </Button>
              ))}
            </div>
          </div>
        </div>
      </DialogContent>
    </Dialog>
  )
}
