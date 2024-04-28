import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

import PlatformVariables from "./PlatformVariables.config"

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    host: PlatformVariables.Frontend.hostAddress,
    port: PlatformVariables.Frontend.hostPort
  }
})
