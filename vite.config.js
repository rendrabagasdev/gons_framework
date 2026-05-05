import { defineConfig } from 'vite';
import tailwindcss from '@tailwindcss/vite';
import fs from 'fs';
import path from 'path';

export default defineConfig({
  plugins: [
    tailwindcss(),
    {
      name: 'php-hot-file',
      configureServer(server) {
        const hotFilePath = path.resolve(__dirname, 'public/hot');
        fs.writeFileSync(hotFilePath, 'http://localhost:5173');
        
        server.httpServer.on('close', () => {
          if (fs.existsSync(hotFilePath)) {
            fs.unlinkSync(hotFilePath);
          }
        });
      }
    }
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