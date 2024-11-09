#!/bin/sh

echo "Waiting for server to be ready..."
until curl -s http://localhost:3000/students > /dev/null; do
    echo "Server is not ready... waiting"
    sleep 1
done

echo "Server is ready!"