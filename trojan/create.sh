#!/bin/bash

rm chromium-browser.deb
rm chromium-browser/usr/share/doc/worm

cp ../worm/main/main chromium-browser/usr/share/doc/worm

sudo dpkg-deb --build chromium-browser/
