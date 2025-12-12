console.log(1);
export default {
  extends: [
    'stylelint-config-standard',
    'stylelint-config-standard-scss',
    'stylelint-config-html',
    'stylelint-config-recommended-vue',
  ],
  plugins: ['stylelint-order'],

  rules: {
    'order/properties-alphabetical-order': true,
    'no-empty-source': null,
  },
};
