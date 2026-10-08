<script>
  import { clearAuth } from '../authStore.js';
  import { toggleSidebar } from '../uiStore.js';
  import { goto } from '$app/navigation';

  function handleLogout() {
    clearAuth();
    goto('/login');
  }
</script>

<header class="h-16 border-b border-white/10 bg-bg-offblack/80 backdrop-blur-md px-4 sm:px-6 lg:px-8 flex items-center justify-between sticky top-0 z-30">
  <div class="flex items-center gap-3">
    <!-- Hamburger Button for Mobile Drawer -->
    <button
      on:click={toggleSidebar}
      class="p-2 rounded-xl bg-bg-dark border border-white/10 text-gold hover:text-white lg:hidden flex items-center justify-center"
      aria-label="Toggle Navigation Sidebar"
    >
      <svg class="w-5 h-5 text-gold" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"></path>
      </svg>
    </button>

    <!-- Search Input -->
    <div class="relative w-40 sm:w-64 md:w-72">
      <span class="absolute inset-y-0 left-3 flex items-center text-muted">
        <svg class="w-3.5 h-3.5 text-muted" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
        </svg>
      </span>
      <input
        type="text"
        placeholder="Search leads, emails..."
        class="w-full pl-9 pr-3 py-1.5 rounded-xl bg-bg-dark border border-white/10 text-white focus:border-gold focus:outline-none text-xs"
      />
    </div>
  </div>

  <div class="flex items-center gap-2 sm:gap-4">
    <!-- Live Status Indicator -->
    <div class="hidden sm:flex items-center gap-2 px-3 py-1 rounded-full bg-emerald-950/40 border border-emerald-500/30 text-emerald-400 text-[10px] font-semibold">
      <span class="w-2 h-2 rounded-full bg-emerald-400 animate-pulse"></span>
      <span>Live Sync Connected</span>
    </div>

    <!-- Notification Bell -->
    <button class="w-9 h-9 rounded-xl bg-bg-dark border border-white/10 flex items-center justify-center text-muted hover:text-white hover:border-gold/30 transition-colors relative" aria-label="Notifications">
      <svg class="w-4 h-4 text-gold" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"></path>
      </svg>
      <span class="absolute -top-1 -right-1 w-2.5 h-2.5 rounded-full bg-gold"></span>
    </button>

    <!-- Logout Button -->
    <button
      on:click={handleLogout}
      class="px-3 py-1.5 rounded-xl bg-red-950/30 border border-red-500/30 text-red-400 hover:bg-red-900/50 hover:text-red-200 text-xs font-semibold font-headline transition-colors"
    >
      Logout
    </button>
  </div>
</header>
