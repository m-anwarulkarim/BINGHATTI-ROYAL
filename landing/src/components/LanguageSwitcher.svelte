<script>
  import { onMount } from 'svelte';
  import { currentLang, setLanguage } from '../lib/langStore.js';

  let selectedLang = 'en';
  let isOpen = false;

  const languages = [
    { code: 'en', name: 'English', label: 'English (EN)' },
    { code: 'ar', name: 'العربية', label: 'العربية (AR)' }
  ];

  currentLang.subscribe(val => {
    selectedLang = val;
    if (typeof document !== 'undefined') {
      const htmlEl = document.documentElement;
      if (val === 'ar') {
        htmlEl.setAttribute('dir', 'rtl');
        htmlEl.setAttribute('lang', 'ar');
        htmlEl.classList.add('font-arabic');
      } else {
        htmlEl.setAttribute('dir', 'ltr');
        htmlEl.setAttribute('lang', 'en');
        htmlEl.classList.remove('font-arabic');
      }
    }
  });

  function toggleDropdown() {
    isOpen = !isOpen;
  }

  function selectLanguage(code) {
    if (selectedLang !== code) {
      setLanguage(code);
      if (typeof window !== 'undefined') {
        window.dispatchEvent(new CustomEvent('languageChange', {
          detail: { lang: code, dir: code === 'ar' ? 'rtl' : 'ltr' }
        }));
      }
    }
    isOpen = false;
  }

  function handleOutsideClick(event) {
    if (isOpen && !event.target.closest('.lang-dropdown-container')) {
      isOpen = false;
    }
  }

  onMount(() => {
    window.addEventListener('click', handleOutsideClick);
    return () => window.removeEventListener('click', handleOutsideClick);
  });
</script>

<div class="relative lang-dropdown-container inline-block text-left">
  <!-- Trigger Button -->
  <button
    type="button"
    on:click|stopPropagation={toggleDropdown}
    class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-xl bg-black/80 backdrop-blur-md border border-gold/40 hover:border-gold text-xs font-headline font-semibold text-white transition-all shadow-md cursor-pointer"
    aria-expanded={isOpen}
    title="Select Language"
  >
    <!-- Globe Vector SVG Icon -->
    <svg class="w-4 h-4 text-gold shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3.055 11H5a2 2 0 012 2v1a2 2 0 002 2 2 2 0 012 2v2.945M8 3.935V5.5A2.5 2.5 0 0010.5 8h.5a2 2 0 012 2 2 2 0 104 0 2 2 0 012-2h1.064M15 20.488V18a2 2 0 012-2h3.064M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
    </svg>

    <span class="text-gold font-bold">{selectedLang === 'en' ? 'EN' : 'AR'}</span>

    <!-- Chevron Arrow SVG Icon -->
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
        Select Language
      </div>

      <div class="py-1">
        {#each languages as item}
          <button
            type="button"
            on:click={() => selectLanguage(item.code)}
            class="w-full flex items-center justify-between px-3 py-2 text-xs text-left transition-colors cursor-pointer {selectedLang === item.code ? 'bg-gold/20 text-gold font-bold' : 'text-white/90 hover:bg-white/10 hover:text-white'}"
          >
            <div class="flex items-center gap-2">
              <span class="text-gold font-bold text-xs">{item.code.toUpperCase()}</span>
              <span>{item.name}</span>
            </div>
            {#if selectedLang === item.code}
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
