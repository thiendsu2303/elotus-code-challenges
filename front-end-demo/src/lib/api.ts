export const API_BASE = process.env.NEXT_PUBLIC_API_BASE_URL || "http://localhost:8080";

export type LoginResponse = {
  status: string;
  message: string;
  data: {
    accessToken: string;
    tokenType: string;
    expiresAt: string;
  };
};

export type RegisterResponse = {
  status: string;
  message: string;
  data: {
    id: number;
    username: string;
    createdAt: string;
  };
};

export async function login(username: string, password: string): Promise<LoginResponse> {
  const res = await fetch(`${API_BASE}/api/v1/auth/login`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ username, password }),
  });
  if (!res.ok) {
    throw new Error("Invalid credentials");
  }
  const json = await res.json();
  // backend uses snake_case; normalize here
  const data = json.data || {};
  return {
    status: json.status,
    message: json.message,
    data: {
      accessToken: data.access_token,
      tokenType: data.token_type,
      expiresAt: data.expires_at,
    },
  };
}

export async function register(username: string, password: string): Promise<RegisterResponse> {
  const res = await fetch(`${API_BASE}/api/v1/auth/register`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ username, password }),
  });
  if (!res.ok) {
    if (res.status === 409) throw new Error("Username already exists");
    throw new Error("Registration failed");
  }
  const json = await res.json();
  return {
    status: json.status,
    message: json.message,
    data: {
      id: json.data?.id,
      username: json.data?.username,
      createdAt: json.data?.created_at,
    },
  };
}

export async function authGet(path: string): Promise<any> {
  const token = typeof window !== "undefined" ? localStorage.getItem("access_token") : null;
  const res = await fetch(`${API_BASE}${path}`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      ...(token ? { Authorization: `Bearer ${token}` } : {}),
    },
  });
  if (!res.ok) {
    throw new Error(`Request failed: ${res.status}`);
  }
  return res.json();
}

export async function logout(): Promise<boolean> {
  const token = typeof window !== "undefined" ? localStorage.getItem("access_token") : null;
  if (!token) return true; // nothing to do
  try {
    const res = await fetch(`${API_BASE}/api/v1/auth/logout`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
    });
    // Even if backend returns non-200, we still proceed to clear token on client.
    return res.ok;
  } catch (e) {
    return false;
  }
}

export type ImageItem = {
  id: number;
  userId?: number | null;
  filename: string;
  contentType: string;
  sizeBytes: number;
  path: string;
  uploadedAt: string;
};

export async function getMyImages(): Promise<ImageItem[]> {
  const token = typeof window !== "undefined" ? localStorage.getItem("access_token") : null;
  const res = await fetch(`${API_BASE}/api/v1/resource/images`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      ...(token ? { Authorization: `Bearer ${token}` } : {}),
    },
  });
  if (!res.ok) {
    throw new Error(`Request failed: ${res.status}`);
  }
  const json = await res.json();
  const items = Array.isArray(json.data) ? json.data : [];
  return items.map((it: any) => ({
    id: it.id,
    userId: it.user_id ?? null,
    filename: it.filename,
    contentType: it.content_type,
    sizeBytes: it.size_bytes,
    path: it.path,
    uploadedAt: it.uploaded_at,
  }));
}

export async function uploadImage(file: File): Promise<ImageItem> {
  const token = typeof window !== "undefined" ? localStorage.getItem("access_token") : null;
  if (!token) throw new Error("Not authenticated");
  const fd = new FormData();
  fd.append("file", file);
  const res = await fetch(`${API_BASE}/api/v1/upload`, {
    method: "POST",
    headers: {
      ...(token ? { Authorization: `Bearer ${token}` } : {}),
    },
    body: fd,
  });
  if (!res.ok) {
    if (res.status === 413) throw new Error("File too large (max 8MB)");
    if (res.status === 415) throw new Error("Only image files are allowed");
    if (res.status === 401) throw new Error("Unauthorized");
    throw new Error(`Upload failed: ${res.status}`);
  }
  const json = await res.json();
  const it = json.data || {};
  return {
    id: it.id,
    userId: it.user_id ?? null,
    filename: it.filename,
    contentType: it.content_type,
    sizeBytes: it.size_bytes,
    path: it.path,
    uploadedAt: it.uploaded_at,
  };
}

export async function uploadImageWithProgress(
  file: File,
  onProgress?: (percent: number) => void
): Promise<ImageItem> {
  const token = typeof window !== "undefined" ? localStorage.getItem("access_token") : null;
  if (!token) throw new Error("Not authenticated");
  const fd = new FormData();
  fd.append("file", file);
  const url = `${API_BASE}/api/v1/upload`;

  return new Promise<ImageItem>((resolve, reject) => {
    const xhr = new XMLHttpRequest();
    xhr.open("POST", url);
    xhr.setRequestHeader("Authorization", `Bearer ${token}`);

    xhr.upload.onprogress = (evt) => {
      if (!onProgress) return;
      if (evt.lengthComputable) {
        const percent = Math.round((evt.loaded / evt.total) * 100);
        onProgress(percent);
      }
    };

    xhr.onreadystatechange = () => {
      if (xhr.readyState !== XMLHttpRequest.DONE) return;
      const status = xhr.status;
      if (status >= 200 && status < 300) {
        try {
          const json = JSON.parse(xhr.responseText);
          const it = json.data || {};
          resolve({
            id: it.id,
            userId: it.user_id ?? null,
            filename: it.filename,
            contentType: it.content_type,
            sizeBytes: it.size_bytes,
            path: it.path,
            uploadedAt: it.uploaded_at,
          });
        } catch (e) {
          reject(new Error("Invalid JSON response"));
        }
      } else {
        if (status === 413) return reject(new Error("File too large (max 8MB)"));
        if (status === 415) return reject(new Error("Only image files are allowed"));
        if (status === 401) return reject(new Error("Unauthorized"));
        return reject(new Error(`Upload failed: ${status}`));
      }
    };

    xhr.onerror = () => reject(new Error("Network error during upload"));
    xhr.onabort = () => reject(new Error("Upload aborted"));

    xhr.send(fd);
  });
}