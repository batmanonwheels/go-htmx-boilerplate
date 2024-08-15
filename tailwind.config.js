/** @type {import('tailwindcss').Config} */
const defaultTheme = require('tailwindcss/defaultTheme');
module.exports = {
	content: ['./templates/*.html', './templates/components/*.html'],
	theme: {
		extend: {},
	},
	plugins: [],
};
