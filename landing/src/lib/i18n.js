import { translations } from './translations.js';

export function getNestedTranslation(obj, path) {
  return path.split('.').reduce((prev, curr) => prev && prev[curr], obj);
}

export function updateDOMTranslations(lang = 'en') {
  if (typeof document === 'undefined') return;

  const t = translations[lang] || translations.en;
  const htmlEl = document.documentElement;

  if (lang === 'ar') {
    htmlEl.setAttribute('dir', 'rtl');
    htmlEl.setAttribute('lang', 'ar');
    htmlEl.classList.add('font-arabic');
  } else {
    htmlEl.setAttribute('dir', 'ltr');
    htmlEl.setAttribute('lang', 'en');
    htmlEl.classList.remove('font-arabic');
  }

  // 1. Update text content for data-i18n
  document.querySelectorAll('[data-i18n]').forEach((el) => {
    const key = el.getAttribute('data-i18n');
    const val = getNestedTranslation(t, key);
    if (val) {
      el.innerText = val;
    }
  });

  // 2. Update placeholders for data-i18n-placeholder
  document.querySelectorAll('[data-i18n-placeholder]').forEach((el) => {
    const key = el.getAttribute('data-i18n-placeholder');
    const val = getNestedTranslation(t, key);
    if (val) {
      el.setAttribute('placeholder', val);
    }
  });

  // 3. Update innerHTML for data-i18n-html
  document.querySelectorAll('[data-i18n-html]').forEach((el) => {
    const key = el.getAttribute('data-i18n-html');
    const val = getNestedTranslation(t, key);
    if (val) {
      el.innerHTML = val;
    }
  });
}
