<script lang="ts">
  import { onMount } from 'svelte';
  import { goto } from '$app/navigation';

  onMount(() => {
    goto('/login');
  });
</script>

<div class="flex justify-center items-center min-h-screen">
  <p class="text-gray-600">Redirecting to login...</p>
</div>