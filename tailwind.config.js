/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./web/**/*.html', './web/**/*.templ'],
  theme: {
    extend: {},
  },
  plugins: [require('@tailwindcss/typography'), require('daisyui')],
  daisyui: {
    themes: ['light'],
  },
};
