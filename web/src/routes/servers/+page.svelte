<script lang="ts">
  import { onMount } from 'svelte';
  import { auth } from '$lib/stores/auth';
  import { getServers, requestAccess } from '$lib/api';
  import type { Server } from '$lib/types';

  let servers: Server[] = [];
  let error = '';
  let loading = true;
  let requestingAccess = false;

  onMount(async () => {
    try {
      if ($auth.token) {
        servers = await getServers($auth.token);
      }
    } catch (e) {
      error = e.message;
    } finally {
      loading = false;
    }
  });

  async function handleRequestAccess(serverId: string) {
    if (!$auth.token) return;
    
    requestingAccess = true;
    error = '';

    try {
      await requestAccess({
        server_id: serverId,
        username: 'current-user' // Replace with actual username from auth store
      }, $auth.token);
      alert('Access request submitted successfully');
    } catch (e) {
      error = e.message;
    } finally {
      requestingAccess = false;
    }
  }
</script>

<div class="space-y-6">
  <div class="flex justify-between items-center">
    <h1 class="text-2xl font-semibold text-gray-900">Servers</h1>
    {#if $auth.role === 'admin'}
      <a 
        href="/servers/add" 
        class="bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded"
      >
        Add Server
      </a>
    {/if}
  </div>

  {#if loading}
    <div class="text-center">Loading servers...</div>
  {:else if error}
    <div class="text-red-500">{error}</div>
  {:else}
    <div class="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
      {#each servers as server}
        <div class="bg-white p-6 rounded-lg shadow">
          <h3 class="text-lg font-semibold text-gray-900">{server.name}</h3>
          <div class="mt-2 space-y-1">
            <p class="text-sm text-gray-600">Environment: {server.environment}</p>
            <p class="text-sm text-gray-600">IP: {server.ip}</p>
          </div>
          <button
            on:click={() => handleRequestAccess(server.id)}
            disabled={requestingAccess}
            class="mt-4 w-full bg-green-500 hover:bg-green-600 text-white px-4 py-2 rounded disabled:opacity-50"
          >
            {requestingAccess ? 'Requesting...' : 'Request Access'}
          </button>
        </div>
      {/each}
    </div>
  {/if}
</div>