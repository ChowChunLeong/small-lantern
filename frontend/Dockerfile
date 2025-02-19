# Use an official Node.js runtime as a base image
FROM node:18-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Install pnpm globally before copying package.json
RUN npm install -g pnpm 

# Copy package.json and pnpm-lock.yaml for dependency installation
COPY package.json pnpm-lock.yaml ./

# Install dependencies with pnpm 
# --frozen-lockfile -> Guarantees that dependencies match exactly what is in pnpm-lock.yaml.
RUN pnpm install --frozen-lockfile

# Copy the rest of the application files
COPY . .

# Build the Next.js app
RUN pnpm build

# Production image
FROM node:18-alpine AS runner

WORKDIR /app

# Install pnpm in the production image
RUN npm install -g pnpm  

# Copy only necessary files from the builder stage
COPY --from=builder /app/package.json /app/pnpm-lock.yaml ./  
COPY --from=builder /app/.next ./.next  
COPY --from=builder /app/public ./public  
COPY --from=builder /app/node_modules ./node_modules  

# Expose the Next.js port
EXPOSE 3000

# Run Next.js in production mode
CMD ["pnpm", "start"]
