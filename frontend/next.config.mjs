/** @type {import('next').NextConfig} */
const nextConfig = {
  experimental: {
    wasm: true,  // Enable WebAssembly support
    esmExternals: true,  // Allow ES modules to be used in externals
  },
  webpack: (config, { isServer }) => {
    // Handle WebAssembly files
    config.experiments = {
      asyncWebAssembly: true,
      layers: true,
    };
    
    // Add rule to handle WASM files
    config.module.rules.push({
      test: /\.wasm$/,
      type: 'webassembly/async',
      loader: 'file-loader',
    });

    // Ensure file-loader is not required in server-side build
    if (isServer) {
      config.externals = [...config.externals, 'file-loader'];
    }

    return config;
  },
};

export default nextConfig;
