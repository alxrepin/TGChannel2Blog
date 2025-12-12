// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2025-07-15',
  devtools: {
    enabled: true,
  },
  server: {
    host: '0.0.0.0',
  },
  css: ['@@/assets/styles/app.scss'],
  plugins: ['@@/plugins/mastercss.client.ts'],
  eslint: {
    config: {
      stylistic: true,
    },
  },
  modules: [
    '@nuxt/eslint',
    '@nuxt/fonts',
    '@nuxt/hints',
    '@nuxt/image',
    '@nuxt/scripts',
    '@nuxt/icon',
    [
      '@nuxtjs/google-fonts',
      {
        families: {
          Inter: '100..900',
        },
      },
    ],
  ],
});
