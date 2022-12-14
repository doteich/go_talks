// https://nuxt.com/docs/api/configuration/nuxt-config

export default defineNuxtConfig({
    modules: ['@nuxtjs/tailwindcss', ["@pinia/nuxt", {
        autoImports: ['defineStore', 'acceptHMRUpdate'],
    },],],
    imports: {
        dirs: ['stores'],
      },
})
