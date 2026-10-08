import { writable } from 'svelte/store';

function getStatusOverrides() {
  if (typeof window === 'undefined') return {};
  try {
    const saved = localStorage.getItem('binghatti_lead_status_overrides');
    return saved ? JSON.parse(saved) : {};
  } catch (e) {
    return {};
  }
}

function applyStatusOverrides(leadsList) {
  const overrides = getStatusOverrides();
  if (!overrides || Object.keys(overrides).length === 0) return leadsList;

  return leadsList.map(l => {
    if (overrides[l.id]) {
      return { ...l, status: overrides[l.id] };
    }
    return l;
  });
}

function loadInitialLeads() {
  if (typeof window === 'undefined') return [];
  try {
    const saved = localStorage.getItem('binghatti_submitted_leads');
    if (saved) {
      const parsed = JSON.parse(saved);
      if (Array.isArray(parsed) && parsed.length > 0) {
        return applyStatusOverrides(parsed);
      }
    }
  } catch (e) {}
  return [];
}

export const leadsStore = writable(loadInitialLeads());

const API_BASE = (typeof import.meta !== 'undefined' && import.meta.env && import.meta.env.PUBLIC_API_URL) 
  ? import.meta.env.PUBLIC_API_URL 
  : (typeof window !== 'undefined' && (window.location.hostname === 'localhost' || window.location.hostname === '127.0.0.1')
    ? 'http://localhost:8085' 
    : 'https://binghatti-royal.onrender.com');

export async function fetchLeadsFromAPI() {
  try {
    const res = await fetch(`${API_BASE}/api/v1/leads`);
    if (res.ok) {
      const json = await res.json();
      if (json.success && Array.isArray(json.data)) {
        const apiLeads = json.data.map(l => ({
          id: l.id || `lead-${Date.now()}`,
          full_name: l.full_name || l.FullName || 'VIP Investor',
          whatsapp_number: l.whatsapp_number || l.WhatsAppNumber || '',
          email: l.email || l.Email || '',
          budget_range: l.budget_range || l.BudgetRange || '$500k-$1M',
          investment_purpose: l.investment_purpose || l.InvestmentPurpose || 'Investment',
          status: l.status || l.Status || 'new',
          created_at: l.created_at || l.CreatedAt || new Date().toISOString()
        }));

        const leadsWithOverrides = applyStatusOverrides(apiLeads);

        leadsStore.update(existing => {
          const apiIds = new Set(leadsWithOverrides.map(l => l.id));
          const rest = existing.filter(l => !apiIds.has(l.id));
          return applyStatusOverrides([...leadsWithOverrides, ...rest]);
        });
      }
    }
  } catch (e) {
    // API offline or fetch error
  }
}

export async function updateLeadStatus(leadId, newStatus) {
  if (!leadId || !newStatus) return;

  try {
    if (typeof window !== 'undefined') {
      const overrides = getStatusOverrides();
      overrides[leadId] = newStatus;
      localStorage.setItem('binghatti_lead_status_overrides', JSON.stringify(overrides));
    }
  } catch (e) {}

  leadsStore.update(leads => {
    return leads.map(l => l.id === leadId ? { ...l, status: newStatus } : l);
  });

  try {
    await fetch(`${API_BASE}/api/v1/leads/${leadId}/status`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ status: newStatus })
    });
  } catch (err) {
    console.warn('Failed to persist lead status update to API:', err);
  }
}

export async function createManualLead(leadData) {
  const newLead = {
    id: `lead-${Date.now()}`,
    full_name: leadData.full_name,
    whatsapp_number: leadData.whatsapp_number,
    email: leadData.email,
    budget_range: leadData.budget_range || '$500k-$1M',
    investment_purpose: leadData.investment_purpose || 'Investment',
    status: 'new',
    created_at: new Date().toISOString()
  };

  try {
    const res = await fetch(`${API_BASE}/api/v1/leads`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(leadData)
    });
    if (res.ok) {
      const json = await res.json();
      if (json.success && json.data) {
        newLead.id = json.data.id || newLead.id;
      }
    }
  } catch (err) {}

  leadsStore.update(existing => [newLead, ...existing]);
  return newLead;
}

if (typeof window !== 'undefined') {
  fetchLeadsFromAPI();
  setInterval(fetchLeadsFromAPI, 3000);
}
