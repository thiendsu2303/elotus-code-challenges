"use client";

import React from "react";
import { useRouter } from "next/navigation";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader } from "@/components/ui/card";
import { useRequireAuth } from "@/lib/auth";
import { getMyImages, uploadImageWithProgress, ImageItem } from "@/lib/api";

export default function ResourcePage() {
  const router = useRouter();
  const { error, loading, logout } = useRequireAuth();
  const [images, setImages] = React.useState<ImageItem[]>([]);
  const [imagesError, setImagesError] = React.useState<string | null>(null);
  const [file, setFile] = React.useState<File | null>(null);
  const [uploading, setUploading] = React.useState(false);
  const [uploadError, setUploadError] = React.useState<string | null>(null);
  const [previewUrl, setPreviewUrl] = React.useState<string | null>(null);
  const [uploadProgress, setUploadProgress] = React.useState<number>(0);
  const fileInputRef = React.useRef<HTMLInputElement | null>(null);

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

  async function onUpload(e: React.FormEvent) {
    e.preventDefault();
    if (!file) {
      // Mở hộp thoại chọn file nếu chưa chọn
      fileInputRef.current?.click();
      return;
    }
    setUploadError(null);
    setUploading(true);
    try {
      const item = await uploadImageWithProgress(file, (p) => setUploadProgress(p));
      setImages((prev) => [item, ...prev]);
      setFile(null);
      setUploadProgress(0);
      if (previewUrl) {
        URL.revokeObjectURL(previewUrl);
        setPreviewUrl(null);
      }
    } catch (err: any) {
      setUploadError(err?.message || "Upload failed");
      // Reset lựa chọn để người dùng có thể upload file khác ngay
      setFile(null);
      setUploadProgress(0);
      if (previewUrl) {
        URL.revokeObjectURL(previewUrl);
        setPreviewUrl(null);
      }
      if (fileInputRef.current) {
        try {
          // Reset giá trị input file để cho phép chọn cùng tên file lần nữa
          (fileInputRef.current as HTMLInputElement).value = "";
        } catch {}
      }
    } finally {
      setUploading(false);
    }
  }

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

          {!loading && !error && (
            <form onSubmit={onUpload} className="space-y-2 mb-4">
              <input
                type="file"
                name="file"
                ref={fileInputRef}
                onChange={(e) => {
                  const f = e.target.files?.[0] || null;
                  if (!f) {
                    setFile(null);
                    if (previewUrl) {
                      URL.revokeObjectURL(previewUrl);
                      setPreviewUrl(null);
                    }
                    return;
                  }
                  setUploadError(null);
                  setFile(f);
                  // Tạo preview nếu là ảnh; nếu không thì bỏ qua preview
                  if (previewUrl) URL.revokeObjectURL(previewUrl);
                  if (f.type && f.type.startsWith("image/")) {
                    setPreviewUrl(URL.createObjectURL(f));
                  } else {
                    setPreviewUrl(null);
                  }
                }}
              />
              {file && (
                <div className="text-xs text-zinc-600">
                  Đã chọn: {file.name} • {(file.size / 1024).toFixed(1)} KB
                </div>
              )}
              {previewUrl && (
                <div className="mt-2">
                  <img src={previewUrl} alt="Preview" className="max-h-48 rounded border" />
                </div>
              )}
              {uploadError && (
                <div className="text-sm text-red-600 flex items-center gap-2">
                  <span>{uploadError}</span>
                  <Button type="button" variant="outline" onClick={() => fileInputRef.current?.click()}>
                    Chọn tệp khác
                  </Button>
                </div>
              )}
              <Button type="submit" disabled={uploading}>
                {uploading ? "Uploading..." : "Upload File"}
              </Button>
              {uploading && (
                <div className="mt-2 w-full bg-zinc-200 h-2 rounded">
                  <div
                    className="h-2 bg-blue-600 rounded"
                    style={{ width: `${uploadProgress}%` }}
                  />
                </div>
              )}
            </form>
          )}
          <Button onClick={logout} variant="outline">Logout</Button>
        </CardContent>
      </Card>
    </div>
  );
}