<script>
  import { page } from '$app/stores';
  import { isSidebarOpen, closeSidebar } from '../uiStore.js';

  const navItems = [
    { label: 'Dashboard Overview', href: '/', svgPath: 'M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6' },
    { label: 'Lead CRM Pipeline', href: '/leads', svgPath: 'M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10' },
    { label: 'Unit Inventory', href: '/inventory', svgPath: 'M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5m0 0v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4' },
    { label: 'Analytics & ROI', href: '/analytics', svgPath: 'M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z' },
    { label: 'Settings', href: '/settings', svgPath: 'M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z' }
  ];

  // Auto-close sidebar on mobile route navigation
  $: if ($page.url.pathname) {
    closeSidebar();
  }
</script>

<!-- Mobile Backdrop Overlay -->
{#if $isSidebarOpen}
  <div
    role="button"
    tabindex="0"
    on:click={closeSidebar}
    on:keydown={(e) => e.key === 'Escape' && closeSidebar()}
    class="fixed inset-0 bg-black/80 backdrop-blur-sm z-40 lg:hidden transition-opacity duration-300"
  ></div>
{/if}

<aside
  class="fixed lg:sticky top-0 left-0 z-50 w-64 bg-bg-dark border-r border-white/10 flex flex-col justify-between shrink-0 h-screen transition-transform duration-300 ease-in-out {$isSidebarOpen ? 'translate-x-0' : '-translate-x-full lg:translate-x-0'}"
>
  <div>
    <!-- Logo Header & Mobile Close Button -->
    <div class="p-6 border-b border-white/10 flex items-center justify-between">
      <div class="flex items-center gap-3">
        <div class="w-8 h-8 rounded-lg bg-gold/20 border border-gold flex items-center justify-center font-bold text-gold font-headline">
          B
        </div>
        <div>
          <span class="font-headline font-bold text-lg text-gold block leading-tight">BINGHATTI</span>
          <span class="text-[9px] uppercase tracking-widest text-muted">Admin CRM Platform</span>
        </div>
      </div>

      <!-- Close button on Mobile -->
      <button
        on:click={closeSidebar}
        class="p-1.5 rounded-lg text-muted hover:text-white hover:bg-white/10 lg:hidden"
        aria-label="Close Sidebar"
      >
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
        </svg>
      </button>
    </div>

    <!-- Nav Items -->
    <nav class="p-4 space-y-1">
      {#each navItems as item}
        <a
          href={item.href}
          class="flex items-center gap-3 px-4 py-3 rounded-xl text-xs font-semibold font-headline transition-colors {$page.url.pathname === item.href ? 'bg-gold/15 text-gold border border-gold/30' : 'text-muted hover:text-white hover:bg-white/5'}"
        >
          <svg class="w-4 h-4 {$page.url.pathname === item.href ? 'text-gold' : 'text-muted'}" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d={item.svgPath}></path>
          </svg>
          <span>{item.label}</span>
        </a>
      {/each}
    </nav>
  </div>

  <!-- Bottom Admin Info Badge -->
  <div class="p-4 border-t border-white/10">
    <div class="p-3 rounded-xl bg-bg-offblack border border-white/5 flex items-center justify-between">
      <div class="flex items-center gap-3">
        <div class="w-8 h-8 rounded-full bg-gold text-black font-bold flex items-center justify-center text-xs font-headline">
          SA
        </div>
        <div>
          <span class="text-xs font-bold text-white block">Senior Advisor</span>
          <span class="text-[10px] text-muted">sales@binghatti.com</span>
        </div>
      </div>
    </div>
  </div>
</aside>
