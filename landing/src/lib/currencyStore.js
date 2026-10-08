import { writable } from 'svelte/store';

export const currencies = {
  AED: { code: 'AED', symbol: 'AED', rate: 1, prefix: true },
  USD: { code: 'USD', symbol: '$', rate: 0.2723, prefix: true },
  EUR: { code: 'EUR', symbol: '€', rate: 0.2510, prefix: true },
  GBP: { code: 'GBP', symbol: '£', rate: 0.2145, prefix: true }
};

export const currentCurrency = writable('AED');
