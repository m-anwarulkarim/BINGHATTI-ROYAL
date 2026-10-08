import { writable, derived } from 'svelte/store';
import { translations } from './translations.js';

const initialLang = typeof localStorage !== 'undefined' ? (localStorage.getItem('binghatti_lang') || 'en') : 'en';

export const currentLang = writable(initialLang);

export const currentDir = derived(currentLang, $lang => $lang === 'ar' ? 'rtl' : 'ltr');

export const t = derived(currentLang, $lang => translations[$lang] || translations.en);

export function setLanguage(lang) {
  currentLang.set(lang);
  if (typeof localStorage !== 'undefined') {
    localStorage.setItem('binghatti_lang', lang);
  }
}
