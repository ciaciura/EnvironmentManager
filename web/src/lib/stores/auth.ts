import { writable } from 'svelte/store';
import type { AuthStore } from '$lib/types';

function createAuthStore() {
  const { subscribe, set, update } = writable<AuthStore>({
    token: null,
    role: null
  });

  return {
    subscribe,
    setAuth: (token: string, role: string) => {
      set({ token, role });
    },
    clearAuth: () => {
      set({ token: null, role: null });
    }
  };
}

export const auth = createAuthStore();