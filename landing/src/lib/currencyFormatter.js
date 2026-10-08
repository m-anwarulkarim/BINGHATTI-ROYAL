import { currencies } from './currencyStore.js';

export function convertAED(aedAmount, currencyCode) {
  const info = currencies[currencyCode] || currencies.AED;
  const converted = Math.round(aedAmount * info.rate);
  
  if (currencyCode === 'AED') {
    return `AED ${aedAmount.toLocaleString('en-US')}`;
  }
  return `${info.symbol}${converted.toLocaleString('en-US')} ${info.code}`;
}

export function convertAEDShort(aedAmount, currencyCode) {
  const info = currencies[currencyCode] || currencies.AED;
  const converted = aedAmount * info.rate;
  
  if (currencyCode === 'AED') {
    return `AED ${(aedAmount / 1000000).toFixed(1)}M`;
  }
  if (converted >= 1000000) {
    return `${info.symbol}${(converted / 1000000).toFixed(2)}M`;
  }
  return `${info.symbol}${Math.round(converted / 1000)}K`;
}
