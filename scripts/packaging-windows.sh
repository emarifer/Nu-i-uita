#!/bin/bash

name=$1

cd build/bin

mkdir dist

cp $name".exe" dist/

# Create the .zip file
zip -9r $name"_windows_x86_64.zip" *.txt dist/

# Delete unnecessary files
rm *.txt && rm -r dist/

cd ../..


# zip -9r wails-todoapp_windows_x86_64.zip *.txt dist/