// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
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
  plugins: ['@@/plugins/mastercss.client.ts'],
  imports: {
    autoImport: false,
  },
  devtools: {
    enabled: true,
  },
  css: ['@@/assets/styles/app.scss'],
  compatibilityDate: '2025-07-15',
  eslint: {
    config: {
      stylistic: true,
    },
  },
});
