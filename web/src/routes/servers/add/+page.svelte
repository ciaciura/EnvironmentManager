<script lang="ts">
  import { auth } from '$lib/stores/auth';
  import { addServer } from '$lib/api';
  import { goto } from '$app/navigation';
  import type { Server } from '$lib/types';

  let server: Server = {
    id: '',
    name: '',
    ip: '',
    environment: ''
  };
  let error = '';
  let loading = false;

  async function handleSubmit() {
    if (!$auth.token || $auth.role !== 'admin') {
      error = 'Unauthorized';
      return;
    }

    loading = true;
    error = '';

    try {
      await addServer(server, $auth.token);
      goto('/servers');
    } catch (e) {
      error = e.message;
    } finally {
      loading = false;
    }
  }
</script>

<div class="max-w-2xl mx-auto">
  <h1 class="text-2xl font-semibold text-gray-900 mb-6">Add New Server</h1>

  <form on:submit|preventDefault={handleSubmit} class="space-y-6">
    <div>
      <label for="name" class="block text-sm font-medium text-gray-700">Server Name</label>
      <input
        bind:value={server.name}
        type="text"
        id="name"
        required
        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
      />
    </div>

    <div>
      <label for="ip" class="block text-sm font-medium text-gray-700">IP Address</label>
      <input
        bind:value={server.ip}
        type="text"
        id="ip"
        required
        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
      />
    </div>

    <div>
      <label for="environment" class="block text-sm font-medium text-gray-700">Environment</label>
      <input
        bind:value={server.environment}
        type="text"
        id="environment"
        required
        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
      />
    </div>

    {#if error}
      <div class="text-red-500">{error}</div>
    {/if}

    <div class="flex justify-end space-x-4">
      <a
        href="/servers"
        class="bg-gray-200 hover:bg-gray-300 px-4 py-2 rounded"
      >
        Cancel
      </a>
      <button
        type="submit"
        disabled={loading}
        class="bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded disabled:opacity-50"
      >
        {loading ? 'Adding...' : 'Add Server'}
      </button>
    </div>
  </form>
</div>