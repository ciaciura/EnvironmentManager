<script lang="ts">
  import { auth } from '$lib/stores/auth';
  import { createUser } from '$lib/api';
  import type { User } from '$lib/types';

  let user: User = {
    username: '',
    password: '',
    role: 'employee'
  };
  let error = '';
  let success = '';
  let loading = false;

  async function handleSubmit() {
    if (!$auth.token || $auth.role !== 'admin') {
      error = 'Unauthorized';
      return;
    }

    loading = true;
    error = '';
    success = '';

    try {
      await createUser(user, $auth.token);
      success = 'User created successfully';
      user = {
        username: '',
        password: '',
        role: 'employee'
      };
    } catch (e) {
      error = e.message;
    } finally {
      loading = false;
    }
  }
</script>

<div class="max-w-2xl mx-auto">
  <h1 class="text-2xl font-semibold text-gray-900 mb-6">Admin Panel - Create User</h1>

  <form on:submit|preventDefault={handleSubmit} class="space-y-6">
    <div>
      <label for="username" class="block text-sm font-medium text-gray-700">Username</label>
      <input
        bind:value={user.username}
        type="text"
        id="username"
        required
        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
      />
    </div>

    <div>
      <label for="password" class="block text-sm font-medium text-gray-700">Password</label>
      <input
        bind:value={user.password}
        type="password"
        id="password"
        required
        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
      />
    </div>

    <div>
      <label for="role" class="block text-sm font-medium text-gray-700">Role</label>
      <select
        bind:value={user.role}
        id="role"
        required
        class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
      >
        <option value="employee">Employee</option>
        <option value="admin">Admin</option>
      </select>
    </div>

    {#if error}
      <div class="text-red-500">{error}</div>
    {/if}

    {#if success}
      <div class="text-green-500">{success}</div>
    {/if}

    <button
      type="submit"
      disabled={loading}
      class="w-full bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded disabled:opacity-50"
    >
      {loading ? 'Creating...' : 'Create User'}
    </button>
  </form>
</div>