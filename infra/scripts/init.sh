#!/bin/bash
set -e

apt-get update
apt-get upgrade -y

apt-get install -y docker.io git

systemctl enable docker
systemctl start docker

useradd -m -s /bin/bash appuser
usermod -aG docker appuser

su - appuser << 'EOF'
cd ~
git clone https://github.com/yourusername/yourrepo.git app
cd app
docker build -t me-app .
EOF

cat > /etc/systemd/system/me-app.service << 'EOF'
[Unit]
Description=Personal Website
After=docker.service
Requires=docker.service

[Service]
Type=simple
User=appuser
ExecStartPre=-/usr/bin/docker stop me-app
ExecStartPre=-/usr/bin/docker rm me-app
ExecStart=/usr/bin/docker run --name me-app -p 3000:3000 me-app
ExecStop=/usr/bin/docker stop me-app
Restart=always
RestartSec=10

[Install]
WantedBy=multi-user.target
EOF

systemctl daemon-reload
systemctl enable me-app
systemctl start me-app

ufw allow 22/tcp
ufw allow 80/tcp
ufw allow 443/tcp
ufw allow 3000/tcp
ufw --force enable

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
