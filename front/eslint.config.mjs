// @ts-check
import withNuxt from './.nuxt/eslint.config.mjs';

import prettier from 'eslint-config-prettier';

export default withNuxt([
  prettier,
  {
    rules: {
      'vue/no-multiple-template-root': 'off',
      'vue/multi-word-component-names': 'off',
    },
  },
]);
