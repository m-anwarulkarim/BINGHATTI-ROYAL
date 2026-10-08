<script>
  import { onMount } from "svelte";
  import { leadsStore, fetchLeadsFromAPI } from "$lib/leadsStore.js";

  let selectedTimeframe = "30";

  onMount(() => {
    fetchLeadsFromAPI();
  });

  // Filter leads based on selected timeframe
  $: filteredLeads = $leadsStore.filter(lead => {
    if (selectedTimeframe === "all") return true;
    const days = parseInt(selectedTimeframe, 10);
    const date = new Date(lead.created_at || Date.now());
    const cutoff = new Date();
    cutoff.setDate(cutoff.getDate() - days);
    return date >= cutoff;
  });

  // 1. Total Leads
  $: totalLeads = filteredLeads.length;

  // 2. Conversion Rate (leads in qualified/contacted/won/negotiation)
  $: convertedCount = filteredLeads.filter(l => 
    ["contacted", "qualified", "negotiation", "won", "completed"].includes((l.status || "").toLowerCase())
  ).length;

  $: conversionRate = totalLeads > 0 
    ? ((convertedCount / totalLeads) * 100).toFixed(1) + "%" 
    : "0%";

  // 3. Revenue Estimation
  $: estimatedRevenue = filteredLeads.reduce((acc, lead) => {
    const budget = parseFloat(lead.budget || lead.budget_range || 0) || 1500000;
    return acc + (lead.status === "won" || lead.status === "completed" ? budget : budget * 0.1);
  }, 0);

  $: formattedRevenue = "$" + (estimatedRevenue / 1000000).toFixed(2) + "M";

  // 4. Average Budget
  $: avgBudgetNum = totalLeads > 0 
    ? filteredLeads.reduce((acc, lead) => acc + (parseFloat(lead.budget || 0) || 2000000), 0) / totalLeads 
    : 0;
  $: avgBudget = "$" + Math.round(avgBudgetNum / 1000).toLocaleString() + "K";

  // Funnel Data Breakdown
  $: funnelData = [
    { label: "Total Inquiries", count: totalLeads, percent: 100, color: "from-amber-500 to-amber-600" },
    { label: "Qualified Leads", count: filteredLeads.filter(l => ["qualified", "contacted", "won"].includes((l.status || "").toLowerCase())).length, percent: totalLeads > 0 ? Math.round((filteredLeads.filter(l => ["qualified", "contacted", "won"].includes((l.status || "").toLowerCase())).length / totalLeads) * 100) : 0, color: "from-amber-400 to-yellow-500" },
    { label: "Contacted & Meeting", count: filteredLeads.filter(l => ["contacted", "won"].includes((l.status || "").toLowerCase())).length, percent: totalLeads > 0 ? Math.round((filteredLeads.filter(l => ["contacted", "won"].includes((l.status || "").toLowerCase())).length / totalLeads) * 100) : 0, color: "from-yellow-400 to-amber-300" },
    { label: "Deals Closed / Won", count: filteredLeads.filter(l => ["won", "completed"].includes((l.status || "").toLowerCase())).length, percent: totalLeads > 0 ? Math.round((filteredLeads.filter(l => ["won", "completed"].includes((l.status || "").toLowerCase())).length / totalLeads) * 100) : 0, color: "from-amber-300 to-emerald-400" }
  ];

  // Budget Breakdown
  $: budgetBreakdown = [
    { range: "$1M - $2M", count: filteredLeads.filter(l => (parseFloat(l.budget) || 1500000) < 2000000).length },
    { range: "$2M - $5M", count: filteredLeads.filter(l => (parseFloat(l.budget) || 1500000) >= 2000000 && (parseFloat(l.budget) || 1500000) < 5000000).length },
    { range: "$5M - $10M", count: filteredLeads.filter(l => (parseFloat(l.budget) || 1500000) >= 5000000 && (parseFloat(l.budget) || 1500000) < 10000000).length },
    { range: "$10M+", count: filteredLeads.filter(l => (parseFloat(l.budget) || 1500000) >= 10000000).length }
  ];

  // Dynamic Chart Points Generation
  $: chartData = Array.from({ length: 30 }, (_, i) => {
    const d = new Date();
    d.setDate(d.getDate() - (29 - i));
    const dateStr = d.toISOString().split("T")[0];
    const count = filteredLeads.filter(l => {
      const lDate = (l.created_at || "").split("T")[0];
      return lDate === dateStr;
    }).length;
    return { date: d.toLocaleDateString("en-US", { month: "short", day: "numeric" }), count };
  });

  $: maxChartVal = Math.max(...chartData.map(d => d.count), 5);
  $: svgPoints = chartData.map((d, i) => {
    const x = (i / 29) * 800;
    const y = 200 - (d.count / maxChartVal) * 160;
    return `${x},${y}`;
  }).join(" ");

  // Export Data to CSV
  function exportCSV() {
    if (filteredLeads.length === 0) {
      alert("No data to export");
      return;
    }
    const headers = ["ID", "Name", "Email", "Phone", "Project", "Budget", "Status", "Created At"];
    const rows = filteredLeads.map(l => [
      l.id || "",
      `"${l.name || ""}"`,
      `"${l.email || ""}"`,
      `"${l.phone || ""}"`,
      `"${l.project_name || l.project || ""}"`,
      `"${l.budget || ""}"`,
      `"${l.status || "New"}"`,
      `"${l.created_at || ""}"`
    ]);
    const csvContent = "data:text/csv;charset=utf-8," + [headers.join(","), ...rows.map(r => r.join(","))].join("\n");
    const encodedUri = encodeURI(csvContent);
    const link = document.createElement("a");
    link.setAttribute("href", encodedUri);
    link.setAttribute("download", `analytics_export_${selectedTimeframe}d.csv`);
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
  }
</script>

<svelte:head>
  <title>Analytics & Insights | Luxury Dashboard</title>
</svelte:head>

<div class="space-y-8 p-6 text-gray-100 min-h-screen">
  <!-- Top Navigation / Header -->
  <div class="flex flex-col md:flex-row justify-between md:items-center gap-4 border-b border-amber-500/20 pb-6">
    <div>
      <h1 class="text-3xl font-bold bg-gradient-to-r from-amber-200 via-yellow-400 to-amber-500 bg-clip-text text-transparent">
        Analytics & Performance
      </h1>
      <p class="text-gray-400 text-sm mt-1">Real-time performance metrics and lead conversion insights</p>
    </div>

    <div class="flex items-center gap-3">
      <!-- Timeframe Filter -->
      <select
        bind:value={selectedTimeframe}
        class="bg-gray-900 border border-amber-500/30 text-amber-300 rounded-lg px-4 py-2 text-sm focus:outline-none focus:border-amber-400 transition"
      >
        <option value="7">Last 7 Days</option>
        <option value="30">Last 30 Days</option>
        <option value="90">Last 90 Days</option>
        <option value="all">All Time</option>
      </select>

      <!-- Export CSV Button -->
      <button
        on:click={exportCSV}
        class="flex items-center gap-2 bg-amber-500/10 hover:bg-amber-500/20 border border-amber-500/40 text-amber-300 px-4 py-2 rounded-lg text-sm transition"
      >
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"/>
        </svg>
        Export CSV
      </button>
    </div>
  </div>

  <!-- Key Metrics Cards -->
  <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
    <!-- Card 1 -->
    <div class="bg-gradient-to-br from-gray-900 via-gray-900 to-gray-950 border border-amber-500/20 p-6 rounded-xl relative overflow-hidden group">
      <div class="absolute -right-4 -bottom-4 w-24 h-24 bg-amber-500/5 rounded-full blur-xl group-hover:bg-amber-500/10 transition"></div>
      <span class="text-xs uppercase font-semibold text-gray-400 tracking-wider">Total Leads</span>
      <div class="text-3xl font-extrabold text-amber-300 mt-2">{totalLeads}</div>
      <div class="text-xs text-amber-500/80 mt-2 flex items-center gap-1">
        <span>Live Synced</span>
      </div>
    </div>

    <!-- Card 2 -->
    <div class="bg-gradient-to-br from-gray-900 via-gray-900 to-gray-950 border border-amber-500/20 p-6 rounded-xl relative overflow-hidden group">
      <div class="absolute -right-4 -bottom-4 w-24 h-24 bg-yellow-500/5 rounded-full blur-xl group-hover:bg-yellow-500/10 transition"></div>
      <span class="text-xs uppercase font-semibold text-gray-400 tracking-wider">Conversion Rate</span>
      <div class="text-3xl font-extrabold text-amber-300 mt-2">{conversionRate}</div>
      <div class="text-xs text-emerald-400 mt-2 flex items-center gap-1">
        <span>{convertedCount} converted / qualified</span>
      </div>
    </div>

    <!-- Card 3 -->
    <div class="bg-gradient-to-br from-gray-900 via-gray-900 to-gray-950 border border-amber-500/20 p-6 rounded-xl relative overflow-hidden group">
      <div class="absolute -right-4 -bottom-4 w-24 h-24 bg-emerald-500/5 rounded-full blur-xl group-hover:bg-emerald-500/10 transition"></div>
      <span class="text-xs uppercase font-semibold text-gray-400 tracking-wider">Estimated Pipeline</span>
      <div class="text-3xl font-extrabold text-amber-300 mt-2">{formattedRevenue}</div>
      <div class="text-xs text-gray-400 mt-2">Weighted property value</div>
    </div>

    <!-- Card 4 -->
    <div class="bg-gradient-to-br from-gray-900 via-gray-900 to-gray-950 border border-amber-500/20 p-6 rounded-xl relative overflow-hidden group">
      <div class="absolute -right-4 -bottom-4 w-24 h-24 bg-blue-500/5 rounded-full blur-xl group-hover:bg-blue-500/10 transition"></div>
      <span class="text-xs uppercase font-semibold text-gray-400 tracking-wider">Avg Lead Value</span>
      <div class="text-3xl font-extrabold text-amber-300 mt-2">{avgBudget}</div>
      <div class="text-xs text-gray-400 mt-2">Average buyer budget</div>
    </div>
  </div>

  <!-- Charts Section -->
  <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
    <!-- Lead Volume Timeline (SVG Area Chart) -->
    <div class="lg:col-span-2 bg-gray-900/60 border border-amber-500/20 rounded-xl p-6 backdrop-blur-sm">
      <div class="flex justify-between items-center mb-6">
        <div>
          <h2 class="text-lg font-bold text-amber-200">Lead Volume Trend</h2>
          <p class="text-xs text-gray-400">Daily lead registrations over past 30 days</p>
        </div>
        <span class="text-xs text-amber-400 bg-amber-500/10 px-3 py-1 rounded-full border border-amber-500/30">
          Real-time
        </span>
      </div>

      <!-- SVG Area Chart -->
      <div class="relative w-full h-64 overflow-x-auto">
        <svg viewBox="0 0 800 200" class="w-full h-full overflow-visible">
          <defs>
            <linearGradient id="chartGradient" x1="0" y1="0" x2="0" y2="1">
              <stop offset="0%" stop-color="#f59e0b" stop-opacity="0.4"/>
              <stop offset="100%" stop-color="#f59e0b" stop-opacity="0.0"/>
            </linearGradient>
          </defs>

          <!-- Grid Lines -->
          <line x1="0" y1="40" x2="800" y2="40" stroke="#374151" stroke-dasharray="4" stroke-width="0.5"/>
          <line x1="0" y1="90" x2="800" y2="90" stroke="#374151" stroke-dasharray="4" stroke-width="0.5"/>
          <line x1="0" y1="140" x2="800" y2="140" stroke="#374151" stroke-dasharray="4" stroke-width="0.5"/>
          <line x1="0" y1="190" x2="800" y2="190" stroke="#374151" stroke-width="1"/>

          <!-- Area -->
          {#if svgPoints}
            <polygon points={`0,200 ${svgPoints} 800,200`} fill="url(#chartGradient)"/>
            <!-- Path Line -->
            <polyline points={svgPoints} fill="none" stroke="#fbbf24" stroke-width="3" stroke-linecap="round"/>
          {/if}

          <!-- Points -->
          {#each chartData as d, i}
            {@const x = (i / 29) * 800}
            {@const y = 200 - (d.count / maxChartVal) * 160}
            <circle cx={x} cy={y} r="4" class="fill-amber-400 hover:r-6 transition-all cursor-pointer">
              <title>{d.date}: {d.count} leads</title>
            </circle>
          {/each}
        </svg>
      </div>
      <div class="flex justify-between text-xs text-gray-500 mt-4 px-2">
        <span>30 Days Ago</span>
        <span>15 Days Ago</span>
        <span>Today</span>
      </div>
    </div>

    <!-- Sales Conversion Funnel -->
    <div class="bg-gray-900/60 border border-amber-500/20 rounded-xl p-6 backdrop-blur-sm">
      <h2 class="text-lg font-bold text-amber-200 mb-2">Lead Conversion Funnel</h2>
      <p class="text-xs text-gray-400 mb-6">Stage progression efficiency</p>

      <div class="space-y-5">
        {#each funnelData as stage}
          <div>
            <div class="flex justify-between text-xs font-semibold mb-1">
              <span class="text-gray-300">{stage.label}</span>
              <span class="text-amber-400">{stage.count} ({stage.percent}%)</span>
            </div>
            <div class="w-full bg-gray-800 rounded-full h-3 overflow-hidden border border-amber-500/10">
              <div
                class={`h-full bg-gradient-to-r ${stage.color} rounded-full transition-all duration-500`}
                style={`width: ${Math.max(stage.percent, 5)}%`}
              ></div>
            </div>
          </div>
        {/each}
      </div>
    </div>
  </div>

  <!-- Lower Section: Budget Distribution & Live Stream -->
  <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
    <!-- Budget Distribution -->
    <div class="bg-gray-900/60 border border-amber-500/20 rounded-xl p-6 backdrop-blur-sm">
      <h2 class="text-lg font-bold text-amber-200 mb-2">Buyer Budget Demographics</h2>
      <p class="text-xs text-gray-400 mb-6">Distribution across price brackets</p>

      <div class="space-y-4">
        {#each budgetBreakdown as item}
          {@const total = totalLeads || 1}
          {@const pct = Math.round((item.count / total) * 100)}
          <div class="flex items-center justify-between gap-4">
            <span class="text-xs text-gray-300 w-24 font-medium">{item.range}</span>
            <div class="flex-1 bg-gray-800 h-2.5 rounded-full overflow-hidden">
              <div class="bg-amber-400 h-full rounded-full" style={`width: ${pct}%`}></div>
            </div>
            <span class="text-xs text-amber-400 font-bold w-12 text-right">{item.count}</span>
          </div>
        {/each}
      </div>
    </div>

    <!-- Live Recent Leads Feed -->
    <div class="lg:col-span-2 bg-gray-900/60 border border-amber-500/20 rounded-xl p-6 backdrop-blur-sm">
      <div class="flex justify-between items-center mb-6">
        <div>
          <h2 class="text-lg font-bold text-amber-200">Recent Lead Submissions</h2>
          <p class="text-xs text-gray-400">Latest active contacts from landing page</p>
        </div>
        <a href="/leads" class="text-xs text-amber-400 hover:text-amber-300 underline">View All Leads →</a>
      </div>

      <div class="overflow-x-auto">
        <table class="w-full text-left text-xs">
          <thead>
            <tr class="text-amber-500/80 border-b border-gray-800">
              <th class="pb-3 font-semibold">Name</th>
              <th class="pb-3 font-semibold">Contact</th>
              <th class="pb-3 font-semibold">Project</th>
              <th class="pb-3 font-semibold">Status</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-800/60 text-gray-300">
            {#each filteredLeads.slice(0, 5) as lead}
              <tr class="hover:bg-amber-500/5 transition">
                <td class="py-3 font-medium text-white">{lead.name || "Anonymous"}</td>
                <td class="py-3 text-gray-400">{lead.phone || lead.email || "N/A"}</td>
                <td class="py-3 text-amber-300/90">{lead.project_name || lead.project || "Binghatti Royal"}</td>
                <td class="py-3">
                  <span class="px-2 py-0.5 rounded text-[10px] uppercase font-bold bg-amber-500/20 text-amber-300 border border-amber-500/30">
                    {lead.status || "New"}
                  </span>
                </td>
              </tr>
            {/each}
            {#if filteredLeads.length === 0}
              <tr>
                <td colspan="4" class="text-center py-6 text-gray-500">No leads found for this period.</td>
              </tr>
            {/if}
          </tbody>
        </table>
      </div>
    </div>
  </div>
</div>
