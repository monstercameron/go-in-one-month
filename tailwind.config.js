/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    'views/pages/**/*.templ',
    'views/partials/**/*.templ',
  ],
  darkMode: 'class',
  theme: {
    extend: {
      fontFamily: {
        mono: ['Courier Prime', 'monospace'],
      }
    },
  },
  plugins: [
  ],
  corePlugins: {
    preflight: true,
  }
}