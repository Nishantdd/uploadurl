// @ts-check
import { defineConfig, envField } from 'astro/config';

import tailwind from '@astrojs/tailwind';

import netlify from "@astrojs/netlify";

import react from "@astrojs/react";

import icon from "astro-icon";

// https://astro.build/config
export default defineConfig({
    integrations: [tailwind(), react(), icon()],
    output: "server",
    env: {
        schema: {
            SERVER_ADDRESS: envField.string({ context: "client", access: "public", optional: false })
        }
    },
    server: {
        port: 5173
    },
    adapter: netlify()
});