// src/components/ui/Badge.tsx

import React from "react"
import { cn } from "../../lib/utils"

export interface BadgeProps {
  variant?: "default" | "success" | "warning" | "danger" | "outline" | "info";
  children?: React.ReactNode;
  className?: string;
  style?: React.CSSProperties;
  key?: string | number;
}

function Badge({ className, variant = "default", ...props }: BadgeProps) {
  return (
    <div
      className={cn(
        "inline-flex items-center rounded-full px-3 py-1 text-[10px] sm:text-xs font-black uppercase tracking-widest transition-colors",
        {
          "bg-dark text-white": variant === "default",
          "bg-green-100 text-green-700": variant === "success",
          "bg-amber-100 text-amber-700": variant === "warning",
          "bg-red-100 text-red-700": variant === "danger",
          "bg-blue-100 text-blue-700": variant === "info",
          "border border-muted text-dark": variant === "outline"
        },
        className
      )}
      {...(props as any)}
    />
  )
}

export { Badge }
