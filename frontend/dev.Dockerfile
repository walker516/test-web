FROM node:23-alpine

WORKDIR /app

# Install dependencies
COPY package.json package-lock.json* .npmrc* ./
RUN npm ci

COPY src ./src
COPY public ./public
COPY next.config.ts .
COPY tsconfig.json .

# Disable Next.js telemetry
ENV NEXT_TELEMETRY_DISABLED 1

CMD npm run dev