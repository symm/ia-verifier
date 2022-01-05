# IA-Verifier

Verify downloads from the internet archive (archive.org)


## Usage

Given a previously downloaded an Item via [ia](https://archive.org/services/docs/api/internetarchive/cli.html) command line tool:

```
➜  ia download Alien_1984_-
Alien_1984_-: dddddddddddddddddddddddd - success
```

Verify the download:

```
➜  tmp cd Alien_1984_-
➜  Alien_1984_- ia-verifier
Missing:
Present:
 00_coverscreenshot.jpg
 00_coverscreenshot_thumb.jpg
 Alien_1984_-.d64
 Alien_1984_-.thumbs/history/files/00_coverscreenshot.jpg_000001.jpg
 Alien_1984_-_archive.torrent
 Alien_1984_-_meta.sqlite
 Alien_1984_-_meta.xml
 Alien_1984_-_screenshot.gif
 __ia_thumb.jpg
 screenshot_00.jpg
 screenshot_00_thumb.jpg
 screenshot_01.jpg
 screenshot_01_thumb.jpg
 screenshot_02.jpg
 screenshot_02_thumb.jpg
 screenshot_03.jpg
 screenshot_03_thumb.jpg
 screenshot_04.jpg
 screenshot_04_thumb.jpg
 screenshot_05.jpg
 screenshot_05_thumb.jpg
 screenshot_06.jpg
 screenshot_06_thumb.jpg
Missing: 0 | Present: 23
```

