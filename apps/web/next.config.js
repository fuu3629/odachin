// /** @type {import('next').NextConfig} */
// const nextConfig = {
//   reactStrictMode: true,
//   output: 'standalone',
//   experimental: {
//     appDir: true /* This is not part of the react-md-editor configuration */,
//     esmExternals: 'loose' /* For react-md-editor */,
//   },
// };

// export default nextConfig;

// eslint-disable-next-line @typescript-eslint/no-require-imports
const removeImports = require('next-remove-imports');

module.exports = removeImports()({
  // ✅  options...
  webpack: function (config) {
    config.module.rules.push({
      test: /\.md$/,
      use: 'raw-loader',
    });
    return config;
  },
  reactStrictMode: true,
  output: 'standalone',
});
