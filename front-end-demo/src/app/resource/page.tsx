"use client";

import React from "react";
import { useRouter } from "next/navigation";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader } from "@/components/ui/card";
import { useRequireAuth } from "@/lib/auth";
import { getMyImages, ImageItem } from "@/lib/api";

export default function ResourcePage() {
  const router = useRouter();
  const { error, loading, logout } = useRequireAuth();
  const [images, setImages] = React.useState<ImageItem[]>([]);
  const [imagesError, setImagesError] = React.useState<string | null>(null);

  React.useEffect(() => {
    if (loading || error) return;
    async function loadImages() {
      try {
        const list = await getMyImages();
        setImages(list);
      } catch (e: any) {
        setImagesError(e?.message || "Failed to load images");
      }
    }
    loadImages();
  }, [loading, error]);

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
          ) : imagesError ? (
            <div className="text-sm text-red-600 mb-4">{imagesError}</div>
          ) : images.length === 0 ? (
            <div className="text-sm text-zinc-600 mb-4">No images found.</div>
          ) : (
            <ul className="space-y-2 mb-4">
              {images.map((img) => (
                <li key={img.id} className="flex items-center justify-between text-sm">
                  <span className="font-medium">{img.filename}</span>
                  <span className="text-zinc-500">{img.contentType} • {(img.sizeBytes/1024).toFixed(1)} KB</span>
                </li>
              ))}
            </ul>
          )}
          <Button onClick={logout} variant="outline">Logout</Button>
        </CardContent>
      </Card>
    </div>
  );
}