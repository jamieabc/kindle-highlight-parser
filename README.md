This program parses kindle highlights. It searches current directory file of `My Clippings.txt`, parse it and output to `kindle.txt`

If it's a mac, kindle clipping file is located under `/Volumes/Kindle/documents/My Clippings.txt`

Note:
    Currently only test chinese version, not sure if other version with same syntax.
    If there's any problem, please send out issue and attach the `My Clippings.txt`

Install:
    `go get -u github.com/jamieabc/kindle-highlight-parser`

Usage:
    `kindle-highlight-parser`
