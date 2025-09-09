"use client"

import { useState } from "react"
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger, DialogClose } from "@/components/ui/dialog"
import { Languages, X } from "lucide-react"

const languages = [
  { code: 'EN', name: 'English' },
  { code: 'DE', name: 'Deutsch' },
  { code: 'FR', name: 'Français' },
  { code: 'ES', name: 'Español' },
  { code: 'IT', name: 'Italiano' },
  { code: 'PT', name: 'Português' },
  { code: 'NL', name: 'Nederlands' },
  { code: 'RU', name: 'Русский' },
  { code: 'JA', name: '日本語' },
  { code: 'KO', name: '한국어' },
  { code: 'ZH', name: '中文' },
  { code: 'AR', name: 'العربية' }
]

export function LanguageSelector() {
  const [currentLang, setCurrentLang] = useState('EN')
  const [languageDialogOpen, setLanguageDialogOpen] = useState(false)

  const selectLanguage = (lang: string) => {
    setCurrentLang(lang)
    setLanguageDialogOpen(false)
  }

  return (
    <Dialog open={languageDialogOpen} onOpenChange={setLanguageDialogOpen}>
      <DialogTrigger asChild>
        <button className="text-gray-400 hover:text-white text-xs sm:text-sm font-light cursor-pointer transition-colors flex items-center gap-1 sm:gap-2">
          <Languages size={14} className="sm:w-4 sm:h-4" />
          <span className="hidden sm:inline">{currentLang}</span>
        </button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[400px] bg-black border-white text-white max-h-[80vh] overflow-y-auto rounded-none p-4 pt-2" showCloseButton={false}>
        <button 
          onClick={() => setLanguageDialogOpen(false)}
          className="absolute top-4 right-4 cursor-pointer hover:opacity-70 transition-opacity"
        >
          <X className="w-4 h-4 text-white" />
        </button>
        <DialogHeader className="flex flex-row items-center justify-between relative mt-2">
          <DialogTitle className="text-white font-light text-xl">Choose Language</DialogTitle>
        </DialogHeader>
        <div className="grid gap-6 py-4">
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-2">
            {languages.map((language) => (
              <button
                key={language.code}
                onClick={() => selectLanguage(language.code)}
                className={`group relative border p-4 hover:bg-white transition-all duration-200 cursor-pointer flex items-center gap-3 ${
                  currentLang === language.code 
                    ? 'bg-white border-white' 
                    : 'bg-transparent border-white'
                }`}
              >
                <div className={`text-sm font-bold min-w-[1.5rem] text-center transition-colors ${
                  currentLang === language.code 
                    ? 'text-black' 
                    : 'text-white group-hover:text-black'
                }`}>
                  {language.code}
                </div>
                <div className={`text-[10px] font-medium transition-colors whitespace-nowrap ${
                  currentLang === language.code 
                    ? 'text-black' 
                    : 'text-white group-hover:text-black'
                }`}>
                  {language.name}
                </div>
              </button>
            ))}
          </div>
        </div>
      </DialogContent>
    </Dialog>
  )
}
