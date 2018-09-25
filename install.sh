#!/bin/bash

echo "Building program..."
go build
echo "Copying to /usr/local/bin/concards, needs root."
sudo cp concards /usr/local/bin/concards;
