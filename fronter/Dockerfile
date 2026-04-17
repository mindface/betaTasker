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

RUN addgroup --system --gid 1001 nodejs && \
    adduser --system --uid 1001 nextjs

COPY --from=builder /app/.next/standalone ./
COPY --from=builder /app/.next/static ./.next/static
COPY --from=builder /app/public ./public

RUN chown -R nextjs:nodejs .next
USER nextjs

EXPOSE 3000

CMD ["node", "server.js"]