FROM gitlab.epay.technology:5050/common/images/go:1.24.5-1

# Install NPM
RUN apk add --update npm

# Prepare GO module for dependency download
COPY project/go.mod ./go.mod
COPY project/go.sum ./go.sum

# Download dependencies
RUN go mod download

# Prepare NPM packages for dependency download
COPY project/Resources/js/package.json ./Resources/js/package.json
COPY project/Resources/js/package-lock.json ./Resources/js/package-lock.json
COPY project/Resources/payment-window-react/package.json ./Resources/payment-window-react/package.json
COPY project/Resources/payment-window-react/package-lock.json ./Resources/payment-window-react/package-lock.json

# Install dependencies for the client side script and build it
WORKDIR /project/Resources/js
RUN npm install

WORKDIR /project/Resources/payment-window-react
RUN npm install

# Reset working directory
WORKDIR /project

# Copy source files
COPY project /project

# Compile JS Clients
WORKDIR /project/Resources/js
RUN npm run build

# Compile Payment Window
WORKDIR /project/Resources/payment-window-react
RUN npm run build

# Compile service
WORKDIR /project
RUN go build -o ./service
