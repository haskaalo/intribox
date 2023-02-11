module.exports = {
    root: true,
    parser: '@typescript-eslint/parser',
    plugins: [
      '@typescript-eslint',
    ],
    extends: [
        'airbnb',
        'airbnb-typescript',
        'prettier'
    ],
    parserOptions: {
        project: './tsconfig.json'
    },
    rules: {
        "react/prefer-stateless-function": 0,
        "import/prefer-default-export": 0,
        "no-alert": 0,
        "@typescript-eslint/default-param-last": 0,
        "no-param-reassign": 0,
        "no-plusplus": 0,
    }
  };