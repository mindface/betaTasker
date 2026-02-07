FROM node:20-alpine AS builder

WORKDIR /app

COPY package.json yarn.lock ./
RUN yarn

COPY . .
RUN yarn build

FROM node:20-alpine AS runner

WORKDIR /app

ENV NODE_ENV=production
ENV PORT=3000
ENV HOSTNAME="0.0.0.0"

# 非rootユーザーの作成（セキュリティ）
RUN addgroup --system --gid 1001 nodejs && \
    adduser --system --uid 1001 nextjs

COPY --from=builder /app ./

# .nextディレクトリの権限を設定
# 既存のディレクトリはそのまま、権限のみ変更
RUN chown -R nextjs:nodejs .next && \
    chmod -R 755 .next

USER nextjs

EXPOSE 3000

CMD ["yarn", "dev"]