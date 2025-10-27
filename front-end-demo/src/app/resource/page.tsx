"use client";

import React from "react";
import { useRouter } from "next/navigation";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader } from "@/components/ui/card";
import { authGet } from "@/lib/api";

export default function ResourcePage() {
  const router = useRouter();
  const [message, setMessage] = React.useState<string>("");
  const [error, setError] = React.useState<string | null>(null);

  React.useEffect(() => {
    async function load() {
      try {
        const res = await authGet("/api/v1/ping-auth");
        setMessage(res.message || "pong_auth");
      } catch (err) {
        setError("Unauthorized. Please login.");
      }
    }
    load();
  }, []);

  function logout() {
    localStorage.removeItem("access_token");
    router.push("/login");
  }

  return (
    <div className="min-h-screen flex items-center justify-center bg-zinc-50">
      <Card className="w-full max-w-md">
        <CardHeader>
          <h2 className="text-xl font-semibold">Protected Resource</h2>
        </CardHeader>
        <CardContent>
          {error ? (
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