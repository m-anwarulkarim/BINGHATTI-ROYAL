import { writable } from 'svelte/store';

export const isSidebarOpen = writable(false);

export function toggleSidebar() {
  isSidebarOpen.update(v => !v);
}

export function closeSidebar() {
  isSidebarOpen.set(false);
}
