<script>
  const defaultUnits = [
    {
      id: "UNIT-101",
      project_name: "Binghatti Mercedes-Benz Places",
      unit_type: "1 Bedroom Suite",
      size_sqft: 850,
      price_aed: 1450000,
      handover_date: "Q4 2027",
      status: "Available",
      floor_plan_pdf: "https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf"
    },
    {
      id: "UNIT-102",
      project_name: "Bugatti Residences by Binghatti",
      unit_type: "2 Bedroom Royal",
      size_sqft: 1450,
      price_aed: 2850000,
      handover_date: "Q2 2027",
      status: "Reserved",
      floor_plan_pdf: "https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf"
    },
    {
      id: "UNIT-103",
      project_name: "Jacob & Co Residences",
      unit_type: "Luxury Sky Penthouse",
      size_sqft: 3800,
      price_aed: 6900000,
      handover_date: "Q1 2028",
      status: "Available",
      floor_plan_pdf: "https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf"
    },
    {
      id: "UNIT-104",
      project_name: "Binghatti Mercedes-Benz Places",
      unit_type: "Studio Luxury",
      size_sqft: 520,
      price_aed: 980000,
      handover_date: "Q4 2027",
      status: "Sold",
      floor_plan_pdf: "https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf"
    },
    {
      id: "UNIT-105",
      project_name: "Bugatti Residences by Binghatti",
      unit_type: "Luxury Sky Penthouse",
      size_sqft: 5400,
      price_aed: 19000000,
      handover_date: "Q3 2027",
      status: "Available",
      floor_plan_pdf: "https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf"
    }
  ];

  function loadStoredUnits() {
    if (typeof window === 'undefined') return defaultUnits;
    try {
      const saved = localStorage.getItem('binghatti_inventory');
      if (saved) {
        const parsed = JSON.parse(saved);
        if (Array.isArray(parsed) && parsed.length > 0) return parsed;
      }
    } catch (e) {}
    return defaultUnits;
  }

  let units = loadStoredUnits();

  function persistUnits(updated) {
    units = updated;
    try {
      if (typeof window !== 'undefined') {
        localStorage.setItem('binghatti_inventory', JSON.stringify(updated));
      }
    } catch (e) {}
  }

  let searchQuery = "";
  let selectedProject = "ALL";
  let selectedType = "ALL";
  let selectedStatus = "ALL";

  // Modal edit state
  let isEditing = false;
  let editingUnit = null;

  // Add new unit modal state
  let isAdding = false;
  let newUnitId = "";
  let newProjectName = "Binghatti Mercedes-Benz Places";
  let newUnitType = "1 Bedroom Suite";
  let newSizeSqft = 900;
  let newPriceAED = 1800000;
  let newHandoverDate = "Q4 2027";
  let newStatus = "Available";
  let newPdfUrl = "https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf";

  $: filteredUnits = units.filter(u => {
    const matchSearch = searchQuery === "" ||
                        u.id.toLowerCase().includes(searchQuery.toLowerCase()) ||
                        u.project_name.toLowerCase().includes(searchQuery.toLowerCase()) ||
                        u.unit_type.toLowerCase().includes(searchQuery.toLowerCase());
    const matchProject = selectedProject === "ALL" || u.project_name === selectedProject;
    const matchType = selectedType === "ALL" || u.unit_type === selectedType;
    const matchStatus = selectedStatus === "ALL" || u.status === selectedStatus;
    return matchSearch && matchProject && matchType && matchStatus;
  });

  $: availableCount = units.filter(u => u.status === 'Available').length;
  $: reservedCount = units.filter(u => u.status === 'Reserved').length;
  $: soldCount = units.filter(u => u.status === 'Sold').length;

  function openEditModal(unit) {
    editingUnit = { ...unit };
    isEditing = true;
  }

  function closeEditModal() {
    isEditing = false;
    editingUnit = null;
  }

  function saveUnit() {
    if (editingUnit) {
      const updated = units.map(u => u.id === editingUnit.id ? { ...editingUnit } : u);
      persistUnits(updated);
      closeEditModal();
    }
  }

  function handleAddUnitSubmit() {
    const nextId = newUnitId ? newUnitId : `UNIT-${100 + units.length + 1}`;

    const created = {
      id: nextId,
      project_name: newProjectName,
      unit_type: newUnitType,
      size_sqft: Number(newSizeSqft),
      price_aed: Number(newPriceAED),
      handover_date: newHandoverDate,
      status: newStatus,
      floor_plan_pdf: newPdfUrl
    };

    persistUnits([created, ...units]);
    
    // Reset modal
    newUnitId = "";
    isAdding = false;
  }

  function exportInventoryCSV() {
    const headers = ['Unit ID', 'Project Name', 'Unit Layout', 'Size (sqft)', 'Price (AED)', 'Handover', 'Status'];
    const rows = filteredUnits.map(u => [
      u.id,
      `"${u.project_name}"`,
      `"${u.unit_type}"`,
      u.size_sqft,
      u.price_aed,
      `"${u.handover_date}"`,
      u.status
    ]);

    const csvContent = "data:text/csv;charset=utf-8," + [headers.join(','), ...rows.map(r => r.join(','))].join('\n');
    const encodedUri = encodeURI(csvContent);
    const link = document.createElement('a');
    link.setAttribute('href', encodedUri);
    link.setAttribute('download', `Binghatti_Inventory_${new Date().toISOString().slice(0,10)}.csv`);
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
  }

  function getStatusBadge(status) {
    switch (status) {
      case 'Available':
        return 'bg-emerald-500/15 text-emerald-400 border-emerald-500/30';
      case 'Reserved':
        return 'bg-amber-500/15 text-amber-400 border-amber-500/30';
      case 'Sold':
        return 'bg-red-500/15 text-red-400 border-red-500/30';
      default:
        return 'bg-white/10 text-white border-white/20';
    }
  }
</script>

<div class="space-y-6">
  <!-- Page Header -->
  <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
    <div>
      <h1 class="text-2xl font-bold font-headline text-white">Unit Inventory & Pricing Matrix</h1>
      <p class="text-xs text-muted">Track property availability, update unit prices, manage floor plan PDFs, and add new luxury units.</p>
    </div>

    <div class="flex items-center gap-3">
      <button
        on:click={() => isAdding = true}
        class="inline-flex items-center gap-2 px-4 py-2.5 rounded-xl bg-gold text-black font-headline font-bold text-xs uppercase tracking-wider hover:bg-gold/90 transition-all shadow-lg shadow-gold/20"
      >
        <svg class="w-4 h-4 text-black" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path>
        </svg>
        <span>+ Add Luxury Unit</span>
      </button>

      <button
        on:click={exportInventoryCSV}
        class="inline-flex items-center gap-2 px-4 py-2.5 rounded-xl bg-white/5 border border-white/10 text-gold hover:bg-gold/15 font-headline font-bold text-xs uppercase tracking-wider transition-all"
      >
        <svg class="w-4 h-4 text-gold" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4"></path>
        </svg>
        <span>Export CSV</span>
      </button>
    </div>
  </div>

  <!-- Inventory Quick Stats Bar -->
  <div class="grid grid-cols-2 sm:grid-cols-4 gap-4">
    <div class="glass-card p-4 rounded-xl border border-white/10">
      <span class="text-[10px] uppercase tracking-wider font-headline text-muted block">Total Catalog Units</span>
      <span class="text-xl font-bold font-headline text-white mt-1 block">{units.length} Units</span>
    </div>
    <div class="glass-card p-4 rounded-xl border border-white/10">
      <span class="text-[10px] uppercase tracking-wider font-headline text-muted block">Available Units</span>
      <span class="text-xl font-bold font-headline text-emerald-400 mt-1 block">{availableCount} Units</span>
    </div>
    <div class="glass-card p-4 rounded-xl border border-white/10">
      <span class="text-[10px] uppercase tracking-wider font-headline text-muted block">Reserved Units</span>
      <span class="text-xl font-bold font-headline text-amber-400 mt-1 block">{reservedCount} Units</span>
    </div>
    <div class="glass-card p-4 rounded-xl border border-white/10">
      <span class="text-[10px] uppercase tracking-wider font-headline text-muted block">Sold Units</span>
      <span class="text-xl font-bold font-headline text-red-400 mt-1 block">{soldCount} Units</span>
    </div>
  </div>

  <!-- Search & Filters Bar -->
  <div class="p-4 rounded-xl glass-card border border-white/10 grid grid-cols-1 sm:grid-cols-4 gap-4">
    <div>
      <label for="inv-search-query" class="block text-[10px] uppercase tracking-wider font-headline text-muted mb-1">Search Keywords</label>
      <input
        id="inv-search-query"
        type="text"
        bind:value={searchQuery}
        placeholder="Filter by Unit ID, project, layout..."
        class="w-full select-luxury py-2 px-3"
      />
    </div>

    <div>
      <label for="inv-filter-project" class="block text-[10px] uppercase tracking-wider font-headline text-muted mb-1">Project</label>
      <select
        id="inv-filter-project"
        bind:value={selectedProject}
        class="w-full select-luxury py-2 px-3"
      >
        <option value="ALL">All Projects</option>
        <option value="Binghatti Mercedes-Benz Places">Binghatti Mercedes-Benz Places</option>
        <option value="Bugatti Residences by Binghatti">Bugatti Residences by Binghatti</option>
        <option value="Jacob & Co Residences">Jacob & Co Residences</option>
      </select>
    </div>

    <div>
      <label for="inv-filter-type" class="block text-[10px] uppercase tracking-wider font-headline text-muted mb-1">Unit Layout Type</label>
      <select
        id="inv-filter-type"
        bind:value={selectedType}
        class="w-full select-luxury py-2 px-3"
      >
        <option value="ALL">All Types</option>
        <option value="Studio Luxury">Studio Luxury</option>
        <option value="1 Bedroom Suite">1 Bedroom Suite</option>
        <option value="2 Bedroom Royal">2 Bedroom Royal</option>
        <option value="Luxury Sky Penthouse">Luxury Sky Penthouse</option>
      </select>
    </div>

    <div>
      <label for="inv-filter-status" class="block text-[10px] uppercase tracking-wider font-headline text-muted mb-1">Availability Status</label>
      <select
        id="inv-filter-status"
        bind:value={selectedStatus}
        class="w-full select-luxury py-2 px-3"
      >
        <option value="ALL">All Statuses</option>
        <option value="Available">Available</option>
        <option value="Reserved">Reserved</option>
        <option value="Sold">Sold</option>
      </select>
    </div>
  </div>

  <!-- Inventory Data Table -->
  <div class="rounded-2xl glass-card border border-white/10 overflow-hidden">
    <div class="overflow-x-auto">
      <table class="w-full text-left text-xs border-collapse">
        <thead class="bg-bg-dark border-b border-white/10 text-muted uppercase tracking-wider font-headline">
          <tr>
            <th class="p-4">Unit ID</th>
            <th class="p-4">Project Name</th>
            <th class="p-4">Unit Layout</th>
            <th class="p-4">Size (sq. ft.)</th>
            <th class="p-4">Price (AED)</th>
            <th class="p-4">Handover</th>
            <th class="p-4">Status</th>
            <th class="p-4 text-right">Actions</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-white/5">
          {#each filteredUnits as unit (unit.id)}
            <tr class="hover:bg-white/5 transition-colors">
              <td class="p-4 font-bold font-headline text-gold">{unit.id}</td>
              <td class="p-4 font-semibold text-white">{unit.project_name}</td>
              <td class="p-4 text-white/80">{unit.unit_type}</td>
              <td class="p-4 text-white/80">{unit.size_sqft.toLocaleString()} sq. ft.</td>
              <td class="p-4 font-bold font-headline text-gold">AED {unit.price_aed.toLocaleString()}</td>
              <td class="p-4 text-white/70">{unit.handover_date}</td>
              <td class="p-4">
                <span class="px-2.5 py-1 rounded-full text-[10px] font-bold border {getStatusBadge(unit.status)}">
                  {unit.status}
                </span>
              </td>
              <td class="p-4 text-right">
                <div class="flex items-center justify-end gap-2">
                  <a
                    href={unit.floor_plan_pdf}
                    target="_blank"
                    class="px-2.5 py-1.5 rounded-lg bg-white/5 border border-white/10 text-white/80 hover:text-gold hover:border-gold/30 text-[11px] font-semibold transition-all flex items-center gap-1"
                    title="View Floor Plan PDF"
                  >
                    <svg class="w-3.5 h-3.5 text-gold" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21h10a2 2 0 002-2V9.414a1 1 0 00-.293-.707l-5.414-5.414A1 1 0 0012.586 3H7a2 2 0 00-2 2v14a2 2 0 002 2z"></path>
                    </svg>
                    <span>PDF</span>
                  </a>

                  <button
                    on:click={() => openEditModal(unit)}
                    class="px-3 py-1.5 rounded-lg bg-gold/15 border border-gold/40 text-gold hover:bg-gold hover:text-black font-semibold text-[11px] transition-all"
                  >
                    Edit Unit
                  </button>
                </div>
              </td>
            </tr>
          {/each}

          {#if filteredUnits.length === 0}
            <tr>
              <td colspan="8" class="p-8 text-center text-muted text-xs">
                No property units found matching your search and filter criteria.
              </td>
            </tr>
          {/if}
        </tbody>
      </table>
    </div>
  </div>
</div>

<!-- Modal 1: Inline Edit Modal -->
{#if isEditing && editingUnit}
  <div class="fixed inset-0 z-50 flex items-center justify-center bg-black/80 backdrop-blur-md p-4">
    <div class="relative w-full max-w-lg rounded-2xl glass-card border-2 border-gold/60 p-6 bg-bg-offblack">
      <div class="flex items-center justify-between mb-4 border-b border-white/10 pb-3">
        <h3 class="text-lg font-bold font-headline text-white">Edit Unit Details: {editingUnit.id}</h3>
        <button on:click={closeEditModal} class="text-muted hover:text-white p-1 rounded-lg hover:bg-white/10 transition-colors" aria-label="Close modal">
          <svg class="w-4 h-4 text-gold" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
        </button>
      </div>

      <form on:submit|preventDefault={saveUnit} class="space-y-4 text-xs">
        <div>
          <label for="inv-edit-project" class="block font-headline text-muted uppercase mb-1">Project Name</label>
          <input
            id="inv-edit-project"
            type="text"
            bind:value={editingUnit.project_name}
            required
            class="w-full select-luxury py-2 px-3"
          />
        </div>

        <div class="grid grid-cols-2 gap-4">
          <div>
            <label for="inv-edit-price" class="block font-headline text-muted uppercase mb-1">Price (AED)</label>
            <input
              id="inv-edit-price"
              type="number"
              bind:value={editingUnit.price_aed}
              required
              class="w-full select-luxury py-2 px-3"
            />
          </div>

          <div>
            <label for="inv-edit-size" class="block font-headline text-muted uppercase mb-1">Size (sq. ft.)</label>
            <input
              id="inv-edit-size"
              type="number"
              bind:value={editingUnit.size_sqft}
              required
              class="w-full select-luxury py-2 px-3"
            />
          </div>
        </div>

        <div class="grid grid-cols-2 gap-4">
          <div>
            <label for="inv-edit-handover" class="block font-headline text-muted uppercase mb-1">Handover Date</label>
            <input
              id="inv-edit-handover"
              type="text"
              bind:value={editingUnit.handover_date}
              required
              class="w-full select-luxury py-2 px-3"
            />
          </div>

          <div>
            <label for="inv-edit-status" class="block font-headline text-muted uppercase mb-1">Availability Status</label>
            <select
              id="inv-edit-status"
              bind:value={editingUnit.status}
              class="w-full select-luxury py-2 px-3"
            >
              <option value="Available">Available</option>
              <option value="Reserved">Reserved</option>
              <option value="Sold">Sold</option>
            </select>
          </div>
        </div>

        <div>
          <label for="inv-edit-pdf" class="block font-headline text-muted uppercase mb-1">Floor Plan PDF URL</label>
          <input
            id="inv-edit-pdf"
            type="text"
            bind:value={editingUnit.floor_plan_pdf}
            class="w-full select-luxury py-2 px-3"
          />
        </div>

        <div class="flex items-center justify-end gap-3 pt-4 border-t border-white/10">
          <button
            type="button"
            on:click={closeEditModal}
            class="px-4 py-2 rounded-xl bg-white/10 text-white text-xs font-semibold hover:bg-white/20"
          >
            Cancel
          </button>
          <button
            type="submit"
            class="px-5 py-2 rounded-xl bg-gold text-black font-headline font-bold text-xs uppercase tracking-wider hover:bg-gold/90 transition-all"
          >
            Save Changes
          </button>
        </div>
      </form>
    </div>
  </div>
{/if}

<!-- Modal 2: Add New Luxury Unit Modal -->
{#if isAdding}
  <div class="fixed inset-0 z-50 flex items-center justify-center bg-black/80 backdrop-blur-md p-4">
    <div class="relative w-full max-w-lg rounded-2xl glass-card border-2 border-gold/60 p-6 bg-bg-offblack">
      <div class="flex items-center justify-between mb-4 border-b border-white/10 pb-3">
        <h3 class="text-lg font-bold font-headline text-white">Add New Luxury Unit</h3>
        <button on:click={() => isAdding = false} class="text-muted hover:text-white p-1 rounded-lg hover:bg-white/10 transition-colors" aria-label="Close modal">
          <svg class="w-4 h-4 text-gold" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
          </svg>
        </button>
      </div>

      <form on:submit|preventDefault={handleAddUnitSubmit} class="space-y-4 text-xs">
        <div class="grid grid-cols-2 gap-4">
          <div>
            <label for="new-unit-id" class="block font-headline text-muted uppercase mb-1">Unit ID *</label>
            <input
              id="new-unit-id"
              type="text"
              bind:value={newUnitId}
              placeholder="e.g. UNIT-106"
              class="w-full select-luxury py-2 px-3"
            />
          </div>

          <div>
            <label for="new-unit-project" class="block font-headline text-muted uppercase mb-1">Project Name</label>
            <select
              id="new-unit-project"
              bind:value={newProjectName}
              class="w-full select-luxury py-2 px-3"
            >
              <option value="Binghatti Mercedes-Benz Places">Binghatti Mercedes-Benz Places</option>
              <option value="Bugatti Residences by Binghatti">Bugatti Residences by Binghatti</option>
              <option value="Jacob & Co Residences">Jacob & Co Residences</option>
            </select>
          </div>
        </div>

        <div class="grid grid-cols-2 gap-4">
          <div>
            <label for="new-unit-type" class="block font-headline text-muted uppercase mb-1">Unit Layout Type</label>
            <select
              id="new-unit-type"
              bind:value={newUnitType}
              class="w-full select-luxury py-2 px-3"
            >
              <option value="Studio Luxury">Studio Luxury</option>
              <option value="1 Bedroom Suite">1 Bedroom Suite</option>
              <option value="2 Bedroom Royal">2 Bedroom Royal</option>
              <option value="Luxury Sky Penthouse">Luxury Sky Penthouse</option>
            </select>
          </div>

          <div>
            <label for="new-unit-price" class="block font-headline text-muted uppercase mb-1">Price (AED) *</label>
            <input
              id="new-unit-price"
              type="number"
              bind:value={newPriceAED}
              required
              class="w-full select-luxury py-2 px-3"
            />
          </div>
        </div>

        <div class="grid grid-cols-3 gap-4">
          <div>
            <label for="new-unit-size" class="block font-headline text-muted uppercase mb-1">Size (sq. ft.)</label>
            <input
              id="new-unit-size"
              type="number"
              bind:value={newSizeSqft}
              required
              class="w-full select-luxury py-2 px-3"
            />
          </div>

          <div>
            <label for="new-unit-handover" class="block font-headline text-muted uppercase mb-1">Handover</label>
            <input
              id="new-unit-handover"
              type="text"
              bind:value={newHandoverDate}
              required
              class="w-full select-luxury py-2 px-3"
            />
          </div>

          <div>
            <label for="new-unit-status" class="block font-headline text-muted uppercase mb-1">Status</label>
            <select
              id="new-unit-status"
              bind:value={newStatus}
              class="w-full select-luxury py-2 px-3"
            >
              <option value="Available">Available</option>
              <option value="Reserved">Reserved</option>
              <option value="Sold">Sold</option>
            </select>
          </div>
        </div>

        <div>
          <label for="new-unit-pdf" class="block font-headline text-muted uppercase mb-1">Floor Plan PDF URL</label>
          <input
            id="new-unit-pdf"
            type="text"
            bind:value={newPdfUrl}
            class="w-full select-luxury py-2 px-3"
          />
        </div>

        <div class="flex items-center justify-end gap-3 pt-4 border-t border-white/10">
          <button
            type="button"
            on:click={() => isAdding = false}
            class="px-4 py-2 rounded-xl bg-white/10 text-white text-xs font-semibold hover:bg-white/20"
          >
            Cancel
          </button>
          <button
            type="submit"
            class="px-5 py-2 rounded-xl bg-gold text-black font-headline font-bold text-xs uppercase tracking-wider hover:bg-gold/90 transition-all"
          >
            Create Unit
          </button>
        </div>
      </form>
    </div>
  </div>
{/if}
