/** @type {import('tailwindcss').Config} */
export default {
    content: ['./src/**/*.{astro,html,js,jsx,md,mdx,svelte,ts,tsx,vue}'],
    theme: {
        extend: {
            colors: {
                background: '#262626',
                special: '#facc15',
                normal: '#a3a3a3',
                button: '#e14d0b',
                error: '#fb4934',
                'button-hover': '#c1410a',
                'gray-dark': '#1c1c1c',
            },
            backgroundImage: {
                'special-gradient': 'linear-gradient(to right, #a3a3a3, #facc15)',
            },
            animation: {
                'fade-in': 'fadeIn 0.6s ease-out forwards',
            },
            backdropBlur: {
                'lg': '16px',
            },
            keyframes: {
                rotation: {
                    '0%': { transform: 'rotate(0deg)' },
                    '100%': { transform: 'rotate(360deg)' },
                },
                rotationBack: {
                    '0%': { transform: 'rotate(0deg)' },
                    '100%': { transform: 'rotate(-360deg)' },
                },
            },
        },
    },
    plugins: [
        function ({ addComponents }) {
            addComponents({
                '.loader': {
                    width: '25px',
                    height: '25px',
                    borderWidth: '3px',
                    borderColor: '#FFF',
                    borderStyle: 'solid solid dotted dotted',
                    borderRadius: '50%',
                    display: 'inline-block',
                    position: 'relative',
                    boxSizing: 'border-box',
                    animation: 'rotation 2s linear infinite',
                },
                '.loader::after': {
                    content: "''",
                    boxSizing: 'border-box',
                    position: 'absolute',
                    left: '0',
                    right: '0',
                    top: '0',
                    bottom: '0',
                    margin: 'auto',
                    borderWidth: '3px',
                    borderColor: '#facc15',
                    borderStyle: 'solid solid dotted',
                    width: '12px',
                    height: '12px',
                    borderRadius: '50%',
                    animation: 'rotationBack 1s linear infinite',
                    transformOrigin: 'center center',
                },
            })
        }
    ],
}
