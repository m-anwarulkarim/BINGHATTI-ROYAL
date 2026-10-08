<script>
  import { setAuth } from '$lib/authStore.js';
  import { goto } from '$app/navigation';
  import { onMount } from 'svelte';

  let email = 'dev.anwarul@gmail.com';
  let password = 'dev.anwarul';
  let isLoading = false;
  let errorMessage = '';

  onMount(() => {
    if (typeof localStorage !== 'undefined' && localStorage.getItem('binghatti_token')) {
      goto('/');
    }
  });

  async function handleLogin() {
    isLoading = true;
    errorMessage = '';

    const cleanEmail = (email || '').trim().toLowerCase();
    const cleanPassword = (password || '').trim();

    // Verify authorized credentials: dev.anwarul@gmail.com / dev.anwarul
    if (cleanEmail === 'dev.anwarul@gmail.com' && cleanPassword === 'dev.anwarul') {
      setTimeout(() => {
        setAuth('binghatti-vip-advisor-jwt-2026', { 
          full_name: 'Anwarul Karim (Senior Advisor)', 
          role: 'super_admin', 
          email: 'dev.anwarul@gmail.com' 
        });
        isLoading = false;
        goto('/');
      }, 400);
      return;
    }

    // Try API auth if live API is available
    try {
      const API_BASE = (typeof import.meta !== 'undefined' && import.meta.env && import.meta.env.PUBLIC_API_URL) 
        ? import.meta.env.PUBLIC_API_URL 
        : 'https://binghatti-royal.onrender.com';

      const response = await fetch(`${API_BASE}/api/v1/auth/login`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email: cleanEmail, password: cleanPassword })
      });

      const data = await response.json();

      if (response.ok && data.success && data.data.token) {
        setAuth(data.data.token, data.data.user || { full_name: 'Anwarul Karim', email: cleanEmail });
        goto('/');
        return;
      }
    } catch (err) {}

    // Invalid Credentials Error
    isLoading = false;
    errorMessage = 'Incorrect Gmail or Password! Please use authorized credentials.';
  }
</script>

<div class="min-h-screen w-full flex items-center justify-center bg-bg-offblack px-4 relative overflow-hidden">
  <div class="absolute inset-0 bg-[radial-gradient(circle_at_center,rgba(212,175,55,0.12)_0,transparent_70%)]"></div>

  <div class="relative w-full max-w-md p-8 rounded-2xl glass-card border border-gold/40 shadow-[0_0_50px_rgba(212,175,55,0.2)] bg-bg-dark space-y-6">
    <div class="text-center">
      <div class="w-14 h-14 rounded-2xl bg-gold/20 border-2 border-gold text-gold font-black font-headline text-3xl flex items-center justify-center mx-auto mb-4 shadow-[0_0_20px_rgba(212,175,55,0.4)]">
        B
      </div>
      <h1 class="text-2xl font-bold font-headline text-white tracking-wide">Binghatti Admin CRM</h1>
      <p class="text-xs text-gold/80 mt-1 font-headline uppercase tracking-widest font-semibold">Authorized Senior Advisor Portal</p>
    </div>

    <form on:submit|preventDefault={handleLogin} class="space-y-4">
      <div>
        <label for="login-email" class="block text-[11px] font-headline text-muted uppercase tracking-wider mb-1.5 font-semibold">Advisor Gmail</label>
        <div class="relative">
          <input
            id="login-email"
            type="email"
            bind:value={email}
            required
            placeholder="dev.anwarul@gmail.com"
            class="w-full px-4 py-3 rounded-xl bg-bg-offblack border border-white/15 text-white focus:border-gold focus:ring-1 focus:ring-gold/30 focus:outline-none text-xs font-mono"
          />
        </div>
      </div>

      <div>
        <label for="login-password" class="block text-[11px] font-headline text-muted uppercase tracking-wider mb-1.5 font-semibold">Password</label>
        <div class="relative">
          <input
            id="login-password"
            type="password"
            bind:value={password}
            required
            placeholder="••••••••"
            class="w-full px-4 py-3 rounded-xl bg-bg-offblack border border-white/15 text-white focus:border-gold focus:ring-1 focus:ring-gold/30 focus:outline-none text-xs font-mono"
          />
        </div>
      </div>

      {#if errorMessage}
        <div class="p-3 rounded-xl bg-red-950/60 border border-red-500/50 text-red-400 text-xs font-semibold text-center animate-pulse">
          {errorMessage}
        </div>
      {/if}

      <button
        type="submit"
        disabled={isLoading}
        class="w-full py-3.5 rounded-xl bg-gold text-black font-headline font-bold uppercase tracking-wider text-xs hover:bg-gold/90 transition-all duration-300 shadow-[0_0_25px_rgba(212,175,55,0.4)] cursor-pointer mt-2"
      >
        {isLoading ? 'Authenticating Credentials...' : 'Authenticate & Enter CRM'}
      </button>
    </form>

    <div class="pt-4 border-t border-white/10 text-center">
      <span class="text-[10px] text-white/40 uppercase tracking-widest font-mono">Restricted Access • Binghatti Royal Platform</span>
    </div>
  </div>
</div>
