#!/bin/sh
cd ..
mv SovyGO michal
git clone -b develop --single-branch https://github.com/Martinhercka/SovyGO.git
mv SovyGO develop
echo "run: screen sh michal/update.sh"
