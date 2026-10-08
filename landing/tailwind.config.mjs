/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{astro,html,js,jsx,md,mdx,svelte,ts,tsx}'],
  theme: {
    extend: {
      colors: {
        bg: {
          offblack: '#0A0A0A',
          dark: '#121212',
          card: '#18181B',
          border: '#27272A',
        },
        gold: {
          DEFAULT: '#D4AF37',
          champagne: '#D4AF37',
          light: '#F3E5AB',
          hover: '#B5942B',
          muted: 'rgba(212, 175, 55, 0.15)',
        },
        chrome: '#E5E5E5',
        muted: '#A1A1AA',
      },
      fontFamily: {
        headline: ['Syne', 'Cinzel', 'serif'],
        body: ['Inter', 'Plus Jakarta Sans', 'sans-serif'],
        arabic: ['Cairo', 'Tajawal', 'sans-serif'],
      },
      backgroundImage: {
        'gold-gradient': 'linear-gradient(135deg, #F3E5AB 0%, #D4AF37 50%, #AA820A 100%)',
        'dark-gradient': 'linear-gradient(180deg, rgba(10, 10, 10, 0.6) 0%, #0A0A0A 100%)',
      },
    },
  },
  plugins: [],
};
