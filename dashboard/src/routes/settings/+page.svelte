<script>
  let fullName = "Senior Advisor";
  let email = "sales@binghatti.com";
  let role = "Super Admin / Luxury Sales Lead";
  let department = "Binghatti VIP Investment Division";
  
  let apiEndpoint = "http://localhost:8085/api/v1";
  let dbConnection = "PostgreSQL (Neon Cloud DB)";
  let whatsappKey = "wh_live_99482710398271a0";
  let autoSyncLeads = true;

  let defaultCurrency = "AED";
  let timezone = "Asia/Dubai (GST +04:00)";
  let defaultLanguage = "en";

  let notifyEmail = true;
  let notifyWhatsapp = true;
  let dailyDigest = false;

  let showToast = false;
  let toastMessage = "";

  function saveSettings() {
    toastMessage = "Platform settings saved successfully!";
    showToast = true;
    setTimeout(() => {
      showToast = false;
    }, 3500);
  }
</script>

<div class="space-y-8 relative pb-12">
  
  <!-- Toast Notification -->
  {#if showToast}
    <div class="fixed bottom-6 right-6 z-50 px-5 py-3 rounded-xl bg-emerald-500 text-black font-headline font-bold text-xs uppercase tracking-wider shadow-2xl flex items-center gap-3 transition-all animate-bounce">
      <svg class="w-5 h-5 text-black" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
      </svg>
      <span>{toastMessage}</span>
    </div>
  {/if}

  <!-- Page Header -->
  <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
    <div>
      <h1 class="text-2xl font-bold font-headline text-white">CRM System Settings</h1>
      <p class="text-xs text-muted">Manage platform preferences, API credentials, notification alerts, and localization for Binghatti CRM.</p>
    </div>

    <button
      on:click={saveSettings}
      class="inline-flex items-center gap-2 px-6 py-2.5 rounded-xl bg-gold text-black font-headline font-bold text-xs uppercase tracking-wider hover:bg-gold/90 transition-all shadow-lg shadow-gold/20"
    >
      <svg class="w-4 h-4 text-black" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
      </svg>
      <span>Save Changes</span>
    </button>
  </div>

  <!-- Settings Grid -->
  <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
    
    <!-- 1. Advisor Profile & Account Credentials -->
    <div class="glass-card p-6 rounded-2xl border border-white/10 space-y-5">
      <div class="flex items-center gap-3 border-b border-white/10 pb-4">
        <div class="p-2.5 rounded-xl bg-gold/15 border border-gold/30 text-gold">
          <svg class="w-5 h-5 text-gold" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
          </svg>
        </div>
        <div>
          <h3 class="text-base font-bold font-headline text-white">Advisor Profile & Account</h3>
          <p class="text-[11px] text-muted">User identity and administrative privileges</p>
        </div>
      </div>

      <div class="space-y-4 text-xs">
        <div>
          <label for="advisor-fullname" class="block text-[10px] uppercase tracking-wider font-headline text-muted mb-1">Full Display Name</label>
          <input
            id="advisor-fullname"
            type="text"
            bind:value={fullName}
            class="w-full select-luxury py-2.5 px-3"
          />
        </div>

        <div>
          <label for="advisor-email" class="block text-[10px] uppercase tracking-wider font-headline text-muted mb-1">Email Address</label>
          <input
            id="advisor-email"
            type="email"
            bind:value={email}
            class="w-full select-luxury py-2.5 px-3"
          />
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <div>
            <label for="advisor-role" class="block text-[10px] uppercase tracking-wider font-headline text-muted mb-1">Role Title</label>
            <input
              id="advisor-role"
              type="text"
              bind:value={role}
              disabled
              class="w-full px-3 py-2.5 rounded-lg bg-white/5 border border-white/10 text-muted cursor-not-allowed"
            />
          </div>
          <div>
            <label for="advisor-dept" class="block text-[10px] uppercase tracking-wider font-headline text-muted mb-1">Department</label>
            <input
              id="advisor-dept"
              type="text"
              bind:value={department}
              class="w-full select-luxury py-2.5 px-3"
            />
          </div>
        </div>
      </div>
    </div>

    <!-- 2. API & Integration Hub -->
    <div class="glass-card p-6 rounded-2xl border border-white/10 space-y-5">
      <div class="flex items-center gap-3 border-b border-white/10 pb-4">
        <div class="p-2.5 rounded-xl bg-gold/15 border border-gold/30 text-gold">
          <svg class="w-5 h-5 text-gold" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 9l3 3-3 3m5 0h3M5 20h14a2 2 0 002-2V6a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
          </svg>
        </div>
        <div>
          <h3 class="text-base font-bold font-headline text-white">API & Integration Hub</h3>
          <p class="text-[11px] text-muted">Backend endpoint, database and WhatsApp keys</p>
        </div>
      </div>

      <div class="space-y-4 text-xs">
        <div>
          <label for="api-endpoint-url" class="block text-[10px] uppercase tracking-wider font-headline text-muted mb-1">REST API Service URL</label>
          <input
            id="api-endpoint-url"
            type="text"
            bind:value={apiEndpoint}
            class="w-full px-3 py-2.5 rounded-lg bg-bg-dark border border-white/10 text-white focus:border-gold focus:outline-none font-mono text-[11px]"
          />
        </div>

        <div>
          <label for="db-connection-string" class="block text-[10px] uppercase tracking-wider font-headline text-muted mb-1">Database Cluster Target</label>
          <input
            id="db-connection-string"
            type="text"
            bind:value={dbConnection}
            class="w-full px-3 py-2.5 rounded-lg bg-bg-dark border border-white/10 text-white focus:border-gold focus:outline-none font-mono text-[11px]"
          />
        </div>

        <div>
          <label for="whatsapp-gateway-key" class="block text-[10px] uppercase tracking-wider font-headline text-muted mb-1">WhatsApp Business API Key</label>
          <input
            id="whatsapp-gateway-key"
            type="password"
            bind:value={whatsappKey}
            class="w-full px-3 py-2.5 rounded-lg bg-bg-dark border border-white/10 text-white focus:border-gold focus:outline-none font-mono text-[11px]"
          />
        </div>

        <div class="flex items-center justify-between pt-2 border-t border-white/5">
          <div>
            <span class="text-xs font-bold text-white block">Auto-Sync Landing Page Leads</span>
            <span class="text-[10px] text-muted">Fetch new registrations into Kanban pipeline automatically</span>
          </div>
          <input
            type="checkbox"
            bind:checked={autoSyncLeads}
            class="w-4 h-4 accent-gold rounded cursor-pointer"
          />
        </div>
      </div>
    </div>

    <!-- 3. Regional & Currency Localization -->
    <div class="glass-card p-6 rounded-2xl border border-white/10 space-y-5">
      <div class="flex items-center gap-3 border-b border-white/10 pb-4">
        <div class="p-2.5 rounded-xl bg-gold/15 border border-gold/30 text-gold">
          <svg class="w-5 h-5 text-gold" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5h12M9 3v2m1.048 9.5A18.022 18.022 0 016.412 9m6.088 9h7M11 21l5-10 5 10M12.751 5C11.783 10.77 8.07 15.61 3 18.129"></path>
          </svg>
        </div>
        <div>
          <h3 class="text-base font-bold font-headline text-white">Regional & Currency Settings</h3>
          <p class="text-[11px] text-muted">Preferred currency formats and regional timezones</p>
        </div>
      </div>

      <div class="space-y-4 text-xs">
        <div>
          <label for="default-currency-select" class="block text-[10px] uppercase tracking-wider font-headline text-muted mb-1">Default Display Currency</label>
          <select
            id="default-currency-select"
            bind:value={defaultCurrency}
            class="w-full select-luxury py-2.5 px-3"
          >
            <option value="AED">AED - United Arab Emirates Dirham</option>
            <option value="USD">USD - United States Dollar ($)</option>
            <option value="EUR">EUR - Euro (€)</option>
            <option value="GBP">GBP - British Pound (£)</option>
          </select>
        </div>

        <div>
          <label for="default-timezone-select" class="block text-[10px] uppercase tracking-wider font-headline text-muted mb-1">Timezone</label>
          <select
            id="default-timezone-select"
            bind:value={timezone}
            class="w-full select-luxury py-2.5 px-3"
          >
            <option value="Asia/Dubai (GST +04:00)">Asia/Dubai (GST +04:00)</option>
            <option value="Europe/London (GMT +00:00)">Europe/London (GMT +00:00)</option>
            <option value="America/New_York (EST -05:00)">America/New_York (EST -05:00)</option>
            <option value="Asia/Riyadh (AST +03:00)">Asia/Riyadh (AST +03:00)</option>
          </select>
        </div>

        <div>
          <label for="default-language-select" class="block text-[10px] uppercase tracking-wider font-headline text-muted mb-1">Interface Language</label>
          <select
            id="default-language-select"
            bind:value={defaultLanguage}
            class="w-full select-luxury py-2.5 px-3"
          >
            <option value="en">English (US)</option>
            <option value="ar">??????? (Arabic)</option>
            <option value="fr">Français (French)</option>
          </select>
        </div>
      </div>
    </div>

    <!-- 4. Automated Notifications & Alerts -->
    <div class="glass-card p-6 rounded-2xl border border-white/10 space-y-5">
      <div class="flex items-center gap-3 border-b border-white/10 pb-4">
        <div class="p-2.5 rounded-xl bg-gold/15 border border-gold/30 text-gold">
          <svg class="w-5 h-5 text-gold" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"></path>
          </svg>
        </div>
        <div>
          <h3 class="text-base font-bold font-headline text-white">Automated Lead Alerts</h3>
          <p class="text-[11px] text-muted">Real-time alerts for VIP high-net-worth inquiries</p>
        </div>
      </div>

      <div class="space-y-4 text-xs">
        <div class="flex items-center justify-between py-2 border-b border-white/5">
          <div>
            <span class="text-xs font-bold text-white block">Instant Email Alert ($1M+ Leads)</span>
            <span class="text-[10px] text-muted">Receive email immediately when a $1M+ budget lead registers</span>
          </div>
          <input
            type="checkbox"
            bind:checked={notifyEmail}
            class="w-4 h-4 accent-gold rounded cursor-pointer"
          />
        </div>

        <div class="flex items-center justify-between py-2 border-b border-white/5">
          <div>
            <span class="text-xs font-bold text-white block">WhatsApp Push Notification</span>
            <span class="text-[10px] text-muted">Ping senior advisor's WhatsApp when a lead changes stage to Qualified</span>
          </div>
          <input
            type="checkbox"
            bind:checked={notifyWhatsapp}
            class="w-4 h-4 accent-gold rounded cursor-pointer"
          />
        </div>

        <div class="flex items-center justify-between py-2">
          <div>
            <span class="text-xs font-bold text-white block">Daily Pipeline Summary Digest</span>
            <span class="text-[10px] text-muted">Send daily performance email digest at 09:00 AM GST</span>
          </div>
          <input
            type="checkbox"
            bind:checked={dailyDigest}
            class="w-4 h-4 accent-gold rounded cursor-pointer"
          />
        </div>
      </div>
    </div>

  </div>

  <!-- Bottom Save Bar -->
  <div class="pt-4 flex justify-end">
    <button
      on:click={saveSettings}
      class="inline-flex items-center gap-2 px-8 py-3.5 rounded-xl bg-gold text-black font-headline font-bold text-xs uppercase tracking-wider hover:bg-gold/90 transition-all shadow-xl shadow-gold/20"
    >
      <svg class="w-4 h-4 text-black" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
      </svg>
      <span>Save All Platform Settings</span>
    </button>
  </div>

</div>
