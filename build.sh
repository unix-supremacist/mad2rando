VARDIR=var
VERSION=$(cat $VARDIR/version)
PYV=$(cat $VARDIR/pyv)
PYF=python-$PYV-embed-amd64
PYDL=https://www.python.org/ftp/python/$PYV/$PYF.zip
CACHEDIR=cache
BLDDIR=bld
OUTDIR=out
NAME=mad2rando

mkdir "$CACHEDIR/" -p
rm -r "${BLDDIR:?}/"
mkdir "$BLDDIR/" -p
mkdir "$OUTDIR/" -p

[ ! -f "$CACHEDIR/$PYF.zip" ] && wget -q "$PYDL" -P "$CACHEDIR/"
[ -f "$CACHEDIR/$PYF.zip" ] && [ ! -d "$CACHEDIR/$PYF" ] && 7z x "$CACHEDIR/$PYF.zip" -o"$CACHEDIR/$PYF" > /dev/null
[ -d "$CACHEDIR/$PYF" ] && cp -r "$CACHEDIR/$PYF"/* "$BLDDIR/"
[ -d src/ ] && cp -r src/* "$BLDDIR/"
[ -f "$OUTDIR/$NAME-$VERSION.tar.lzma" ] && rm "$OUTDIR/$NAME-$VERSION.tar.lzma"
[ -d "$BLDDIR/" ] && cd "$BLDDIR/" && tar cvf "../$OUTDIR/$NAME-$VERSION.tar" . > /dev/null && cd ..
[ -f "$OUTDIR/$NAME-$VERSION.tar" ] && lzma "$OUTDIR/$NAME-$VERSION.tar"