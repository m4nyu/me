"use client"

import { useState } from "react"
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger, DialogClose } from "@/components/ui/dialog"
import { Settings, X } from "lucide-react"

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
    name: "Modern Minimal",
    colors: {
      primary: "#ffffff",
      secondary: "#6b7280",
      accent: "#3b82f6",
      background: "#000000",
      foreground: "#ffffff",
      muted: "#374151",
      border: "#ffffff"
    }
  },
  {
    name: "Twitter",
    colors: {
      primary: "#1da1f2",
      secondary: "#14171a",
      accent: "#657786",
      background: "#ffffff",
      foreground: "#14171a",
      muted: "#f7f9fa",
      border: "#e1e8ed"
    }
  },
  {
    name: "Amethyst Haze",
    colors: {
      primary: "#8b5cf6",
      secondary: "#a78bfa",
      accent: "#06ffa5",
      background: "#0c0a09",
      foreground: "#fafafa",
      muted: "#27272a",
      border: "#8b5cf6"
    }
  },
  {
    name: "Catppuccin",
    colors: {
      primary: "#cba6f7",
      secondary: "#f5c2e7",
      accent: "#a6e3a1",
      background: "#1e1e2e",
      foreground: "#cdd6f4",
      muted: "#313244",
      border: "#cba6f7"
    }
  },
  {
    name: "Kodama Grove",
    colors: {
      primary: "#10b981",
      secondary: "#34d399",
      accent: "#f59e0b",
      background: "#064e3b",
      foreground: "#ecfdf5",
      muted: "#065f46",
      border: "#10b981"
    }
  },
  {
    name: "Quantum Rose",
    colors: {
      primary: "#ec4899",
      secondary: "#f472b6",
      accent: "#06ffa5",
      background: "#831843",
      foreground: "#fdf2f8",
      muted: "#9d174d",
      border: "#ec4899"
    }
  },
  {
    name: "Elegant Luxury",
    colors: {
      primary: "#d4af37",
      secondary: "#f4e4bc",
      accent: "#2c3e50",
      background: "#1a1a1a",
      foreground: "#f5f5f5",
      muted: "#2d2d2d",
      border: "#d4af37"
    }
  },
  {
    name: "Neo Brutali",
    colors: {
      primary: "#ff0000",
      secondary: "#00ff00",
      accent: "#0000ff",
      background: "#000000",
      foreground: "#ffffff",
      muted: "#333333",
      border: "#ff0000"
    }
  },
  {
    name: "Violet Bloom",
    colors: {
      primary: "#7c3aed",
      secondary: "#a855f7",
      accent: "#fbbf24",
      background: "#1e1b4b",
      foreground: "#f3f4f6",
      muted: "#312e81",
      border: "#7c3aed"
    }
  },
  {
    name: "Mocha Mousse",
    colors: {
      primary: "#d2b48c",
      secondary: "#f5deb3",
      accent: "#8b4513",
      background: "#2f1b14",
      foreground: "#f5deb3",
      muted: "#3d2817",
      border: "#d2b48c"
    }
  },
  {
    name: "Notebook",
    colors: {
      primary: "#f59e0b",
      secondary: "#fbbf24",
      accent: "#ef4444",
      background: "#fef3c7",
      foreground: "#92400e",
      muted: "#fde68a",
      border: "#f59e0b"
    }
  },
  {
    name: "Graphite",
    colors: {
      primary: "#6b7280",
      secondary: "#9ca3af",
      accent: "#3b82f6",
      background: "#111827",
      foreground: "#f9fafb",
      muted: "#374151",
      border: "#6b7280"
    }
  },
  {
    name: "Cosmic Night",
    colors: {
      primary: "#06b6d4",
      secondary: "#22d3ee",
      accent: "#fbbf24",
      background: "#083344",
      foreground: "#ecfeff",
      muted: "#0e7490",
      border: "#06b6d4"
    }
  },
  {
    name: "Nature",
    colors: {
      primary: "#16a34a",
      secondary: "#22c55e",
      accent: "#eab308",
      background: "#052e16",
      foreground: "#f0fdf4",
      muted: "#166534",
      border: "#16a34a"
    }
  },
  {
    name: "Amber Minimal",
    colors: {
      primary: "#f59e0b",
      secondary: "#fbbf24",
      accent: "#ef4444",
      background: "#451a03",
      foreground: "#fffbeb",
      muted: "#92400e",
      border: "#f59e0b"
    }
  },
  {
    name: "Solar Dusk",
    colors: {
      primary: "#f97316",
      secondary: "#fb923c",
      accent: "#06ffa5",
      background: "#7c2d12",
      foreground: "#fff7ed",
      muted: "#9a3412",
      border: "#f97316"
    }
  },
  {
    name: "Chat",
    colors: {
      primary: "#3b82f6",
      secondary: "#60a5fa",
      accent: "#fbbf24",
      background: "#0f172a",
      foreground: "#f8fafc",
      muted: "#1e293b",
      border: "#3b82f6"
    }
  },
  {
    name: "Bubblegum",
    colors: {
      primary: "#ec4899",
      secondary: "#f472b6",
      accent: "#06ffa5",
      background: "#831843",
      foreground: "#fdf2f8",
      muted: "#9d174d",
      border: "#ec4899"
    }
  },
  {
    name: "Doom 64",
    colors: {
      primary: "#ef4444",
      secondary: "#f87171",
      accent: "#fbbf24",
      background: "#7f1d1d",
      foreground: "#fef2f2",
      muted: "#991b1b",
      border: "#ef4444"
    }
  },
  {
    name: "Perpetuity",
    colors: {
      primary: "#8b5cf6",
      secondary: "#a78bfa",
      accent: "#06ffa5",
      background: "#0c0a09",
      foreground: "#fafafa",
      muted: "#27272a",
      border: "#8b5cf6"
    }
  },
  {
    name: "Tangerine",
    colors: {
      primary: "#f97316",
      secondary: "#fb923c",
      accent: "#06ffa5",
      background: "#7c2d12",
      foreground: "#fff7ed",
      muted: "#9a3412",
      border: "#f97316"
    }
  },
  {
    name: "Bold Tech",
    colors: {
      primary: "#00ff00",
      secondary: "#00cc00",
      accent: "#ff0000",
      background: "#000000",
      foreground: "#00ff00",
      muted: "#003300",
      border: "#00ff00"
    }
  },
  {
    name: "Supabase",
    colors: {
      primary: "#3ecf8e",
      secondary: "#4ade80",
      accent: "#f59e0b",
      background: "#0f172a",
      foreground: "#f8fafc",
      muted: "#1e293b",
      border: "#3ecf8e"
    }
  },
  {
    name: "Claymorphism",
    colors: {
      primary: "#fbbf24",
      secondary: "#fcd34d",
      accent: "#ef4444",
      background: "#fef3c7",
      foreground: "#92400e",
      muted: "#fde68a",
      border: "#fbbf24"
    }
  },
  {
    name: "Clean S",
    colors: {
      primary: "#ffffff",
      secondary: "#6b7280",
      accent: "#3b82f6",
      background: "#000000",
      foreground: "#ffffff",
      muted: "#374151",
      border: "#ffffff"
    }
  }
]

export function ThemeSelector() {
  const [themeDialogOpen, setThemeDialogOpen] = useState(false)
  const [selectedTheme, setSelectedTheme] = useState('Modern Minimal')

  const applyTheme = (theme: Theme) => {
    const root = document.documentElement
    Object.entries(theme.colors).forEach(([key, value]) => {
      root.style.setProperty(`--color-${key}`, value)
    })
    setSelectedTheme(theme.name)
    setThemeDialogOpen(false)
  }

  return (
    <Dialog open={themeDialogOpen} onOpenChange={setThemeDialogOpen}>
      <DialogTrigger asChild>
        <button className="text-gray-400 hover:text-white text-xs sm:text-sm font-light cursor-pointer transition-colors flex items-center gap-1 sm:gap-2">
          <Settings size={14} className="sm:w-4 sm:h-4" />
          <span className="hidden sm:inline">Themes</span>
        </button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[600px] bg-black border-white text-white max-h-[80vh] overflow-y-auto rounded-none p-4 pt-2" showCloseButton={false}>
        <button 
          onClick={() => setThemeDialogOpen(false)}
          className="absolute top-4 right-4 cursor-pointer hover:opacity-70 transition-opacity"
        >
          <X className="w-4 h-4 text-white" />
        </button>
        <DialogHeader className="flex flex-row items-center justify-between relative mt-2">
          <DialogTitle className="text-white font-light text-xl">Choose a Theme</DialogTitle>
        </DialogHeader>
        <div className="grid gap-6 py-4">
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-2">
            {themes.map((theme) => (
              <button
                key={theme.name}
                onClick={() => applyTheme(theme)}
                className={`group relative border border-white p-4 hover:bg-white transition-all duration-200 cursor-pointer flex items-center gap-3 ${
                  selectedTheme === theme.name 
                    ? 'bg-white' 
                    : 'bg-transparent'
                }`}
              >
                <div className="flex gap-1">
                  <div 
                    className="w-3 h-3" 
                    style={{ backgroundColor: theme.colors.primary }}
                  />
                  <div 
                    className="w-3 h-3" 
                    style={{ backgroundColor: theme.colors.secondary }}
                  />
                  <div 
                    className="w-3 h-3" 
                    style={{ backgroundColor: theme.colors.accent }}
                  />
                  <div 
                    className="w-3 h-3" 
                    style={{ backgroundColor: theme.colors.background }}
                  />
                </div>
                <div className={`text-[10px] font-medium transition-colors whitespace-nowrap ${
                  selectedTheme === theme.name 
                    ? 'text-black' 
                    : 'text-white group-hover:text-black'
                }`}>
                  {theme.name}
                </div>
              </button>
            ))}
          </div>
        </div>
      </DialogContent>
    </Dialog>
  )
}
