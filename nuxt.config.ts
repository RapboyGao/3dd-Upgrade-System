import { defineNuxtConfig } from 'nuxt/config';
import config from './config';

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  ...config,
  modules: ['@nuxt/ui'],
});