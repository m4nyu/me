"use client"

import type { ReactNode } from "react"

interface SectionProps {
  children: ReactNode
  className?: string
  hasNav?: boolean
}

export function Wrap({ children, className = "", hasNav = false }: SectionProps) {
  if (hasNav) {
    return (
      <div
        className={`h-screen w-full ${className}`}
        style={{
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
          paddingTop: "100px",
          paddingBottom: "100px",
          boxSizing: "border-box",
        }}
      >
        <div className="w-full max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">{children}</div>
      </div>
    )
  }

  return (
    <div className={`h-screen w-full flex items-center justify-center ${className}`}>
      <div className="w-full max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">{children}</div>
    </div>
  )
}
