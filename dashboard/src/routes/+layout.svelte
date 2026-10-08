<script>
  import '../app.css';
  import { page } from '$app/stores';
  import Sidebar from '$lib/components/Sidebar.svelte';
  import Navbar from '$lib/components/Navbar.svelte';
</script>

{#if $page.url.pathname === '/login'}
  <slot />
{:else}
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
