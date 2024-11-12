<script lang="ts">
  import { auth } from '$lib/stores/auth';
  import { login } from '$lib/api';
  import { goto } from '$app/navigation';
  import type { User } from '$lib/types';

  let username = '';
  let password = '';
  let error = '';
  let loading = false;

  async function handleSubmit() {
    error = '';
    loading = true;

    try {
      const credentials: User = { username, password };
      const response = await login(credentials);
      
      if (response.token && response.role) {
        auth.setAuth(response.token[0], response.role[0]);
        goto('/');
      } else {
        error = 'Invalid response from server';
      }
    } catch (e) {
      error = e.message || 'Login failed';
    } finally {
      loading = false;
    }
  }
</script>

<div class="flex min-h-full flex-col justify-center px-6 py-12 lg:px-8">
  <div class="sm:mx-auto sm:w-full sm:max-w-sm">
    <h2 class="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900">
      Sign in to your account
    </h2>
  </div>

  <div class="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
    <form class="space-y-6" on:submit|preventDefault={handleSubmit}>
      <div>
        <label for="username" class="block text-sm font-medium leading-6 text-gray-900">
          Username
        </label>
        <div class="mt-2">
          <input
            bind:value={username}
            id="username"
            name="username"
            type="text"
            required
            class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-blue-600 sm:text-sm sm:leading-6"
          />
        </div>
      </div>

      <div>
        <label for="password" class="block text-sm font-medium leading-6 text-gray-900">
          Password
        </label>
        <div class="mt-2">
          <input
            bind:value={password}
            id="password"
            name="password"
            type="password"
            required
            class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-blue-600 sm:text-sm sm:leading-6"
          />
        </div>
      </div>

      {#if error}
        <div class="text-red-500 text-sm">{error}</div>
      {/if}

      <div>
        <button
          type="submit"
          disabled={loading}
          class="flex w-full justify-center rounded-md bg-blue-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-blue-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-blue-600 disabled:opacity-50"
        >
          {loading ? 'Signing in...' : 'Sign in'}
        </button>
      </div>
    </form>
  </div>
</div>