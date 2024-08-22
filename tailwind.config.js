/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./views/**/*.{html,templ,go}",
  ],
  theme: {
    fontFamily: {
      'montserrat': ['Montserrat', 'Arial', 'sans-serif'],
      'raleway': ["Raleway", 'sans-serif']
    },
    extend: {},
  },
  plugins: [
    require('flowbite/plugin')
  ],
}

