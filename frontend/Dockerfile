# Base image
FROM node:20-alpine

# Set working directory
WORKDIR /app

# Copy files and install deps
COPY package.json package-lock.json ./
RUN npm install

# Copy rest of the app
COPY . .

# Build the Next.js app
RUN npm run build

# Expose port and run Next.js
EXPOSE 3000
CMD ["npm", "start"]
