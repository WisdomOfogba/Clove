import * as React from "react"
import { cn } from "../../lib/utils"

export interface InputProps
  extends React.InputHTMLAttributes<HTMLInputElement> {
    icon?: React.ReactNode;
  }

const Input = React.forwardRef<HTMLInputElement, InputProps>(
  ({ className, type, icon, ...props }, ref) => {
    return (
      <div className="relative flex items-center">
        {icon && (
          <div className="absolute left-4 opacity-50 flex items-center justify-center">
            {icon}
          </div>
        )}
        <input
          type={type}
          className={cn(
            "flex h-14 w-full rounded-2xl border-2 border-muted bg-white px-5 py-3 text-sm font-bold text-dark transition-colors file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-dark/30 placeholder:font-medium focus-visible:outline-none focus-visible:border-primary disabled:cursor-not-allowed disabled:opacity-50 hover:border-slate-300",
            icon && "pl-12",
            className
          )}
          ref={ref}
          {...props}
        />
      </div>
    )
  }
)
Input.displayName = "Input"

export { Input }
