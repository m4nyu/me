"use client"

import { useTheme } from "next-themes"
import { useEffect } from "react"

export function Favicon() {
  const { resolvedTheme } = useTheme()

  useEffect(() => {
    const favicon = document.querySelector("link[rel~='icon']") as HTMLLinkElement
    
    if (!favicon) return

    // Match logo: black bg + white M in light mode, white bg + black M in dark mode  
    const faviconPath = resolvedTheme === "dark" ? "/dark.svg" : "/light.svg"
    
    if (favicon.href !== window.location.origin + faviconPath) {
      favicon.href = faviconPath
    }
  }, [resolvedTheme])

  return null
}