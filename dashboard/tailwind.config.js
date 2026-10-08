/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],
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
        },
        muted: '#A1A1AA',
      },
      fontFamily: {
        headline: ['Syne', 'Inter', 'sans-serif'],
        body: ['Inter', 'sans-serif'],
      },
    },
  },
  plugins: [],
};
