#!/bin/bash

year=`date +'%Y'`
app=$1
name=$2
exec=$2
icon="$2.png"

cd build/bin

# Generate files
cat > LICENSE.txt << EOF
The MIT License (MIT)

Copyright © $year Enrique Marín

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
EOF
cat > README.txt << EOF
INSTALLING $app:
=========================

Since the supplied executable contains everything needed to run on Windows,
it can be considered portable software.
This means that you can run it from any location by simply double-clicking on it.

However, if you want to use it as an "installed" application
you can place the executable in the directory "C:\Program Files".
You will then need to add it to the system PATH as explained here:
https://www.eukhost.com/kb/how-to-add-to-the-path-on-windows-10-and-windows-11/

================================================================================

Enjoy it!!
EOF

echo "# If PREFIX isn't provided, we check for \$(DESTDIR)/usr/local and use that if it exists." > Makefile
echo '# Otherwice we fall back to using /usr.' >> Makefile
echo '' >> Makefile
echo 'LOCAL != test -d $(DESTDIR)/usr/local && echo -n "/local" || echo -n ""' >> Makefile
echo 'LOCAL ?= $(shell test -d $(DESTDIR)/usr/local && echo "/local" || echo "")' >> Makefile
echo 'PREFIX ?= /usr$(LOCAL)' >> Makefile
echo '' >> Makefile
echo "App := $app" >> Makefile
echo "Name := \"$name\"" >> Makefile
echo "Exec := \"$exec\"" >> Makefile
echo "Icon := \"$icon\"" >> Makefile
echo '' >> Makefile
echo 'default:' >> Makefile
echo '	# User install' >> Makefile
echo '	# Run "make user-install" to install in ~/.local/' >> Makefile
echo '	# Run "make user-uninstall" to uninstall from ~/.local/' >> Makefile
echo '	#' >> Makefile
echo '	# System install' >> Makefile
echo '	# Run "sudo make install" to install the application.' >> Makefile
echo '	# Run "sudo make uninstall" to uninstall the application.' >> Makefile
echo '' >> Makefile
echo 'install:' >> Makefile
echo '	install -Dm00644 usr/local/share/applications/$(Name).desktop $(DESTDIR)$(PREFIX)/share/applications/$(Name).desktop' >> Makefile
echo '	install -Dm00755 usr/local/bin/$(Exec) $(DESTDIR)$(PREFIX)/bin/$(Exec)' >> Makefile
echo '	install -Dm00644 usr/local/share/pixmaps/$(Icon) $(DESTDIR)$(PREFIX)/share/pixmaps/$(Icon)' >> Makefile
echo 'uninstall:' >> Makefile
echo '	-rm $(DESTDIR)$(PREFIX)/share/applications/$(Name).desktop' >> Makefile
echo '	-rm $(DESTDIR)$(PREFIX)/bin/$(Exec)' >> Makefile
echo '	-rm $(DESTDIR)$(PREFIX)/share/pixmaps/$(Icon)' >> Makefile
echo '' >> Makefile
echo 'user-install:' >> Makefile
echo '	install -Dm00644 usr/local/share/applications/$(Name).desktop $(DESTDIR)$(HOME)/.local/share/applications/$(Name).desktop' >> Makefile
echo '	install -Dm00755 usr/local/bin/$(Exec) $(DESTDIR)$(HOME)/.local/bin/$(Exec)' >> Makefile
echo '	install -Dm00644 usr/local/share/pixmaps/$(Icon) $(DESTDIR)$(HOME)/.local/share/icons/$(Icon)' >> Makefile
echo '	sed -i -e "s,Exec=$(Exec),Exec=$(DESTDIR)$(HOME)/.local/bin/$(Exec),g" $(DESTDIR)$(HOME)/.local/share/applications/$(Name).desktop' >> Makefile
echo '	sed -i -e "s,Name=$(Name),Name=$(App),g" $(DESTDIR)$(HOME)/.local/share/applications/$(Name).desktop' >> Makefile
echo '' >> Makefile
echo 'user-uninstall:' >> Makefile
echo '	-rm $(DESTDIR)$(HOME)/.local/share/applications/$(Name).desktop' >> Makefile
echo '	-rm $(DESTDIR)$(HOME)/.local/bin/$(Exec)' >> Makefile
echo '	-rm $(DESTDIR)$(HOME)/.local/share/icons/$(Icon)' >> Makefile

cat > $name.desktop << EOF
[Desktop Entry]
Type=Application
Name=$name
Exec=$name
Icon=$name
Comment=Minimalist password store.
Categories=Development;
Keywords=wails;
EOF

# Generate directories
mkdir -p usr/local/bin && mkdir -p usr/local/share/applications && mkdir -p usr/local/share/pixmaps

cd ..

# Places files in their destination
cp appicon.png bin/usr/local/share/pixmaps/ && mv bin/usr/local/share/pixmaps/appicon.png bin/usr/local/share/pixmaps/$name.png
cp bin/$name.desktop bin/usr/local/share/applications
cp bin/$name bin/usr/local/bin/

cd bin/

# Create the .tar.xz file
tar -czvf $name"_linux_x86_64.tar.xz" usr/ Makefile LICENSE.txt

# Delete unnecessary files
rm *.desktop Makefile && rm -r usr/

cd ../..