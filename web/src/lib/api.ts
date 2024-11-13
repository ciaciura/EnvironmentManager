import { browser } from '$app/environment';
import type { User, Server, AccessRequest, ApiResponse } from '$lib/types';

const API_URL = import.meta.env.VITE_API_BASE_URL 

async function handleResponse(response: Response): Promise<any> {
  const data = await response.json();
  if (!response.ok) {
    throw new Error(data.message || 'API request failed');
  }
  return data;
}

export async function login(credentials: User): Promise<ApiResponse> {
  const response = await fetch(`${API_URL}/login`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Access-Control-Allow-Origin': `${API_URL}`,
      'Access-Control-Allow-Credentials': 'true'
    },
    body: JSON.stringify(credentials),
  });
  return handleResponse(response);
}

export async function createUser(user: User, token: string): Promise<ApiResponse> {
  const response = await fetch(`${API_URL}/admin/user`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': token,
    },
    body: JSON.stringify(user),
  });
  return handleResponse(response);
}

export async function getServers(token: string): Promise<Server[]> {
  const response = await fetch(`${API_URL}/servers`, {
    headers: {
      'Authorization': token,
    },
  });
  return handleResponse(response);
}

export async function addServer(server: Server, token: string): Promise<ApiResponse> {
  const response = await fetch(`${API_URL}/servers/add`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': token,
    },
    body: JSON.stringify(server),
  });
  return handleResponse(response);
}

export async function requestAccess(request: AccessRequest, token: string): Promise<ApiResponse> {
  const response = await fetch(`${API_URL}/servers/request-access`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': token,
    },
    body: JSON.stringify(request),
  });
  return handleResponse(response);
}