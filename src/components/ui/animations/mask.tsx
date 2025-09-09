"use client"
import { useState, useRef, useCallback } from "react"
import type React from "react"

export const Mask = ({
  children,
  revealText,
  revealSize = 400,
  className,
}: {
  children?: string | React.ReactNode
  revealText?: string | React.ReactNode
  revealSize?: number
  className?: string
}) => {
  const [mousePosition, setMousePosition] = useState({ x: 0, y: 0 })
  const [isHovered, setIsHovered] = useState(false)
  const containerRef = useRef<HTMLDivElement>(null)

  const handleMouseMove = useCallback((e: React.MouseEvent) => {
    if (!containerRef.current) return
    const rect = containerRef.current.getBoundingClientRect()
    setMousePosition({
      x: e.clientX - rect.left,
      y: e.clientY - rect.top,
    })
  }, [])

  const handleMouseEnter = useCallback(() => {
    setIsHovered(true)
  }, [])

  const handleMouseLeave = useCallback(() => {
    setIsHovered(false)
  }, [])

  return (
    <div
      ref={containerRef}
      className={`relative h-screen w-full bg-black overflow-visible ${className || ""}`}
      onMouseMove={handleMouseMove}
      onMouseEnter={handleMouseEnter}
      onMouseLeave={handleMouseLeave}
    >
      {/* Default content - white text on black background */}
      <div className="absolute inset-0 flex items-center justify-center text-white z-10">{revealText}</div>

      {/* Masked content - black text on white background with higher z-index */}
      <div
        className="absolute inset-0 flex items-center justify-center bg-white text-black transition-all duration-300 ease-out z-20 overflow-visible"
        style={{
          clipPath: isHovered
            ? `polygon(${mousePosition.x - revealSize/2}px ${mousePosition.y - revealSize/2}px, ${mousePosition.x + revealSize/2}px ${mousePosition.y - revealSize/2}px, ${mousePosition.x + revealSize/2}px ${mousePosition.y + revealSize/2}px, ${mousePosition.x - revealSize/2}px ${mousePosition.y + revealSize/2}px)`
            : `polygon(${mousePosition.x}px ${mousePosition.y}px, ${mousePosition.x}px ${mousePosition.y}px, ${mousePosition.x}px ${mousePosition.y}px, ${mousePosition.x}px ${mousePosition.y}px)`,
        }}
      >
        {children}
      </div>
    </div>
  )
}
