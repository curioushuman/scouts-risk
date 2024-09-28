/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './web/**/*.html',
    './web/**/*.templ',
    './internal/ui/**/*.templ',
    './internal/icons/**/*.templ',
  ],
  theme: {
    extend: {},
  },
  plugins: [require('@tailwindcss/typography'), require('daisyui')],
  daisyui: {
    themes: ['light', 'lemonade'],
  },
};
