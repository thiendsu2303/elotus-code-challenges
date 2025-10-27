import * as React from "react";
import clsx from "clsx";

export function Alert({ className, ...props }: React.HTMLAttributes<HTMLDivElement>) {
  return (
    <div className={clsx("rounded-md border border-green-300 bg-green-50 text-green-800 p-3 text-sm", className)} {...props} />
  );
}