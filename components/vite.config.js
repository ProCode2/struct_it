import { defineConfig } from "vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";
import { resolve } from "path";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [svelte({})],
  build: {
    outDir: "../static/build", // Specify the output directory for the build files
    rollupOptions: {
      input: {
        component1: resolve(__dirname, "src", "component1.js"),
        component2: resolve(__dirname, "src", "component2.js"),
      },
      output: {
        entryFileNames: "[name].js", // Specify the format for the entry files
        chunkFileNames: "chunks/[name].js", // Specify the format for the chunk files (optional)
        assetFileNames: "assets/[name].[ext]", // Specify the format for the asset files (optional)
      },
    },
  },
});
