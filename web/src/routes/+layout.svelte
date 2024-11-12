<script lang="ts">
  import '../app.css';
  import { auth } from '$lib/stores/auth';
  import { page } from '$app/stores';
  import { goto } from '$app/navigation';

  function logout() {
    auth.clearAuth();
    goto('/login');
  }
</script>

<div class="min-h-screen bg-gray-100">
  <nav class="bg-white shadow-lg">
    <div class="max-w-7xl mx-auto px-4">
      <div class="flex justify-between h-16">
        <div class="flex">
          <div class="flex-shrink-0 flex items-center">
            <a href="/" class="text-xl font-bold text-gray-800">Environment Manager</a>
          </div>
        </div>
        
        {#if $auth.token}
          <div class="flex items-center space-x-4">
            <a href="/servers" class="text-gray-700 hover:text-gray-900">Servers</a>
            {#if $auth.role === 'admin'}
              <a href="/admin" class="text-gray-700 hover:text-gray-900">Admin</a>
            {/if}
            <button 
              on:click={logout}
              class="bg-red-500 hover:bg-red-600 text-white px-4 py-2 rounded"
            >
              Logout
            </button>
          </div>
        {/if}
      </div>
    </div>
  </nav>

  <main class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
    <slot />
  </main>
</div>