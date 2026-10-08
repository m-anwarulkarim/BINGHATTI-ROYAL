<script>
  import { leadsStore } from '$lib/leadsStore';

  const metrics = [
    {
      title: "Total Pipeline Value",
      value: "$142.5M",
      change: "+14.2% vs last month",
      trend: "up",
      svgPath: "M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V6m0 8v2m0-6c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
    },
    {
      title: "Active VIP Leads",
      value: `${$leadsStore.length} Leads`,
      change: "+24.5% new inquiries",
      trend: "up",
      svgPath: "M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"
    },
    {
      title: "Available Inventory",
      value: "42 Units",
      change: "12 Reserved, 8 Sold",
      trend: "neutral",
      svgPath: "M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5m0 0v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4"
    },
    {
      title: "Avg Conversion Rate",
      value: "18.5%",
      change: "+3.2% stage progression",
      trend: "up",
      svgPath: "M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"
    }
  ];

  const flagshipProjects = [
    {
      name: "Mercedes-Benz Places",
      location: "Downtown Dubai",
      unitsAvailable: "18 Units Available",
      startingPrice: "AED 2,200,000",
      type: "Luxury 1, 2 & 3 BR Residences",
      badgeColor: "bg-emerald-500/10 text-emerald-400 border-emerald-500/30"
    },
    {
      name: "Bugatti Residences",
      location: "Business Bay, Dubai",
      unitsAvailable: "8 Units Available",
      startingPrice: "AED 19,000,000",
      type: "Ultra-Luxury Penthouse Collection",
      badgeColor: "bg-gold/10 text-gold border-gold/30"
    },
    {
      name: "Jacob & Co Residences",
      location: "Burj Crown District",
      unitsAvailable: "16 Units Available",
      startingPrice: "AED 8,500,000",
      type: "Branded Sky Mansions & Penthouses",
      badgeColor: "bg-blue-500/10 text-blue-400 border-blue-500/30"
    }
  ];

  function getStatusBadge(status) {
    switch (status) {
      case 'new':
        return { label: 'New Lead', color: 'bg-blue-500/15 text-blue-400 border-blue-500/30' };
      case 'contacted':
        return { label: 'Contacted', color: 'bg-amber-500/15 text-amber-400 border-amber-500/30' };
      case 'qualified':
        return { label: 'Qualified', color: 'bg-purple-500/15 text-purple-400 border-purple-500/30' };
      case 'viewing_scheduled':
        return { label: 'Viewing Scheduled', color: 'bg-gold/15 text-gold border-gold/30' };
      case 'closed_won':
        return { label: 'Closed Won', color: 'bg-emerald-500/15 text-emerald-400 border-emerald-500/30' };
      default:
        return { label: status, color: 'bg-white/10 text-white border-white/20' };
    }
  }
</script>

<div class="space-y-8">
  <!-- Executive Welcome Banner -->
  <div class="glass-card p-6 sm:p-8 rounded-3xl border border-white/10 relative overflow-hidden bg-gradient-to-r from-bg-dark via-bg-dark to-gold/10">
    <div class="relative z-10 flex flex-col md:flex-row md:items-center justify-between gap-6">
      <div class="space-y-2 max-w-2xl">
        <div class="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-gold/10 border border-gold/30 text-gold text-[11px] font-headline font-bold uppercase tracking-wider">
          <svg class="w-3.5 h-3.5 text-gold" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
          </svg>
          <span>Executive CRM Command Center</span>
        </div>
        <h1 class="text-3xl font-extrabold font-headline text-white tracking-tight">
          Welcome Back, <span class="text-gold">Senior Advisor</span>
        </h1>
        <p class="text-xs sm:text-sm text-muted leading-relaxed">
          Monitor your high-net-worth sales pipeline, track active VIP buyer inquiries, and review luxury unit availability in real time.
        </p>
      </div>

      <div class="flex flex-wrap items-center gap-3">
        <a
          href="/leads"
          class="inline-flex items-center gap-2 px-5 py-3 rounded-xl bg-gold text-black font-headline font-bold text-xs uppercase tracking-wider hover:bg-gold/90 transition-all shadow-lg shadow-gold/20"
        >
          <svg class="w-4 h-4 text-black" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"></path>
          </svg>
          <span>Lead Pipeline</span>
        </a>
        <a
          href="/inventory"
          class="inline-flex items-center gap-2 px-5 py-3 rounded-xl bg-white/5 border border-white/10 text-white font-headline font-bold text-xs uppercase tracking-wider hover:border-gold/50 hover:text-gold transition-all"
        >
          <svg class="w-4 h-4 text-gold" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5m0 0v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4"></path>
          </svg>
          <span>Unit Inventory</span>
        </a>
      </div>
    </div>
  </div>

  <!-- KPI Metrics Grid -->
  <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4">
    {#each metrics as m}
      <div class="glass-card p-6 rounded-2xl border border-white/10 hover:border-gold/40 transition-all group flex flex-col justify-between">
        <div class="flex items-center justify-between mb-4">
          <span class="text-[10px] uppercase font-headline tracking-widest text-muted">{m.title}</span>
          <div class="p-2.5 rounded-xl bg-gold/15 border border-gold/30 text-gold flex items-center justify-center group-hover:scale-105 transition-transform">
            <svg class="w-5 h-5 text-gold" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.8" d={m.svgPath}></path>
            </svg>
          </div>
        </div>
        <div>
          <span class="text-2xl font-extrabold font-headline text-gold block mb-1 tracking-tight">{m.value}</span>
          <span class="text-[11px] font-semibold text-emerald-400">{m.change}</span>
        </div>
      </div>
    {/each}
  </div>

  <!-- Two Column Main Content -->
  <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
    
    <!-- Left Column (2 Cols): Recent VIP Lead Registrations -->
    <div class="lg:col-span-2 space-y-4">
      <div class="flex items-center justify-between">
        <div>
          <h2 class="text-lg font-bold font-headline text-white">Recent VIP Lead Registrations</h2>
          <p class="text-xs text-muted">Latest investor & buyer inquiries submitted through the landing page.</p>
        </div>
        <a
          href="/leads"
          class="text-xs font-bold font-headline text-gold hover:underline flex items-center gap-1"
        >
          <span>View All Pipeline</span>
          <svg class="w-4 h-4 text-gold" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 5l7 7m0 0l-7 7m7-7H3"></path>
          </svg>
        </a>
      </div>

      <div class="glass-card rounded-2xl border border-white/10 overflow-hidden">
        <div class="overflow-x-auto">
          <table class="w-full text-left border-collapse">
            <thead>
              <tr class="border-b border-white/10 bg-white/5 text-[10px] uppercase tracking-wider font-headline text-muted">
                <th class="p-4">Investor Name</th>
                <th class="p-4">Budget Range</th>
                <th class="p-4">Purpose</th>
                <th class="p-4">Pipeline Status</th>
                <th class="p-4 text-right">Quick Contact</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-white/5 text-xs">
              {#each $leadsStore.slice(0, 5) as lead}
                {@const statusBadge = getStatusBadge(lead.status)}
                <tr class="hover:bg-white/[0.02] transition-colors">
                  <td class="p-4">
                    <div class="font-bold text-white font-headline">{lead.full_name}</div>
                    <div class="text-[10px] text-muted">{lead.email}</div>
                  </td>
                  <td class="p-4">
                    <span class="px-2.5 py-1 rounded bg-gold/10 text-gold font-bold border border-gold/20 text-[11px]">
                      {lead.budget_range}
                    </span>
                  </td>
                  <td class="p-4 text-white/80 font-medium">
                    {lead.investment_purpose}
                  </td>
                  <td class="p-4">
                    <span class="px-2.5 py-1 rounded-full text-[10px] font-bold border {statusBadge.color}">
                      {statusBadge.label}
                    </span>
                  </td>
                  <td class="p-4 text-right">
                    <a
                      href="https://wa.me/{lead.whatsapp_number.replace(/[^0-9]/g, '')}"
                      target="_blank"
                      class="inline-flex items-center gap-1.5 px-3 py-1.5 rounded-lg bg-emerald-500/10 text-emerald-400 hover:bg-emerald-500 hover:text-black font-semibold text-[11px] transition-all"
                    >
                      <svg class="w-3.5 h-3.5 fill-current" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                        <path d="M.057 24l1.687-6.163c-1.041-1.804-1.588-3.849-1.587-5.946.003-6.556 5.338-11.891 11.893-11.891 3.181.001 6.167 1.24 8.413 3.488 2.245 2.248 3.481 5.236 3.48 8.414-.003 6.557-5.338 11.892-11.893 11.892-1.99-.001-3.951-.5-5.688-1.448l-6.305 1.654zm6.597-3.807c1.676.995 3.276 1.591 5.392 1.592 5.448 0 9.886-4.434 9.889-9.885.002-5.462-4.415-9.89-9.881-9.892-5.452 0-9.887 4.434-9.889 9.884-.001 2.225.651 3.891 1.746 5.634l-1.157 4.228 4.398-1.153z"/>
                      </svg>
                      <span>Chat</span>
                    </a>
                  </td>
                </tr>
              {/each}
            </tbody>
          </table>
        </div>
      </div>
    </div>

    <!-- Right Column (1 Col): Quick Shortcuts & System Status -->
    <div class="space-y-6">
      
      <!-- Quick Navigation Cards -->
      <div class="space-y-3">
        <h2 class="text-lg font-bold font-headline text-white">Platform Modules</h2>
        
        <a
          href="/leads"
          class="glass-card p-4 rounded-2xl border border-white/10 hover:border-gold/40 flex items-center justify-between group transition-all"
        >
          <div class="flex items-center gap-3">
            <div class="p-3 rounded-xl bg-gold/10 text-gold border border-gold/20 group-hover:scale-105 transition-transform">
              <svg class="w-5 h-5 text-gold" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"></path>
              </svg>
            </div>
            <div>
              <h3 class="text-sm font-bold font-headline text-white group-hover:text-gold transition-colors">Lead CRM Pipeline</h3>
              <p class="text-[11px] text-muted">Kanban sales board & status updates</p>
            </div>
          </div>
          <svg class="w-4 h-4 text-muted group-hover:text-gold group-hover:translate-x-1 transition-all" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
          </svg>
        </a>

        <a
          href="/inventory"
          class="glass-card p-4 rounded-2xl border border-white/10 hover:border-gold/40 flex items-center justify-between group transition-all"
        >
          <div class="flex items-center gap-3">
            <div class="p-3 rounded-xl bg-gold/10 text-gold border border-gold/20 group-hover:scale-105 transition-transform">
              <svg class="w-5 h-5 text-gold" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5m0 0v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4"></path>
              </svg>
            </div>
            <div>
              <h3 class="text-sm font-bold font-headline text-white group-hover:text-gold transition-colors">Unit Inventory</h3>
              <p class="text-[11px] text-muted">Luxury properties & pricing matrix</p>
            </div>
          </div>
          <svg class="w-4 h-4 text-muted group-hover:text-gold group-hover:translate-x-1 transition-all" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
          </svg>
        </a>

        <a
          href="/analytics"
          class="glass-card p-4 rounded-2xl border border-white/10 hover:border-gold/40 flex items-center justify-between group transition-all"
        >
          <div class="flex items-center gap-3">
            <div class="p-3 rounded-xl bg-gold/10 text-gold border border-gold/20 group-hover:scale-105 transition-transform">
              <svg class="w-5 h-5 text-gold" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 012-2h2a2 2 0 012 2v6a2 2 0 01-2 2h-2a2 2 0 01-2-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
              </svg>
            </div>
            <div>
              <h3 class="text-sm font-bold font-headline text-white group-hover:text-gold transition-colors">Analytics & ROI</h3>
              <p class="text-[11px] text-muted">Conversion rates & inquiry trends</p>
            </div>
          </div>
          <svg class="w-4 h-4 text-muted group-hover:text-gold group-hover:translate-x-1 transition-all" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
          </svg>
        </a>
      </div>

      <!-- Engine & Backend Status -->
      <div class="glass-card p-5 rounded-2xl border border-white/10 space-y-4">
        <div class="flex items-center justify-between">
          <span class="text-xs font-bold font-headline text-white">Engine Status</span>
          <span class="inline-flex items-center gap-1.5 px-2.5 py-0.5 rounded-full bg-emerald-500/15 text-emerald-400 border border-emerald-500/30 text-[10px] font-bold">
            <span class="w-1.5 h-1.5 rounded-full bg-emerald-400 animate-pulse"></span>
            <span>Online & Active</span>
          </span>
        </div>

        <div class="space-y-2 text-xs">
          <div class="flex justify-between items-center py-1 border-b border-white/5">
            <span class="text-muted">Go REST Backend</span>
            <span class="text-white font-mono text-[11px]">:8085 Connected</span>
          </div>
          <div class="flex justify-between items-center py-1 border-b border-white/5">
            <span class="text-muted">Database Storage</span>
            <span class="text-white font-mono text-[11px]">PostgreSQL Active</span>
          </div>
          <div class="flex justify-between items-center py-1">
            <span class="text-muted">Frontend Platform</span>
            <span class="text-white font-mono text-[11px]">SvelteKit SSR</span>
          </div>
        </div>
      </div>

    </div>

  </div>

  <!-- Flagship Projects Overview -->
  <div class="space-y-4 pt-4">
    <div>
      <h2 class="text-lg font-bold font-headline text-white">Flagship Development Portfolio</h2>
      <p class="text-xs text-muted">Overview of active Binghatti luxury architectural residences currently listed in CRM.</p>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
      {#each flagshipProjects as project}
        <div class="glass-card p-6 rounded-2xl border border-white/10 hover:border-gold/40 transition-all flex flex-col justify-between space-y-4">
          <div class="space-y-2">
            <div class="flex items-center justify-between">
              <span class="text-[10px] uppercase font-headline tracking-wider text-muted">{project.location}</span>
              <span class="px-2.5 py-0.5 rounded-full text-[10px] font-bold border {project.badgeColor}">
                {project.unitsAvailable}
              </span>
            </div>
            <h3 class="text-base font-extrabold font-headline text-white">{project.name}</h3>
            <p class="text-xs text-muted">{project.type}</p>
          </div>

          <div class="pt-3 border-t border-white/10 flex items-center justify-between">
            <span class="text-[10px] uppercase tracking-wider text-muted font-headline">Starting Price</span>
            <span class="text-sm font-bold font-headline text-gold">{project.startingPrice}</span>
          </div>
        </div>
      {/each}
    </div>
  </div>

</div>
