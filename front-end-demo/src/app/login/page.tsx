"use client";

import React, { Suspense } from "react";
import Link from "next/link";
import { useRouter, useSearchParams } from "next/navigation";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Card, CardContent, CardHeader } from "@/components/ui/card";
import { Alert } from "@/components/ui/alert";
import { login } from "@/lib/api";

function LoginContent() {
  const router = useRouter();
  const params = useSearchParams();
  const [username, setUsername] = React.useState("");
  const [password, setPassword] = React.useState("");
  const [error, setError] = React.useState<string | null>(null);

  const registered = params.get("registered") === "1";

  async function onSubmit(e: React.FormEvent) {
    e.preventDefault();
    setError(null);
    try {
      const res = await login(username, password);
      localStorage.setItem("access_token", res.data.accessToken);
      router.push("/resource");
    } catch (err: any) {
      setError(err.message || "Login failed");
    }
  }

  return (
    <div className="min-h-screen flex items-center justify-center bg-zinc-50">
      <Card className="w-full max-w-md">
        <CardHeader>
          <h2 className="text-xl font-semibold">Login</h2>
        </CardHeader>
        <CardContent>
          {registered && (
            <Alert className="mb-4">Registered successfully. Please login.</Alert>
          )}
          {error && (
            <div className="mb-4 text-sm text-red-600">{error}</div>
          )}
          <form onSubmit={onSubmit} className="space-y-4">
            <div>
              <Label htmlFor="username">Username</Label>
              <Input
                id="username"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                placeholder="Enter your username"
                required
              />
            </div>
            <div>
              <Label htmlFor="password">Password</Label>
              <Input
                id="password"
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                placeholder="Enter your password"
                required
              />
            </div>
            <Button type="submit" className="w-full">Login</Button>
          </form>

          <div className="mt-4 text-sm text-zinc-600">
            Don&apos;t have an account? {" "}
            <Link href="/register" className="text-blue-600 hover:underline">
              Register
            </Link>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}

export default function LoginPage() {
  return (
    <Suspense fallback={null}>
      <LoginContent />
    </Suspense>
  );
}