const defaultTheme = require('tailwindcss/defaultTheme')

module.exports = {
  content: ['layouts/**/*.html', 'content/**/*.html'],
  theme: {
    extend: {
      fontFamily: {
        sans: ['Inter', ...defaultTheme.fontFamily.sans],
        display: ['Montserrat', 'Inter', ...defaultTheme.fontFamily.sans],
      }
    }
  },
  plugins: [
    require('@tailwindcss/typography'),
  ],
}
