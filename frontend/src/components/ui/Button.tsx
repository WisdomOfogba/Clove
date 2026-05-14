import * as React from "react"
import { cn } from "../../lib/utils"

export interface ButtonProps
  extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: "default" | "primary" | "secondary" | "outline" | "ghost" | "danger"
  size?: "default" | "sm" | "lg" | "icon"
}

const Button = React.forwardRef<HTMLButtonElement, ButtonProps>(
  ({ className, variant = "default", size = "default", ...props }, ref) => {
    return (
      <button
        ref={ref}
        className={cn(
          "inline-flex items-center justify-center whitespace-nowrap rounded-2xl sm:rounded-full text-sm font-black transition-all focus-visible:outline-none disabled:pointer-events-none disabled:opacity-50 active:scale-95",
          {
            "bg-dark text-white hover:bg-primary shadow-xl": variant === "default",
            "bg-primary text-dark hover:bg-dark hover:text-white": variant === "primary",
            "bg-muted text-dark hover:bg-slate-200": variant === "secondary",
            "border-2 border-muted text-dark hover:border-dark hover:bg-muted/50": variant === "outline",
            "bg-red-50 text-red-600 hover:bg-red-100": variant === "danger",
            "hover:bg-slate-100 text-dark": variant === "ghost",
            "h-12 sm:h-14 px-8 py-3": size === "default",
            "h-10 px-5 text-xs": size === "sm",
            "h-16 px-10 text-base": size === "lg",
            "h-12 w-12": size === "icon",
          },
          className
        )}
        {...props}
      />
    )
  }
)
Button.displayName = "Button"

export { Button }
