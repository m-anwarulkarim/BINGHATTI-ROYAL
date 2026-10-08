<script>
  import { onMount } from 'svelte';

  export let isOpen = false;
  export let unitTitle = 'Luxury Floor Plan PDF';

  let fullName = '';
  let whatsappNumber = '';
  let email = '';
  let budgetRange = '$500k-$1M';
  let isLoading = false;
  let errorMessage = '';
  let successMessage = '';

  function closeModal() {
    isOpen = false;
    errorMessage = '';
    successMessage = '';
  }

  onMount(() => {
    // Listen for custom trigger event across the page
    const handleTrigger = (e) => {
      if (e.detail?.unitTitle) {
        unitTitle = e.detail.unitTitle;
      }
      isOpen = true;
    };

    window.addEventListener('openFloorPlanModal', handleTrigger);
    return () => window.removeEventListener('openFloorPlanModal', handleTrigger);
  });

  async function handleSubmit() {
    isLoading = true;
    errorMessage = '';
    successMessage = '';

    const payload = {
      full_name: fullName,
      whatsapp_number: whatsappNumber,
      email: email,
      budget_range: budgetRange,
      investment_purpose: 'Investment',
      lead_source: `Floor Plan Modal - ${unitTitle}`
    };

    try {
      const res = await fetch('http://localhost:8080/api/v1/leads', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload)
      });

      const data = await res.json();

      if (res.ok && data.success) {
        successMessage = 'Lead Verified! Downloading high-resolution Floor Plan PDF...';
        
        // Trigger direct browser PDF download simulation
        setTimeout(() => {
          const dummyPdfUrl = 'https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf';
          const link = document.createElement('a');
          link.href = dummyPdfUrl;
          link.download = `Binghatti_${unitTitle.replace(/\s+/g, '_')}_FloorPlan.pdf`;
          link.target = '_blank';
          document.body.appendChild(link);
          link.click();
          document.body.removeChild(link);
          
          setTimeout(() => {
            closeModal();
          }, 2000);
        }, 1000);
      } else {
        throw new Error(data.error || 'Failed to verify lead.');
      }
    } catch (err) {
      errorMessage = err.message || 'Submission error. Please check backend connection.';
    } finally {
      isLoading = false;
    }
  }
</script>

{#if isOpen}
  <!-- Backdrop -->
  <div
    class="fixed inset-0 z-50 flex items-center justify-center bg-black/80 backdrop-blur-md p-4 animate-fade-in"
    on:click|self={closeModal}
  >
    <!-- Modal Dialog -->
    <div class="relative w-full max-w-lg rounded-2xl glass-card border-2 border-gold/60 p-8 shadow-[0_0_50px_rgba(212,175,55,0.3)] bg-bg-offblack">
      
      <!-- Close Button -->
      <button
        on:click={closeModal}
        class="absolute top-4 right-4 text-muted hover:text-white p-2 rounded-lg hover:bg-white/10 transition-colors"
        aria-label="Close modal"
      >
        <svg class="w-5 h-5 text-gold" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
        </svg>
      </button>

      <!-- Modal Header -->
      <div class="text-center mb-6">
        <span class="text-[10px] uppercase tracking-[0.3em] text-gold font-headline font-bold block mb-1">
          Gated Luxury Access
        </span>
        <h3 class="text-2xl font-bold font-headline text-white">
          Unlock {unitTitle}
        </h3>
        <p class="text-xs text-muted mt-1 font-light">
          Enter your contact details to download official 2D/3D high-res floor plans.
        </p>
      </div>

      <!-- Form -->
      <form on:submit|preventDefault={handleSubmit} class="space-y-4">
        <div>
          <label class="block text-xs font-headline text-muted uppercase mb-1">Full Name *</label>
          <input
            type="text"
            bind:value={fullName}
            required
            placeholder="e.g. Marcus Aurelius"
            class="w-full px-4 py-3 rounded-xl bg-bg-dark border border-white/10 text-white focus:border-gold focus:outline-none text-sm"
          />
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <div>
            <label class="block text-xs font-headline text-muted uppercase mb-1">WhatsApp Number *</label>
            <input
              type="tel"
              bind:value={whatsappNumber}
              required
              placeholder="+971 50 000 0000"
              class="w-full px-4 py-3 rounded-xl bg-bg-dark border border-white/10 text-white focus:border-gold focus:outline-none text-sm"
            />
          </div>

          <div>
            <label class="block text-xs font-headline text-muted uppercase mb-1">Email Address *</label>
            <input
              type="email"
              bind:value={email}
              required
              placeholder="marcus@vip.com"
              class="w-full px-4 py-3 rounded-xl bg-bg-dark border border-white/10 text-white focus:border-gold focus:outline-none text-sm"
            />
          </div>
        </div>

        <div>
          <label class="block text-xs font-headline text-muted uppercase mb-1">Estimated Budget</label>
          <select
            bind:value={budgetRange}
            class="w-full px-4 py-3 rounded-xl bg-bg-dark border border-white/10 text-white focus:border-gold focus:outline-none text-sm"
          >
            <option value="$300k-$500k">$300,000 – $500,000</option>
            <option value="$500k-$1M">$500,000 – $1,000,000</option>
            <option value="$1M+">$1,000,000+</option>
          </select>
        </div>

        {#if errorMessage}
          <div class="p-3 rounded-xl bg-red-900/30 border border-red-500/40 text-red-400 text-xs font-semibold text-center">
            {errorMessage}
          </div>
        {/if}

        {#if successMessage}
          <div class="p-3 rounded-xl bg-gold/20 border border-gold/40 text-gold text-xs font-semibold text-center">
            {successMessage}
          </div>
        {/if}

        <button
          type="submit"
          disabled={isLoading}
          class="w-full py-4 rounded-xl bg-gold-gradient text-black font-headline font-bold uppercase tracking-wider text-xs hover:scale-[1.02] transition-transform duration-300 shadow-[0_0_20px_rgba(212,175,55,0.4)]"
        >
          {isLoading ? 'Verifying & Unlocking PDF...' : 'Download Floor Plan PDF Now'}
        </button>
      </form>

    </div>
  </div>
{/if}
