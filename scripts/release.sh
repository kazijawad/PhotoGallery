#!/bin/bash

cd "$GOPATH/src/PhotoGallery"

echo "[Releasing PhotoGallery]"
echo "  Deleting the local binary if it exists (so it isn't uploaded)..."
rm PhotoGallery
echo "  Done!"

echo "  Deleting existing code..."
ssh root@123.123.12.12 "rm -rf /root/go/src/PhotoGallery"
echo "  Code deleted successfully!"

echo "  Uploading code..."
rsync -avr --exclude '.git/*' --exclude 'images/*' ./ root@123.123.12.12:/root/go/src/PhotoGallery/
echo "  Code uploaded successfully!"

echo "  Go getting deps..."
ssh root@123.123.12.12 "export GOPATH=/root/go; /usr/local/go/bin/go get golang.org/x/crypto/bcrypt"
ssh root@123.123.12.12 "export GOPATH=/root/go; /usr/local/go/bin/go get github.com/gorilla/mux"
ssh root@123.123.12.12 "export GOPATH=/root/go; /usr/local/go/bin/go get github.com/gorilla/schema"
ssh root@123.123.12.12 "export GOPATH=/root/go; /usr/local/go/bin/go get github.com/lib/pq"
ssh root@123.123.12.12 "export GOPATH=/root/go; /usr/local/go/bin/go get github.com/jinzhu/gorm"
ssh root@123.123.12.12 "export GOPATH=/root/go; /usr/local/go/bin/go get github.com/gorilla/csrf"

echo "  Building the code on remote server..."
ssh root@123.123.12.12 "export GOPATH=/root/go; cd /root/app; /usr/local/go/bin/go build -o ./server $GOPATH/src/PhotoGallery/*.go"
echo "  Code built successfully!"

echo "  Moving assets..."
ssh root@123.123.12.12 "cd /root/app; cp -R /root/go/src/PhotoGallery/assets ."
echo "  Assets moved successfully!"

echo "  Moving views..."
ssh root@123.123.12.12 "cd /root/app; cp -R /root/go/src/PhotoGallery/views ."
echo "  Views moved successfully!"

echo "  Moving Caddyfile..."
ssh root@123.123.12.12 "cd /root/app; cp /root/go/src/PhotoGallery/Caddyfile ."
echo "  Caddyfile moved successfully!"

echo "  Restarting the server..."
ssh root@123.123.12.12 "sudo service photogallery restart"
echo "  Server restarted successfully!"

echo "  Restarting Caddy server..."
ssh root@123.123.12.12 "sudo service caddy restart"
echo "  Caddy restarted successfully!"

echo "[Done Releasing PhotoGallery]"
