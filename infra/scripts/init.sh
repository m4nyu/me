#!/bin/bash
set -e

# Update system
apt-get update
apt-get upgrade -y

# Install Go
wget https://go.dev/dl/go1.24.7.linux-arm64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.24.7.linux-arm64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> /etc/profile
export PATH=$PATH:/usr/local/go/bin

# Install Git
apt-get install -y git

# Create app user
useradd -m -s /bin/bash appuser

# Clone and setup application
su - appuser << 'EOF'
cd ~
git clone https://github.com/yourusername/yourrepo.git app
cd app
/usr/local/go/bin/go mod download
/usr/local/go/bin/go build -o /home/appuser/app/server ./src/cmd/server
EOF

# Create systemd service
cat > /etc/systemd/system/me-app.service << 'EOF'
[Unit]
Description=Personal Website
After=network.target

[Service]
Type=simple
User=appuser
WorkingDirectory=/home/appuser/app
ExecStart=/home/appuser/app/server
Restart=always
RestartSec=10
Environment="PORT=3000"

[Install]
WantedBy=multi-user.target
EOF

# Enable and start service
systemctl daemon-reload
systemctl enable me-app
systemctl start me-app

# Setup firewall
ufw allow 22/tcp
ufw allow 80/tcp
ufw allow 443/tcp
ufw allow 3000/tcp
ufw --force enable

# Install and setup nginx as reverse proxy
apt-get install -y nginx
cat > /etc/nginx/sites-available/default << 'EOF'
server {
    listen 80;
    server_name _;

    location / {
        proxy_pass http://localhost:3000;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_cache_bypass $http_upgrade;
    }
}
EOF

systemctl restart nginx
