"use client"

import { useEffect, useMemo, useRef, useState } from "react"

interface Tile {
  id: string
  x: number
  y: number
  width: number
  height: number
  delay: number
}

interface AnimatedLogoProps {
  size?: "small" | "medium" | "large" | "xlarge"
  className?: string
  autoStart?: boolean
}

export function Logo({ size = "medium", className = "", autoStart = true }: AnimatedLogoProps) {
  const [isVisible, setIsVisible] = useState(false)
  const [visibleTileIds, setVisibleTileIds] = useState<Set<string>>(new Set())
  const sectionRef = useRef<HTMLDivElement>(null)

  const sizeConfig = {
    small: "w-8 h-8",
    medium: "w-16 h-16 sm:w-20 sm:h-20 md:w-24 md:h-24 lg:w-28 lg:h-28 xl:w-32 xl:h-32",
    large: "w-20 h-20 sm:w-24 sm:h-24 md:w-28 md:h-28 lg:w-32 lg:h-32 xl:w-36 xl:h-36",
    xlarge: "w-24 h-24 sm:w-28 sm:h-28 md:w-32 md:h-32 lg:w-36 lg:h-36 xl:w-40 xl:h-40",
  }

  const tiles = useMemo(() => {
    const tileList: Tile[] = []
    const gridSize = 6
    const tileBaseSize = 100 / gridSize

    let seed = 12345
    const seededRandom = () => {
      seed = (seed * 9301 + 49297) % 233280
      return seed / 233280
    }

    for (let row = 0; row < gridSize; row++) {
      for (let col = 0; col < gridSize; col++) {
        const delay = seededRandom() * 2000

        tileList.push({
          id: `tile-${row}-${col}`,
          x: col * tileBaseSize,
          y: row * tileBaseSize,
          width: tileBaseSize,
          height: tileBaseSize,
          delay: delay,
        })
      }
    }

    return tileList
  }, [])

  useEffect(() => {
    if (autoStart) {
      setIsVisible(true)

      tiles.forEach((tile) => {
        setTimeout(() => {
          setVisibleTileIds((prev) => new Set([...prev, tile.id]))
        }, tile.delay)
      })
    } else {
      const observer = new IntersectionObserver(
        ([entry]) => {
          if (entry.isIntersecting) {
            setIsVisible(false)
            setVisibleTileIds(new Set())

            setTimeout(() => {
              setIsVisible(true)

              tiles.forEach((tile) => {
                setTimeout(() => {
                  setVisibleTileIds((prev) => new Set([...prev, tile.id]))
                }, tile.delay)
              })
            }, 100)
          }
        },
        { threshold: 0.3 }
      )

      if (sectionRef.current) {
        observer.observe(sectionRef.current)
      }

      return () => observer.disconnect()
    }
  }, [tiles, autoStart])

  return (
    <div ref={sectionRef} className={`${sizeConfig[size]} flex-shrink-0 relative ${className}`}>
      {tiles.map((tile) => {
        const isThisTileVisible = visibleTileIds.has(tile.id)

        return (
          <div
            key={tile.id}
            className="absolute"
            style={{
              width: `${tile.width}%`,
              height: `${tile.height}%`,
              left: `${tile.x}%`,
              top: `${tile.y}%`,
              overflow: "hidden",
              opacity: isThisTileVisible ? 1 : 0,
              transition: "opacity 0.1s ease-in",
            }}
          >
            <svg
              width="100%"
              height="100%"
              viewBox="0 0 32 32"
              style={{
                position: "absolute",
                width: `${(100 / tile.width) * 100}%`,
                height: `${(100 / tile.height) * 100}%`,
                left: `-${(tile.x / tile.width) * 100}%`,
                top: `-${(tile.y / tile.height) * 100}%`,
              }}
            >
              <rect width="32" height="32" className="fill-foreground" />
              <g transform="translate(16,16) rotate(40.5) translate(-16,-16)">
                <path
                  d="M8.5 24 L8.5 8 L10.5 8 L15.2 18 L16.8 18 L21.5 8 L23.5 8 L23.5 24 L21.5 24 L21.5 10.5 L17.2 20.5 L14.8 20.5 L10.5 10.5 L10.5 24 L8.5 24 Z"
                  className="fill-background"
                />
              </g>
            </svg>
          </div>
        )
      })}
    </div>
  )
}
