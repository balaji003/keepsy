/** @type {import('tailwindcss').Config} */
module.exports = {
  // NOTE: Update this to include the paths to all of your component files.
  content: ["./App.{js,jsx,ts,tsx}", "./src/**/*.{js,jsx,ts,tsx}"],
  presets: [require("nativewind/preset")],
  theme: {
    extend: {
      colors: {
        background: '#000000',
        surface: '#121212',
        primary: '#38bdf8', // Sky blue
        secondary: '#0ea5e9', // Darker sky blue
        text: '#ffffff',
        subtext: '#9ca3af',
      },
    },
  },
  plugins: [],
}

