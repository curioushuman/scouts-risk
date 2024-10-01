/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './web/**/*.html',
    './web/**/*.templ',
    './internal/ui/**/*.templ',
    './internal/icons/**/*.templ',
  ],
  theme: {
    extend: {
      animation: {
        'bounce-slow': 'bounce-slow 10s infinite',
        fade: 'fadeOut 10s 1',
      },
      keyframes: {
        'bounce-slow': {
          '0%, 100%': {
            transform: 'translateY(0)',
            'animation-timing-function': 'cubic-bezier(0, 0, 0.2, 1)',
          },
          '95%': {
            transform: 'translateY(-25%)',
            'animation-timing-function': 'cubic-bezier(0, 0, 0.2, 1)',
          },
        },
        fadeOut: {
          '0%': {
            opacity: 1,
          },
          '90%': {
            opacity: 1,
          },
          '100%': {
            opacity: 0,
          },
        },
      },
    },
  },
  plugins: [require('@tailwindcss/typography'), require('daisyui')],
  daisyui: {
    themes: ['light', 'lemonade'],
  },
};
