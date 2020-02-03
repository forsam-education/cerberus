// https://eslint.org/docs/user-guide/configuring

module.exports = {
  root: true,
  parserOptions: {
    parser: 'babel-eslint'
  },
  env: {
    browser: true
  },
  extends: ['plugin:vue/essential', 'standard', 'prettier/vue'],
  // required to lint *.vue files
  plugins: ['vue'],
  // add your custom rules here
  rules: {
    // allow async-await
    'generator-star-spacing': 'off',
    // allow debugger during development
    'no-debugger': process.env.NODE_ENV === 'production' ? 'error' : 'off',
    'object-shorthand': 'error',
    semi: ['error', 'always'],
    'comma-dangle': ['error', 'always-multiline'],
    quotes: ['error', 'single', { allowTemplateLiterals: true }],
    'prefer-template': 'error',
    'no-useless-concat': 'error',
    'no-unused-vars': 'error',
    'no-undef-init': 'error',
    'no-undef': 'error',
    'no-const-assign': 'error',
    'prefer-const': 'error',
    'no-var': 'error',
    'eol-last': ['error', 'always'],
    'padding-line-between-statements': [
      'error',
      { blankLine: 'always', prev: '*', next: 'return' },
      { blankLine: 'always', prev: 'directive', next: '*' },
      { blankLine: 'always', prev: '*', next: 'class' },
      { blankLine: 'always', prev: '*', next: 'export' },
      { blankLine: 'any', prev: 'export', next: 'export' },
      { blankLine: 'always', prev: '*', next: 'const' },
      { blankLine: 'any', prev: 'const', next: 'const' },
      { blankLine: 'always', prev: '*', next: 'function' },
      { blankLine: 'any', prev: 'const', next: 'function' },
      { blankLine: 'always', prev: '*', next: 'if' },
      { blankLine: 'any', prev: 'const', next: 'if' }
    ]
  }
};
