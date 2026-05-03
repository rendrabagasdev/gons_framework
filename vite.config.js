import { defineConfig } from 'vite';
import tailwindcss from '@tailwindcss/vite';

export default defineConfig({
  plugins: [
    tailwindcss(),
  ],
  server: {
    port: 5173,
    strictPort: true,
    cors: true,
  },
  build: {
    outDir: 'public',
    emptyOutDir: false,
    rollupOptions: {
      input: {
        'app': 'resources/js/app.js',
        'style': 'resources/css/app.css'
      },
      output: {
        entryFileNames: 'js/[name].js',
        assetFileNames: 'css/[name].[ext]'
      }
    }
  }
});