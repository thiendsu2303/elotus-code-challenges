"use client";

import React from "react";
import { useRouter } from "next/navigation";
import { authGet, logout as logoutApi } from "@/lib/api";

export function useRequireAuth() {
  const router = useRouter();
  const [message, setMessage] = React.useState<string>("");
  const [error, setError] = React.useState<string | null>(null);
  const [loading, setLoading] = React.useState<boolean>(true);

  React.useEffect(() => {
    async function verify() {
      const token = typeof window !== "undefined" ? localStorage.getItem("access_token") : null;
      if (!token) {
        router.replace("/login");
        return;
      }
      try {
        const res = await authGet("/api/v1/ping-auth");
        setMessage(res.message || "pong_auth");
      } catch (err) {
        setError("Unauthorized. Please login.");
        router.replace("/login");
        return;
      } finally {
        setLoading(false);
      }
    }
    verify();
  }, [router]);

  async function logout() {
    try {
      await logoutApi();
    } catch (e) {
      // ignore errors
    }
    localStorage.removeItem("access_token");
    router.push("/login");
  }

  return { message, error, loading, logout };
}