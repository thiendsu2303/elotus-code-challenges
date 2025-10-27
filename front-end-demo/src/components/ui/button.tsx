"use client";

import React from "react";
import clsx from "clsx";

type ButtonProps = React.ButtonHTMLAttributes<HTMLButtonElement> & {
  variant?: "default" | "outline";
};

export function Button({ className, variant = "default", ...props }: ButtonProps) {
  const base = "inline-flex items-center justify-center rounded-md text-sm font-medium transition-colors focus-visible:outline-none disabled:opacity-50 disabled:pointer-events-none h-10 px-4 py-2";
  const variants = {
    default: "bg-black text-white hover:bg-zinc-800",
    outline: "border border-zinc-300 hover:bg-zinc-100",
  };
  return <button className={clsx(base, variants[variant], className)} {...props} />;
}