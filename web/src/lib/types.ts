export interface User {
  username: string;
  password: string;
  role?: string;
}

export interface Server {
  id: string;
  name: string;
  ip: string;
  environment: string;
}

export interface AccessRequest {
  server_id: string;
  username: string;
}

export interface ApiResponse {
  [key: string]: string[];
}

export interface AuthStore {
  token: string | null;
  role: string | null;
}