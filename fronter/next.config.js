/** @type {import('next').NextConfig} */
module.exports = {
  output: 'export',
  reactStrictMode: true,
  async headers() {
    if (process.env.NODE_ENV === 'development') {
      return [
        {
          source: '/(.*)',
          headers: [
            {
              key: 'Content-Security-Policy',
              value: "default-src 'self' http: https: data: blob: 'unsafe-inline' 'unsafe-eval'",
            },
          ],
        },
      ]
    }
    return []
  },
}
