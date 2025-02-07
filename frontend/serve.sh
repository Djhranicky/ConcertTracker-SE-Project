#!/bin/sh

ng serve &
gin --port 4201 --path . --build ../backend/cmd/ --i --all &

wait