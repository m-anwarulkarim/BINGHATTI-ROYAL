<script>
  import { leadsStore, updateLeadStatus, createManualLead } from '$lib/leadsStore.js';

  const columns = [
    { key: 'new', title: 'New Leads', color: 'bg-blue-500 text-black border-blue-400 font-extrabold' },
    { key: 'contacted', title: 'Contacted', color: 'bg-yellow-400 text-black border-yellow-300 font-extrabold' },
    { key: 'qualified', title: 'Qualified', color: 'bg-purple-500 text-white border-purple-400 font-extrabold' },
    { key: 'viewing_scheduled', title: 'Viewing Scheduled', color: 'bg-amber-400 text-black border-amber-300 font-extrabold' },
    { key: 'closed_won', title: 'Closed Won', color: 'bg-emerald-500 text-black border-emerald-400 font-extrabold' },
    { key: 'closed_lost', title: 'Closed Lost', color: 'bg-red-500 text-white border-red-400 font-extrabold' }
  ];

  let searchQuery = '';
  let selectedBudget = 'ALL';
  let selectedPurpose = 'ALL';

  // Add Lead Modal State
  let showAddModal = false;
  let newFullName = '';
  let newWhatsApp = '';
  let newEmail = '';
  let newBudget = '$500k-$1M';
  let newPurpose = 'Investment';

  // Selected Lead Detail Modal State
  let selectedLeadModal = null;

  // Drag and Drop State
  let draggedLeadId = null;
  let dragOverColumnKey = null;

  $: filteredLeads = $leadsStore.filter(lead => {
    const matchesSearch = (lead.full_name || '').toLowerCase().includes(searchQuery.toLowerCase()) ||
                          (lead.email || '').toLowerCase().includes(searchQuery.toLowerCase()) ||
                          (lead.whatsapp_number || '').includes(searchQuery);
    const matchesBudget = selectedBudget === 'ALL' || lead.budget_range === selectedBudget;
    const matchesPurpose = selectedPurpose === 'ALL' || lead.investment_purpose === selectedPurpose;

    return matchesSearch && matchesBudget && matchesPurpose;
  });

  function exportToCSV() {
    const headers = ['ID', 'Full Name', 'WhatsApp', 'Email', 'Budget Range', 'Purpose', 'Status', 'Created At'];
    const rows = filteredLeads.map(l => [
      l.id,
      `"${l.full_name}"`,
      `"${l.whatsapp_number}"`,
      `"${l.email}"`,
      `"${l.budget_range}"`,
      `"${l.investment_purpose}"`,
      l.status,
      l.created_at
    ]);

    const csvContent = "data:text/csv;charset=utf-8," + [headers.join(','), ...rows.map(r => r.join(','))].join('\n');
    const encodedUri = encodeURI(csvContent);
    const link = document.createElement('a');
    link.setAttribute('href', encodedUri);
    link.setAttribute('download', `Binghatti_Leads_${new Date().toISOString().slice(0,10)}.csv`);
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
  }

  function handleStatusChange(leadId, newStatus) {
    if (!leadId || !newStatus) return;
    updateLeadStatus(leadId, newStatus);
    if (selectedLeadModal && selectedLeadModal.id === leadId) {
      selectedLeadModal.status = newStatus;
    }
  }

  async function handleAddLeadSubmit() {
    if (!newFullName || !newWhatsApp || !newEmail) return;

    await createManualLead({
      full_name: newFullName,
      whatsapp_number: newWhatsApp,
      email: newEmail,
      budget_range: newBudget,
      investment_purpose: newPurpose,
      lead_source: 'Manual Advisor Input'
    });

    // Reset Form & Close Modal
    newFullName = '';
    newWhatsApp = '';
    newEmail = '';
    newBudget = '$500k-$1M';
    newPurpose = 'Investment';
    showAddModal = false;
  }

  // Robust HTML5 Drag and Drop handlers
  function handleDragStart(e, leadId) {
    draggedLeadId = leadId;
    e.dataTransfer.effectAllowed = 'move';
    try {
      e.dataTransfer.setData('text/plain', leadId);
    } catch (err) {}
  }

  function handleDragOver(e, colKey) {
    e.preventDefault();
    if (e.dataTransfer) {
      e.dataTransfer.dropEffect = 'move';
    }
    dragOverColumnKey = colKey;
  }

  function handleDragLeave(e, colKey) {
    if (dragOverColumnKey === colKey) {
      dragOverColumnKey = null;
    }
  }

  function handleDrop(e, targetStatus) {
    e.preventDefault();
    dragOverColumnKey = null;

    let id = draggedLeadId;
    try {
      const dataId = e.dataTransfer.getData('text/plain');
      if (dataId) id = dataId;
    } catch (err) {}

    if (id) {
      handleStatusChange(id, targetStatus);
    }
    draggedLeadId = null;
  }
</script>

<div class="space-y-6">
  
  <!-- Page Header -->
  <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
    <div>
      <h1 class="text-2xl font-bold font-headline text-white">VIP Lead Kanban Pipeline</h1>
      <p class="text-xs text-muted">Manage, filter, drag & drop, and track high-net-worth real estate leads across sales stages.</p>
    </div>

    <div class="flex items-center gap-3">
      <button
        on:click={() => showAddModal = true}
        class="inline-flex items-center gap-2 px-4 py-2.5 rounded-xl bg-gold text-black font-headline font-bold text-xs uppercase tracking-wider hover:bg-gold/90 transition-all shadow-lg shadow-gold/20"
      >
        <svg class="w-4 h-4 text-black" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path>
        </svg>
        <span>+ Add VIP Lead</span>
      </button>

      <button
        on:click={exportToCSV}
        class="inline-flex items-center gap-2 px-4 py-2.5 rounded-xl bg-white/5 border border-white/10 text-gold hover:bg-gold/15 font-headline font-bold text-xs uppercase tracking-wider transition-all"
      >
        <svg class="w-4 h-4 text-gold" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4"></path>
        </svg>
        <span>Export CSV</span>
      </button>
    </div>
  </div>

  <!-- Filters Bar -->
  <div class="p-4 rounded-xl glass-card border border-white/10 grid grid-cols-1 sm:grid-cols-3 gap-4">
    <div>
      <label for="lead-search-input" class="block text-[10px] uppercase tracking-wider font-headline text-muted mb-1">Search Keywords</label>
      <input
        id="lead-search-input"
        type="text"
        bind:value={searchQuery}
        placeholder="Filter by name, email, whatsapp..."
        class="w-full select-luxury py-2 px-3"
      />
    </div>

    <div>
      <label for="lead-budget-select" class="block text-[10px] uppercase tracking-wider font-headline text-muted mb-1">Budget Range</label>
      <select
        id="lead-budget-select"
        bind:value={selectedBudget}
        class="w-full select-luxury py-2 px-3"
      >
        <option value="ALL">All Budgets</option>
        <option value="$300k-$500k">$300k - $500k</option>
        <option value="$500k-$1M">$500k - $1M</option>
        <option value="$1M+">$1M+</option>
      </select>
    </div>

    <div>
      <label for="lead-purpose-select" class="block text-[10px] uppercase tracking-wider font-headline text-muted mb-1">Investment Purpose</label>
      <select
        id="lead-purpose-select"
        bind:value={selectedPurpose}
        class="w-full select-luxury py-2 px-3"
      >
        <option value="ALL">All Purposes</option>
        <option value="Investment">Investment</option>
        <option value="Self-use">Self-use</option>
      </select>
    </div>
  </div>

  <!-- Kanban Board Grid -->
  <div class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-6 gap-4 overflow-x-auto pb-4">
    {#each columns as col}
      {@const colLeads = filteredLeads.filter(l => l.status === col.key)}
      <!-- svelte-ignore a11y_no_static_element_interactions -->
      <div
        on:dragover|preventDefault={(e) => handleDragOver(e, col.key)}
        on:dragenter|preventDefault={(e) => handleDragOver(e, col.key)}
        on:dragleave={(e) => handleDragLeave(e, col.key)}
        on:drop={(e) => handleDrop(e, col.key)}
        class="flex flex-col rounded-2xl bg-bg-dark border transition-all duration-200 p-3 min-w-[250px] {dragOverColumnKey === col.key ? 'border-gold bg-gold/10 shadow-[0_0_20px_rgba(212,175,55,0.2)] scale-[1.01]' : 'border-white/10'}"
      >
        
        <!-- Column Header -->
        <div class="flex items-center gap-2 pb-3 border-b border-white/10 mb-3 px-1">
          <span class="text-xs font-bold font-headline text-white shrink-0">{col.title}</span>
          <!-- High-Contrast Bright Badge Counter Directly Beside Title -->
          <span 
            class="min-w-[24px] h-6 px-2 flex items-center justify-center rounded-full text-xs font-black font-mono shadow-lg transition-transform hover:scale-110"
            style={
              col.key === 'new' ? 'background-color: #3b82f6; color: #ffffff; border: 1.5px solid #93c5fd; box-shadow: 0 0 10px rgba(59,130,246,0.6);' :
              col.key === 'contacted' ? 'background-color: #eab308; color: #000000; border: 1.5px solid #fef08a; box-shadow: 0 0 10px rgba(234,179,8,0.6);' :
              col.key === 'qualified' ? 'background-color: #a855f7; color: #ffffff; border: 1.5px solid #e9d5ff; box-shadow: 0 0 10px rgba(168,85,247,0.6);' :
              col.key === 'viewing_scheduled' ? 'background-color: #f97316; color: #ffffff; border: 1.5px solid #ffedd5; box-shadow: 0 0 10px rgba(249,115,22,0.6);' :
              col.key === 'closed_won' ? 'background-color: #10b981; color: #000000; border: 1.5px solid #a7f3d0; box-shadow: 0 0 10px rgba(16,185,129,0.6);' :
              'background-color: #ef4444; color: #ffffff; border: 1.5px solid #fca5a5; box-shadow: 0 0 10px rgba(239,68,68,0.7);'
            }
          >
            {colLeads.length}
          </span>
        </div>

        <!-- Cards Container -->
        <div class="space-y-3 flex-1 min-h-[220px]">
          {#each colLeads as lead (lead.id)}
            <!-- svelte-ignore a11y_no_static_element_interactions -->
            <div
              draggable="true"
              on:dragstart={(e) => handleDragStart(e, lead.id)}
              class="glass-card p-4 rounded-xl border border-white/10 hover:border-gold/50 transition-all space-y-3 cursor-grab active:cursor-grabbing hover:scale-[1.01]"
            >
              
              <div class="flex items-start justify-between">
                <div>
                  <button
                    on:click={() => selectedLeadModal = lead}
                    class="text-sm font-bold font-headline text-white hover:text-gold text-left transition-colors"
                  >
                    {lead.full_name}
                  </button>
                  <span class="text-[10px] text-muted block">{lead.email}</span>
                </div>
              </div>

              <div class="flex items-center justify-between text-[11px]">
                <span class="px-2 py-0.5 rounded bg-gold/10 text-gold font-semibold border border-gold/20">
                  {lead.budget_range}
                </span>
                <span class="text-white/70">{lead.investment_purpose}</span>
              </div>

              <div class="pt-2 border-t border-white/5 flex items-center justify-between">
                <a
                  href="https://wa.me/{lead.whatsapp_number.replace(/[^0-9]/g, '')}"
                  target="_blank"
                  class="text-[10px] text-emerald-400 font-semibold hover:underline flex items-center gap-1.5"
                >
                  <svg class="w-3 h-3 fill-current text-emerald-400" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                    <path d="M.057 24l1.687-6.163c-1.041-1.804-1.588-3.849-1.587-5.946.003-6.556 5.338-11.891 11.893-11.891 3.181.001 6.167 1.24 8.413 3.488 2.245 2.248 3.481 5.236 3.48 8.414-.003 6.557-5.338 11.892-11.893 11.892-1.99-.001-3.951-.5-5.688-1.448l-6.305 1.654zm6.597-3.807c1.676.995 3.276 1.591 5.392 1.592 5.448 0 9.886-4.434 9.889-9.885.002-5.462-4.415-9.89-9.881-9.892-5.452 0-9.887 4.434-9.889 9.884-.001 2.225.651 3.891 1.746 5.634l-1.157 4.228 4.398-1.153z"/>
                  </svg>
                  <span>WhatsApp</span>
                </a>

                <!-- Move Stage Dropdown -->
                <select
                  value={lead.status}
                  on:change={(e) => handleStatusChange(lead.id, e.target.value)}
                  class="select-card-stage"
                >
                  {#each columns as targetCol}
                    <option value={targetCol.key}>{targetCol.title}</option>
                  {/each}
                </select>
              </div>

            </div>
          {/each}

          {#if colLeads.length === 0}
            <div class="h-28 rounded-xl border border-dashed border-white/10 flex flex-col items-center justify-center text-center p-3 text-muted">
              <span class="text-xs font-semibold">Drop leads here</span>
              <span class="text-[10px] text-white/40 mt-1">Move to {col.title}</span>
            </div>
          {/if}
        </div>

      </div>
    {/each}
  </div>

</div>

<!-- Modal 1: Add Manual VIP Lead -->
{#if showAddModal}
  <div class="fixed inset-0 z-50 flex items-center justify-center bg-black/80 backdrop-blur-md p-4">
    <div class="relative w-full max-w-lg rounded-2xl glass-card border-2 border-gold/60 p-6 bg-bg-offblack">
      <div class="flex items-center justify-between mb-4 border-b border-white/10 pb-3">
        <h3 class="text-lg font-bold font-headline text-white">Add New VIP Lead</h3>
        <button on:click={() => showAddModal = false} class="text-muted hover:text-white p-1 rounded-lg hover:bg-white/10 transition-colors" aria-label="Close modal">
          <svg class="w-4 h-4 text-gold" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
        </button>
      </div>

      <form on:submit|preventDefault={handleAddLeadSubmit} class="space-y-4 text-xs">
        <div>
          <label for="new-lead-name" class="block font-headline text-muted uppercase mb-1">Full Name *</label>
          <input
            id="new-lead-name"
            type="text"
            bind:value={newFullName}
            required
            placeholder="e.g. Sheikh Mohammed Al-Qasimi"
            class="w-full select-luxury py-2 px-3"
          />
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <div>
            <label for="new-lead-whatsapp" class="block font-headline text-muted uppercase mb-1">WhatsApp Number *</label>
            <input
              id="new-lead-whatsapp"
              type="text"
              bind:value={newWhatsApp}
              required
              placeholder="+971 50 999 8877"
              class="w-full select-luxury py-2 px-3"
            />
          </div>

          <div>
            <label for="new-lead-email" class="block font-headline text-muted uppercase mb-1">Email Address *</label>
            <input
              id="new-lead-email"
              type="email"
              bind:value={newEmail}
              required
              placeholder="m.qasimi@investor.ae"
              class="w-full select-luxury py-2 px-3"
            />
          </div>
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <div>
            <label for="new-lead-budget" class="block font-headline text-muted uppercase mb-1">Budget Range</label>
            <select
              id="new-lead-budget"
              bind:value={newBudget}
              class="w-full select-luxury py-2 px-3"
            >
              <option value="$300k-$500k">$300k - $500k</option>
              <option value="$500k-$1M">$500k - $1M</option>
              <option value="$1M+">$1M+</option>
            </select>
          </div>

          <div>
            <label for="new-lead-purpose" class="block font-headline text-muted uppercase mb-1">Investment Purpose</label>
            <select
              id="new-lead-purpose"
              bind:value={newPurpose}
              class="w-full select-luxury py-2 px-3"
            >
              <option value="Investment">Investment</option>
              <option value="Self-use">Self-use</option>
            </select>
          </div>
        </div>

        <div class="flex items-center justify-end gap-3 pt-4 border-t border-white/10">
          <button
            type="button"
            on:click={() => showAddModal = false}
            class="px-4 py-2 rounded-xl bg-white/10 text-white text-xs font-semibold hover:bg-white/20"
          >
            Cancel
          </button>
          <button
            type="submit"
            class="px-5 py-2 rounded-xl bg-gold text-black font-headline font-bold text-xs uppercase tracking-wider hover:bg-gold/90 transition-all"
          >
            Create Lead
          </button>
        </div>
      </form>
    </div>
  </div>
{/if}

<!-- Modal 2: Lead Details Modal -->
{#if selectedLeadModal}
  <div class="fixed inset-0 z-50 flex items-center justify-center bg-black/80 backdrop-blur-md p-4">
    <div class="relative w-full max-w-md rounded-2xl glass-card border-2 border-gold/60 p-6 bg-bg-offblack space-y-5">
      <div class="flex items-center justify-between border-b border-white/10 pb-3">
        <div>
          <h3 class="text-lg font-bold font-headline text-white">{selectedLeadModal.full_name}</h3>
          <span class="text-xs text-gold font-mono">{selectedLeadModal.id}</span>
        </div>
        <button on:click={() => selectedLeadModal = null} class="text-muted hover:text-white p-1 rounded-lg hover:bg-white/10 transition-colors" aria-label="Close modal">
          <svg class="w-4 h-4 text-gold" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
        </button>
      </div>

      <div class="space-y-3 text-xs">
        <div class="flex justify-between py-1 border-b border-white/5">
          <span class="text-muted">Email Address</span>
          <span class="text-white font-semibold">{selectedLeadModal.email}</span>
        </div>
        <div class="flex justify-between py-1 border-b border-white/5">
          <span class="text-muted">WhatsApp Number</span>
          <span class="text-emerald-400 font-semibold">{selectedLeadModal.whatsapp_number}</span>
        </div>
        <div class="flex justify-between py-1 border-b border-white/5">
          <span class="text-muted">Budget Range</span>
          <span class="text-gold font-bold">{selectedLeadModal.budget_range}</span>
        </div>
        <div class="flex justify-between py-1 border-b border-white/5">
          <span class="text-muted">Investment Purpose</span>
          <span class="text-white">{selectedLeadModal.investment_purpose}</span>
        </div>
        <div class="flex justify-between py-1 border-b border-white/5">
          <span class="text-muted">Lead Source</span>
          <span class="text-white/80">{selectedLeadModal.lead_source || 'Website VIP Form'}</span>
        </div>

        <div class="pt-2">
          <label for="modal-status-select" class="block text-[10px] uppercase font-headline text-muted mb-1">Update Lead Pipeline Stage</label>
          <select
            id="modal-status-select"
            value={selectedLeadModal.status}
            on:change={(e) => handleStatusChange(selectedLeadModal.id, e.target.value)}
            class="w-full select-luxury py-2 px-3"
          >
            {#each columns as col}
              <option value={col.key}>{col.title}</option>
            {/each}
          </select>
        </div>
      </div>

      <div class="pt-4 border-t border-white/10 flex items-center justify-between">
        <a
          href="https://wa.me/{selectedLeadModal.whatsapp_number.replace(/[^0-9]/g, '')}"
          target="_blank"
          class="inline-flex items-center gap-2 px-4 py-2 rounded-xl bg-emerald-500/15 border border-emerald-500/40 text-emerald-400 hover:bg-emerald-500 hover:text-black font-semibold text-xs transition-all"
        >
          <svg class="w-4 h-4 fill-current" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
            <path d="M.057 24l1.687-6.163c-1.041-1.804-1.588-3.849-1.587-5.946.003-6.556 5.338-11.891 11.893-11.891 3.181.001 6.167 1.24 8.413 3.488 2.245 2.248 3.481 5.236 3.48 8.414-.003 6.557-5.338 11.892-11.893 11.892-1.99-.001-3.951-.5-5.688-1.448l-6.305 1.654zm6.597-3.807c1.676.995 3.276 1.591 5.392 1.592 5.448 0 9.886-4.434 9.889-9.885.002-5.462-4.415-9.89-9.881-9.892-5.452 0-9.887 4.434-9.889 9.884-.001 2.225.651 3.891 1.746 5.634l-1.157 4.228 4.398-1.153z"/>
          </svg>
          <span>Open WhatsApp Chat</span>
        </a>

        <button
          on:click={() => selectedLeadModal = null}
          class="px-4 py-2 rounded-xl bg-white/10 text-white text-xs font-semibold hover:bg-white/20"
        >
          Close
        </button>
      </div>
    </div>
  </div>
{/if}
