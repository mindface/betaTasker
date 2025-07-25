# ✅ Build stage
FROM node:20-alpine AS builder

WORKDIR /app

# 必要ファイルを先にコピーして依存関係を解決
COPY package.json yarn.lock ./
RUN yarn install --frozen-lockfile

# 残りのソースコードをコピーしてビルド
COPY . .
RUN yarn build

# ✅ Production stage
FROM node:20-alpine AS runner

WORKDIR /app

ENV NODE_ENV=production
ENV PORT=3000
ENV HOSTNAME="0.0.0.0"

# アプリ本体をコピー
COPY --from=builder /app ./

# 非rootユーザーの作成（セキュリティ）
RUN addgroup --system --gid 1001 nodejs && \
    adduser --system --uid 1001 nextjs
USER nextjs

EXPOSE 3000

# ✅ `next start` を実行（yarn 経由）
CMD ["yarn", "start"]
