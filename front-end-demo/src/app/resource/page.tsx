"use client";

import React from "react";
import { useRouter } from "next/navigation";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader } from "@/components/ui/card";
import { useRequireAuth } from "@/lib/auth";

export default function ResourcePage() {
  const router = useRouter();
  const { message, error, loading, logout } = useRequireAuth();

  return (
    <div className="min-h-screen flex items-center justify-center bg-zinc-50">
      <Card className="w-full max-w-md">
        <CardHeader>
          <h2 className="text-xl font-semibold">Protected Resource</h2>
        </CardHeader>
        <CardContent>
          {loading ? (
            <div className="text-sm text-zinc-500 mb-4">Loading...</div>
          ) : error ? (
            <div className="text-sm text-red-600 mb-4">{error}</div>
          ) : (
            <div className="text-sm text-zinc-800 mb-4">Response: {message}</div>
          )}
          <Button onClick={logout} variant="outline">Logout</Button>
        </CardContent>
      </Card>
    </div>
  );
}