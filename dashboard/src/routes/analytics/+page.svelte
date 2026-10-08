<script>
  const metrics = [
    { title: "Total Leads This Month", value: "148", change: "+24.5%", trend: "up", svgPath: "M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" },
    { title: "Lead Conversion Rate", value: "18.5%", change: "+3.2%", trend: "up", svgPath: "M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" },
    { title: "Top Lead Budget Range", value: "$500k – $1M", change: "42% of total", trend: "neutral", svgPath: "M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V6m0 8v2m0-6c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z" },
    { title: "Most Popular Project", value: "Mercedes-Benz Places", change: "38 inquiries", trend: "up", svgPath: "M5 3v4M3 5h4M6 17v4m-2-2h4m5-16l2.286 6.857L21 12l-5.714 2.143L13 21l-2.286-6.857L5 12l5.714-2.143L13 3z" }
  ];

  // Daily lead data points (30 days)
  const dailyData = [
    3, 5, 2, 8, 6, 9, 12, 7, 10, 15, 11, 14, 18, 13, 16, 22, 19, 21, 25, 20, 24, 28, 26, 30, 27, 32, 29, 35, 31, 38
  ];

  // SVG chart path calculations
  const maxVal = Math.max(...dailyData);
  const chartHeight = 160;
  const chartWidth = 700;

  const points = dailyData.map((val, idx) => {
    const x = (idx / (dailyData.length - 1)) * chartWidth;
    const y = chartHeight - (val / maxVal) * chartHeight;
    return `${x},${y}`;
  }).join(' ');
</script>

<div class="space-y-8">
  <!-- Page Header -->
  <div>
    <h1 class="text-2xl font-bold font-headline text-white">Conversion & Traffic Analytics</h1>
    <p class="text-xs text-muted">Real-time breakdown of VIP lead volume, marketing ROI, and property inquiry trends.</p>
  </div>

  <!-- KPI Metrics Grid -->
  <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4">
    {#each metrics as m}
      <div class="glass-card p-6 rounded-2xl border border-white/10 hover:border-gold/40 transition-colors flex flex-col justify-between">
        <div class="flex items-center justify-between mb-4">
          <span class="text-[10px] uppercase font-headline tracking-widest text-muted">{m.title}</span>
          <div class="p-2.5 rounded-xl bg-gold/15 border border-gold/30 text-gold flex items-center justify-center">
            <svg class="w-5 h-5 text-gold" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.8" d={m.svgPath}></path>
            </svg>
          </div>
        </div>
        <div>
          <span class="text-2xl font-bold font-headline text-gold block mb-1">{m.value}</span>
          <span class="text-[11px] font-semibold text-emerald-400">{m.change}</span>
        </div>
      </div>
    {/each}
  </div>

  <!-- Daily Lead Volume Line Chart -->
  <div class="glass-card p-6 rounded-2xl border border-white/10 space-y-6">
    <div class="flex items-center justify-between">
      <div>
        <h3 class="text-lg font-bold font-headline text-white">Daily Lead Volume (Last 30 Days)</h3>
        <p class="text-xs text-muted">Daily track of organic and VIP registration inquiries.</p>
      </div>
      <div class="px-3 py-1 rounded-full bg-gold/10 text-gold text-xs font-headline font-bold border border-gold/30">
        30-Day Total: 593 Leads
      </div>
    </div>

    <!-- SVG Area Chart -->
    <div class="w-full overflow-hidden pt-4">
      <svg viewBox="0 0 700 180" class="w-full h-48 overflow-visible">
        <!-- Grid lines -->
        <line x1="0" y1="0" x2="700" y2="0" stroke="rgba(255,255,255,0.05)" stroke-dasharray="4" />
        <line x1="0" y1="60" x2="700" y2="60" stroke="rgba(255,255,255,0.05)" stroke-dasharray="4" />
        <line x1="0" y1="120" x2="700" y2="120" stroke="rgba(255,255,255,0.05)" stroke-dasharray="4" />
        <line x1="0" y1="180" x2="700" y2="180" stroke="rgba(255,255,255,0.1)" />

        <!-- Gradient Area Fill -->
        <defs>
          <linearGradient id="chartGradient" x1="0" y1="0" x2="0" y2="1">
            <stop offset="0%" stop-color="#D4AF37" stop-opacity="0.4" />
            <stop offset="100%" stop-color="#D4AF37" stop-opacity="0" />
          </linearGradient>
        </defs>

        <polygon points={`0,180 ${points} 700,180`} fill="url(#chartGradient)" />
        <polyline points={points} fill="none" stroke="#D4AF37" stroke-width="3" stroke-linecap="round" stroke-linejoin="round" />
      </svg>

      <!-- Chart X-Axis Labels -->
      <div class="flex justify-between text-[10px] text-muted font-headline border-t border-white/10 pt-2">
        <span>Day 1</span>
        <span>Day 10</span>
        <span>Day 20</span>
        <span>Day 30 (Today)</span>
      </div>
    </div>
  </div>
</div>
