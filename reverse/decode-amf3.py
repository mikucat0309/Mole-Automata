import sys
import zlib
import json
from pyamf import amf3
from pathlib import Path

def main():
    if len(sys.argv) < 2:
        return
    f = Path(sys.argv[1])
    df = Path(sys.argv[2]) if len(sys.argv) >= 3 else f.parent / f'{f.stem}.json'

    with open(f, 'rb') as fp:
        buf = fp.read()
    buf = zlib.decompress(buf)
    ba = amf3.ByteArray(buf)
    obj = ba.readObject()
    
    with open(df, 'w') as fp:
        json.dump(obj, fp, ensure_ascii=False)

if __name__ == '__main__':
    main()
