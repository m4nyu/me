"use client"

import { Bot, Languages, Moon, Settings, Sun, X } from "lucide-react"
import { useTheme } from "next-themes"
import { useCallback, useEffect, useState } from "react"
import { Dialog, DialogClose, DialogContent, DialogHeader, DialogTitle, DialogTrigger } from "@/components/ui/dialog"

interface Theme {
  name: string
  colors: {
    primary: string
    secondary: string
    accent: string
    background: string
    foreground: string
    muted: string
    border: string
  }
}

const themes: Theme[] = [
  {
    name: "Default",
    colors: {
      primary: "#ffffff",
      secondary: "#6b7280",
      accent: "#3b82f6",
      background: "#000000",
      foreground: "#ffffff",
      muted: "#374151",
      border: "#ffffff",
    },
  },
  {
    name: "Dark Blue",
    colors: {
      primary: "#60a5fa",
      secondary: "#3b82f6",
      accent: "#1d4ed8",
      background: "#0f172a",
      foreground: "#f8fafc",
      muted: "#1e293b",
      border: "#3b82f6",
    },
  },
  {
    name: "Navy Blue",
    colors: {
      primary: "#2563eb",
      secondary: "#1e40af",
      accent: "#1d4ed8",
      background: "#0c1427",
      foreground: "#f1f5f9",
      muted: "#1e2a4a",
      border: "#2563eb",
    },
  },
  {
    name: "Steel Blue",
    colors: {
      primary: "#0ea5e9",
      secondary: "#0284c7",
      accent: "#0369a1",
      background: "#0c1421",
      foreground: "#f0f9ff",
      muted: "#1e3a5f",
      border: "#0ea5e9",
    },
  },
  {
    name: "Midnight Blue",
    colors: {
      primary: "#1e40af",
      secondary: "#1e3a8a",
      accent: "#1d4ed8",
      background: "#0a0f1c",
      foreground: "#e0e7ff",
      muted: "#1e293b",
      border: "#1e40af",
    },
  },
  {
    name: "Ocean Deep",
    colors: {
      primary: "#06b6d4",
      secondary: "#0891b2",
      accent: "#0e7490",
      background: "#0c1618",
      foreground: "#f0fdff",
      muted: "#164e63",
      border: "#06b6d4",
    },
  },
  {
    name: "Forest Green",
    colors: {
      primary: "#10b981",
      secondary: "#059669",
      accent: "#047857",
      background: "#0f1419",
      foreground: "#f0fdf4",
      muted: "#1f2937",
      border: "#10b981",
    },
  },
  {
    name: "Deep Purple",
    colors: {
      primary: "#8b5cf6",
      secondary: "#7c3aed",
      accent: "#6d28d9",
      background: "#1e1b4b",
      foreground: "#faf5ff",
      muted: "#312e81",
      border: "#8b5cf6",
    },
  },
  {
    name: "Crimson Red",
    colors: {
      primary: "#ef4444",
      secondary: "#dc2626",
      accent: "#b91c1c",
      background: "#1f1315",
      foreground: "#fef2f2",
      muted: "#374151",
      border: "#ef4444",
    },
  },
  {
    name: "Amber Gold",
    colors: {
      primary: "#f59e0b",
      secondary: "#d97706",
      accent: "#b45309",
      background: "#1c1917",
      foreground: "#fffbeb",
      muted: "#78716c",
      border: "#f59e0b",
    },
  },
  {
    name: "Rose Pink",
    colors: {
      primary: "#f43f5e",
      secondary: "#e11d48",
      accent: "#be123c",
      background: "#1f1114",
      foreground: "#fff1f2",
      muted: "#4c1d24",
      border: "#f43f5e",
    },
  },
  {
    name: "Slate Gray",
    colors: {
      primary: "#64748b",
      secondary: "#475569",
      accent: "#334155",
      background: "#0f172a",
      foreground: "#f8fafc",
      muted: "#1e293b",
      border: "#64748b",
    },
  },
  {
    name: "Teal Mint",
    colors: {
      primary: "#14b8a6",
      secondary: "#0d9488",
      accent: "#0f766e",
      background: "#042f2e",
      foreground: "#f0fdfa",
      muted: "#134e4a",
      border: "#14b8a6",
    },
  },
  {
    name: "Indigo Night",
    colors: {
      primary: "#6366f1",
      secondary: "#4f46e5",
      accent: "#4338ca",
      background: "#1e1b4b",
      foreground: "#f0f4ff",
      muted: "#312e81",
      border: "#6366f1",
    },
  },
  {
    name: "Copper Bronze",
    colors: {
      primary: "#ea580c",
      secondary: "#c2410c",
      accent: "#9a3412",
      background: "#1c1917",
      foreground: "#fff7ed",
      muted: "#78716c",
      border: "#ea580c",
    },
  },
  {
    name: "Lime Electric",
    colors: {
      primary: "#84cc16",
      secondary: "#65a30d",
      accent: "#4d7c0f",
      background: "#0f1419",
      foreground: "#f7fee7",
      muted: "#365314",
      border: "#84cc16",
    },
  },
  {
    name: "Violet Dream",
    colors: {
      primary: "#a855f7",
      secondary: "#9333ea",
      accent: "#7e22ce",
      background: "#2d1b69",
      foreground: "#faf5ff",
      muted: "#4c1d95",
      border: "#a855f7",
    },
  },
  {
    name: "Cyan Ice",
    colors: {
      primary: "#22d3ee",
      secondary: "#06b6d4",
      accent: "#0891b2",
      background: "#083344",
      foreground: "#ecfeff",
      muted: "#155e75",
      border: "#22d3ee",
    },
  },
  {
    name: "Emerald Forest",
    colors: {
      primary: "#34d399",
      secondary: "#10b981",
      accent: "#059669",
      background: "#022c22",
      foreground: "#ecfdf5",
      muted: "#065f46",
      border: "#34d399",
    },
  },
  {
    name: "High Contrast",
    colors: {
      primary: "#ffffff",
      secondary: "#000000",
      accent: "#0066cc",
      background: "#000000",
      foreground: "#ffffff",
      muted: "#333333",
      border: "#ffffff",
    },
  },
]

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

interface PickerProps {
  type: "theme" | "language" | "mode"
}

export function Picker({ type }: PickerProps) {
  const [dialogOpen, setDialogOpen] = useState(false)
  const [selectedTheme, setSelectedTheme] = useState("Default")
  const [selectedLang, setSelectedLang] = useState("EN")
  const { theme, setTheme } = useTheme()
  const [mounted, setMounted] = useState(false)

  useEffect(() => {
    setMounted(true)
  }, [])

  const applyCSSVariables = useCallback((colors: Theme['colors']) => {
    const root = document.documentElement
    const cssVariableMap = {
      background: '--background',
      foreground: '--foreground',
      primary: '--primary',
      secondary: '--secondary',
      accent: '--accent',
      muted: '--muted',
      border: '--border'
    }
    
    Object.entries(colors).forEach(([key, value]) => {
      // Set both the prefixed version for custom use and the mapped version for Tailwind
      root.style.setProperty(`--color-${key}`, value)
      if (cssVariableMap[key as keyof typeof cssVariableMap]) {
        root.style.setProperty(cssVariableMap[key as keyof typeof cssVariableMap], value)
      }
    })
  }, [])

  const applyTheme = (selectedTheme: Theme) => {
    applyCSSVariables(selectedTheme.colors)

    // Store theme selection in localStorage
    localStorage.setItem("selected-theme", selectedTheme.name)

    setSelectedTheme(selectedTheme.name)
    setDialogOpen(false)
  }

  // Load saved theme on mount and apply it
  useEffect(() => {
    if (mounted) {
      const savedTheme = localStorage.getItem("selected-theme")
      if (savedTheme) {
        const themeObj = themes.find((t) => t.name === savedTheme)
        if (themeObj) {
          setSelectedTheme(themeObj.name)
          applyCSSVariables(themeObj.colors)
        }
      } else {
        // Apply default theme if none saved
        const defaultTheme = themes.find((t) => t.name === "Default")
        if (defaultTheme) {
          applyCSSVariables(defaultTheme.colors)
        }
      }
    }
  }, [mounted, applyCSSVariables])

  const selectLanguage = (lang: string) => {
    setSelectedLang(lang)
    setDialogOpen(false)
  }

  const updateFavicon = (isDark: boolean) => {
    const favicon = document.querySelector('link[rel="icon"]') as HTMLLinkElement
    if (favicon) {
      favicon.href = isDark ? "/favicon.svg" : "/favicon-light.svg"
    } else {
      // Create favicon link if it doesn't exist
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

  // Update favicon when theme changes
  useEffect(() => {
    if (mounted) {
      const isDark =
        theme === "dark" || (theme === "system" && window.matchMedia("(prefers-color-scheme: dark)").matches)
      updateFavicon(isDark)
    }
  }, [theme, mounted, updateFavicon])

  // Listen for system theme changes
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

  if (type === "theme") {
    return (
      <Dialog open={dialogOpen} onOpenChange={setDialogOpen}>
        <DialogTrigger asChild>
          <button
            type="button"
            className="text-muted-foreground hover:text-foreground text-xs sm:text-sm font-light cursor-pointer transition-colors flex items-center gap-1 sm:gap-2"
          >
            <Settings size={14} className="sm:w-4 sm:h-4" />
            <span className="hidden sm:inline">Themes</span>
          </button>
        </DialogTrigger>
        <DialogContent
          className="sm:max-w-[600px] bg-background border-foreground text-foreground max-h-[80vh] rounded-none p-0"
          showCloseButton={false}
        >
          <DialogClose asChild>
            <button
              type="button"
              className="absolute top-4 right-4 z-50 cursor-pointer hover:opacity-70 transition-opacity p-1"
              aria-label="Close dialog"
            >
              <X className="w-4 h-4 text-foreground pointer-events-none" />
            </button>
          </DialogClose>
          <div className="overflow-y-auto max-h-[80vh] p-4 pt-2">
            <DialogHeader className="flex flex-row items-center justify-between relative mt-2">
              <DialogTitle className="text-foreground font-light text-xl">Choose a Theme</DialogTitle>
            </DialogHeader>
            <div className="grid gap-6 py-4">
              <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-2">
                {themes.map((theme) => (
                  <button
                    key={theme.name}
                    type="button"
                    onClick={() => applyTheme(theme)}
                    className={`group relative border border-foreground p-4 hover:bg-foreground transition-all duration-200 cursor-pointer flex items-center gap-3 ${
                      selectedTheme === theme.name ? "bg-foreground" : "bg-transparent"
                    }`}
                  >
                    <div className="flex gap-1">
                      <div className="w-3 h-3" style={{ backgroundColor: theme.colors.primary }} />
                      <div className="w-3 h-3" style={{ backgroundColor: theme.colors.secondary }} />
                      <div className="w-3 h-3" style={{ backgroundColor: theme.colors.accent }} />
                      <div className="w-3 h-3" style={{ backgroundColor: theme.colors.background }} />
                    </div>
                    <div
                      className={`text-[10px] font-medium transition-colors whitespace-nowrap ${
                        selectedTheme === theme.name ? "text-background" : "text-foreground group-hover:text-background"
                      }`}
                    >
                      {theme.name}
                    </div>
                  </button>
                ))}
              </div>
            </div>
          </div>
        </DialogContent>
      </Dialog>
    )
  }

  if (type === "mode") {
    if (!mounted) {
      return null
    }

    return (
      <button
        type="button"
        onClick={cycleMode}
        className="text-muted-foreground hover:text-foreground text-xs sm:text-sm font-light cursor-pointer transition-colors flex items-center gap-1 sm:gap-2"
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
      </button>
    )
  }

  return (
    <Dialog open={dialogOpen} onOpenChange={setDialogOpen}>
      <DialogTrigger asChild>
        <button
          type="button"
          className="text-muted-foreground hover:text-foreground text-xs sm:text-sm font-light cursor-pointer transition-colors flex items-center gap-1 sm:gap-2"
        >
          <Languages size={14} className="sm:w-4 sm:h-4" />
          <span className="hidden sm:inline">{selectedLang}</span>
        </button>
      </DialogTrigger>
      <DialogContent
        className="sm:max-w-[400px] bg-background border-foreground text-foreground max-h-[80vh] rounded-none p-0"
        showCloseButton={false}
      >
        <DialogClose asChild>
          <button
            type="button"
            className="absolute top-4 right-4 z-50 cursor-pointer hover:opacity-70 transition-opacity p-1"
            aria-label="Close dialog"
          >
            <X className="w-4 h-4 text-foreground pointer-events-none" />
          </button>
        </DialogClose>
        <div className="overflow-y-auto max-h-[80vh] p-4 pt-2">
          <DialogHeader className="flex flex-row items-center justify-between relative mt-2">
            <DialogTitle className="text-foreground font-light text-xl">Choose Language</DialogTitle>
          </DialogHeader>
          <div className="grid gap-6 py-4">
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-2">
              {languages.map((language) => (
                <button
                  key={language.code}
                  type="button"
                  onClick={() => selectLanguage(language.code)}
                  className={`group relative border p-4 hover:bg-foreground transition-all duration-200 cursor-pointer flex items-center gap-3 ${
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
                </button>
              ))}
            </div>
          </div>
        </div>
      </DialogContent>
    </Dialog>
  )
}
