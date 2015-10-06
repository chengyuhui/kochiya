@echo off

cd %3\video_utils

ffmpeg -i "%~1" -f wav - | qaac\qaac --quality 2 --tvbr 127 --ignorelength -o "%~dpn1_a.m4a" -

avs4x264mod --x264-binary x264_64_tMod-8bit-all --level 4.1 --crf 21 --preset veryslow --tune animation  --threads 16 --vbv-bufsize 15000 --vbv-maxrate 13000 --ref 6 --opt 1 --deblock -1:-1 --b-adapt 2 --bframes 5 --keyint 300 --min-keyint 1 --direct auto --qcomp 0.6 --rc-lookahead 50 --me tesa --merange 40 -m 10 -t 2 --weightp 2  --aq-mode 3  --aq-strength 0.8 --psy-rd 0.0:2.0 --fade-compensate 0.2 --input-depth 8  --colormatrix bt709 --slow-firstpass  --acodec none --output "%~dpn1_v.mp4" "%~1"

remuxer -i "%~dpn1_v.mp4" -i "%~dpn1_a.m4a" -o "%~2"
