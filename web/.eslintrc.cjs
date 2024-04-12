/* eslint-env node */
require('@rushstack/eslint-patch/modern-module-resolution')

module.exports = {
  root: true,
  'extends': [
    'plugin:vue/vue3-essential',
    'eslint:recommended',
    '@vue/eslint-config-prettier/skip-formatting'
  ],
  global: {
  },
  parserOptions: {
    ecmaVersion: 'latest'
  },
  rules: {
    'no-console': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
    'no-debugger': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
    'linebreak-style': ['off', 'windows'],
    'no-plusplus': 'off',
    'eqeqeq': 'off',
    'comma-dangle': 'off',
    'arrow-parens': 'off',
    'implicit-arrow-linebreak': 'off',
    'function-paren-newline': 'off',
    'prefer-promise-reject-errors': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
    'no-else-return': 'off',
    "vue/multi-word-component-names": 'off',
  },
}
