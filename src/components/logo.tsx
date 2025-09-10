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
    medium: "w-20 h-20 sm:w-28 sm:h-28 md:w-40 md:h-40 lg:w-52 lg:h-52 xl:w-64 xl:h-64",
    large: "w-32 h-32 sm:w-40 sm:h-40 md:w-48 md:h-48 lg:w-56 lg:h-56 xl:w-64 xl:h-64",
    xlarge: "w-40 h-40 sm:w-48 sm:h-48 md:w-56 md:h-56 lg:w-64 lg:h-64 xl:w-72 xl:h-72",
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

