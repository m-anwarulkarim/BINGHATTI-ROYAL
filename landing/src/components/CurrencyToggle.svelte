<script>
  import { onMount } from 'svelte';
  import { currentCurrency, currencies } from '../lib/currencyStore.js';

  let selectedCurrency = 'AED';
  let isOpen = false;

  const currencyList = [
    { code: 'AED', name: 'UAE Dirham', symbol: 'AED' },
    { code: 'USD', name: 'US Dollar', symbol: '$' },
    { code: 'EUR', name: 'Euro', symbol: '€' },
    { code: 'GBP', name: 'British Pound', symbol: '£' }
  ];

  currentCurrency.subscribe(val => {
    selectedCurrency = val;
  });

  function toggleDropdown() {
    isOpen = !isOpen;
  }

  function selectCurrency(code) {
    selectedCurrency = code;
    currentCurrency.set(code);
    if (typeof localStorage !== 'undefined') {
      localStorage.setItem('binghatti_currency', code);
    }
    isOpen = false;

    if (typeof window !== 'undefined') {
      window.dispatchEvent(new CustomEvent('currencyChange', {
        detail: {
          currency: code,
          info: currencies[code]
        }
      }));
    }
  }

  function handleOutsideClick(event) {
    if (isOpen && !event.target.closest('.currency-dropdown-container')) {
      isOpen = false;
    }
  }

  onMount(() => {
    const savedCurrency = localStorage.getItem('binghatti_currency') || 'AED';
    if (savedCurrency && savedCurrency !== selectedCurrency) {
      selectedCurrency = savedCurrency;
      currentCurrency.set(savedCurrency);
    }
    
    // Broadcast currency state on mount
    window.dispatchEvent(new CustomEvent('currencyChange', {
      detail: {
        currency: selectedCurrency,
        info: currencies[selectedCurrency]
      }
    }));

    window.addEventListener('click', handleOutsideClick);
    return () => window.removeEventListener('click', handleOutsideClick);
  });
</script>

<div class="relative currency-dropdown-container inline-block text-left">
  <!-- Trigger Button -->
  <button
    type="button"
    on:click|stopPropagation={toggleDropdown}
    class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-xl bg-black/80 backdrop-blur-md border border-gold/40 hover:border-gold text-xs font-headline font-semibold text-white transition-all shadow-md cursor-pointer"
    aria-expanded={isOpen}
  >
    <!-- Coins SVG Icon -->
    <svg class="w-4 h-4 text-gold shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
    </svg>

    <span class="text-gold font-bold">{selectedCurrency}</span>
    <svg class="w-3.5 h-3.5 text-gold/80 transition-transform duration-200 {isOpen ? 'rotate-180' : ''}" fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
    </svg>
  </button>

  <!-- Custom Dropdown Menu -->
  {#if isOpen}
    <div
      class="absolute right-0 mt-2 w-44 rounded-2xl bg-black/95 backdrop-blur-2xl border-2 border-gold/60 shadow-[0_10px_35px_rgba(0,0,0,0.9)] py-2 z-50 animate-fade-in divide-y divide-white/10"
    >
      <div class="px-3 py-1.5 text-[9px] uppercase tracking-widest text-gold font-bold font-headline border-b border-white/10">
        Select Currency
      </div>

      <div class="py-1">
        {#each currencyList as item}
          <button
            type="button"
            on:click={() => selectCurrency(item.code)}
            class="w-full flex items-center justify-between px-3 py-2 text-xs text-left transition-colors cursor-pointer {selectedCurrency === item.code ? 'bg-gold/20 text-gold font-bold' : 'text-white/90 hover:bg-white/10 hover:text-white'}"
          >
            <div class="flex items-center gap-2">
              <span class="text-gold font-bold text-xs">{item.symbol}</span>
              <span>{item.code}</span>
              <span class="text-[10px] text-white/60">({item.name})</span>
            </div>
            {#if selectedCurrency === item.code}
              <svg class="w-3.5 h-3.5 text-gold" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M5 13l4 4L19 7"></path>
              </svg>
            {/if}
          </button>
        {/each}
      </div>
    </div>
  {/if}
</div>
