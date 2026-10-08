<script>
  import '../app.css';
  import { page } from '$app/stores';
  import { goto } from '$app/navigation';
  import { onMount } from 'svelte';
  import { authToken } from '$lib/authStore.js';
  import Sidebar from '$lib/components/Sidebar.svelte';
  import Navbar from '$lib/components/Navbar.svelte';

  $: isLoginPage = $page.url.pathname === '/login';

  onMount(() => {
    checkAuth();
  });

  $: if (typeof window !== 'undefined' && $page.url.pathname) {
    checkAuth();
  }

  function checkAuth() {
    if (typeof window === 'undefined') return;
    const token = localStorage.getItem('binghatti_token');
    
    if (!token && !isLoginPage) {
      goto('/login');
    } else if (token && isLoginPage) {
      goto('/');
    }
  }
</script>

{#if isLoginPage}
  <slot />
{:else if $authToken || (typeof window !== 'undefined' && localStorage.getItem('binghatti_token'))}
  <div class="flex h-screen bg-bg-offblack overflow-hidden">
    <Sidebar />
    <div class="flex-1 flex flex-col min-w-0 h-screen overflow-hidden">
      <Navbar />
      <main class="flex-1 p-4 sm:p-6 lg:p-8 overflow-y-auto">
        <slot />
      </main>
    </div>
  </div>
{/if}
