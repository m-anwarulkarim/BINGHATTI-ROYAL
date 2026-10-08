import { writable } from 'svelte/store';

const initialToken = typeof localStorage !== 'undefined' ? localStorage.getItem('binghatti_token') : null;
const initialUser = typeof localStorage !== 'undefined' ? JSON.parse(localStorage.getItem('binghatti_user') || 'null') : null;

export const authToken = writable(initialToken);
export const currentUser = writable(initialUser);

export function setAuth(token, user) {
  authToken.set(token);
  currentUser.set(user);
  if (typeof localStorage !== 'undefined') {
    localStorage.setItem('binghatti_token', token);
    localStorage.setItem('binghatti_user', JSON.stringify(user));
  }
}

export function clearAuth() {
  authToken.set(null);
  currentUser.set(null);
  if (typeof localStorage !== 'undefined') {
    localStorage.removeItem('binghatti_token');
    localStorage.removeItem('binghatti_user');
  }
}
