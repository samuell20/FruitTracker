/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
}
const redirects = {
  async redirects() {
    return [
      {
        source: '/admin',
        destination: '/admin/dashboard',
        permanent: true,
      },
    ]
  },
}
module.exports = nextConfig
module.exports = redirects 